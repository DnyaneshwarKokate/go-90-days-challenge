package main

import (
	"context"
	"log"

	"day56/service"
)

func main() {
	svc := service.NewUserService()
	log.Println("Starting Day 56 gRPC Server Demonstration...")

	res, err := svc.GetUser(context.Background(), &service.GetUserRequest{UserID: "usr_101"})
	if err != nil {
		log.Fatalf("gRPC Unary Call Failed: %v", err)
	}

	log.Printf("gRPC Unary Response: ID=%s, Name=%s, Role=%s", res.ID, res.Name, res.Role)

	log.Println("Streaming Users:")
	_ = svc.StreamUsers(3, func(u *service.UserResponse) error {
		log.Printf(" -> Streamed User: %s (%s)", u.Name, u.Email)
		return nil
	})
}
