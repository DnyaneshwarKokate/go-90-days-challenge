package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"day53/domain"
	"day53/repository"
	"day53/service"
)

func setupProductTestServer() *ProductHandler {
	repo := repository.NewInMemoryProductRepository()
	svc := service.NewProductService(repo)
	return NewProductHandler(svc)
}

func TestProductEndpoints(t *testing.T) {
	h := setupProductTestServer()

	// 1. Create Product
	createPayload := `{"sku":"PROD-101","name":"Mechanical Keyboard","category":"electronics","price":129.99,"stock":50}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBufferString(createPayload))
	rr := httptest.NewRecorder()

	h.HandleProducts(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d", rr.Code)
	}

	var prod domain.Product
	_ = json.Unmarshal(rr.Body.Bytes(), &prod)

	if prod.ID == "" || prod.SKU != "PROD-101" {
		t.Errorf("Unexpected product creation output: %+v", prod)
	}

	// 2. Adjust Stock
	stockPayload := `{"quantity_delta":-10}`
	reqStock := httptest.NewRequest(http.MethodPatch, "/api/v1/products/"+prod.ID+"/stock", bytes.NewBufferString(stockPayload))
	rrStock := httptest.NewRecorder()

	h.HandleProductByID(rrStock, reqStock)

	if rrStock.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK for stock update, got %d", rrStock.Code)
	}

	var updatedProd domain.Product
	_ = json.Unmarshal(rrStock.Body.Bytes(), &updatedProd)

	if updatedProd.Stock != 40 {
		t.Errorf("Expected updated stock to be 40, got %d", updatedProd.Stock)
	}
}
