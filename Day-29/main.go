package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"day-29/domain"
	"day-29/handler"
	"day-29/logger"
	"day-29/middleware"
	"day-29/repository"
	"day-29/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("🚀 Day 29: Production-Grade Logging System in Go")
	fmt.Println("==================================================")

	logFilePath := "app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()
	zapLog.Info(ctx, "Logger initialized successfully", "log_file", logFilePath, "environment", "development")

	slogLog := logger.NewSlogLogger("development")
	slogLog.Info(ctx, "Go stdlib slog logger initialized as secondary logger")

	userRepo := repository.NewMemoryUserRepository(zapLog)
	userUC := usecase.NewUserUseCase(userRepo, zapLog)
	userHandler := handler.NewUserHandler(userUC, zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.RegisterUser)
		api.GET("/users", userHandler.ListUsers)
		api.GET("/users/:id", userHandler.GetUserByID)
	}

	fmt.Println("\n--- 🛠️ Demonstrating Structured HTTP Requests & Context Tracing ---")

	createUser(router, "Dnyaneshwar Kokate", "dnyaneshwar@example.com", "ADMIN", "req-trace-101")
	createUser(router, "Rahul Sharma", "rahul@example.com", "DEVELOPER", "req-trace-102")

	fmt.Println("\n--- ⚠️ Demonstrating Validation Warning & Contextual Error Logs ---")
	createUser(router, "Duplicate User", "dnyaneshwar@example.com", "USER", "req-trace-103")

	fmt.Println("\n--- 🔍 Demonstrating Get User & List Users Requests ---")
	listUsers(router, "req-trace-104")

	fmt.Println("\n✅ Demonstration completed successfully! Structured logs recorded in terminal & app.log")
}

func createUser(router *gin.Engine, name, email, role, requestID string) {
	payload := domain.CreateUserInput{
		Name:  name,
		Email: email,
		Role:  role,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if requestID != "" {
		req.Header.Set("X-Request-ID", requestID)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("Response Status: %d | Body: %s\n", w.Code, w.Body.String())
}

func listUsers(router *gin.Engine, requestID string) {
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	if requestID != "" {
		req.Header.Set("X-Request-ID", requestID)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("Response Status: %d | Body: %s\n", w.Code, w.Body.String())
}
