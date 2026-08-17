package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("StatusHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resp StatusResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if resp.Service != "github-actions-demo" {
		t.Errorf("Expected service 'github-actions-demo', got %s", resp.Service)
	}
}
