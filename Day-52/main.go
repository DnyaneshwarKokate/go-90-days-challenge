package main

import (
	"log"
	"net/http"

	"day52/handler"
	"day52/repository"
	"day52/service"
)

func main() {
	repo := repository.NewInMemoryUserRepository()
	svc := service.NewUserService(repo, "day52-jwt-secret-key-change-in-prod")
	userHandler := handler.NewUserHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/users/register", userHandler.Register)
	mux.HandleFunc("/api/v1/users/login", userHandler.Login)
	mux.HandleFunc("/api/v1/users/profile", userHandler.Profile)

	log.Println("Starting Day 52 User Microservice on :8081...")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("User service crashed: %v", err)
	}
}
