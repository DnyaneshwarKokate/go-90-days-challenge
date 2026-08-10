package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"time"

	"day-35/client"
	"day-35/domain"
	"day-35/handler"
	"day-35/logger"
	"day-35/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("==========================================================================")
	fmt.Println("🚀 Day 35: REST Client & External APIs (Pooling, Timeouts, Exponential Backoff)")
	fmt.Println("==========================================================================")

	logFilePath := "external_app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()

	// 1. Setup Mock Remote External HTTP Server
	var failAttempts int32 = 2
	mockRemoteServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		current := atomic.AddInt32(&failAttempts, -1)
		if current >= 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"temporary service unavailable"}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":101,"name":"Dnyaneshwar Kokate","username":"dnyanesh0902","email":"dnyanesh@example.com","company":"Google","source":"Mock Remote Server"}`))
	}))
	defer mockRemoteServer.Close()

	zapLog.Info(ctx, "Mock External HTTP Server started", "url", mockRemoteServer.URL)

	// 2. Setup Resilient REST Client
	clientCfg := domain.ClientConfig{
		BaseURL:        mockRemoteServer.URL,
		Timeout:        2 * time.Second,
		MaxRetries:     3,
		InitialBackoff: 100 * time.Millisecond,
	}

	restClient := client.NewResilientHTTPClient(clientCfg, zapLog)
	extHandler := handler.NewExternalHandler(restClient, zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	v1 := router.Group("/api/v1/external")
	{
		v1.GET("/users/:id", extHandler.GetUser)
	}

	fmt.Println("\n--- 1️⃣ Testing Resilient Client with Retry Logic (Initial 2 Failures -> Success on Attempt 3) ---")
	fetchExternalUser(router, "/api/v1/external/users/101")

	fmt.Println("\n✅ Day 35 REST Client & External APIs executed successfully! Check external_app.log for logs.")
}

func fetchExternalUser(router *gin.Engine, url string) {
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("GET %s -> Status %d | Body: %s\n", url, w.Code, w.Body.String())
}
