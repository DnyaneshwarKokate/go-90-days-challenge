package domain

import (
	"context"
	"time"
)

type ContextKey string

const (
	RequestIDKey ContextKey = "X-Request-ID"
)

type Task struct {
	ID        int       `json:"id"`
	JobType   string    `json:"job_type"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskResult struct {
	TaskID      int           `json:"task_id"`
	WorkerID    int           `json:"worker_id"`
	Result      string        `json:"result"`
	ExecutionMs time.Duration `json:"execution_ms"`
	Err         error         `json:"-"`
}

type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}
