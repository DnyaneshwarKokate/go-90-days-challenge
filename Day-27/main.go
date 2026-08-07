package main

import (
	"fmt"
	"log"

	"day-27/domain"
	"day-27/handler"
	"day-27/repository"
	"day-27/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 Day 27: Repository Pattern in Go")

	// 1. Initialize Storage Engine (Swap between Memory or GORM easily)
	// Example using In-Memory Repository implementation:
	var productRepo domain.ProductRepository = repository.NewMemoryProductRepository()

	// Alternative GORM SQLite implementation:
	// db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	// if err != nil { log.Fatal(err) }
	// db.AutoMigrate(&domain.Product{})
	// var productRepo domain.ProductRepository = repository.NewGORMProductRepository(db)

	// 2. Inject Repository into Business Use Case
	productUseCase := usecase.NewProductUseCase(productRepo)

	// 3. Inject Use Case into HTTP Handler
	productHandler := handler.NewProductHandler(productUseCase)

	// 4. Register HTTP Router & Endpoints
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products", productHandler.ListProducts)
		api.GET("/products/:id", productHandler.GetProductByID)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	log.Println("⚡ Server listening on http://localhost:8087")
	if err := router.Run(":8087"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
