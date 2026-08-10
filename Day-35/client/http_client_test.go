package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"day-35/client"
	"day-35/domain"
	"day-35/logger"
)

func TestResilientHTTPClient_RetryAndSuccess(t *testing.T) {
	zapLog, _ := logger.NewZapLogger("test", "")

	var attempts int32 = 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		current := atomic.AddInt32(&attempts, 1)
		if current == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":1,"name":"Test User","email":"test@test.com"}`))
	}))
	defer server.Close()

	cfg := domain.ClientConfig{
		BaseURL:        server.URL,
		Timeout:        1 * time.Second,
		MaxRetries:     3,
		InitialBackoff: 50 * time.Millisecond,
	}

	restClient := client.NewResilientHTTPClient(cfg, zapLog)
	user, err := restClient.GetUser(context.Background(), 1)

	if err != nil {
		t.Fatalf("expected successful response after retry, got error: %v", err)
	}

	if user.Name != "Test User" {
		t.Errorf("expected user name 'Test User', got '%s'", user.Name)
	}

	if atomic.LoadInt32(&attempts) != 2 {
		t.Errorf("expected exactly 2 attempts, got %d", attempts)
	}
}
