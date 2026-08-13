package config

import (
	"os"
)

// Config holds the application runtime configuration populated from env vars, ConfigMaps, Secrets, and Downward API.
type Config struct {
	Port         string
	AppEnv       string
	DataPath     string
	APIKey       string
	PodName      string
	PodNamespace string
	NodeName     string
	PodIP        string
}

// LoadConfig fetches values from environment variables with fallback defaults.
func LoadConfig() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		AppEnv:       getEnv("APP_ENV", "production"),
		DataPath:     getEnv("DATA_PATH", "/data/orders.json"),
		APIKey:       getEnv("API_KEY", "default-secret-key-12345"),
		PodName:      getEnv("POD_NAME", "go-app-pod-local"),
		PodNamespace: getEnv("POD_NAMESPACE", "default"),
		NodeName:     getEnv("NODE_NAME", "node-local"),
		PodIP:        getEnv("POD_IP", "127.0.0.1"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}
