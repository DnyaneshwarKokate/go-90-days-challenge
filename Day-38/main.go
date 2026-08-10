package main

import (
	"context"
	"fmt"
	"net/http/httptest"
	"os"
	"strings"
	"time"

	"day-38/domain"
	"day-38/handler"
	"day-38/logger"
	"day-38/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("==========================================================================")
	fmt.Println("🚀 Day 38: WebSockets in Go (Gorilla WebSocket, Upgrading, Read/Write Pumps)")
	fmt.Println("==========================================================================")

	logFilePath := "ws_app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()
	zapLog.Info(ctx, "WebSocket Server initializing")

	wsHandler := handler.NewWSHandler(zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	router.GET("/ws", wsHandler.HandleWebSocket)

	// Start local test server
	server := httptest.NewServer(router)
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	fmt.Printf("\n--- 1️⃣ Dialing WebSocket Endpoint (%s) ---\n", wsURL)

	wsClient, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		fmt.Printf("Failed to dial WebSocket server: %v\n", err)
		os.Exit(1)
	}
	defer wsClient.Close()

	fmt.Println("\n--- 2️⃣ Sending JSON Message via WebSocket ---")
	msg := domain.WSMessage{
		Type:      "MESSAGE",
		Sender:    "Client_User_1",
		Content:   "Hello Go 90 Days WebSocket Server!",
		Timestamp: time.Now(),
	}
	wsClient.WriteJSON(msg)

	fmt.Println("\n--- 3️⃣ Reading Server Echo Response ---")
	var response domain.WSMessage
	wsClient.ReadJSON(&response)
	fmt.Printf("Received Echo Response from Server: Type=%s | Sender=%s | Content=\"%s\"\n", response.Type, response.Sender, response.Content)

	time.Sleep(200 * time.Millisecond)

	fmt.Println("\n✅ Day 38 WebSockets in Go executed successfully! Check ws_app.log for logs.")
}
