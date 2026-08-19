package main

import (
	"log"
	"net/http"
	"time"

	"day54/middleware"
	"day54/proxy"
)

func main() {
	userServiceURL := "http://localhost:8081"
	productServiceURL := "http://localhost:8082"
	jwtSecret := "day52-jwt-secret-key-change-in-prod"

	gatewayProxy, err := proxy.NewGatewayProxy(userServiceURL, productServiceURL)
	if err != nil {
		log.Fatalf("Failed to initialize API gateway proxy: %v", err)
	}

	authMw := middleware.NewAuthMiddleware(jwtSecret)
	rateLimiter := middleware.NewRateLimiter(100, time.Minute)

	handlerStack := rateLimiter.Limit(authMw.ValidateJWT(gatewayProxy))

	log.Println("Starting Day 54 API Gateway on :8000...")
	log.Printf("Routing /api/v1/users -> %s", userServiceURL)
	log.Printf("Routing /api/v1/products -> %s", productServiceURL)

	if err := http.ListenAndServe(":8000", handlerStack); err != nil {
		log.Fatalf("API Gateway crashed: %v", err)
	}
}
