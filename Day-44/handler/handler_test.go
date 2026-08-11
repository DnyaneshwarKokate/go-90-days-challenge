package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-44/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-44/handler"
)

func TestHandleHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	handler.HandleHealth(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", status)
	}

	var resp handler.HealthResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal health response: %v", err)
	}

	if resp.Status != "UP" {
		t.Errorf("expected status 'UP', got '%s'", resp.Status)
	}
}

func TestHandleInfo(t *testing.T) {
	cfg := config.Config{
		Port:        "8080",
		AppEnv:      "testing",
		ServiceName: "docker-test-svc",
		Version:     "1.0.0",
	}
	svc := handler.NewItemService(cfg)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/info", nil)
	rr := httptest.NewRecorder()

	svc.HandleInfo(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", status)
	}

	var resp handler.InfoResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal info response: %v", err)
	}

	if resp.ServiceName != "docker-test-svc" {
		t.Errorf("expected service name 'docker-test-svc', got '%s'", resp.ServiceName)
	}
	if resp.Environment != "testing" {
		t.Errorf("expected environment 'testing', got '%s'", resp.Environment)
	}
}

func TestHandleItems_GetAndPost(t *testing.T) {
	cfg := config.Config{AppEnv: "testing"}
	svc := handler.NewItemService(cfg)

	// 1. GET items
	reqGet := httptest.NewRequest(http.MethodGet, "/api/v1/items", nil)
	rrGet := httptest.NewRecorder()
	svc.HandleItems(rrGet, reqGet)

	if rrGet.Code != http.StatusOK {
		t.Errorf("expected GET items to return 200, got %d", rrGet.Code)
	}

	var items []handler.Item
	if err := json.Unmarshal(rrGet.Body.Bytes(), &items); err != nil {
		t.Fatalf("failed to unmarshal items: %v", err)
	}
	if len(items) < 2 {
		t.Errorf("expected at least 2 default items, got %d", len(items))
	}

	// 2. POST new item
	newItem := handler.Item{
		ID:    "3",
		Name:  "Kubernetes Pod",
		Price: 99.99,
	}
	bodyBytes, _ := json.Marshal(newItem)
	reqPost := httptest.NewRequest(http.MethodPost, "/api/v1/items", bytes.NewBuffer(bodyBytes))
	rrPost := httptest.NewRecorder()

	svc.HandleItems(rrPost, reqPost)

	if rrPost.Code != http.StatusCreated {
		t.Errorf("expected POST item to return 201 Created, got %d", rrPost.Code)
	}

	var createdItem handler.Item
	if err := json.Unmarshal(rrPost.Body.Bytes(), &createdItem); err != nil {
		t.Fatalf("failed to unmarshal created item: %v", err)
	}
	if createdItem.Name != "Kubernetes Pod" {
		t.Errorf("expected name 'Kubernetes Pod', got '%s'", createdItem.Name)
	}
}

func TestHandleItems_InvalidInputs(t *testing.T) {
	cfg := config.Config{AppEnv: "testing"}
	svc := handler.NewItemService(cfg)

	// Invalid JSON
	reqBadJSON := httptest.NewRequest(http.MethodPost, "/api/v1/items", bytes.NewBufferString("{invalid-json"))
	rrBadJSON := httptest.NewRecorder()
	svc.HandleItems(rrBadJSON, reqBadJSON)

	if rrBadJSON.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 for bad JSON, got %d", rrBadJSON.Code)
	}

	// Missing Name
	reqMissingName := httptest.NewRequest(http.MethodPost, "/api/v1/items", bytes.NewBufferString(`{"id":"4"}`))
	rrMissingName := httptest.NewRecorder()
	svc.HandleItems(rrMissingName, reqMissingName)

	if rrMissingName.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 for missing fields, got %d", rrMissingName.Code)
	}
}
