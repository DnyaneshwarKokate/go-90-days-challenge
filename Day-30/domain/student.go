package domain

import (
	"context"
	"errors"
	"time"
)

type ContextKey string

const (
	RequestIDKey ContextKey = "X-Request-ID"
	UserIDKey    ContextKey = "user_id"
	UserRoleKey  ContextKey = "user_role"
	UserEmailKey ContextKey = "user_email"
)

const (
	RoleAdmin   = "ADMIN"
	RoleTeacher = "TEACHER"
	RoleStudent = "STUDENT"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrStudentNotFound    = errors.New("student record not found")
	ErrEmailExists        = errors.New("email address is already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUnauthorized       = errors.New("authentication token is missing or invalid")
	ErrForbidden          = errors.New("access forbidden: insufficient permissions")
	ErrInvalidInput       = errors.New("invalid request payload")
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Student struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	GPA        float64   `json:"gpa"`
	Status     string    `json:"status"` // ACTIVE, INACTIVE, GRADUATED
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	User      *User     `json:"user"`
}

type CreateStudentInput struct {
	FullName   string  `json:"full_name" binding:"required"`
	Email      string  `json:"email" binding:"required,email"`
	Department string  `json:"department" binding:"required"`
	GPA        float64 `json:"gpa" binding:"min=0,max=4"`
}

type UpdateStudentInput struct {
	FullName   string  `json:"full_name"`
	Department string  `json:"department"`
	GPA        float64 `json:"gpa" binding:"min=0,max=4"`
	Status     string  `json:"status"`
}

type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}

type UserRepository interface {
	SaveUser(ctx context.Context, user *User) error
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	FindUserByID(ctx context.Context, id string) (*User, error)
}

type StudentRepository interface {
	Save(ctx context.Context, student *Student) error
	FindByID(ctx context.Context, id string) (*Student, error)
	FindByEmail(ctx context.Context, email string) (*Student, error)
	FindAll(ctx context.Context, department string, status string) ([]*Student, error)
	Update(ctx context.Context, student *Student) error
	Delete(ctx context.Context, id string) error
}
