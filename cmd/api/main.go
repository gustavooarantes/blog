package main

import (
	"log"

	"github.com/gustavooarantes/blog/internal/env"
	"github.com/gustavooarantes/blog/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
	}

	storage := storage.NewStorage(nil)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
