package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"day-31/domain"
	"day-31/handler"
	"day-31/logger"
	"day-31/middleware"
	"day-31/repository"
	"day-31/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("==========================================================================")
	fmt.Println("🚀 Day 31: Advanced CRUD APIs in Go (Bulk, PATCH, Soft Delete, OCC)")
	fmt.Println("==========================================================================")

	logFilePath := "product_app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()
	zapLog.Info(ctx, "Advanced CRUD API Server starting", "log_file", logFilePath)

	productRepo := repository.NewMemoryProductRepository(zapLog)
	productUC := usecase.NewProductUseCase(productRepo, zapLog)
	productHandler := handler.NewProductHandler(productUC, zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
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

	fmt.Println("\n--- 1️⃣ Demonstrating Single & Bulk Product Creations ---")
	p1ID := createSingleProduct(router, "LAPTOP-01", "MacBook Pro M3", "ELECTRONICS", 1999.99, 15)
	p2IDs := createBulkProducts(router, []domain.CreateProductInput{
		{SKU: "PHONE-01", Name: "iPhone 15 Pro", Category: "ELECTRONICS", Price: 999.99, Stock: 30},
		{SKU: "AUDIO-01", Name: "AirPods Pro 2", Category: "AUDIO", Price: 249.99, Stock: 50},
	})

	fmt.Println("\n--- 2️⃣ Demonstrating Partial Update (PATCH) & Version Increment ---")
	newPrice := 1899.99
	newStock := 12
	expectedVer := 1
	patchProduct(router, p1ID, &newPrice, &newStock, nil, &expectedVer) // Version becomes 2

	fmt.Println("\n--- 3️⃣ Demonstrating Optimistic Concurrency Conflict (Outdated Version) ---")
	staleVersion := 1
	patchProduct(router, p1ID, &newPrice, &newStock, nil, &staleVersion) // Expected: 409 Conflict

	fmt.Println("\n--- 4️⃣ Demonstrating Soft Deletion ---")
	softDeleteProduct(router, p1ID)

	fmt.Println("\n--- 5️⃣ Demonstrating Product Listing (Standard vs ?include_deleted=true) ---")
	listProducts(router, false)
	listProducts(router, true)

	fmt.Println("\n--- 6️⃣ Demonstrating Product Restoration ---")
	restoreProduct(router, p1ID)

	fmt.Println("\n--- 7️⃣ Demonstrating Bulk Soft Deletion ---")
	if len(p2IDs) > 0 {
		bulkSoftDelete(router, p2IDs)
	}

	fmt.Println("\n✅ Day 31 Advanced CRUD APIs executed successfully! Check product_app.log for logs.")
}

func createSingleProduct(router *gin.Engine, sku, name, category string, price float64, stock int) string {
	payload := domain.CreateProductInput{SKU: sku, Name: name, Category: category, Price: price, Stock: stock}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /products -> Status %d | Body: %s\n", w.Code, w.Body.String())

	var res map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &res)
	if data, ok := res["data"].(map[string]interface{}); ok {
		id, _ := data["id"].(string)
		return id
	}
	return ""
}

func createBulkProducts(router *gin.Engine, items []domain.CreateProductInput) []string {
	payload := domain.BulkCreateInput{Items: items}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/products/bulk", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /products/bulk -> Status %d | Body: %s\n", w.Code, w.Body.String())

	ids := make([]string, 0)
	var res map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &res)
	if dataList, ok := res["data"].([]interface{}); ok {
		for _, item := range dataList {
			if m, ok := item.(map[string]interface{}); ok {
				if id, ok := m["id"].(string); ok {
					ids = append(ids, id)
				}
			}
		}
	}
	return ids
}

func patchProduct(router *gin.Engine, id string, price *float64, stock *int, name *string, expectedVersion *int) {
	payload := domain.PatchProductInput{
		Name:            name,
		Price:           price,
		Stock:           stock,
		ExpectedVersion: expectedVersion,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PATCH", "/api/v1/products/"+id, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if expectedVersion != nil {
		req.Header.Set("If-Match", fmt.Sprintf("%d", *expectedVersion))
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("PATCH /products/%s -> Status %d | Body: %s\n", id, w.Code, w.Body.String())
}

func softDeleteProduct(router *gin.Engine, id string) {
	req, _ := http.NewRequest("DELETE", "/api/v1/products/"+id, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("DELETE /products/%s -> Status %d | Body: %s\n", id, w.Code, w.Body.String())
}

func listProducts(router *gin.Engine, includeDeleted bool) {
	url := "/api/v1/products"
	if includeDeleted {
		url += "?include_deleted=true"
	}
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("GET %s -> Status %d | Body: %s\n", url, w.Code, w.Body.String())
}

func restoreProduct(router *gin.Engine, id string) {
	req, _ := http.NewRequest("POST", "/api/v1/products/"+id+"/restore", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /products/%s/restore -> Status %d | Body: %s\n", id, w.Code, w.Body.String())
}

func bulkSoftDelete(router *gin.Engine, ids []string) {
	payload := domain.BulkDeleteInput{IDs: ids}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/products/bulk-delete", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /products/bulk-delete -> Status %d | Body: %s\n", w.Code, w.Body.String())
}
