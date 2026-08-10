package handler

import (
	"time"

	"day-38/domain"
	"day-38/ws"

	"github.com/gin-gonic/gin"
)

type WSHandler struct {
	logger domain.Logger
}

func NewWSHandler(logger domain.Logger) *WSHandler {
	return &WSHandler{logger: logger}
}

func (h *WSHandler) HandleWebSocket(c *gin.Context) {
	conn, err := ws.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		h.logger.Error(c.Request.Context(), "Failed to upgrade connection to WebSocket", "error", err)
		return
	}

	clientConn := ws.NewClientConnection(conn, h.logger)
	h.logger.Info(c.Request.Context(), "WebSocket connection established", "client_id", clientConn.ID)

	go clientConn.WriteLoop(c.Request.Context())

	clientConn.ReadLoop(c.Request.Context(), func(msg domain.WSMessage) {
		echoMsg := domain.WSMessage{
			Type:      "ECHO_RESPONSE",
			Sender:    "SERVER",
			Content:   "Server Echo: " + msg.Content,
			Timestamp: time.Now(),
		}
		clientConn.Send <- echoMsg
	})
}
