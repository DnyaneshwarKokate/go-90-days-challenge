package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWT secret key
var secretKey = []byte("rbac-super-secret-key-2026")

// Role constants
const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleUser    = "user"
)

// Claims structure for JWT
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// Mock User DB
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var mockUsers = map[string]User{
	"admin_user":   {Username: "admin_user", Password: "admin123", Role: RoleAdmin},
	"manager_user": {Username: "manager_user", Password: "manager123", Role: RoleManager},
	"regular_user": {Username: "regular_user", Password: "user123", Role: RoleUser},
}

// Generate JWT token containing role
func GenerateToken(username, role string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// AuthMiddleware authenticates JWT token and stores claims in context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenStr := parts[1]
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Store username and role in Gin context
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}

// AuthorizeRole checks if user role matches one of allowed roles
func AuthorizeRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "User role not found in context"})
			c.Abort()
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid role type"})
			c.Abort()
			return
		}

		// Check if user's role is permitted
		for _, allowedRole := range allowedRoles {
			if roleStr == allowedRole {
				c.Next()
				return
			}
		}

		// Role not permitted
		c.JSON(http.StatusForbidden, gin.H{
			"error": fmt.Sprintf("Access denied: role '%s' does not have permission to access this resource", roleStr),
		})
		c.Abort()
	}
}

func main() {
	router := gin.Default()

	// 1. Login Endpoint - Authenticates user and returns JWT with role
	router.POST("/login", func(c *gin.Context) {
		var input struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
			return
		}

		user, exists := mockUsers[input.Username]
		if !exists || user.Password != input.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		token, err := GenerateToken(user.Username, user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful!",
			"role":    user.Role,
			"token":   token,
		})
	})

	// 2. Public Endpoint
	router.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Public API! No authentication required."})
	})

	// 3. User Protected Routes
	userRoutes := router.Group("/user")
	userRoutes.Use(AuthMiddleware(), AuthorizeRole(RoleUser, RoleManager, RoleAdmin))
	{
		userRoutes.GET("/profile", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			role := c.MustGet("role").(string)
			c.JSON(http.StatusOK, gin.H{
				"message":  "Welcome to User Profile",
				"username": username,
				"role":     role,
			})
		})
	}

	// 4. Manager Protected Routes
	managerRoutes := router.Group("/manager")
	managerRoutes.Use(AuthMiddleware(), AuthorizeRole(RoleManager, RoleAdmin))
	{
		managerRoutes.GET("/reports", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			role := c.MustGet("role").(string)
			c.JSON(http.StatusOK, gin.H{
				"message":  "Manager Reports & Analytics Access Granted",
				"username": username,
				"role":     role,
			})
		})
	}

	// 5. Admin Protected Routes
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(AuthMiddleware(), AuthorizeRole(RoleAdmin))
	{
		adminRoutes.GET("/settings", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			role := c.MustGet("role").(string)
			c.JSON(http.StatusOK, gin.H{
				"message":  "Admin System Settings Access Granted",
				"username": username,
				"role":     role,
			})
		})
	}

	fmt.Println("Role-Based Authorization API running on http://localhost:8080")
	router.Run(":8080")
}
