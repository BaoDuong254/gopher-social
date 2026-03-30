package main

import (
	"log"

	"github.com/baoduong254/gopher-social/internal/env"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()
	log.Printf("Starting API server on %s", cfg.addr)
	err = app.run(mux)
	if err != nil {
		log.Fatal(err)
	}
}
