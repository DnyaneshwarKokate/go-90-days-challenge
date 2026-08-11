package config

import (
	"os"
)

// Config holds runtime configuration for the containerized application.
type Config struct {
	Port        string `json:"port"`
	AppEnv      string `json:"app_env"`
	ServiceName string `json:"service_name"`
	Version     string `json:"version"`
}

// LoadConfig reads configuration from environment variables with fallback defaults.
func LoadConfig() Config {
	return Config{
		Port:        getEnv("PORT", "8080"),
		AppEnv:      getEnv("APP_ENV", "development"),
		ServiceName: getEnv("SERVICE_NAME", "go-docker-basics-service"),
		Version:     getEnv("VERSION", "1.0.0"),
	}
}

func getEnv(key, fallback string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return fallback
}
