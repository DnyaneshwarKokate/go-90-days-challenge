package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"day-34/domain"
	"day-34/handler"
	"day-34/logger"
	"day-34/middleware"
	"day-34/service"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("==========================================================================")
	fmt.Println("🚀 Day 34: Email Sending System in Go (SMTP, HTML Templates, Async Queue)")
	fmt.Println("==========================================================================")

	logFilePath := "email_app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()
	zapLog.Info(ctx, "Email Service starting", "smtp_host", "smtp.example.com", "mock_mode", true)

	emailSvc := service.NewEmailService("smtp.example.com", "587", "no-reply@example.com", "secret", true, zapLog)
	defer emailSvc.Close()

	emailHandler := handler.NewEmailHandler(emailSvc, zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	v1 := router.Group("/api/v1/email")
	{
		v1.POST("/send", emailHandler.SendSync)
		v1.POST("/queue", emailHandler.SendAsync)
	}

	fmt.Println("\n--- 1️⃣ Testing Synchronous HTML Email Sending ---")
	sendSyncEmail(router, "WELCOME", "Dnyaneshwar Kokate", "dnyaneshwar@example.com", "Welcome to Go 90 Days Challenge!")

	fmt.Println("\n--- 2️⃣ Testing Asynchronous Background Worker Queue Email ---")
	sendAsyncEmail(router, "ORDER_CONFIRMATION", "Rahul Sharma", "rahul@example.com", "Your Order #9090 has been confirmed")

	time.Sleep(500 * time.Millisecond) // Allow background worker to process queue

	fmt.Println("\n✅ Day 34 Email Sending System executed successfully! Check email_app.log for logs.")
}

func sendSyncEmail(router *gin.Engine, tpl, name, email, subject string) {
	payload := domain.SendEmailInput{
		To:       []string{email},
		Subject:  subject,
		Template: tpl,
		Name:     name,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/email/send", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /email/send -> Status %d | Body: %s\n", w.Code, w.Body.String())
}

func sendAsyncEmail(router *gin.Engine, tpl, name, email, subject string) {
	payload := domain.SendEmailInput{
		To:       []string{email},
		Subject:  subject,
		Template: tpl,
		Name:     name,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/email/queue", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /email/queue -> Status %d | Body: %s\n", w.Code, w.Body.String())
}
