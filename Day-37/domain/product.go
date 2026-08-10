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
	ErrProductNotFound = errors.New("product record not found")
)

type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Category  string    `json:"category"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductInput struct {
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Category string  `json:"category" binding:"required"`
}

type UpdateProductInput struct {
	Name  *string  `json:"name,omitempty"`
	Price *float64 `json:"price,omitempty"`
}

type ProductRepository interface {
	Save(ctx context.Context, p *Product) error
	FindByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, p *Product) error
}

type ProductCache interface {
	Get(ctx context.Context, id string) (*Product, error)
	Set(ctx context.Context, p *Product, ttl time.Duration) error
	Delete(ctx context.Context, id string) error
}

type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}
