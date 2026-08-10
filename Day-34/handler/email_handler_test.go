package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"day-34/domain"
	"day-34/handler"
	"day-34/logger"
	"day-34/middleware"
	"day-34/service"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() (*gin.Engine, *service.EmailService) {
	gin.SetMode(gin.TestMode)
	zapLog, _ := logger.NewZapLogger("test", "")

	emailSvc := service.NewEmailService("smtp.test.com", "587", "test@test.com", "pass", true, zapLog)
	emailHandler := handler.NewEmailHandler(emailSvc, zapLog)

	router := gin.New()
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	v1 := router.Group("/api/v1/email")
	{
		v1.POST("/send", emailHandler.SendSync)
		v1.POST("/queue", emailHandler.SendAsync)
	}

	return router, emailSvc
}

func TestEmailSending_SyncSuccess(t *testing.T) {
	router, emailSvc := setupTestRouter()
	defer emailSvc.Close()

	payload := domain.SendEmailInput{
		To:       []string{"alice@test.com"},
		Subject:  "Test Email Subject",
		Template: "WELCOME",
		Name:     "Alice",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/email/send", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d. Body: %s", w.Code, w.Body.String())
	}
}

func TestEmailSending_AsyncQueue(t *testing.T) {
	router, emailSvc := setupTestRouter()
	defer emailSvc.Close()

	payload := domain.SendEmailInput{
		To:       []string{"bob@test.com"},
		Subject:  "Async Email Subject",
		Template: "PASSWORD_RESET",
		Name:     "Bob",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/email/queue", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected 202 Accepted, got %d", w.Code)
	}

	time.Sleep(100 * time.Millisecond)
}
