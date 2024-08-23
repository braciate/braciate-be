package main

import (
	"github.com/braciate/braciate-be/internal/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	logger := config.NewLogger()
	fiber := config.NewFiber(logger)

	app, err := config.NewServer(fiber, logger)
	if err != nil {
		log.Fatal("error creating server")
	}

	app.RegisterHandler()

	if err := app.Run(); err != nil {
		log.Fatal("error running server")
	}
}
