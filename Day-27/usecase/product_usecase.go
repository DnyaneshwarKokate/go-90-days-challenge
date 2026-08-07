package usecase

import (
	"day-27/domain"
)

type productUseCase struct {
	repo domain.ProductRepository
}

// NewProductUseCase constructs the Product Business Logic layer, depending strictly on ProductRepository contract.
func NewProductUseCase(repo domain.ProductRepository) domain.ProductUseCase {
	return &productUseCase{repo: repo}
}

func (u *productUseCase) CreateProduct(name, sku string, price float64, stock int) (*domain.Product, error) {
	if name == "" || sku == "" || price <= 0 || stock < 0 {
		return nil, domain.ErrInvalidProduct
	}

	p := &domain.Product{
		Name:  name,
		SKU:   sku,
		Price: price,
		Stock: stock,
	}

	return u.repo.Create(p)
}

func (u *productUseCase) GetProductByID(id int) (*domain.Product, error) {
	return u.repo.GetByID(id)
}

func (u *productUseCase) GetProductBySKU(sku string) (*domain.Product, error) {
	return u.repo.GetBySKU(sku)
}

func (u *productUseCase) ListProducts() ([]domain.Product, error) {
	return u.repo.GetAll()
}

func (u *productUseCase) UpdateProduct(id int, name string, price float64, stock int) (*domain.Product, error) {
	p := &domain.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}
	return u.repo.Update(id, p)
}

func (u *productUseCase) DeleteProduct(id int) error {
	return u.repo.Delete(id)
}
