package repository

import (
	"sync"
	"time"

	"day-27/domain"
)

type memoryProductRepository struct {
	mu       sync.RWMutex
	products map[int]domain.Product
	nextID   int
}

// NewMemoryProductRepository instantiates a thread-safe in-memory ProductRepository.
func NewMemoryProductRepository() domain.ProductRepository {
	return &memoryProductRepository{
		products: make(map[int]domain.Product),
		nextID:   1,
	}
}

func (r *memoryProductRepository) Create(p *domain.Product) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, item := range r.products {
		if item.SKU == p.SKU {
			return nil, domain.ErrDuplicateSKU
		}
	}

	p.ID = r.nextID
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
	r.products[p.ID] = *p
	r.nextID++

	return p, nil
}

func (r *memoryProductRepository) GetByID(id int) (*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, exists := r.products[id]
	if !exists {
		return nil, domain.ErrProductNotFound
	}
	return &p, nil
}

func (r *memoryProductRepository) GetBySKU(sku string) (*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, p := range r.products {
		if p.SKU == sku {
			return &p, nil
		}
	}
	return nil, domain.ErrProductNotFound
}

func (r *memoryProductRepository) GetAll() ([]domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]domain.Product, 0, len(r.products))
	for _, p := range r.products {
		list = append(list, p)
	}
	return list, nil
}

func (r *memoryProductRepository) Update(id int, p *domain.Product) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, exists := r.products[id]
	if !exists {
		return nil, domain.ErrProductNotFound
	}

	if p.Name != "" {
		existing.Name = p.Name
	}
	if p.Price > 0 {
		existing.Price = p.Price
	}
	if p.Stock >= 0 {
		existing.Stock = p.Stock
	}
	existing.UpdatedAt = time.Now()

	r.products[id] = existing
	return &existing, nil
}

func (r *memoryProductRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[id]; !exists {
		return domain.ErrProductNotFound
	}

	delete(r.products, id)
	return nil
}
