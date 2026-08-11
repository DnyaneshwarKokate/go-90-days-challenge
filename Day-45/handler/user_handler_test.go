package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-45/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-45/repository"
)

func setupTestHandler() *APIHandler {
	cfg := &config.Config{
		Port:       "8080",
		AppEnv:     "test",
		AppVersion: "1.0.0-test",
		DBHost:     "localhost",
	}
	repo := repository.NewMemoryUserRepository()
	return NewAPIHandler(cfg, repo)
}

func TestHealthzEndpoint(t *testing.T) {
	h := setupTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()

	h.Healthz(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse json response: %v", err)
	}

	if resp["status"] != "healthy" {
		t.Errorf("expected status healthy, got %v", resp["status"])
	}
}

func TestReadyEndpoint(t *testing.T) {
	h := setupTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	w := httptest.NewRecorder()

	h.Ready(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestInfoEndpoint(t *testing.T) {
	h := setupTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/info", nil)
	w := httptest.NewRecorder()

	h.Info(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestGetUsers(t *testing.T) {
	h := setupTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	w := httptest.NewRecorder()

	h.HandleUsers(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestGetUserByID_Found(t *testing.T) {
	h := setupTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/1", nil)
	w := httptest.NewRecorder()

	h.HandleUsers(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestGetUserByID_NotFound(t *testing.T) {
	h := setupTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/999", nil)
	w := httptest.NewRecorder()

	h.HandleUsers(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestCreateUser_Success(t *testing.T) {
	h := setupTestHandler()
	newUser := repository.User{
		ID:    "3",
		Name:  "Alice Smith",
		Email: "alice@example.com",
		Role:  "Engineer",
	}
	body, _ := json.Marshal(newUser)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	h.HandleUsers(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestCreateUser_Duplicate(t *testing.T) {
	h := setupTestHandler()
	existingUser := repository.User{
		ID:    "1",
		Name:  "Duplicate User",
		Email: "dup@example.com",
		Role:  "Admin",
	}
	body, _ := json.Marshal(existingUser)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	h.HandleUsers(w, req)

	if w.Code != http.StatusConflict {
		t.Errorf("expected status %d, got %d", http.StatusConflict, w.Code)
	}
}
