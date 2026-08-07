package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User struct to store registered users
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // Stores hashed password
}

// In-memory database of users
var users = make(map[string]User)
var userIDCounter = 1

// HashPassword generates a bcrypt hash from password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares hashed password with plain password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	router := gin.Default()

	// 1. User Registration Endpoint
	router.POST("/register", func(c *gin.Context) {
		var input struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
			return
		}

		// Check if user already exists
		if _, exists := users[input.Username]; exists {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		// Hash password before saving
		hashedPassword, err := HashPassword(input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		// Save user with hashed password
		newUser := User{
			ID:       userIDCounter,
			Username: input.Username,
			Password: hashedPassword,
		}
		users[input.Username] = newUser
		userIDCounter++

		c.JSON(http.StatusCreated, gin.H{
			"message":  "User registered successfully!",
			"user_id":  newUser.ID,
			"username": newUser.Username,
		})
	})

	// 2. User Login Endpoint
	router.POST("/login", func(c *gin.Context) {
		var input struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
			return
		}

		// Find user by username
		user, exists := users[input.Username]
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// Compare stored hash with incoming password
		if !CheckPasswordHash(input.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Login successful!",
			"username": user.Username,
		})
	})

	// 3. View Users (Debugging / Demo)
	router.GET("/users", func(c *gin.Context) {
		userList := []gin.H{}
		for _, u := range users {
			userList = append(userList, gin.H{
				"id":          u.ID,
				"username":    u.Username,
				"hashed_pass": u.Password,
			})
		}
		c.JSON(http.StatusOK, userList)
	})

	fmt.Println("Password Hashing API running on http://localhost:8080")
	router.Run(":8080")
}
