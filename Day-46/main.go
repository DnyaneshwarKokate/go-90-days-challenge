package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-46/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-46/db"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-46/handler"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-46/repository"
)

func main() {
	log.Println("Starting Day-46 Docker Compose Go API Microservice...")

	cfg := config.LoadConfig()
	log.Printf("Loaded configuration for environment [%s] on port :%s", cfg.AppEnv, cfg.Port)

	// Initialize PostgreSQL Database Connection with retries
	postgresConn, err := db.InitPostgres(cfg)
	if err != nil {
		log.Fatalf("Fatal: Unable to connect to PostgreSQL: %v", err)
	}
	defer postgresConn.Close()

	// Initialize Redis Cache Client with retries
	redisClient, err := db.InitRedis(cfg)
	if err != nil {
		log.Printf("Warning: Redis initialization failed: %v. Continuing without cache layer.", err)
	} else {
		defer redisClient.Close()
	}

	repo := repository.NewUserRepository(postgresConn, redisClient)
	apiHandler := handler.NewAPIHandler(cfg, repo, postgresConn, redisClient)

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", apiHandler.Healthz)
	mux.HandleFunc("/ready", apiHandler.Ready)
	mux.HandleFunc("/api/v1/users", apiHandler.HandleUsers)
	mux.HandleFunc("/api/v1/users/", apiHandler.HandleUserByID)
	mux.HandleFunc("/api/v1/stats", apiHandler.Stats)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Intercept termination signals for graceful container shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on port :%s (Container IP)", cfg.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	sig := <-sigChan
	log.Printf("Received termination signal [%v]. Initiating graceful shutdown...", sig)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Forced server shutdown error: %v", err)
	}

	log.Println("Go API Microservice stopped gracefully.")
}
