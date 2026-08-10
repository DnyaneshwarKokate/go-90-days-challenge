package usecase_test

import (
	"context"
	"testing"
	"time"

	"day-37/cache"
	"day-37/domain"
	"day-37/logger"
	"day-37/repository"
	"day-37/usecase"
)

func TestCacheAsidePattern_Lifecycle(t *testing.T) {
	zapLog, _ := logger.NewZapLogger("test", "")
	repo := repository.NewMemoryProductRepository(zapLog)
	prodCache := cache.NewRedisProductCache("localhost:6379", zapLog)
	uc := usecase.NewProductUseCase(repo, prodCache, 5*time.Minute, zapLog)

	ctx := context.Background()

	// 1. Create Product
	p, err := uc.CreateProduct(ctx, domain.CreateProductInput{
		Name:     "Test Tablet",
		Price:    499.99,
		Category: "ELECTRONICS",
	})
	if err != nil {
		t.Fatalf("failed to create product: %v", err)
	}

	// 2. First fetch (Cache Miss -> DB Query -> Set Cache)
	p1, err := uc.GetProductByID(ctx, p.ID)
	if err != nil || p1.Name != "Test Tablet" {
		t.Errorf("expected product 'Test Tablet', got %v, err: %v", p1, err)
	}

	// 3. Second fetch (Cache Hit)
	p2, err := uc.GetProductByID(ctx, p.ID)
	if err != nil || p2.ID != p.ID {
		t.Errorf("expected product ID %s from cache, got %v", p.ID, p2)
	}

	// 4. Update Product (Cache Invalidation)
	newPrice := 449.99
	updated, err := uc.UpdateProduct(ctx, p.ID, domain.UpdateProductInput{Price: &newPrice})
	if err != nil || updated.Price != 449.99 {
		t.Errorf("expected price 449.99 after update, got %v", updated)
	}
}
