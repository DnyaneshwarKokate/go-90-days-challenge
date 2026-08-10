package repository

import (
	"context"
	"strings"
	"sync"
	"time"

	"day-31/domain"
)

type memoryProductRepository struct {
	mu       sync.RWMutex
	products map[string]*domain.Product
	skus     map[string]string
	logger   domain.Logger
}

func NewMemoryProductRepository(logger domain.Logger) domain.ProductRepository {
	return &memoryProductRepository{
		products: make(map[string]*domain.Product),
		skus:     make(map[string]string),
		logger:   logger,
	}
}

func (r *memoryProductRepository) Save(ctx context.Context, product *domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug(ctx, "DB Save Product requested", "sku", product.SKU)

	if existingID, exists := r.skus[product.SKU]; exists && existingID != product.ID {
		r.logger.Warn(ctx, "DB SKU Conflict", "sku", product.SKU)
		return domain.ErrSKUExists
	}

	r.products[product.ID] = product
	r.skus[product.SKU] = product.ID
	r.logger.Info(ctx, "DB Product saved", "product_id", product.ID, "version", product.Version)
	return nil
}

func (r *memoryProductRepository) BulkSave(ctx context.Context, products []*domain.Product) ([]*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Info(ctx, "DB BulkSave requested", "count", len(products))

	// Check SKU unique conflicts before mutating state
	for _, p := range products {
		if existingID, exists := r.skus[p.SKU]; exists && existingID != p.ID {
			r.logger.Warn(ctx, "DB BulkSave aborted: SKU conflict", "sku", p.SKU)
			return nil, domain.ErrSKUExists
		}
	}

	created := make([]*domain.Product, 0, len(products))
	for _, p := range products {
		r.products[p.ID] = p
		r.skus[p.SKU] = p.ID
		created = append(created, p)
	}

	r.logger.Info(ctx, "DB BulkSave completed successfully", "count", len(created))
	return created, nil
}

func (r *memoryProductRepository) FindByID(ctx context.Context, id string, includeDeleted bool) (*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "DB FindByID requested", "product_id", id, "include_deleted", includeDeleted)

	product, exists := r.products[id]
	if !exists {
		return nil, domain.ErrProductNotFound
	}

	if product.IsDeleted && !includeDeleted {
		return nil, domain.ErrProductNotFound
	}

	return product, nil
}

func (r *memoryProductRepository) FindBySKU(ctx context.Context, sku string) (*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	id, exists := r.skus[strings.ToUpper(strings.TrimSpace(sku))]
	if !exists {
		return nil, domain.ErrProductNotFound
	}
	return r.products[id], nil
}

func (r *memoryProductRepository) FindAll(ctx context.Context, includeDeleted bool) ([]*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "DB FindAll requested", "include_deleted", includeDeleted)

	result := make([]*domain.Product, 0)
	for _, p := range r.products {
		if p.IsDeleted && !includeDeleted {
			continue
		}
		result = append(result, p)
	}

	return result, nil
}

func (r *memoryProductRepository) PatchUpdate(ctx context.Context, id string, input domain.PatchProductInput) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug(ctx, "DB PatchUpdate requested", "product_id", id)

	product, exists := r.products[id]
	if !exists || product.IsDeleted {
		return nil, domain.ErrProductNotFound
	}

	// Optimistic Concurrency Control check
	if input.ExpectedVersion != nil {
		if product.Version != *input.ExpectedVersion {
			r.logger.Warn(ctx, "DB Concurrency Conflict detected", "product_id", id, "current_version", product.Version, "expected_version", *input.ExpectedVersion)
			return nil, domain.ErrConcurrencyConflict
		}
	}

	// Mutate only supplied non-nil fields
	if input.Name != nil && strings.TrimSpace(*input.Name) != "" {
		product.Name = strings.TrimSpace(*input.Name)
	}
	if input.Category != nil && strings.TrimSpace(*input.Category) != "" {
		product.Category = strings.ToUpper(strings.TrimSpace(*input.Category))
	}
	if input.Price != nil && *input.Price > 0 {
		product.Price = *input.Price
	}
	if input.Stock != nil && *input.Stock >= 0 {
		product.Stock = *input.Stock
	}

	product.Version++
	product.UpdatedAt = time.Now()

	r.products[id] = product
	r.logger.Info(ctx, "DB Product patched & version incremented", "product_id", id, "new_version", product.Version)
	return product, nil
}

func (r *memoryProductRepository) SoftDelete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug(ctx, "DB SoftDelete requested", "product_id", id)

	product, exists := r.products[id]
	if !exists {
		return domain.ErrProductNotFound
	}

	if product.IsDeleted {
		return domain.ErrAlreadyDeleted
	}

	now := time.Now()
	product.IsDeleted = true
	product.DeletedAt = &now
	product.UpdatedAt = now

	r.products[id] = product
	r.logger.Info(ctx, "DB Product soft-deleted", "product_id", id, "deleted_at", now)
	return nil
}

func (r *memoryProductRepository) Restore(ctx context.Context, id string) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug(ctx, "DB Restore requested", "product_id", id)

	product, exists := r.products[id]
	if !exists {
		return nil, domain.ErrProductNotFound
	}

	if !product.IsDeleted {
		return nil, domain.ErrNotDeleted
	}

	product.IsDeleted = false
	product.DeletedAt = nil
	product.UpdatedAt = time.Now()

	r.products[id] = product
	r.logger.Info(ctx, "DB Product restored", "product_id", id)
	return product, nil
}

func (r *memoryProductRepository) BulkSoftDelete(ctx context.Context, ids []string) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Info(ctx, "DB BulkSoftDelete requested", "count", len(ids))

	count := 0
	now := time.Now()
	for _, id := range ids {
		product, exists := r.products[id]
		if exists && !product.IsDeleted {
			product.IsDeleted = true
			product.DeletedAt = &now
			product.UpdatedAt = now
			r.products[id] = product
			count++
		}
	}

	r.logger.Info(ctx, "DB BulkSoftDelete completed", "deleted_count", count)
	return count, nil
}
