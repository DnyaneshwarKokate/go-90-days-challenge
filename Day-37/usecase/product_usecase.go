package usecase

import (
	"context"
	"fmt"
	"time"

	"day-37/domain"

	"github.com/google/uuid"
)

type ProductUseCase struct {
	repo   domain.ProductRepository
	cache  domain.ProductCache
	ttl    time.Duration
	logger domain.Logger
}

func NewProductUseCase(repo domain.ProductRepository, cache domain.ProductCache, ttl time.Duration, logger domain.Logger) *ProductUseCase {
	return &ProductUseCase{
		repo:   repo,
		cache:  cache,
		ttl:    ttl,
		logger: logger,
	}
}

func (u *ProductUseCase) CreateProduct(ctx context.Context, input domain.CreateProductInput) (*domain.Product, error) {
	now := time.Now()
	p := &domain.Product{
		ID:        fmt.Sprintf("prd_%s", uuid.New().String()[:8]),
		Name:      input.Name,
		Price:     input.Price,
		Category:  input.Category,
		UpdatedAt: now,
	}

	if err := u.repo.Save(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (u *ProductUseCase) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	// 1. Read-Through: Check Redis Cache First
	cachedProduct, err := u.cache.Get(ctx, id)
	if err == nil && cachedProduct != nil {
		return cachedProduct, nil
	}

	// 2. Cache Miss: Query Underlying DB
	dbProduct, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 3. Populate Redis Cache asynchronously or synchronously for subsequent queries
	if err := u.cache.Set(ctx, dbProduct, u.ttl); err != nil {
		u.logger.Warn(ctx, "Failed to set product in Redis cache", "id", id, "error", err)
	}

	return dbProduct, nil
}

func (u *ProductUseCase) UpdateProduct(ctx context.Context, id string, input domain.UpdateProductInput) (*domain.Product, error) {
	product, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Price != nil {
		product.Price = *input.Price
	}

	// 1. Update DB
	if err := u.repo.Update(ctx, product); err != nil {
		return nil, err
	}

	// 2. Cache Invalidation: Delete old cached entry in Redis
	if err := u.cache.Delete(ctx, id); err != nil {
		u.logger.Warn(ctx, "Failed to invalidate product cache", "id", id, "error", err)
	}

	return product, nil
}
