package domain

import (
	"context"
	"errors"
)

type ContextKey string

const (
	RequestIDKey ContextKey = "X-Request-ID"
)

var (
	ErrKeyNotFound = errors.New("redis key not found")
)

type KeyValueInput struct {
	Key        string        `json:"key" binding:"required"`
	Value      string        `json:"value" binding:"required"`
	TTLSeconds int           `json:"ttl_seconds,omitempty"`
}

type HashInput struct {
	Key    string            `json:"key" binding:"required"`
	Fields map[string]string `json:"fields" binding:"required"`
}

type ListInput struct {
	Key   string   `json:"key" binding:"required"`
	Items []string `json:"items" binding:"required"`
}

type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}
