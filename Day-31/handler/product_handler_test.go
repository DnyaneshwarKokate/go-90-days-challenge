package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"day-31/domain"
	"day-31/handler"
	"day-31/logger"
	"day-31/middleware"
	"day-31/repository"
	"day-31/usecase"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	zapLog, _ := logger.NewZapLogger("test", "")

	productRepo := repository.NewMemoryProductRepository(zapLog)
	productUC := usecase.NewProductUseCase(productRepo, zapLog)
	productHandler := handler.NewProductHandler(productUC, zapLog)

	router := gin.New()
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	v1 := router.Group("/api/v1/products")
	{
		v1.POST("", productHandler.CreateProduct)
		v1.POST("/bulk", productHandler.BulkCreateProducts)
		v1.GET("", productHandler.ListProducts)
		v1.GET("/:id", productHandler.GetProductByID)
		v1.PATCH("/:id", productHandler.PatchUpdateProduct)
		v1.POST("/:id/restore", productHandler.RestoreProduct)
		v1.DELETE("/:id", productHandler.SoftDeleteProduct)
		v1.POST("/bulk-delete", productHandler.BulkSoftDeleteProducts)
	}

	return router
}

func TestAdvancedCRUD_Lifecycle(t *testing.T) {
	router := setupTestRouter()

	// 1. Create Product
	createInput := domain.CreateProductInput{
		SKU:      "TEST-SKU-01",
		Name:     "Test Keyboard",
		Category: "PERIPHERALS",
		Price:    89.99,
		Stock:    20,
	}
	body, _ := json.Marshal(createInput)
	req1, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	if w1.Code != http.StatusCreated {
		t.Fatalf("expected status 201 Created, got %d. Body: %s", w1.Code, w1.Body.String())
	}

	var res1 map[string]interface{}
	json.Unmarshal(w1.Body.Bytes(), &res1)
	productData := res1["data"].(map[string]interface{})
	productID := productData["id"].(string)
	version := int(productData["version"].(float64))

	if version != 1 {
		t.Errorf("expected initial version 1, got %d", version)
	}

	// 2. Partial Update (PATCH) with correct Expected Version
	newPrice := 79.99
	patchInput := domain.PatchProductInput{
		Price:           &newPrice,
		ExpectedVersion: &version,
	}
	patchBody, _ := json.Marshal(patchInput)
	req2, _ := http.NewRequest("PATCH", "/api/v1/products/"+productID, bytes.NewBuffer(patchBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK for PATCH update, got %d. Body: %s", w2.Code, w2.Body.String())
	}

	var res2 map[string]interface{}
	json.Unmarshal(w2.Body.Bytes(), &res2)
	patchedData := res2["data"].(map[string]interface{})
	newVersion := int(patchedData["version"].(float64))

	if newVersion != 2 {
		t.Errorf("expected version to increment to 2, got %d", newVersion)
	}

	// 3. Optimistic Concurrency Conflict (Passing outdated version 1)
	staleVersion := 1
	patchInputStale := domain.PatchProductInput{
		Price:           &newPrice,
		ExpectedVersion: &staleVersion,
	}
	staleBody, _ := json.Marshal(patchInputStale)
	req3, _ := http.NewRequest("PATCH", "/api/v1/products/"+productID, bytes.NewBuffer(staleBody))
	req3.Header.Set("Content-Type", "application/json")
	w3 := httptest.NewRecorder()
	router.ServeHTTP(w3, req3)

	if w3.Code != http.StatusConflict {
		t.Errorf("expected status 409 Conflict on stale version, got %d", w3.Code)
	}

	// 4. Soft Delete
	req4, _ := http.NewRequest("DELETE", "/api/v1/products/"+productID, nil)
	w4 := httptest.NewRecorder()
	router.ServeHTTP(w4, req4)

	if w4.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK for soft delete, got %d", w4.Code)
	}

	// 5. Verify Soft Deleted item is hidden from standard GET
	req5, _ := http.NewRequest("GET", "/api/v1/products/"+productID, nil)
	w5 := httptest.NewRecorder()
	router.ServeHTTP(w5, req5)

	if w5.Code != http.StatusNotFound {
		t.Errorf("expected status 404 Not Found for soft deleted item, got %d", w5.Code)
	}

	// 6. Restore Soft Deleted item
	req6, _ := http.NewRequest("POST", "/api/v1/products/"+productID+"/restore", nil)
	w6 := httptest.NewRecorder()
	router.ServeHTTP(w6, req6)

	if w6.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK for product restore, got %d", w6.Code)
	}
}

func TestBulkOperations(t *testing.T) {
	router := setupTestRouter()

	bulkInput := domain.BulkCreateInput{
		Items: []domain.CreateProductInput{
			{SKU: "BULK-01", Name: "Item 1", Category: "CAT-A", Price: 10.0, Stock: 5},
			{SKU: "BULK-02", Name: "Item 2", Category: "CAT-B", Price: 20.0, Stock: 10},
		},
	}
	body, _ := json.Marshal(bulkInput)

	req1, _ := http.NewRequest("POST", "/api/v1/products/bulk", bytes.NewBuffer(body))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	if w1.Code != http.StatusCreated {
		t.Fatalf("expected status 201 Created for bulk create, got %d. Body: %s", w1.Code, w1.Body.String())
	}
}
