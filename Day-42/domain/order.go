package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrInvalidOrderAmount = errors.New("order amount must be greater than zero")
	ErrOrderNotFound      = errors.New("order not found")
)

type Order struct {
	ID        string    `json:"id"`
	Customer  string    `json:"customer"`
	Subtotal  float64   `json:"subtotal"`
	Discount  float64   `json:"discount"`
	Tax       float64   `json:"tax"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderRepository interface {
	Save(ctx context.Context, order *Order) error
	FindByID(ctx context.Context, id string) (*Order, error)
}
