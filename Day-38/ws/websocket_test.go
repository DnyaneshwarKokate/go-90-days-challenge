package ws_test

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"day-38/domain"
	"day-38/handler"
	"day-38/logger"
	"day-38/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func TestWebSocket_UpgradeAndEcho(t *testing.T) {
	gin.SetMode(gin.TestMode)
	zapLog, _ := logger.NewZapLogger("test", "")
	wsHandler := handler.NewWSHandler(zapLog)

	router := gin.New()
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))
	router.GET("/ws", wsHandler.HandleWebSocket)

	server := httptest.NewServer(router)
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	wsClient, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("failed to connect to websocket server: %v", err)
	}
	defer wsClient.Close()

	input := domain.WSMessage{
		Type:      "CHAT",
		Sender:    "Tester",
		Content:   "Unit Test Message",
		Timestamp: time.Now(),
	}

	if err := wsClient.WriteJSON(input); err != nil {
		t.Fatalf("failed to write json to websocket: %v", err)
	}

	var resp domain.WSMessage
	if err := wsClient.ReadJSON(&resp); err != nil {
		t.Fatalf("failed to read json from websocket: %v", err)
	}

	if !strings.Contains(resp.Content, "Server Echo: Unit Test Message") {
		t.Errorf("unexpected echo message: %s", resp.Content)
	}
}
