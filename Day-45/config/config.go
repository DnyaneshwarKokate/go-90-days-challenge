package config

import (
	"os"
)

// Config holds environment configurations for the Go REST API container
type Config struct {
	Port       string
	AppEnv     string
	AppVersion string
	DBHost     string
}

// LoadConfig retrieves environment variables or applies production-safe defaults
func LoadConfig() *Config {
	return &Config{
		Port:       getEnv("PORT", "8080"),
		AppEnv:     getEnv("APP_ENV", "development"),
		AppVersion: getEnv("APP_VERSION", "1.0.0"),
		DBHost:     getEnv("DB_HOST", "localhost"),
	}
}

func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return defaultValue
}
