package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain text password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a plain text password with a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	fmt.Println("--- Day 23: Password Hashing in Go ---")

	password := "SuperSecretPassword123"

	// 1. Generate Hash
	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	fmt.Printf("Original Password: %s\n", password)
	fmt.Printf("Hashed Password:   %s\n\n", hashedPassword)

	// 2. Verify Correct Password
	isValid := CheckPasswordHash("SuperSecretPassword123", hashedPassword)
	fmt.Printf("Testing with correct password ('SuperSecretPassword123'): Valid = %t\n", isValid)

	// 3. Verify Wrong Password
	isWrongValid := CheckPasswordHash("WrongPassword456", hashedPassword)
	fmt.Printf("Testing with wrong password ('WrongPassword456'):   Valid = %t\n", isWrongValid)
}
