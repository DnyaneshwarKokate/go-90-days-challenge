package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"day-30/config"
	"day-30/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo domain.UserRepository
	config   config.Config
	logger   domain.Logger
}

func NewAuthUseCase(userRepo domain.UserRepository, cfg config.Config, logger domain.Logger) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
		config:   cfg,
		logger:   logger,
	}
}

func (u *AuthUseCase) Register(ctx context.Context, input domain.RegisterInput) (*domain.User, error) {
	u.logger.Info(ctx, "Processing user registration", "email", input.Email, "role", input.Role)

	email := strings.ToLower(strings.TrimSpace(input.Email))
	if email == "" || !strings.Contains(email, "@") {
		u.logger.Warn(ctx, "Validation failed: Invalid email format", "email", input.Email)
		return nil, domain.ErrInvalidInput
	}

	if len(input.Password) < 6 {
		u.logger.Warn(ctx, "Validation failed: Password too short")
		return nil, domain.ErrInvalidInput
	}

	existingUser, err := u.userRepo.FindUserByEmail(ctx, email)
	if err == nil && existingUser != nil {
		u.logger.Warn(ctx, "Registration rejected: Email already registered", "email", email)
		return nil, domain.ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error(ctx, "Failed to hash password", "error", err)
		return nil, fmt.Errorf("password hashing failed: %w", err)
	}

	role := strings.ToUpper(strings.TrimSpace(input.Role))
	if role != domain.RoleAdmin && role != domain.RoleTeacher && role != domain.RoleStudent {
		role = domain.RoleStudent
	}

	user := &domain.User{
		ID:        fmt.Sprintf("usr_%s", uuid.New().String()[:8]),
		Email:     email,
		Password:  string(hashedPassword),
		Role:      role,
		CreatedAt: time.Now(),
	}

	if err := u.userRepo.SaveUser(ctx, user); err != nil {
		u.logger.Error(ctx, "Failed to save user in repository", "user_id", user.ID, "error", err)
		return nil, err
	}

	u.logger.Info(ctx, "User account created successfully", "user_id", user.ID, "email", user.Email, "role", user.Role)
	return user, nil
}

func (u *AuthUseCase) Login(ctx context.Context, input domain.LoginInput) (*domain.TokenResponse, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	u.logger.Info(ctx, "Processing login authentication", "email", email)

	user, err := u.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		u.logger.Warn(ctx, "Login failed: User not found", "email", email)
		return nil, domain.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		u.logger.Warn(ctx, "Login failed: Invalid password", "user_id", user.ID)
		return nil, domain.ErrInvalidCredentials
	}

	expiresAt := time.Now().Add(u.config.JWTExpiresIn)
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   expiresAt.Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.config.JWTSecret))
	if err != nil {
		u.logger.Error(ctx, "Failed to sign JWT token", "user_id", user.ID, "error", err)
		return nil, fmt.Errorf("token signing failed: %w", err)
	}

	u.logger.Info(ctx, "User authenticated & JWT token issued", "user_id", user.ID, "role", user.Role)
	return &domain.TokenResponse{
		Token:     tokenString,
		ExpiresAt: expiresAt,
		User:      user,
	}, nil
}
