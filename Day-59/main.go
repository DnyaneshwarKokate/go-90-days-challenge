package main

import (
	"log"
	"net/http"

	"day59/metrics"
)

func main() {
	col := metrics.NewMetricsCollector()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/data", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	mux.HandleFunc("/metrics", col.Handler())

	handlerStack := col.MetricsMiddleware(mux)

	log.Println("Starting Day 59 Prometheus Monitoring Service on :8080...")
	log.Println("Prometheus metrics exporter live at http://localhost:8080/metrics")

	if err := http.ListenAndServe(":8080", handlerStack); err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
