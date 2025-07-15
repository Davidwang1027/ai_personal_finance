package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	PlaidClientID     string
	PlaidSecret       string
	PlaidEnv          string
	DatabaseURL       string
	Port              string
	JWTSecret         string
}

// Load reads configuration from .env file
func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	config := &Config{
		PlaidClientID:     getEnv("PLAID_CLIENT_ID", ""),
		PlaidSecret:       getEnv("PLAID_SECRET", ""),
		PlaidEnv:          getEnv("PLAID_ENV", "sandbox"),
		DatabaseURL:       getEnv("DATABASE_URL", ""),
		Port:              getEnv("PORT", "8080"),
		JWTSecret:         getEnv("JWT_SECRET", "your-secret-key"),
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
