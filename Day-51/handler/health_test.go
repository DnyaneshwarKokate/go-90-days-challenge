package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"day51/config"
	"day51/middleware"
)

func TestHealthCheckHandler(t *testing.T) {
	cfg := &config.Config{
		ServiceName: "test-microservice",
		Port:        "8081",
		Environment: "testing",
	}
	h := NewHealthHandler(cfg)

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	wrappedHandler := middleware.CorrelationIDMiddleware(http.HandlerFunc(h.HealthCheck))
	wrappedHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HealthCheck returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resp HealthResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if resp.ServiceName != "test-microservice" {
		t.Errorf("Expected ServiceName 'test-microservice', got %s", resp.ServiceName)
	}

	if resp.CorrelationID == "" {
		t.Error("Expected CorrelationID to be set in context and response, got empty string")
	}

	correlationHeader := rr.Header().Get(middleware.HeaderCorrelationID)
	if correlationHeader == "" {
		t.Error("Expected X-Correlation-ID header in response")
	}
}
