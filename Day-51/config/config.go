package config

import (
	"os"
)

type Config struct {
	ServiceName string
	Port        string
	Environment string
}

func LoadConfig() *Config {
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		serviceName = "user-account-service"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	return &Config{
		ServiceName: serviceName,
		Port:        port,
		Environment: env,
	}
}
