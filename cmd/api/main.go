package main

import (
	"fmt"
	"log"
	"time"

	"github.com/baoduong254/gopher-social/internal/db"
	"github.com/baoduong254/gopher-social/internal/env"
	"github.com/baoduong254/gopher-social/internal/store"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const version = "0.0.1"

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
	// Load environment variables from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("API_URL", "localhost:8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:password@localhost:5432/gopher_social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp: time.Hour * 24 * 3, // 3 days
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
	logger.Info("Database connection pool established")
	store := store.NewStorage(db)

	// Initialize a new instance of our application struct, containing the config and store objects.
	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}
	mux := app.mount()
	logger.Info(fmt.Sprintf("Starting API server on http://localhost%s", cfg.addr))
	err = app.run(mux)
	if err != nil {
		log.Fatal(err)
	}
}
