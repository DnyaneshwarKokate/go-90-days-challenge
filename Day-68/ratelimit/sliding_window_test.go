package ratelimit_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"day68/ratelimit"
)

func TestSlidingWindowLimiterEnforcement(t *testing.T) {
	// Limit: 3 requests per 100 milliseconds
	limiter := ratelimit.NewSlidingWindowLimiter(3, 100*time.Millisecond)

	clientID := "192.168.1.50"

	// First 3 requests must be allowed
	for i := 1; i <= 3; i++ {
		if !limiter.Allow(clientID) {
			t.Fatalf("Request %d should be allowed", i)
		}
	}

	// 4th request must be denied
	if limiter.Allow(clientID) {
		t.Fatalf("4th request within window should be rejected")
	}

	// Wait for sliding window to expire
	time.Sleep(120 * time.Millisecond)

	// Next request after window reset must be allowed
	if !limiter.Allow(clientID) {
		t.Fatalf("Request after window expiration should be allowed")
	}
}

func TestRateLimiterHTTPMiddleware(t *testing.T) {
	limiter := ratelimit.NewSlidingWindowLimiter(1, 500*time.Millisecond)

	handler := limiter.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Request 1: Allowed
	req1 := httptest.NewRequest("GET", "/api/data", nil)
	req1.Header.Set("X-Forwarded-For", "203.0.113.195")
	rec1 := httptest.NewRecorder()
	handler.ServeHTTP(rec1, req1)

	if rec1.Code != http.StatusOK {
		t.Fatalf("Expected HTTP 200, got %d", rec1.Code)
	}

	// Request 2: Rejected (HTTP 429)
	req2 := httptest.NewRequest("GET", "/api/data", nil)
	req2.Header.Set("X-Forwarded-For", "203.0.113.195")
	rec2 := httptest.NewRecorder()
	handler.ServeHTTP(rec2, req2)

	if rec2.Code != http.StatusTooManyRequests {
		t.Fatalf("Expected HTTP 429 Too Many Requests, got %d", rec2.Code)
	}
}
