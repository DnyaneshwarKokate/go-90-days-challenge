package hub_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"day-39/domain"
	"day-39/handler"
	"day-39/hub"
	"day-39/logger"
	"day-39/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func TestChatHub_Broadcasting(t *testing.T) {
	gin.SetMode(gin.TestMode)
	zapLog, _ := logger.NewZapLogger("test", "")

	chatHub := hub.NewChatHub(zapLog)
	chatHandler := handler.NewChatHandler(chatHub, zapLog)

	router := gin.New()
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))
	router.GET("/ws/chat", chatHandler.JoinRoom)

	server := httptest.NewServer(router)
	defer server.Close()

	baseURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws/chat"

	c1, _, err := websocket.DefaultDialer.Dial(baseURL+"?room=test-room&user=User1", nil)
	if err != nil {
		t.Fatalf("failed c1 dial: %v", err)
	}
	defer c1.Close()

	var join1 domain.ChatEvent
	c1.ReadJSON(&join1) // User1 system join message

	c2, _, err := websocket.DefaultDialer.Dial(baseURL+"?room=test-room&user=User2", nil)
	if err != nil {
		t.Fatalf("failed c2 dial: %v", err)
	}
	defer c2.Close()

	var join2 domain.ChatEvent
	c1.ReadJSON(&join2) // User1 receives User2 join message

	if join2.Sender != "SYSTEM" || !strings.Contains(join2.Content, "User2 joined") {
		t.Errorf("expected User2 join event, got %v", join2)
	}

	c2.WriteJSON(domain.ChatEvent{
		Content: "Broadcast Test",
	})

	var msg domain.ChatEvent
	c1.ReadJSON(&msg)

	if msg.Sender != "User2" || msg.Content != "Broadcast Test" {
		t.Errorf("expected broadcast message from User2, got %v", msg)
	}
}
