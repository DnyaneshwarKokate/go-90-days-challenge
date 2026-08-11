package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-44/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-44/handler"
)

func main() {
	cfg := config.LoadConfig()

	itemSvc := handler.NewItemService(cfg)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HandleHealth)
	mux.HandleFunc("/api/v1/info", itemSvc.HandleInfo)
	mux.HandleFunc("/api/v1/items", itemSvc.HandleItems)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// Server shutdown channel
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("🚀 Starting [%s v%s] on port :%s (Environment: %s)",
			cfg.ServiceName, cfg.Version, cfg.Port, cfg.AppEnv)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("❌ Server failed to start: %v", err)
		}
	}()

	sig := <-shutdown
	log.Printf("⚠️ Received shutdown signal (%s). Initiating graceful shutdown...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}

	fmt.Println("✅ Server exited cleanly.")
}
