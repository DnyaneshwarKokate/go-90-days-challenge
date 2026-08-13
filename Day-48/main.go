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

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-48/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-48/handler"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-48/repository"
)

func main() {
	log.Println("Initializing Day-48 Production Kubernetes Go Microservice...")

	cfg := config.LoadConfig()
	log.Printf("Pod [%s] in Namespace [%s] on Node [%s] (Pod IP: %s, Env: %s)",
		cfg.PodName, cfg.PodNamespace, cfg.NodeName, cfg.PodIP, cfg.AppEnv)

	store, err := repository.NewStore(cfg.DataPath)
	if err != nil {
		log.Fatalf("Failed to initialize persistent volume store at [%s]: %v", cfg.DataPath, err)
	}

	apiHandler := handler.NewAPIHandler(cfg, store)

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", apiHandler.Healthz)
	mux.HandleFunc("/ready", apiHandler.Ready)
	mux.HandleFunc("/api/v1/info", apiHandler.Info)
	mux.HandleFunc("/api/v1/orders", apiHandler.Orders)
	mux.HandleFunc("/api/v1/metrics", apiHandler.Metrics)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on port :%s inside container...", cfg.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server failure: %v", err)
		}
	}()

	sig := <-sigChan
	log.Printf("Received signal [%v] from Kubelet/OS. Starting zero-downtime termination sequence...", sig)

	// Step 1: Set readiness probe to unready (503 Service Unavailable)
	apiHandler.SetNotReady()
	log.Println("Readiness probe set to UNREADY. Sleeping 5s for Kube-Proxy endpoint propagation...")
	time.Sleep(5 * time.Second)

	// Step 2: Shutdown HTTP server gracefully
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Forced server shutdown error: %v", err)
	}

	log.Println("Kubernetes Pod process exited cleanly.")
}
