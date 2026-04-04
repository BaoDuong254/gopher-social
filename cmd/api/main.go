package main

import (
	"expvar"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/baoduong254/gopher-social/internal/auth"
	"github.com/baoduong254/gopher-social/internal/db"
	"github.com/baoduong254/gopher-social/internal/env"
	"github.com/baoduong254/gopher-social/internal/mailer"
	"github.com/baoduong254/gopher-social/internal/ratelimiter"
	"github.com/baoduong254/gopher-social/internal/store"
	"github.com/baoduong254/gopher-social/internal/store/cache"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const version = "1.3.0"

type mailTrapSender interface {
	Send(string, string, string, any, bool) (int, error)
}

type mailTrapClientAdapter struct {
	client mailTrapSender
}

func (a mailTrapClientAdapter) Send(to string, subject string, templateFile string, data any, isSandbox bool) error {
	_, err := a.client.Send(to, subject, templateFile, data, isSandbox)
	return err
}

//	@title			Gopher Social API
//	@version		1.0
//	@description	This is the API documentation for Gopher Social, a simple social media platform built with Go.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token.
func main() {
	// In cloud runtimes, environment variables are injected by the platform and .env may not exist.
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file loaded, using runtime environment variables: %v", err)
	}

	port := env.GetString("PORT", "8080")
	defaultAddr := ":" + port

	var err error
	cfg := config{
		addr:        env.GetString("ADDR", defaultAddr),
		apiURL:      env.GetString("API_URL", "localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:5173"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:password@localhost:5432/gopher_social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp:       time.Hour * 24 * 3, // 3 days
			fromEmail: env.GetString("SENDGRID_FROM_EMAIL", ""),
			sendGrid: sendGridConfig{
				apiKey: env.GetString("SENDGRID_API_KEY", ""),
			},
			mailTrap: mailTrapConfig{
				apiKey: env.GetString("MAILTRAP_API_KEY", ""),
			},
		},
		auth: authConfig{
			basic: basicAuthConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "password"),
			},
			token: tokenAuthConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "supersecretkey"),
				exp:    time.Hour * 24 * 3, // 3 days
				iss:    "gopher-social.com",
			},
		},
		redisCfg: redisConfig{
			addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			pw:      env.GetString("REDIS_PASSWORD", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", true),
		},
		rateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUESTS_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Println("failed to sync logger:", err)
		}
	}()
	logger.Infof("Starting application in %s environment", cfg.env)

	// Initialize a connection pool to the database, passing in all the relevant configuration settings.
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("failed to close db:", err)
		}
	}()
	// Cache
	var rdb *redis.Client
	if cfg.redisCfg.enabled {
		rdb = cache.NewRedisClient(cfg.redisCfg.addr, cfg.redisCfg.pw, cfg.redisCfg.db)
		logger.Info("Redis cache enabled")
	}
	logger.Info("Database connection pool established")
	store := store.NewStorage(db)
	cacheStorage := cache.NewRedisStorage(rdb)

	// Rate Limiter
	ratelimiter := ratelimiter.NewFixedWindowRateLimiter(
		cfg.rateLimiter.RequestsPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)

	// Mailer
	// mailer := mailer.NewSendGrid(cfg.mail.sendGrid.apiKey, cfg.mail.fromEmail)
	mailtrap, err := mailer.NewMailTrapClient(cfg.mail.mailTrap.apiKey, cfg.mail.fromEmail)
	if err != nil {
		log.Panic(err)
	}

	// JWT Authenticator
	jwtAuthenticator := auth.NewJWTAuthenticator(cfg.auth.token.secret, cfg.auth.token.iss, cfg.auth.token.iss)

	// Initialize a new instance of our application struct, containing the config and store objects.
	app := &application{
		config:        cfg,
		store:         store,
		logger:        logger,
		mailer:        mailTrapClientAdapter{client: mailtrap},
		authenticator: jwtAuthenticator,
		cacheStorage:  cacheStorage,
		rateLimiter:   ratelimiter,
	}

	// Metrics collection
	expvar.NewString("version").Set(version)
	expvar.Publish("database", expvar.Func(func() any {
		return db.Stats()
	}))
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	mux := app.mount()
	serverURL := "http://" + cfg.addr
	if strings.HasPrefix(cfg.addr, ":") {
		serverURL = "http://localhost" + cfg.addr
	}
	logger.Info(fmt.Sprintf("Starting API server on %s", serverURL))
	err = app.run(mux)
	if err != nil {
		log.Fatal(err)
	}
}
