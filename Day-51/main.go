package main

import (
	"log"
	"net/http"

	"day51/config"
	"day51/handler"
	"day51/middleware"
)

func main() {
	cfg := config.LoadConfig()
	healthHandler := handler.NewHealthHandler(cfg)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler.HealthCheck)

	handlerStack := middleware.CorrelationIDMiddleware(mux)

	log.Printf("Starting Microservice [%s] on port :%s (Env: %s)...\n", cfg.ServiceName, cfg.Port, cfg.Environment)
	if err := http.ListenAndServe(":"+cfg.Port, handlerStack); err != nil {
		log.Fatalf("Microservice terminated unexpectedly: %v", err)
	}
}
