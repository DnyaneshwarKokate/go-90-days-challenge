package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Define role constants
const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleUser    = "user"
)

// Claims struct including custom Role field
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecretKey = []byte("super-secret-rbac-key")

// GenerateJWT creates a JWT signed token with assigned role
func GenerateJWT(username, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

// ValidateJWT verifies the token string and returns parsed Claims
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return claims, nil
}

// CheckPermission checks if a given role is allowed to perform an action
func CheckPermission(userRole string, allowedRoles []string) bool {
	for _, role := range allowedRoles {
		if userRole == role {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("--- Day 24: Role-Based Authorization (RBAC) in Go ---")

	// 1. Generate JWT tokens for different user roles
	adminToken, _ := GenerateJWT("alice_admin", RoleAdmin)
	managerToken, _ := GenerateJWT("bob_manager", RoleManager)
	userToken, _ := GenerateJWT("charlie_user", RoleUser)

	fmt.Println("Generated Admin Token:  ", adminToken[:25]+"...")
	fmt.Println("Generated Manager Token:", managerToken[:25]+"...")
	fmt.Println("Generated User Token:   ", userToken[:25]+"...\n")

	// 2. Validate tokens and check authorization for restricted resources
	tokensToTest := map[string]string{
		"Alice (Admin)":    adminToken,
		"Bob (Manager)":    managerToken,
		"Charlie (User)":   userToken,
	}

	// Define resource access policies
	adminOnlyPolicy := []string{RoleAdmin}
	managerOrAdminPolicy := []string{RoleAdmin, RoleManager}

	fmt.Println("--- Checking Policy: Admin Only Settings ---")
	for name, tokenStr := range tokensToTest {
		claims, err := ValidateJWT(tokenStr)
		if err != nil {
			fmt.Printf("%s: Token Error: %v\n", name, err)
			continue
		}

		isAllowed := CheckPermission(claims.Role, adminOnlyPolicy)
		if isAllowed {
			fmt.Printf("✅ %s (Role: %s) -> Access GRANTED to Admin Settings\n", name, claims.Role)
		} else {
			fmt.Printf("❌ %s (Role: %s) -> Access DENIED to Admin Settings\n", name, claims.Role)
		}
	}

	fmt.Println("\n--- Checking Policy: Manager & Admin Reports ---")
	for name, tokenStr := range tokensToTest {
		claims, err := ValidateJWT(tokenStr)
		if err != nil {
			fmt.Printf("%s: Token Error: %v\n", name, err)
			continue
		}

		isAllowed := CheckPermission(claims.Role, managerOrAdminPolicy)
		if isAllowed {
			fmt.Printf("✅ %s (Role: %s) -> Access GRANTED to Reports\n", name, claims.Role)
		} else {
			fmt.Printf("❌ %s (Role: %s) -> Access DENIED to Reports\n", name, claims.Role)
		}
	}
}
