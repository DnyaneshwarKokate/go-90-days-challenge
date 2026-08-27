package service

import (
	"context"
	"fmt"
)

type UserResponse struct {
	ID    string
	Name  string
	Email string
	Role  string
}

type GetUserRequest struct {
	UserID string
}

type UserServiceServer interface {
	GetUser(ctx context.Context, req *GetUserRequest) (*UserResponse, error)
	StreamUsers(count int, sendFunc func(*UserResponse) error) error
}

type UserService struct {
	users map[string]*UserResponse
}

func NewUserService() *UserService {
	return &UserService{
		users: map[string]*UserResponse{
			"usr_101": {ID: "usr_101", Name: "Alice Smith", Email: "alice@example.com", Role: "admin"},
			"usr_102": {ID: "usr_102", Name: "Bob Jones", Email: "bob@example.com", Role: "user"},
			"usr_103": {ID: "usr_103", Name: "Charlie Brown", Email: "charlie@example.com", Role: "user"},
		},
	}
}

func (s *UserService) GetUser(ctx context.Context, req *GetUserRequest) (*UserResponse, error) {
	if req == nil || req.UserID == "" {
		return nil, fmt.Errorf("user_id cannot be empty")
	}

	user, ok := s.users[req.UserID]
	if !ok {
		return nil, fmt.Errorf("user %s not found", req.UserID)
	}

	return user, nil
}

func (s *UserService) StreamUsers(count int, sendFunc func(*UserResponse) error) error {
	idx := 0
	for _, u := range s.users {
		if idx >= count {
			break
		}
		if err := sendFunc(u); err != nil {
			return err
		}
		idx++
	}
	return nil
}
