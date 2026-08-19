package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"day55/circuitbreaker"
)

func TestResilientHTTPClientAndCircuitBreaker(t *testing.T) {
	failureCount := 0
	// Backend server that fails twice then succeeds
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		failureCount++
		if failureCount <= 2 {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer ts.Close()

	// Client configured with 2 retries (so 3rd attempt succeeds)
	resilientClient := NewResilientHTTPClient(2*time.Second, 2, 5, 100*time.Millisecond)

	ctx := context.Background()
	body, err := resilientClient.GetWithRetry(ctx, ts.URL)
	if err != nil {
		t.Fatalf("Expected successful GET after retries, got error: %v", err)
	}

	if string(body) != `{"status":"ok"}` {
		t.Errorf("Expected body `{\"status\":\"ok\"}`, got %s", string(body))
	}
}

func TestCircuitBreakerTripToOpen(t *testing.T) {
	// Failing backend
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	// Circuit breaker trips after 2 failures
	resilientClient := NewResilientHTTPClient(500*time.Millisecond, 0, 2, 200*time.Millisecond)
	ctx := context.Background()

	// Call 1 -> fails
	_, _ = resilientClient.GetWithRetry(ctx, ts.URL)
	// Call 2 -> fails, trips circuit breaker to OPEN
	_, _ = resilientClient.GetWithRetry(ctx, ts.URL)

	if resilientClient.CircuitState() != circuitbreaker.StateOpen {
		t.Errorf("Expected Circuit Breaker state to be OPEN, got %s", resilientClient.CircuitState())
	}

	// Call 3 -> Rejected immediately by Circuit Breaker
	_, err := resilientClient.GetWithRetry(ctx, ts.URL)
	if err != circuitbreaker.ErrCircuitOpen {
		t.Errorf("Expected ErrCircuitOpen, got %v", err)
	}
}
