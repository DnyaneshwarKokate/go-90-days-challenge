package main

import (
	"context"
	"fmt"
	"net/http/httptest"
	"os"
	"strings"
	"time"

	"day-39/domain"
	"day-39/handler"
	"day-39/hub"
	"day-39/logger"
	"day-39/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("==========================================================================")
	fmt.Println("🚀 Day 39: Real-Time Chat Backend in Go (Multi-Room Hub, Channels, Broadcasting)")
	fmt.Println("==========================================================================")

	logFilePath := "chat_app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()
	zapLog.Info(ctx, "Real-Time Chat Hub initializing")

	chatHub := hub.NewChatHub(zapLog)
	chatHandler := handler.NewChatHandler(chatHub, zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	router.GET("/ws/chat", chatHandler.JoinRoom)

	server := httptest.NewServer(router)
	defer server.Close()

	baseURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws/chat"

	fmt.Println("\n--- 1️⃣ Alice Joining Room 'go-devs' ---")
	aliceWS, _, _ := websocket.DefaultDialer.Dial(baseURL+"?room=go-devs&user=Alice", nil)
	defer aliceWS.Close()

	var event1 domain.ChatEvent
	aliceWS.ReadJSON(&event1)
	fmt.Printf("Alice Received: Type=%s | Content=\"%s\"\n", event1.Type, event1.Content)

	fmt.Println("\n--- 2️⃣ Bob Joining Room 'go-devs' ---")
	bobWS, _, _ := websocket.DefaultDialer.Dial(baseURL+"?room=go-devs&user=Bob", nil)
	defer bobWS.Close()

	var event2 domain.ChatEvent
	aliceWS.ReadJSON(&event2) // System notification to Alice that Bob joined
	fmt.Printf("Alice Received Notification: Sender=%s | Content=\"%s\"\n", event2.Sender, event2.Content)

	fmt.Println("\n--- 3️⃣ Bob Broadcasting Chat Message to Room 'go-devs' ---")
	bobWS.WriteJSON(domain.ChatEvent{
		Type:    "CHAT",
		Content: "Hey Alice, welcome to Go 90 Days Challenge!",
	})

	var event3 domain.ChatEvent
	aliceWS.ReadJSON(&event3)
	fmt.Printf("Alice Received Message from %s: \"%s\"\n", event3.Sender, event3.Content)

	time.Sleep(200 * time.Millisecond)

	fmt.Println("\n✅ Day 39 Real-Time Chat Backend executed successfully! Check chat_app.log for logs.")
}
