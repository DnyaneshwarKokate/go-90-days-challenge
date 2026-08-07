package domain

import (
	"errors"
	"time"
)

// Product represents the core domain model for a product in inventory.
type Product struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	SKU       string    `json:"sku" gorm:"uniqueIndex;not null"`
	Price     float64   `json:"price" gorm:"not null"`
	Stock     int       `json:"stock" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Domain Error Contracts
var (
	ErrProductNotFound  = errors.New("product not found")
	ErrDuplicateSKU     = errors.New("product with this SKU already exists")
	ErrInvalidProduct   = errors.New("product details are invalid")
)

// ProductRepository defines the data access abstraction contract for Products.
// Any storage engine (In-Memory, PostgreSQL/GORM, MongoDB) must satisfy this interface.
type ProductRepository interface {
	Create(product *Product) (*Product, error)
	GetByID(id int) (*Product, error)
	GetBySKU(sku string) (*Product, error)
	GetAll() ([]Product, error)
	Update(id int, product *Product) (*Product, error)
	Delete(id int) error
}

// ProductUseCase defines the business logic abstraction contract.
type ProductUseCase interface {
	CreateProduct(name, sku string, price float64, stock int) (*Product, error)
	GetProductByID(id int) (*Product, error)
	GetProductBySKU(sku string) (*Product, error)
	ListProducts() ([]Product, error)
	UpdateProduct(id int, name string, price float64, stock int) (*Product, error)
	DeleteProduct(id int) error
}
