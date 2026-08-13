package domain

import (
	"errors"
	"time"
)

// Order represents an e-commerce order managed in Kubernetes.
type Order struct {
	ID        string    `json:"id"`
	Customer  string    `json:"customer"`
	Item      string    `json:"item"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateOrderRequest defines payload for creating an order.
type CreateOrderRequest struct {
	Customer string  `json:"customer"`
	Item     string  `json:"item"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Validate checks request input parameters.
func (req *CreateOrderRequest) Validate() error {
	if req.Customer == "" {
		return errors.New("customer name is required")
	}
	if req.Item == "" {
		return errors.New("item name is required")
	}
	if req.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	if req.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}
