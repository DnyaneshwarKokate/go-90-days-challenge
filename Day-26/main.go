package main

import (
	"fmt"
	"log"

	"day-26/handler"
	"day-26/repository"
	"day-26/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("--- Day 26: Clean Architecture in Go ---")

	// 1. Initialize Repository (Data Access Layer)
	userRepo := repository.NewInMemoryUserRepository()

	// 2. Initialize UseCase (Business Logic Layer) with Repository dependency
	userUseCase := usecase.NewUserUseCase(userRepo)

	// 3. Initialize Handler (Delivery Layer) with UseCase dependency
	userHandler := handler.NewUserHandler(userUseCase)

	// 4. Setup Gin Router
	router := gin.Default()

	// 5. Register Routes
	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.RegisterUser)
		api.GET("/users", userHandler.ListUsers)
		api.GET("/users/:id", userHandler.GetUserByID)
	}

	fmt.Println("🚀 Clean Architecture User API running on http://localhost:8086")
	log.Fatal(router.Run(":8086"))
}
