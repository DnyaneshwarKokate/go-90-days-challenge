package main

import (
	"encoding/json"
	"log"
	"net/http"

	"day55/service"
)

func main() {
	orderSvc := service.NewOrderService("http://localhost:8081", "http://localhost:8082")

	http.HandleFunc("/api/v1/orders/aggregate", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		productID := r.URL.Query().Get("product_id")

		order, err := orderSvc.GetAggregatedOrder(r.Context(), userID, productID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(order)
	})

	log.Println("Starting Day 55 Inter-Service Aggregator Service on :8083...")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatalf("Server crash: %v", err)
	}
}
