package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 1. Custom Logger Middleware
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		fmt.Printf("[LOG] Incoming Request: %s %s\n", c.Request.Method, c.Request.URL.Path)

		// Process request by passing to the next handler
		c.Next()

		// Code executed after request handler finishes
		latency := time.Since(startTime)
		status := c.Writer.Status()

		fmt.Printf("[LOG] Completed Request: %s %s | Status: %d | Time: %v\n",
			c.Request.Method, c.Request.URL.Path, status, latency)
	}
}

// 2. Custom Header / Key Injection Middleware
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-ID", "REQ-90DAYS-12345")
		c.Set("requestID", "REQ-90DAYS-12345")
		c.Next()
	}
}

// 3. Simple Authentication Middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != "Bearer secrettoken123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized Access - Invalid Token",
			})
			c.Abort() // Stop subsequent handlers from executing
			return
		}

		// Store user identity in context
		c.Set("user", "Dnyaneshwar")
		c.Next()
	}
}

func main() {
	// Create router without default middlewares
	router := gin.New()

	// Apply global middlewares
	router.Use(gin.Recovery())
	router.Use(LoggerMiddleware())
	router.Use(RequestIDMiddleware())

	// Public Route
	router.GET("/public", func(c *gin.Context) {
		reqID, _ := c.Get("requestID")
		c.JSON(http.StatusOK, gin.H{
			"message":    "Public endpoint accessible by everyone!",
			"request_id": reqID,
		})
	})

	// Protected Route Group (Requires AuthMiddleware)
	protected := router.Group("/api/v1")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/dashboard", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Welcome to Dashboard, %s!", user),
			})
		})

		protected.GET("/profile", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(http.StatusOK, gin.H{
				"user":   user,
				"role":   "Software Engineer",
				"status": "Active",
			})
		})
	}

	fmt.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
