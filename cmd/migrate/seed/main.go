package main

import (
	"log"

	"github.com/baoduong254/gopher-social/internal/db"
	"github.com/baoduong254/gopher-social/internal/env"
	"github.com/baoduong254/gopher-social/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	addr := env.GetString("DB_ADDR", "postgres://postgres:password@localhost:5432/gopher_social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	store := store.NewStorage(conn)
	if err := db.Seed(store, conn); err != nil {
		log.Fatal("Error seeding database:", err)
	}
}
