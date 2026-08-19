package domain

import (
	"errors"
	"time"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrInsufficientStock = errors.New("insufficient inventory stock")
)

type Product struct {
	ID          string    `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

type UpdateStockRequest struct {
	QuantityDelta int `json:"quantity_delta"`
}

type ProductRepository interface {
	Create(p *Product) error
	FindByID(id string) (*Product, error)
	List(category string) ([]*Product, error)
	UpdateStock(id string, delta int) (*Product, error)
}
