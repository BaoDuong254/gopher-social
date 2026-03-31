package main

import (
	"log"

	"github.com/baoduong254/gopher-social/internal/db"
	"github.com/baoduong254/gopher-social/internal/env"
	"github.com/baoduong254/gopher-social/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:password@localhost:5432/gopher_social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

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
	log.Printf("Database connection pool established")
	store := store.NewStorage(db)

	// Initialize a new instance of our application struct, containing the config and store objects.
	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Printf("Starting API server on %s", cfg.addr)
	err = app.run(mux)
	if err != nil {
		log.Fatal(err)
	}
}
