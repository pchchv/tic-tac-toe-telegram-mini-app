package config

import "os"

// getEnvValue retrieves an environment variable or returns a fallback value if not set
func getEnvValue(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}
