package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DSN  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		if os.IsNotExist(err) {
			log.Println("No .env file found. Reading configuration from environment variables only.")
		} else {
			// Log other errors loading .env, but don't fatally exit.
			// This allows the app to potentially run if crucial vars are set elsewhere.
			log.Printf("Warning: Error loading .env file: %v. Continuing without it.\n", err)
		}
	}

	cfg := &Config{}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	cfg.Port = port
	cfg.DSN = os.Getenv("DSN")
	return cfg, err
}
