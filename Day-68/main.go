package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"day68/ratelimit"
)

func main() {
	fmt.Println("=== Day 68: Distributed Rate Limiting (Sliding Window) ===")

	// Limit: Max 2 requests per 500 milliseconds window
	limiter := ratelimit.NewSlidingWindowLimiter(2, 500*time.Millisecond)

	apiHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"SUCCESS","data":"Protected Payload"}`))
	})

	protectedApp := limiter.Middleware(apiHandler)
	clientIP := "10.0.4.15"

	fmt.Println("\n--- Simulating Burst Traffic from Client IP (10.0.4.15) ---")
	for i := 1; i <= 4; i++ {
		req := httptest.NewRequest("GET", "/api/v1/resource", nil)
		req.Header.Set("X-Forwarded-For", clientIP)
		rec := httptest.NewRecorder()

		protectedApp.ServeHTTP(rec, req)

		fmt.Printf("Request #%d: Status Code = %d | Body = %s\n", i, rec.Code, rec.Body.String())
	}

	fmt.Println("\n--- Waiting for Sliding Window Expiration (600ms) ---")
	time.Sleep(600 * time.Millisecond)

	reqReset := httptest.NewRequest("GET", "/api/v1/resource", nil)
	reqReset.Header.Set("X-Forwarded-For", clientIP)
	recReset := httptest.NewRecorder()
	protectedApp.ServeHTTP(recReset, reqReset)

	fmt.Printf("Request After Expiration: Status Code = %d | Body = %s\n", recReset.Code, recReset.Body.String())
}
