package handler

import (
	"net/http"
	"strconv"

	"day-35/client"
	"day-35/domain"

	"github.com/gin-gonic/gin"
)

type ExternalHandler struct {
	restClient *client.ResilientHTTPClient
	logger     domain.Logger
}

func NewExternalHandler(restClient *client.ResilientHTTPClient, logger domain.Logger) *ExternalHandler {
	return &ExternalHandler{
		restClient: restClient,
		logger:     logger,
	}
}

func (h *ExternalHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID parameter"})
		return
	}

	user, err := h.restClient.GetUser(c.Request.Context(), userID)
	if err != nil {
		h.logger.Error(c.Request.Context(), "External API request failed", "user_id", userID, "error", err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch external user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "External user retrieved successfully",
		"data":    user,
	})
}
