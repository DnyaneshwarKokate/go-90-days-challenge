package domain

import "errors"

// User entity model
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// Custom domain errors
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user with this email already exists")
	ErrInvalidUserData   = errors.New("name, email, and role are required")
)

// UserRepository interface defines data persistence operations
type UserRepository interface {
	Create(user *User) (*User, error)
	GetByID(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]User, error)
}

// UserUseCase interface defines core business logic operations
type UserUseCase interface {
	RegisterUser(name, email, role string) (*User, error)
	GetUserByID(id int) (*User, error)
	ListUsers() ([]User, error)
}
