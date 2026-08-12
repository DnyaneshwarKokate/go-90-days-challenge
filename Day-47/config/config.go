package config

import (
	"os"
)

// Config represents application configuration injected via Kubernetes ConfigMaps and Secrets.
type Config struct {
	Port         string
	AppEnv       string
	AppVersion   string
	LogLevel     string
	APIKey       string
	PodName      string
	PodNamespace string
	NodeName     string
}

// LoadConfig fetches environment variables with fallback defaults.
func LoadConfig() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		AppEnv:       getEnv("APP_ENV", "development"),
		AppVersion:   getEnv("APP_VERSION", "1.0.0"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		APIKey:       getEnv("API_KEY", "default-secret-key-123"),
		PodName:      getEnv("POD_NAME", "localhost-pod"),
		PodNamespace: getEnv("POD_NAMESPACE", "default"),
		NodeName:     getEnv("NODE_NAME", "minikube"),
	}
}

func getEnv(key, fallback string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return fallback
}
