package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"day-41/logger"
	"day-41/middleware"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func TestRateLimiterMiddleware_BurstAndThrottle(t *testing.T) {
	gin.SetMode(gin.TestMode)
	zapLog, _ := logger.NewZapLogger("test", "")

	limiter := middleware.NewIPRateLimiter(rate.Limit(1), 2, zapLog)

	router := gin.New()
	router.Use(middleware.RateLimiterMiddleware(limiter))
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Request 1: Burst 1 -> 200 OK
	req1, _ := http.NewRequest("GET", "/test", nil)
	req1.RemoteAddr = "10.0.0.1:5555"
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)
	if w1.Code != http.StatusOK {
		t.Fatalf("expected 200 OK for req 1, got %d", w1.Code)
	}

	// Request 2: Burst 2 -> 200 OK
	req2, _ := http.NewRequest("GET", "/test", nil)
	req2.RemoteAddr = "10.0.0.1:5555"
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)
	if w2.Code != http.StatusOK {
		t.Fatalf("expected 200 OK for req 2, got %d", w2.Code)
	}

	// Request 3: Burst exceeded -> 429 Too Many Requests
	req3, _ := http.NewRequest("GET", "/test", nil)
	req3.RemoteAddr = "10.0.0.1:5555"
	w3 := httptest.NewRecorder()
	router.ServeHTTP(w3, req3)
	if w3.Code != http.StatusTooManyRequests {
		t.Errorf("expected 429 Too Many Requests for req 3, got %d", w3.Code)
	}
}
