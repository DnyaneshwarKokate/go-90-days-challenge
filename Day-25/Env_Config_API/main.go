package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Config struct holds application configurations
type Config struct {
	Port        string
	AppName     string
	Environment string
	DBHost      string
	DBPort      string
	DBUser      string
	DBName      string
	JWTSecret   string
}

// Global config instance
var AppConfig Config

// getEnv retrieves env value or returns fallback default
func getEnv(key, defaultValue string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return defaultValue
}

// LoadConfig initializes application configuration from .env and environment
func LoadConfig() {
	// Load .env file if available
	err := godotenv.Load("Env_Config_API/.env")
	if err != nil {
		// Try loading from current directory if path differs
		err = godotenv.Load(".env")
		if err != nil {
			log.Println("⚠️  Warning: .env file not found, using system environment variables and defaults")
		}
	} else {
		log.Println("✅ Successfully loaded .env file")
	}

	AppConfig = Config{
		Port:        getEnv("PORT", "8085"),
		AppName:     getEnv("APP_NAME", "Default_Go_API"),
		Environment: getEnv("ENVIRONMENT", "development"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBName:      getEnv("DB_NAME", "default_db"),
		JWTSecret:   getEnv("JWT_SECRET", "default_fallback_secret"),
	}

	fmt.Println("-------------------------------------------")
	fmt.Printf("🚀 Configuration Loaded:\n")
	fmt.Printf("   App Name:    %s\n", AppConfig.AppName)
	fmt.Printf("   Environment: %s\n", AppConfig.Environment)
	fmt.Printf("   Port:        %s\n", AppConfig.Port)
	fmt.Printf("   DB Host:     %s:%s\n", AppConfig.DBHost, AppConfig.DBPort)
	fmt.Println("-------------------------------------------")
}

func main() {
	LoadConfig()

	router := gin.Default()

	// 1. Health Check Endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      "UP",
			"app_name":    AppConfig.AppName,
			"environment": AppConfig.Environment,
		})
	})

	// 2. Application Info Endpoint (Sanitized Config)
	router.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app_name":    AppConfig.AppName,
			"environment": AppConfig.Environment,
			"port":        AppConfig.Port,
			"database": gin.H{
				"host": AppConfig.DBHost,
				"port": AppConfig.DBPort,
				"user": AppConfig.DBUser,
				"name": AppConfig.DBName,
			},
		})
	})

	serverAddr := ":" + AppConfig.Port
	fmt.Printf("Server starting on http://localhost%s\n", serverAddr)
	log.Fatal(router.Run(serverAddr))
}
