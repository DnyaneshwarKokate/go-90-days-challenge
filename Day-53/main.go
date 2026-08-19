package main

import (
	"log"
	"net/http"

	"day53/handler"
	"day53/repository"
	"day53/service"
)

func main() {
	repo := repository.NewInMemoryProductRepository()
	svc := service.NewProductService(repo)
	h := handler.NewProductHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/products", h.HandleProducts)
	mux.HandleFunc("/api/v1/products/", h.HandleProductByID)

	log.Println("Starting Day 53 Product Microservice on :8082...")
	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatalf("Product service failure: %v", err)
	}
}
