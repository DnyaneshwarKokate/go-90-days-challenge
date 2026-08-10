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
	ErrProductNotFound     = errors.New("product record not found")
	ErrSKUExists           = errors.New("product SKU already exists")
	ErrAlreadyDeleted      = errors.New("product is already soft deleted")
	ErrNotDeleted          = errors.New("product is not deleted")
	ErrConcurrencyConflict = errors.New("concurrency conflict: record was modified by another request")
	ErrInvalidInput        = errors.New("invalid request input parameters")
	ErrEmptyBulkRequest    = errors.New("bulk payload cannot be empty")
)

type Product struct {
	ID        string     `json:"id"`
	SKU       string     `json:"sku"`
	Name      string     `json:"name"`
	Category  string     `json:"category"`
	Price     float64    `json:"price"`
	Stock     int        `json:"stock"`
	Version   int        `json:"version"` // Optimistic Concurrency Control counter
	IsDeleted bool       `json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CreateProductInput struct {
	SKU      string  `json:"sku" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Stock    int     `json:"stock" binding:"gte=0"`
}

type BulkCreateInput struct {
	Items []CreateProductInput `json:"items" binding:"required,dive"`
}

type PatchProductInput struct {
	Name            *string  `json:"name,omitempty"`
	Category        *string  `json:"category,omitempty"`
	Price           *float64 `json:"price,omitempty"`
	Stock           *int     `json:"stock,omitempty"`
	ExpectedVersion *int     `json:"expected_version,omitempty"` // For optimistic concurrency check
}

type BulkDeleteInput struct {
	IDs []string `json:"ids" binding:"required"`
}

type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}

type ProductRepository interface {
	Save(ctx context.Context, product *Product) error
	BulkSave(ctx context.Context, products []*Product) ([]*Product, error)
	FindByID(ctx context.Context, id string, includeDeleted bool) (*Product, error)
	FindBySKU(ctx context.Context, sku string) (*Product, error)
	FindAll(ctx context.Context, includeDeleted bool) ([]*Product, error)
	PatchUpdate(ctx context.Context, id string, input PatchProductInput) (*Product, error)
	SoftDelete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) (*Product, error)
	BulkSoftDelete(ctx context.Context, ids []string) (int, error)
}
