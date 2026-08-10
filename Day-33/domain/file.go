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
	ErrFileTooLarge    = errors.New("file size exceeds maximum allowed limit (5MB)")
	ErrInvalidFileType = errors.New("invalid file extension or MIME type (allowed: .png, .jpg, .jpeg, .pdf)")
	ErrNoFileProvided  = errors.New("no file was provided in form data")
)

type FileMeta struct {
	ID          string    `json:"id"`
	OriginalName string   `json:"original_name"`
	StoredName   string   `json:"stored_name"`
	ContentType  string   `json:"content_type"`
	SizeBytes   int64     `json:"size_bytes"`
	URL         string    `json:"url"`
	UploadedAt  time.Time `json:"uploaded_at"`
}

type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}
