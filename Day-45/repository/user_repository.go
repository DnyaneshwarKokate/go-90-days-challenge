package repository

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user with this ID already exists")
)

// User represents a domain model in the system
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// UserRepository defines operations for managing users
type UserRepository interface {
	GetAll() []User
	GetByID(id string) (User, error)
	Create(user User) error
}

type memoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]User
}

// NewMemoryUserRepository creates an initialized in-memory repository with seed data
func NewMemoryUserRepository() UserRepository {
	repo := &memoryUserRepository{
		users: make(map[string]User),
	}
	// Seed sample data
	repo.users["1"] = User{ID: "1", Name: "Dnyaneshwar Kokate", Email: "dnyaneshwar@example.com", Role: "Admin"}
	repo.users["2"] = User{ID: "2", Name: "Jane Doe", Email: "jane@example.com", Role: "Developer"}
	return repo
}

func (r *memoryUserRepository) GetAll() []User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]User, 0, len(r.users))
	for _, u := range r.users {
		list = append(list, u)
	}
	return list
}

func (r *memoryUserRepository) GetByID(id string) (User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, found := r.users[id]
	if !found {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

func (r *memoryUserRepository) Create(user User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, found := r.users[user.ID]; found {
		return ErrUserExists
	}
	r.users[user.ID] = user
	return nil
}

func (r *memoryUserRepository) String() string {
	return fmt.Sprintf("MemoryUserRepository with %d users", len(r.users))
}
