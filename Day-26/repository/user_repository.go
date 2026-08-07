package repository

import (
	"sync"

	"day-26/domain"
)

type inMemoryUserRepository struct {
	mu     sync.RWMutex
	users  map[int]domain.User
	nextID int
}

// NewInMemoryUserRepository creates a new repository instance
func NewInMemoryUserRepository() domain.UserRepository {
	repo := &inMemoryUserRepository{
		users:  make(map[int]domain.User),
		nextID: 1,
	}

	// Seed initial data
	repo.users[1] = domain.User{ID: 1, Name: "Dnyaneshwar Kokate", Email: "dnyaneshwar@example.com", Role: "Admin"}
	repo.users[2] = domain.User{ID: 2, Name: "Rahul Sharma", Email: "rahul@example.com", Role: "Developer"}
	repo.nextID = 3

	return repo
}

func (r *inMemoryUserRepository) Create(user *domain.User) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.users[user.ID] = *user
	r.nextID++

	return user, nil
}

func (r *inMemoryUserRepository) GetByID(id int) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, domain.ErrUserNotFound
	}
	return &user, nil
}

func (r *inMemoryUserRepository) GetByEmail(email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *inMemoryUserRepository) GetAll() ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	userList := make([]domain.User, 0, len(r.users))
	for _, user := range r.users {
		userList = append(userList, user)
	}
	return userList, nil
}
