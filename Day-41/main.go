package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"day-41/logger"
	"day-41/middleware"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func main() {
	fmt.Println("==========================================================================")
	fmt.Println("🚀 Day 41: Rate Limiting in Go (Token Bucket Algorithm, IP Limiter, 429)")
	fmt.Println("==========================================================================")

	logFilePath := "ratelimit_app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	// Rate limit: 2 requests per second with burst size of 3
	rateLimiter := middleware.NewIPRateLimiter(rate.Limit(2), 3, zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RateLimiterMiddleware(rateLimiter))

	router.GET("/api/v1/resource", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Resource accessed successfully",
			"time":    time.Now(),
		})
	})

	fmt.Println("\n--- 1️⃣ Bursting 6 Requests Concurrently from IP 192.168.1.50 (Limit = 3) ---")
	for i := 1; i <= 6; i++ {
		req, _ := http.NewRequest("GET", "/api/v1/resource", nil)
		req.RemoteAddr = "192.168.1.50:12345"

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		remHeader := w.Header().Get("X-RateLimit-Remaining")
		fmt.Printf("Req #%d -> Status %d | Remaining Tokens: %s | Body: %s\n", i, w.Code, remHeader, w.Body.String())
	}

	fmt.Println("\n--- 2️⃣ Waiting 1 Second for Token Refill ---")
	time.Sleep(1 * time.Second)

	fmt.Println("\n--- 3️⃣ Retrying Request after Refill ---")
	req, _ := http.NewRequest("GET", "/api/v1/resource", nil)
	req.RemoteAddr = "192.168.1.50:12345"
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("Req #7 -> Status %d | Remaining Tokens: %s | Body: %s\n", w.Code, w.Header().Get("X-RateLimit-Remaining"), w.Body.String())

	fmt.Println("\n✅ Day 41 Rate Limiting executed successfully! Check ratelimit_app.log for logs.")
}
