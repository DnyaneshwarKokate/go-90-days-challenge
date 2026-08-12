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

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-47/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-47/handler"
)

func main() {
	log.Println("Starting Day-47 Kubernetes Basics Go Microservice...")

	cfg := config.LoadConfig()
	log.Printf("Pod [%s] in namespace [%s] running on Node [%s] (Env: %s)",
		cfg.PodName, cfg.PodNamespace, cfg.NodeName, cfg.AppEnv)

	apiHandler := handler.NewAPIHandler(cfg)

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", apiHandler.Healthz)
	mux.HandleFunc("/ready", apiHandler.Ready)
	mux.HandleFunc("/api/v1/info", apiHandler.Info)
	mux.HandleFunc("/api/v1/data", apiHandler.Data)

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
		log.Printf("Server listening on port :%s inside container", cfg.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server failure: %v", err)
		}
	}()

	sig := <-sigChan
	log.Printf("Received termination signal [%v] from Kubelet. Initiating zero-downtime shutdown...", sig)

	// Step 1: Mark readiness probe as unready so Kube-Proxy removes pod from Service endpoints
	apiHandler.SetNotReady()
	log.Println("Readiness probe marked UNREADY. Waiting 5s for Kube-Proxy endpoint updates...")
	time.Sleep(5 * time.Second)

	// Step 2: Shutdown HTTP server cleanly
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Forced server shutdown error: %v", err)
	}

	log.Println("Kubernetes Pod process exited cleanly.")
}
