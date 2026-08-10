package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"day-29/domain"

	"github.com/google/uuid"
)

type UserUseCase struct {
	repo   domain.UserRepository
	logger domain.Logger
}

func NewUserUseCase(repo domain.UserRepository, logger domain.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, input domain.CreateUserInput) (*domain.User, error) {
	u.logger.Info(ctx, "Processing user registration request", "name", input.Name, "email", input.Email)

	if strings.TrimSpace(input.Name) == "" {
		u.logger.Warn(ctx, "Validation failed: Empty user name")
		return nil, domain.ErrEmptyName
	}

	if !strings.Contains(input.Email, "@") {
		u.logger.Warn(ctx, "Validation failed: Invalid email format", "email", input.Email)
		return nil, domain.ErrInvalidEmail
	}

	existingUser, err := u.repo.FindByEmail(ctx, input.Email)
	if err == nil && existingUser != nil {
		u.logger.Warn(ctx, "User registration rejected: Email already registered", "email", input.Email)
		return nil, domain.ErrEmailAlreadyExists
	}

	role := input.Role
	if role == "" {
		role = "USER"
	}

	now := time.Now()
	user := &domain.User{
		ID:        fmt.Sprintf("usr_%s", uuid.New().String()[:8]),
		Name:      strings.TrimSpace(input.Name),
		Email:     strings.ToLower(strings.TrimSpace(input.Email)),
		Role:      strings.ToUpper(role),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.repo.Save(ctx, user); err != nil {
		u.logger.Error(ctx, "Failed to persist new user to repository", "user_id", user.ID, "error", err.Error())
		return nil, err
	}

	u.logger.Info(ctx, "User successfully registered", "user_id", user.ID, "email", user.Email, "role", user.Role)
	return user, nil
}

func (u *UserUseCase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	u.logger.Debug(ctx, "Fetching user profile by ID", "user_id", id)

	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if err == domain.ErrUserNotFound {
			u.logger.Warn(ctx, "User profile requested but not found", "user_id", id)
		} else {
			u.logger.Error(ctx, "Unexpected error retrieving user", "user_id", id, "error", err.Error())
		}
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) ListUsers(ctx context.Context) ([]*domain.User, error) {
	u.logger.Info(ctx, "Listing all active users")
	return u.repo.FindAll(ctx)
}
