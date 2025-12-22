package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	BackPort             string
	BackPostgresPassword string
	BackPostgresUser     string
	BackPostgresDB       string
	BackPostgresPort     string
	BackPostgresHost     string
	SecretToken          string
}

func Load() *Config {
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("Successfully loaded .env file")
	}

	backPort := ":" + getEnv("BACK_PORT", "")
	backPostgresPassword := getEnv("BACK_POSTGRES_PASSWORD", "")
	backPostgresUser := getEnv("BACK_POSTGRES_USER", "")
	backPostgresDB := getEnv("BACK_POSTGRES_DB", "")
	backPostgresPort := getEnv("BACK_POSTGRES_PORT", "")
	secretToken := getEnv("BACK_SECRET_TOKEN", "")
	backPostgresHost := getEnv("BACK_POSTGRES_HOST", "")

	return &Config{
		BackPort:             backPort,
		BackPostgresPassword: backPostgresPassword,
		BackPostgresUser:     backPostgresUser,
		BackPostgresDB:       backPostgresDB,
		BackPostgresPort:     backPostgresPort,
		BackPostgresHost:     backPostgresHost,
		SecretToken:          secretToken,
	}
}

func getEnv(key, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
