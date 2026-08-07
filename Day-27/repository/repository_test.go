package repository_test

import (
	"testing"

	"day-27/domain"
	"day-27/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func runRepositoryTests(t *testing.T, repo domain.ProductRepository) {
	// 1. Create Product
	p1 := &domain.Product{
		Name:  "Mechanical Keyboard",
		SKU:   "KB-101",
		Price: 99.99,
		Stock: 15,
	}

	created, err := repo.Create(p1)
	if err != nil {
		t.Fatalf("expected no error on Create, got %v", err)
	}
	if created.ID == 0 {
		t.Errorf("expected product ID to be assigned, got 0")
	}

	// 2. Duplicate SKU error
	pDuplicate := &domain.Product{
		Name:  "Another Keyboard",
		SKU:   "KB-101",
		Price: 80.00,
		Stock: 5,
	}
	_, err = repo.Create(pDuplicate)
	if err != domain.ErrDuplicateSKU {
		t.Errorf("expected ErrDuplicateSKU, got %v", err)
	}

	// 3. GetByID
	fetched, err := repo.GetByID(created.ID)
	if err != nil {
		t.Fatalf("expected to find product by ID, got error %v", err)
	}
	if fetched.Name != p1.Name {
		t.Errorf("expected name %s, got %s", p1.Name, fetched.Name)
	}

	// 4. Update
	updated, err := repo.Update(created.ID, &domain.Product{Price: 109.99, Stock: 20})
	if err != nil {
		t.Fatalf("expected successful update, got %v", err)
	}
	if updated.Price != 109.99 {
		t.Errorf("expected updated price 109.99, got %f", updated.Price)
	}

	// 5. GetAll
	all, err := repo.GetAll()
	if err != nil {
		t.Fatalf("expected GetAll to succeed, got %v", err)
	}
	if len(all) != 1 {
		t.Errorf("expected 1 product in repo, got %d", len(all))
	}

	// 6. Delete
	err = repo.Delete(created.ID)
	if err != nil {
		t.Fatalf("expected Delete to succeed, got %v", err)
	}

	// 7. Get after Delete
	_, err = repo.GetByID(created.ID)
	if err != domain.ErrProductNotFound {
		t.Errorf("expected ErrProductNotFound after deletion, got %v", err)
	}
}

func TestMemoryProductRepository(t *testing.T) {
	repo := repository.NewMemoryProductRepository()
	runRepositoryTests(t, repo)
}

func TestGORMProductRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to sqlite memory db: %v", err)
	}

	if err := db.AutoMigrate(&domain.Product{}); err != nil {
		t.Fatalf("failed to auto migrate domain.Product: %v", err)
	}

	repo := repository.NewGORMProductRepository(db)
	runRepositoryTests(t, repo)
}
