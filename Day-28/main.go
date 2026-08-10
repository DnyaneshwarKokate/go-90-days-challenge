package main

import (
	"fmt"
	"log"

	"day-28/di"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 Day 28: Dependency Injection in Go")

	// 1. Initialize central DI Container
	cfg := di.Config{Env: "development"}
	container := di.NewContainer(cfg)

	// 2. Initialize Gin web engine
	router := gin.Default()

	// 3. Register HTTP endpoints from DI container
	container.SetupRoutes(router)

	// 4. Start HTTP Server
	port := ":8088"
	log.Printf("⚡ Server running on http://localhost%s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Server launch failed: %v", err)
	}
}
