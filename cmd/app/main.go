package main

import (
	"flag"
	"github.com/braciate/braciate-be/database/seeder"
	"github.com/braciate/braciate-be/internal/config"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	handleFlagArgs()

	logger := config.NewLogger()
	fiber := config.NewFiber(logger)

	app, err := config.NewServer(fiber, logger)
	if err != nil {
		log.Fatalf("error creating server: %v", err)
	}

	app.RegisterHandler()

	if err := app.Run(); err != nil {
		log.Fatal("error running server")
	}
}

func handleFlagArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		switch args[0] {
		case "seed":
			seeder.Seed()
			os.Exit(0)
		}
	}
}
