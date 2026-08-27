package service

import (
	"context"
	"testing"
)

func TestGRPCServerUnaryAndStreaming(t *testing.T) {
	svc := NewUserService()

	// Test 1: Unary RPC GetUser
	req := &GetUserRequest{UserID: "usr_101"}
	res, err := svc.GetUser(context.Background(), req)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	if res.Name != "Alice Smith" {
		t.Errorf("Expected name 'Alice Smith', got %s", res.Name)
	}

	// Test 2: Server Streaming RPC StreamUsers
	var streamedUsers []*UserResponse
	err = svc.StreamUsers(2, func(u *UserResponse) error {
		streamedUsers = append(streamedUsers, u)
		return nil
	})

	if err != nil {
		t.Fatalf("StreamUsers failed: %v", err)
	}

	if len(streamedUsers) != 2 {
		t.Errorf("Expected 2 streamed users, got %d", len(streamedUsers))
	}
}
