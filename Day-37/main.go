package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"day-37/cache"
	"day-37/domain"
	"day-37/handler"
	"day-37/logger"
	"day-37/middleware"
	"day-37/repository"
	"day-37/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("==========================================================================")
	fmt.Println("🚀 Day 37: Redis Caching Pattern in Go (Cache-Aside, Hit vs Miss, Invalidation)")
	fmt.Println("==========================================================================")

	logFilePath := "cache_app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()
	zapLog.Info(ctx, "Redis Cache Application initializing", "target_host", "localhost:6379")

	productRepo := repository.NewMemoryProductRepository(zapLog)
	productCache := cache.NewRedisProductCache("localhost:6379", zapLog)
	productUC := usecase.NewProductUseCase(productRepo, productCache, 5*time.Minute, zapLog)
	productHandler := handler.NewProductHandler(productUC, zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	v1 := router.Group("/api/v1/products")
	{
		v1.POST("", productHandler.CreateProduct)
		v1.GET("/:id", productHandler.GetProductByID)
		v1.PUT("/:id", productHandler.UpdateProduct)
	}

	fmt.Println("\n--- 1️⃣ Creating a New Product in DB ---")
	productID := createProduct(router, "MacBook Pro M3 Max", 3499.99, "LAPTOPS")

	fmt.Println("\n--- 2️⃣ Request 1: GET Product (Expect CACHE MISS -> DB Query -> Populate Cache) ---")
	getProduct(router, productID)

	fmt.Println("\n--- 3️⃣ Request 2: GET Product (Expect ⚡ CACHE HIT -> Instant Redis Response) ---")
	getProduct(router, productID)

	fmt.Println("\n--- 4️⃣ Request 3: UPDATE Product (Expect 🔥 CACHE INVALIDATION -> Delete Redis Key) ---")
	newPrice := 3299.99
	updateProduct(router, productID, &newPrice)

	fmt.Println("\n--- 5️⃣ Request 4: GET Product (Expect CACHE MISS -> Fresh DB Query -> Re-populate Cache) ---")
	getProduct(router, productID)

	fmt.Println("\n✅ Day 37 Redis Caching executed successfully! Check cache_app.log for logs.")
}

func createProduct(router *gin.Engine, name string, price float64, cat string) string {
	payload := domain.CreateProductInput{Name: name, Price: price, Category: cat}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /products -> Status %d | Body: %s\n", w.Code, w.Body.String())

	var res map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &res)
	if data, ok := res["data"].(map[string]interface{}); ok {
		return data["id"].(string)
	}
	return ""
}

func getProduct(router *gin.Engine, id string) {
	req, _ := http.NewRequest("GET", "/api/v1/products/"+id, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("GET /products/%s -> Status %d | Body: %s\n", id, w.Code, w.Body.String())
}

func updateProduct(router *gin.Engine, id string, price *float64) {
	payload := domain.UpdateProductInput{Price: price}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "/api/v1/products/"+id, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("PUT /products/%s -> Status %d | Body: %s\n", id, w.Code, w.Body.String())
}
