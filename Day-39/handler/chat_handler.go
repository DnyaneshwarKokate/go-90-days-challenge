package handler

import (
	"day-39/domain"
	"day-39/hub"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	chatHub *hub.ChatHub
	logger  domain.Logger
}

func NewChatHandler(chatHub *hub.ChatHub, logger domain.Logger) *ChatHandler {
	return &ChatHandler{
		chatHub: chatHub,
		logger:  logger,
	}
}

func (h *ChatHandler) JoinRoom(c *gin.Context) {
	roomID := c.DefaultQuery("room", "general")
	username := c.DefaultQuery("user", "Anonymous")

	h.chatHub.HandleWS(c.Writer, c.Request, roomID, username)
}
