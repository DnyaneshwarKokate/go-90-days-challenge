package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-45/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-45/handler"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-45/repository"
)

func main() {
	cfg := config.LoadConfig()
	userRepo := repository.NewMemoryUserRepository()
	apiHandler := handler.NewAPIHandler(cfg, userRepo)

	mux := http.NewServeMux()

	// Probes & Meta
	mux.HandleFunc("/healthz", apiHandler.Healthz)
	mux.HandleFunc("/ready", apiHandler.Ready)
	mux.HandleFunc("/api/v1/info", apiHandler.Info)

	// Domain Endpoints
	mux.HandleFunc("/api/v1/users", apiHandler.HandleUsers)
	mux.HandleFunc("/api/v1/users/", apiHandler.HandleUsers)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Server run context for graceful shutdown
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for OS signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		log.Println("Received termination signal. Shutting down container process gracefully...")

		// Shutdown signal with a 5-second deadline
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
		serverStopCtx()
	}()

	log.Printf("🚀 Starting Go REST API container (Env: %s, Version: %s) on port %s...", cfg.AppEnv, cfg.AppVersion, cfg.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server failure: %v", err)
	}

	// Wait for server context to finish
	<-serverCtx.Done()
	log.Println("Server gracefully stopped. Container exit success.")
}
