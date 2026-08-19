package proxy

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"day54/middleware"
)

func generateTestToken(secret string) string {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	payload := base64.RawURLEncoding.EncodeToString([]byte(`{"sub":"usr_123","email":"test@example.com"}`))
	unsigned := header + "." + payload

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsigned))
	sig := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	return unsigned + "." + sig
}

func TestAPIGatewayRoutingAndAuth(t *testing.T) {
	secret := "gateway-secret-key"

	// Mock User Microservice Backend
	userBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"service":"user_backend_ok"}`))
	}))
	defer userBackend.Close()

	// Mock Product Microservice Backend
	productBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"service":"product_backend_ok"}`))
	}))
	defer productBackend.Close()

	gp, err := NewGatewayProxy(userBackend.URL, productBackend.URL)
	if err != nil {
		t.Fatalf("Failed to initialize Gateway Proxy: %v", err)
	}

	authMw := middleware.NewAuthMiddleware(secret)
	handlerStack := authMw.ValidateJWT(gp)

	// Test 1: Public route bypasses auth
	reqPub := httptest.NewRequest("POST", "/api/v1/users/login", nil)
	rrPub := httptest.NewRecorder()
	handlerStack.ServeHTTP(rrPub, reqPub)

	if rrPub.Code != http.StatusOK {
		t.Errorf("Expected public route login to return 200, got %d", rrPub.Code)
	}

	// Test 2: Protected route without header fails
	reqProtNoHeader := httptest.NewRequest("GET", "/api/v1/products", nil)
	rrProtNoHeader := httptest.NewRecorder()
	handlerStack.ServeHTTP(rrProtNoHeader, reqProtNoHeader)

	if rrProtNoHeader.Code != http.StatusUnauthorized {
		t.Errorf("Expected unauthorized 401, got %d", rrProtNoHeader.Code)
	}

	// Test 3: Protected route with valid token succeeds and routes to Product Backend
	validToken := generateTestToken(secret)
	reqProtValid := httptest.NewRequest("GET", "/api/v1/products", nil)
	reqProtValid.Header.Set("Authorization", fmt.Sprintf("Bearer %s", validToken))
	rrProtValid := httptest.NewRecorder()

	handlerStack.ServeHTTP(rrProtValid, reqProtValid)

	if rrProtValid.Code != http.StatusOK {
		t.Errorf("Expected 200 OK for valid JWT routed request, got %d", rrProtValid.Code)
	}

	if rrProtValid.Header().Get("X-Gateway-ProxiedBy") == "" {
		t.Error("Expected X-Gateway-ProxiedBy header in proxied response")
	}
}
