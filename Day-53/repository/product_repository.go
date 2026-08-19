package repository

import (
	"sync"
	"time"

	"day53/domain"
)

type InMemoryProductRepository struct {
	mu       sync.RWMutex
	products map[string]*domain.Product
}

func NewInMemoryProductRepository() domain.ProductRepository {
	return &InMemoryProductRepository{
		products: make(map[string]*domain.Product),
	}
}

func (r *InMemoryProductRepository) Create(p *domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.products[p.ID] = p
	return nil
}

func (r *InMemoryProductRepository) FindByID(id string) (*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if p, ok := r.products[id]; ok {
		cp := *p
		return &cp, nil
	}
	return nil, domain.ErrProductNotFound
}

func (r *InMemoryProductRepository) List(category string) ([]*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*domain.Product
	for _, p := range r.products {
		if category == "" || p.Category == category {
			cp := *p
			result = append(result, &cp)
		}
	}
	return result, nil
}

func (r *InMemoryProductRepository) UpdateStock(id string, delta int) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, ok := r.products[id]
	if !ok {
		return nil, domain.ErrProductNotFound
	}

	newStock := p.Stock + delta
	if newStock < 0 {
		return nil, domain.ErrInsufficientStock
	}

	p.Stock = newStock
	p.UpdatedAt = time.Now().UTC()
	cp := *p
	return &cp, nil
}
