package ws

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"day-38/domain"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow cross-origin requests for demo
	},
}

type ClientConnection struct {
	ID     string
	Conn   *websocket.Conn
	Send   chan domain.WSMessage
	logger domain.Logger
	mu     sync.Mutex
}

func NewClientConnection(conn *websocket.Conn, logger domain.Logger) *ClientConnection {
	return &ClientConnection{
		ID:     uuid.New().String()[:8],
		Conn:   conn,
		Send:   make(chan domain.WSMessage, 256),
		logger: logger,
	}
}

func (c *ClientConnection) ReadLoop(ctx context.Context, onMessage func(msg domain.WSMessage)) {
	defer func() {
		c.Conn.Close()
		c.logger.Info(ctx, "WebSocket client connection closed", "client_id", c.ID)
	}()

	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, messageBytes, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.Warn(ctx, "WebSocket read error", "client_id", c.ID, "error", err)
			}
			break
		}

		var msg domain.WSMessage
		if err := json.Unmarshal(messageBytes, &msg); err != nil {
			msg = domain.WSMessage{
				Type:      "ECHO",
				Sender:    c.ID,
				Content:   string(messageBytes),
				Timestamp: time.Now(),
			}
		}

		c.logger.Info(ctx, "WebSocket received message", "client_id", c.ID, "content", msg.Content)
		onMessage(msg)
	}
}

func (c *ClientConnection) WriteLoop(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			json.NewEncoder(w).Encode(msg)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
