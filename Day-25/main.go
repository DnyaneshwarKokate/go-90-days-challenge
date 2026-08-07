package main

import (
	"fmt"
	"os"
)

// getEnvWithDefault returns the value of environment variable or fallback if empty
func getEnvWithDefault(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}

func main() {
	fmt.Println("--- Day 25: Environment Variables in Go (Standard Library os) ---")

	// 1. Setting Environment Variables programmatically
	_ = os.Setenv("APP_NAME", "Go90DaysChallenge")
	_ = os.Setenv("APP_PORT", "8080")
	_ = os.Setenv("DB_HOST", "localhost")

	// 2. Reading Environment Variables with os.Getenv()
	appName := os.Getenv("APP_NAME")
	appPort := os.Getenv("APP_PORT")

	fmt.Println("\n1. Reading set environment variables:")
	fmt.Printf("   APP_NAME: %s\n", appName)
	fmt.Printf("   APP_PORT: %s\n", appPort)

	// 3. Checking Existence with os.LookupEnv()
	fmt.Println("\n2. Checking existence with os.LookupEnv():")
	if dbHost, exists := os.LookupEnv("DB_HOST"); exists {
		fmt.Printf("   DB_HOST exists with value: '%s'\n", dbHost)
	} else {
		fmt.Println("   DB_HOST does not exist")
	}

	if secretKey, exists := os.LookupEnv("API_SECRET_KEY"); exists {
		fmt.Printf("   API_SECRET_KEY exists with value: '%s'\n", secretKey)
	} else {
		fmt.Println("   API_SECRET_KEY is missing/unset")
	}

	// 4. Using Fallback Default Values
	fmt.Println("\n3. Using fallback default values for missing env vars:")
	dbPort := getEnvWithDefault("DB_PORT", "5432")
	appEnv := getEnvWithDefault("APP_ENV", "development")

	fmt.Printf("   DB_PORT: %s (Fallback active)\n", dbPort)
	fmt.Printf("   APP_ENV: %s (Fallback active)\n", appEnv)

	// 5. Cleanup / Unsetting environment variables
	_ = os.Unsetenv("APP_NAME")
	_ = os.Unsetenv("APP_PORT")
	_ = os.Unsetenv("DB_HOST")
}
