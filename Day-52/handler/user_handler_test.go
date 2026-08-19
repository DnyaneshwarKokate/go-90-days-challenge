package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"day52/domain"
	"day52/repository"
	"day52/service"
)

func setupTestServer() (*UserHandler, *service.UserService) {
	repo := repository.NewInMemoryUserRepository()
	svc := service.NewUserService(repo, "secret-key-12345")
	h := NewUserHandler(svc)
	return h, svc
}

func TestRegisterAndLoginFlow(t *testing.T) {
	h, _ := setupTestServer()

	// 1. Register User
	regPayload := `{"email":"alice@example.com","name":"Alice","password":"securepassword"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBufferString(regPayload))
	rr := httptest.NewRecorder()

	h.Register(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("Expected status 201 Created, got %d. Body: %s", rr.Code, rr.Body.String())
	}

	var authResp domain.AuthResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &authResp); err != nil {
		t.Fatalf("Failed to unmarshal register response: %v", err)
	}

	if authResp.Token == "" {
		t.Error("Expected non-empty token in register response")
	}

	// 2. Login User
	loginPayload := `{"email":"alice@example.com","password":"securepassword"}`
	reqLogin := httptest.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBufferString(loginPayload))
	rrLogin := httptest.NewRecorder()

	h.Login(rrLogin, reqLogin)

	if rrLogin.Code != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", rrLogin.Code)
	}

	var loginResp domain.AuthResponse
	_ = json.Unmarshal(rrLogin.Body.Bytes(), &loginResp)

	// 3. Access Profile with JWT
	reqProfile := httptest.NewRequest(http.MethodGet, "/api/v1/users/profile", nil)
	reqProfile.Header.Set("Authorization", "Bearer "+loginResp.Token)
	rrProfile := httptest.NewRecorder()

	h.Profile(rrProfile, reqProfile)

	if rrProfile.Code != http.StatusOK {
		t.Fatalf("Expected status 200 OK for Profile, got %d", rrProfile.Code)
	}

	var user domain.User
	_ = json.Unmarshal(rrProfile.Body.Bytes(), &user)

	if user.Email != "alice@example.com" {
		t.Errorf("Expected profile email alice@example.com, got %s", user.Email)
	}
}
