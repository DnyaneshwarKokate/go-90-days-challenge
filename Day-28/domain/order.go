package domain

import (
	"errors"
	"time"
)

// Sentinel Errors
var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrInvalidOrderAmount = errors.New("order amount must be greater than zero")
	ErrEmptyCustomerEmail = errors.New("customer email cannot be empty")
)

// Order represents an order entity in the system.
type Order struct {
	ID            string    `json:"id"`
	CustomerEmail string    `json:"customer_email"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

// CreateOrderInput holds data required to place an order.
type CreateOrderInput struct {
	CustomerEmail string  `json:"customer_email" binding:"required,email"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
}

// Logger interface decouples logging concerns from business logic.
type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// NotificationService interface abstracts email/SMS delivery.
type NotificationService interface {
	SendOrderConfirmation(order *Order) error
}

// OrderRepository interface abstracts database persistence logic.
type OrderRepository interface {
	Save(order *Order) error
	FindByID(id string) (*Order, error)
	FindAll() ([]*Order, error)
}
