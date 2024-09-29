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
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using environment variables.")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" {
		log.Fatal("Required environment variables are missing!")
	}

	log.Printf("Connecting to DB at %s:%s as user %s", dbHost, dbPort, dbUser)

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
