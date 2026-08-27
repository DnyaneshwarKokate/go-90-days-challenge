package main

import (
	"log"
	"net/http"

	"day60/handler"
)

func main() {
	h := handler.NewOrderHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", h.HealthCheck)
	mux.HandleFunc("/api/v1/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreateOrder(w, r)
			return
		}
		if r.Method == http.MethodGet {
			h.GetOrder(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	log.Println("Starting Day 60 Production API Service on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Production API failure: %v", err)
	}
}
