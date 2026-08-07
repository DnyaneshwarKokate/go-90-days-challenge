package repository

import (
	"errors"

	"day-27/domain"

	"gorm.io/gorm"
)

type gormProductRepository struct {
	db *gorm.DB
}

// NewGORMProductRepository creates a GORM-backed implementation of domain.ProductRepository.
func NewGORMProductRepository(db *gorm.DB) domain.ProductRepository {
	return &gormProductRepository{db: db}
}

func (r *gormProductRepository) Create(p *domain.Product) (*domain.Product, error) {
	var existing domain.Product
	err := r.db.Where("sku = ?", p.SKU).First(&existing).Error
	if err == nil {
		return nil, domain.ErrDuplicateSKU
	}

	if err := r.db.Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (r *gormProductRepository) GetByID(id int) (*domain.Product, error) {
	var p domain.Product
	if err := r.db.First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (r *gormProductRepository) GetBySKU(sku string) (*domain.Product, error) {
	var p domain.Product
	if err := r.db.Where("sku = ?", sku).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (r *gormProductRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *gormProductRepository) Update(id int, p *domain.Product) (*domain.Product, error) {
	existing, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	updates := map[string]interface{}{}
	if p.Name != "" {
		updates["name"] = p.Name
	}
	if p.Price > 0 {
		updates["price"] = p.Price
	}
	if p.Stock >= 0 {
		updates["stock"] = p.Stock
	}

	if err := r.db.Model(existing).Updates(updates).Error; err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *gormProductRepository) Delete(id int) error {
	result := r.db.Delete(&domain.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrProductNotFound
	}
	return nil
}
