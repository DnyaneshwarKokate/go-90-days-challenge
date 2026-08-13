package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-48/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-48/domain"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-48/handler"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-48/repository"
)

func setupTestHandler(t *testing.T) (*handler.APIHandler, string) {
	t.Helper()
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test_orders.json")

	cfg := &config.Config{
		Port:         "8080",
		AppEnv:       "testing",
		DataPath:     filePath,
		APIKey:       "test-key",
		PodName:      "test-pod-0",
		PodNamespace: "test-namespace",
		NodeName:     "test-node",
		PodIP:        "10.244.0.5",
	}

	store, err := repository.NewStore(filePath)
	if err != nil {
		t.Fatalf("failed to initialize test store: %v", err)
	}

	return handler.NewAPIHandler(cfg, store), filePath
}

func TestHealthz(t *testing.T) {
	h, _ := setupTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	h.Healthz(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}

	if resp["status"] != "healthy" {
		t.Errorf("expected status healthy, got %v", resp["status"])
	}
}

func TestReadyAndSetNotReady(t *testing.T) {
	h, _ := setupTestHandler(t)

	// Initial state: Ready (200 OK)
	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	rec := httptest.NewRecorder()

	h.Ready(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	// Dynamic transition: SetNotReady() -> 503 Service Unavailable
	h.SetNotReady()
	rec = httptest.NewRecorder()
	h.Ready(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Errorf("expected status 503, got %d", rec.Code)
	}
}

func TestInfo(t *testing.T) {
	h, _ := setupTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/info", nil)
	rec := httptest.NewRecorder()

	h.Info(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	var resp map[string]interface{}
	_ = json.Unmarshal(rec.Body.Bytes(), &resp)

	if resp["pod_name"] != "test-pod-0" {
		t.Errorf("expected pod_name test-pod-0, got %v", resp["pod_name"])
	}
	if resp["pod_namespace"] != "test-namespace" {
		t.Errorf("expected pod_namespace test-namespace, got %v", resp["pod_namespace"])
	}
}

func TestOrdersGETandPOST(t *testing.T) {
	h, filePath := setupTestHandler(t)

	// 1. Initial GET - empty list
	req := httptest.NewRequest(http.MethodGet, "/api/v1/orders", nil)
	rec := httptest.NewRecorder()

	h.Orders(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}

	// 2. Invalid POST payload
	req = httptest.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewBufferString(`invalid-json`))
	rec = httptest.NewRecorder()
	h.Orders(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for invalid json, got %d", rec.Code)
	}

	// 3. Unprocessable Entity (validation failure)
	payload := domain.CreateOrderRequest{Customer: "", Item: "Laptop", Quantity: 1, Price: 1200.00}
	body, _ := json.Marshal(payload)
	req = httptest.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewBuffer(body))
	rec = httptest.NewRecorder()
	h.Orders(rec, req)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422 for missing customer, got %d", rec.Code)
	}

	// 4. Valid POST creation
	payload = domain.CreateOrderRequest{Customer: "Alice", Item: "MacBook Pro", Quantity: 1, Price: 2400.00}
	body, _ = json.Marshal(payload)
	req = httptest.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewBuffer(body))
	rec = httptest.NewRecorder()
	h.Orders(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", rec.Code)
	}

	var created domain.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &created)
	if created.Customer != "Alice" || created.ID == "" {
		t.Errorf("unexpected created order object: %+v", created)
	}

	// 5. Verify disk persistence
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("expected disk file [%s] to exist after save", filePath)
	}

	// 6. Method Not Allowed
	req = httptest.NewRequest(http.MethodDelete, "/api/v1/orders", nil)
	rec = httptest.NewRecorder()
	h.Orders(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405 Method Not Allowed, got %d", rec.Code)
	}
}

func TestMetrics(t *testing.T) {
	h, _ := setupTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/metrics", nil)
	rec := httptest.NewRecorder()

	h.Metrics(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}
