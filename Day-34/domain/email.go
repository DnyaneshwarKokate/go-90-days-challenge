package domain

import (
	"context"
	"errors"
	"time"
)

type ContextKey string

const (
	RequestIDKey ContextKey = "X-Request-ID"
)

var (
	ErrInvalidEmailAddress = errors.New("invalid email address")
	ErrEmptySubject        = errors.New("email subject cannot be empty")
)

type Attachment struct {
	Filename string `json:"filename"`
	Content  []byte `json:"-"`
}

type EmailMessage struct {
	ID          string       `json:"id"`
	To          []string     `json:"to"`
	Subject     string       `json:"subject"`
	BodyHTML    string       `json:"body_html"`
	Attachments []Attachment `json:"attachments,omitempty"`
	Status      string       `json:"status"` // SENT, QUEUED, FAILED
	SentAt      time.Time    `json:"sent_at"`
}

type SendEmailInput struct {
	To       []string `json:"to" binding:"required,gt=0"`
	Subject  string   `json:"subject" binding:"required"`
	Template string   `json:"template"` // WELCOME, PASSWORD_RESET, ORDER_CONFIRMATION
	Name     string   `json:"name"`
}

type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}
