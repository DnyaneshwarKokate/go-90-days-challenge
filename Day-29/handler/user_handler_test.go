package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"day-29/domain"
	"day-29/handler"
	"day-29/logger"
	"day-29/middleware"
	"day-29/repository"
	"day-29/usecase"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() (*gin.Engine, domain.Logger) {
	gin.SetMode(gin.TestMode)
	zapLog, _ := logger.NewZapLogger("development", "")

	userRepo := repository.NewMemoryUserRepository(zapLog)
	userUC := usecase.NewUserUseCase(userRepo, zapLog)
	userHandler := handler.NewUserHandler(userUC, zapLog)

	router := gin.New()
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.RegisterUser)
		api.GET("/users", userHandler.ListUsers)
		api.GET("/users/:id", userHandler.GetUserByID)
	}

	return router, zapLog
}

func TestRegisterUser_Success(t *testing.T) {
	router, _ := setupTestRouter()

	input := domain.CreateUserInput{
		Name:  "Test User",
		Email: "test@example.com",
		Role:  "USER",
	}
	body, _ := json.Marshal(input)

	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	customReqID := "test-request-id-123"
	req.Header.Set("X-Request-ID", customReqID)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d. Body: %s", http.StatusCreated, w.Code, w.Body.String())
	}

	resHeaderID := w.Header().Get("X-Request-ID")
	if resHeaderID != customReqID {
		t.Errorf("expected X-Request-ID header %s, got %s", customReqID, resHeaderID)
	}
}

func TestRegisterUser_DuplicateEmail(t *testing.T) {
	router, _ := setupTestRouter()

	input := domain.CreateUserInput{
		Name:  "User One",
		Email: "duplicate@example.com",
		Role:  "USER",
	}
	body, _ := json.Marshal(input)

	req1, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	if w1.Code != http.StatusCreated {
		t.Fatalf("first registration failed, got status %d", w1.Code)
	}

	req2, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	if w2.Code != http.StatusConflict {
		t.Errorf("expected status %d for duplicate email, got %d", http.StatusConflict, w2.Code)
	}
}

func TestGetUserByID_NotFound(t *testing.T) {
	router, _ := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/v1/users/non-existent-id", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}
