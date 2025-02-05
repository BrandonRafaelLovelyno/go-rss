package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port  string
	DBUrl string
}

func Load() *Config {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("PG_URL")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if dbUrl == "" {
		log.Fatal("$PG_URL must be set")
	}

	return &Config{Port: port, DBUrl: dbUrl}
}
