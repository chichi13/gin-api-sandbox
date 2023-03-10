package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Env string

const (
	Development Env = "development"
	// Production  Env = "production"
	// Testing     Env = "testing"
)

type Config struct {
	GinMode string `env:"gin_mode"`
	Env     Env    `env:"env" envDefault:"production"`
}

var Settings = loadEnv()

func loadEnv() Config {
	// Load .env file if it exists, otherwise ignore
	if _, err := os.Stat(".env"); err != nil {
		log.Println("Warning: .env file does not exist, skipping loading it")
		return Config{}
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Unable to parse ennvironment variables: %e", err)
	}

	return cfg
}
