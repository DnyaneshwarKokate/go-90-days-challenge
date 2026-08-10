package config

import (
	"os"
	"time"
)

type Config struct {
	Port         string
	Environment  string
	JWTSecret    string
	JWTExpiresIn time.Duration
	LogFilePath  string
}

func LoadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "super-secret-student-mgmt-key-90-days-challenge"
	}

	logPath := os.Getenv("LOG_FILE_PATH")
	if logPath == "" {
		logPath = "student_app.log"
	}

	return Config{
		Port:         port,
		Environment:  env,
		JWTSecret:    secret,
		JWTExpiresIn: 24 * time.Hour,
		LogFilePath:  logPath,
	}
}
