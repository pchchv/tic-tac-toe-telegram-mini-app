package config

import (
	"log"
	"os"

	"github.com/pchchv/env"
)

// Config holds all configuration variables for the application
type Config struct {
	ListenAddress string
	BotToken      string
	WebAppURL     string
}

// Load initializes environment variables and returns a populated Config struct
func Load() *Config {
	if err := env.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		ListenAddress: getEnvValue("LISTEN_ADDRESS", "0.0.0.0:3000"),
		BotToken:      getEnvValue("TG_BOT_TOKEN", ""),
		WebAppURL:     getEnvValue("WEBAPP_URL", ""),
	}
}

// getEnvValue retrieves an environment variable or returns a fallback value if not set
func getEnvValue(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}
