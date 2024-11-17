package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

type DatabaseConfig struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
	ELASTIC_URL string
	ELASTIC_USER string
	ELASTIC_PASS string
}
func LoadDBConfig() *DatabaseConfig {
	var res = new(DatabaseConfig)

	if val, found := os.LookupEnv("DB_USER"); found {
		res.DB_USER = val
	}

	if val, found := os.LookupEnv("DB_PASS"); found {
		res.DB_PASS = val
	}

	if val, found := os.LookupEnv("DB_HOST"); found {
		res.DB_HOST = val
	}

	if val, found := os.LookupEnv("DB_PORT"); found {
		res.DB_PORT = val
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		res.DB_NAME = val
	}

	if val, found := os.LookupEnv("ELASTIC_URL"); found {
		res.ELASTIC_URL = val
	}

	if val, found := os.LookupEnv("ELASTIC_USER"); found {
		res.ELASTIC_USER = val
	}

	if val, found := os.LookupEnv("ELASTIC_PASS"); found {
		res.ELASTIC_PASS = val
	}

	return res
}