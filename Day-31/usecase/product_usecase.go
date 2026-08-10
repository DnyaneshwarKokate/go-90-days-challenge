package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"day-31/domain"

	"github.com/google/uuid"
)

type ProductUseCase struct {
	repo   domain.ProductRepository
	logger domain.Logger
}

func NewProductUseCase(repo domain.ProductRepository, logger domain.Logger) *ProductUseCase {
	return &ProductUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (u *ProductUseCase) CreateProduct(ctx context.Context, input domain.CreateProductInput) (*domain.Product, error) {
	sku := strings.ToUpper(strings.TrimSpace(input.SKU))
	u.logger.Info(ctx, "Processing Single Product Create", "sku", sku, "name", input.Name)

	if sku == "" || strings.TrimSpace(input.Name) == "" {
		return nil, domain.ErrInvalidInput
	}

	existing, err := u.repo.FindBySKU(ctx, sku)
	if err == nil && existing != nil {
		u.logger.Warn(ctx, "Product creation rejected: SKU already exists", "sku", sku)
		return nil, domain.ErrSKUExists
	}

	now := time.Now()
	product := &domain.Product{
		ID:        fmt.Sprintf("prd_%s", uuid.New().String()[:8]),
		SKU:       sku,
		Name:      strings.TrimSpace(input.Name),
		Category:  strings.ToUpper(strings.TrimSpace(input.Category)),
		Price:     input.Price,
		Stock:     input.Stock,
		Version:   1, // Initial version counter
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.repo.Save(ctx, product); err != nil {
		return nil, err
	}

	u.logger.Info(ctx, "Product created successfully", "product_id", product.ID, "sku", product.SKU)
	return product, nil
}

func (u *ProductUseCase) BulkCreateProducts(ctx context.Context, input domain.BulkCreateInput) ([]*domain.Product, error) {
	u.logger.Info(ctx, "Processing Bulk Product Create", "item_count", len(input.Items))

	if len(input.Items) == 0 {
		return nil, domain.ErrEmptyBulkRequest
	}

	now := time.Now()
	products := make([]*domain.Product, 0, len(input.Items))
	seenSKUs := make(map[string]bool)

	for _, item := range input.Items {
		sku := strings.ToUpper(strings.TrimSpace(item.SKU))
		if sku == "" || strings.TrimSpace(item.Name) == "" {
			return nil, domain.ErrInvalidInput
		}
		if seenSKUs[sku] {
			u.logger.Warn(ctx, "Bulk create failed: Duplicate SKU in payload", "sku", sku)
			return nil, domain.ErrSKUExists
		}
		seenSKUs[sku] = true

		product := &domain.Product{
			ID:        fmt.Sprintf("prd_%s", uuid.New().String()[:8]),
			SKU:       sku,
			Name:      strings.TrimSpace(item.Name),
			Category:  strings.ToUpper(strings.TrimSpace(item.Category)),
			Price:     item.Price,
			Stock:     item.Stock,
			Version:   1,
			IsDeleted: false,
			CreatedAt: now,
			UpdatedAt: now,
		}
		products = append(products, product)
	}

	created, err := u.repo.BulkSave(ctx, products)
	if err != nil {
		return nil, err
	}

	u.logger.Info(ctx, "Bulk Product Create completed successfully", "count", len(created))
	return created, nil
}

func (u *ProductUseCase) GetProductByID(ctx context.Context, id string, includeDeleted bool) (*domain.Product, error) {
	u.logger.Debug(ctx, "Fetching product by ID", "product_id", id, "include_deleted", includeDeleted)
	return u.repo.FindByID(ctx, id, includeDeleted)
}

func (u *ProductUseCase) ListProducts(ctx context.Context, includeDeleted bool) ([]*domain.Product, error) {
	u.logger.Info(ctx, "Listing all products", "include_deleted", includeDeleted)
	return u.repo.FindAll(ctx, includeDeleted)
}

func (u *ProductUseCase) PatchUpdateProduct(ctx context.Context, id string, input domain.PatchProductInput) (*domain.Product, error) {
	u.logger.Info(ctx, "Processing Partial PATCH Update", "product_id", id)

	patched, err := u.repo.PatchUpdate(ctx, id, input)
	if err != nil {
		if err == domain.ErrConcurrencyConflict {
			u.logger.Warn(ctx, "PATCH update aborted: Concurrency conflict detected", "product_id", id)
		}
		return nil, err
	}

	u.logger.Info(ctx, "Product patched successfully", "product_id", id, "new_version", patched.Version)
	return patched, nil
}

func (u *ProductUseCase) SoftDeleteProduct(ctx context.Context, id string) error {
	u.logger.Info(ctx, "Processing Soft Delete", "product_id", id)
	return u.repo.SoftDelete(ctx, id)
}

func (u *ProductUseCase) RestoreProduct(ctx context.Context, id string) (*domain.Product, error) {
	u.logger.Info(ctx, "Processing Product Restore", "product_id", id)
	return u.repo.Restore(ctx, id)
}

func (u *ProductUseCase) BulkSoftDeleteProducts(ctx context.Context, ids []string) (int, error) {
	u.logger.Info(ctx, "Processing Bulk Soft Delete", "count", len(ids))
	if len(ids) == 0 {
		return 0, domain.ErrEmptyBulkRequest
	}
	return u.repo.BulkSoftDelete(ctx, ids)
}
