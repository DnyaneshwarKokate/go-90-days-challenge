package handler

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Commit  string `json:"commit"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := StatusResponse{
		Service: "github-actions-demo",
		Status:  "operational",
		Commit:  "a1b2c3d4e5",
	}
	_ = json.NewEncoder(w).Encode(res)
}
