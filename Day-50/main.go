package main

import (
	"log"
	"net/http"

	"day50/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", handler.StatusHandler)

	log.Println("Starting Day 50 GitHub Actions Service on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
