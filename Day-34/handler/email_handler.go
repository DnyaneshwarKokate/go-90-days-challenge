package handler

import (
	"net/http"
	"time"

	"day-34/domain"
	"day-34/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmailHandler struct {
	emailService *service.EmailService
	logger       domain.Logger
}

func NewEmailHandler(emailService *service.EmailService, logger domain.Logger) *EmailHandler {
	return &EmailHandler{
		emailService: emailService,
		logger:       logger,
	}
}

func (h *EmailHandler) SendSync(c *gin.Context) {
	var input domain.SendEmailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warn(c.Request.Context(), "Invalid email request payload", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	name := input.Name
	if name == "" {
		name = "User"
	}

	htmlBody, err := h.emailService.RenderTemplate(input.Template, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render email template"})
		return
	}

	msg := domain.EmailMessage{
		ID:       uuid.New().String(),
		To:       input.To,
		Subject:  input.Subject,
		BodyHTML: htmlBody,
		Status:   "SENT",
		SentAt:   time.Now(),
	}

	if err := h.emailService.SendEmail(c.Request.Context(), msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email sent successfully",
		"data":    msg,
	})
}

func (h *EmailHandler) SendAsync(c *gin.Context) {
	var input domain.SendEmailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	name := input.Name
	if name == "" {
		name = "User"
	}

	htmlBody, err := h.emailService.RenderTemplate(input.Template, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render email template"})
		return
	}

	msg := domain.EmailMessage{
		To:       input.To,
		Subject:  input.Subject,
		BodyHTML: htmlBody,
	}

	h.emailService.EnqueueEmail(c.Request.Context(), msg)

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Email queued for background delivery",
	})
}
