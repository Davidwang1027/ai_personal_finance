package config

import (
	"log"
	"os"

	"github.com/davidwang/go-finance-api/go-finance-api/db"
	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	PlaidClientID string
	PlaidSecret   string
	PlaidEnv      string
	Port          string
	JWTSecret     string
	DB            db.DBConfig
}

// Load reads configuration from .env file
func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	config := &Config{
		PlaidClientID: getEnv("PLAID_CLIENT_ID", ""),
		PlaidSecret:   getEnv("PLAID_SECRET", ""),
		PlaidEnv:      getEnv("PLAID_ENV", "sandbox"),
		Port:          getEnv("PORT", "8080"),
		JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
		DB: db.DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "finance"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}

	return config
}

// getEnv reads an environment variable with a default fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
