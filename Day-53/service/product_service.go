package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"day53/domain"
)

type ProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(req domain.CreateProductRequest) (*domain.Product, error) {
	if req.Name == "" || req.SKU == "" || req.Price <= 0 {
		return nil, fmt.Errorf("invalid product input parameters")
	}

	idBytes := make([]byte, 8)
	_, _ = rand.Read(idBytes)
	pID := "prod_" + hex.EncodeToString(idBytes)

	now := time.Now().UTC()
	p := &domain.Product{
		ID:        pID,
		SKU:       req.SKU,
		Name:      req.Name,
		Category:  req.Category,
		Price:     req.Price,
		Stock:     req.Stock,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) GetProduct(id string) (*domain.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) ListProducts(category string) ([]*domain.Product, error) {
	return s.repo.List(category)
}

func (s *ProductService) AdjustStock(id string, delta int) (*domain.Product, error) {
	return s.repo.UpdateStock(id, delta)
}
