package main

import (
	"log"
	"net/http"

	"day49/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/info", handler.InfoHandler)

	log.Println("Starting Day 49 CI/CD Baseline Service on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
