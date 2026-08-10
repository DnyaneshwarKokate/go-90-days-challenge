package repository

import (
	"context"
	"sync"
	"time"

	"day-29/domain"
)

type memoryUserRepository struct {
	mu     sync.RWMutex
	users  map[string]*domain.User
	emails map[string]string
	logger domain.Logger
}

func NewMemoryUserRepository(logger domain.Logger) domain.UserRepository {
	return &memoryUserRepository{
		users:  make(map[string]*domain.User),
		emails: make(map[string]string),
		logger: logger,
	}
}

func (r *memoryUserRepository) Save(ctx context.Context, user *domain.User) error {
	start := time.Now()
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug(ctx, "Executing DB query: Save User", "user_id", user.ID, "email", user.Email)

	if existingID, exists := r.emails[user.Email]; exists && existingID != user.ID {
		r.logger.Warn(ctx, "DB Conflict: Email already exists", "email", user.Email, "existing_user_id", existingID)
		return domain.ErrEmailAlreadyExists
	}

	r.users[user.ID] = user
	r.emails[user.Email] = user.ID

	r.logger.Info(ctx, "DB Write successful", "user_id", user.ID, "duration_ms", time.Since(start).Milliseconds())
	return nil
}

func (r *memoryUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	start := time.Now()
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "Executing DB query: FindByID", "user_id", id)

	user, exists := r.users[id]
	if !exists {
		r.logger.Warn(ctx, "DB Record not found", "user_id", id, "duration_ms", time.Since(start).Milliseconds())
		return nil, domain.ErrUserNotFound
	}

	r.logger.Debug(ctx, "DB Record retrieved", "user_id", id, "duration_ms", time.Since(start).Milliseconds())
	return user, nil
}

func (r *memoryUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "Executing DB query: FindByEmail", "email", email)

	id, exists := r.emails[email]
	if !exists {
		return nil, domain.ErrUserNotFound
	}

	return r.users[id], nil
}

func (r *memoryUserRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	start := time.Now()
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "Executing DB query: FindAll")

	result := make([]*domain.User, 0, len(r.users))
	for _, user := range r.users {
		result = append(result, user)
	}

	r.logger.Info(ctx, "DB Records fetched", "count", len(result), "duration_ms", time.Since(start).Milliseconds())
	return result, nil
}
