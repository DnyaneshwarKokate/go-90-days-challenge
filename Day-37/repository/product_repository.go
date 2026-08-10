package repository

import (
	"context"
	"sync"
	"time"

	"day-37/domain"
)

type memoryProductRepository struct {
	mu       sync.RWMutex
	products map[string]*domain.Product
	logger   domain.Logger
}

func NewMemoryProductRepository(logger domain.Logger) domain.ProductRepository {
	return &memoryProductRepository{
		products: make(map[string]*domain.Product),
		logger:   logger,
	}
}

func (r *memoryProductRepository) Save(ctx context.Context, p *domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.products[p.ID] = p
	r.logger.Info(ctx, "DB Save Product successful", "id", p.ID)
	return nil
}

func (r *memoryProductRepository) FindByID(ctx context.Context, id string) (*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Info(ctx, "🐢 DB QUERY EXECUTING (Slow Disk Look Up)", "id", id)
	time.Sleep(20 * time.Millisecond) // Simulate DB I/O latency

	product, exists := r.products[id]
	if !exists {
		return nil, domain.ErrProductNotFound
	}

	return product, nil
}

func (r *memoryProductRepository) Update(ctx context.Context, p *domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[p.ID]; !exists {
		return domain.ErrProductNotFound
	}

	p.UpdatedAt = time.Now()
	r.products[p.ID] = p
	r.logger.Info(ctx, "DB Update Product successful", "id", p.ID)
	return nil
}
