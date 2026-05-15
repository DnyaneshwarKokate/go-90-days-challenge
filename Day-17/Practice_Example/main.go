package main

import (
	"fmt"
	"net/http"
	"strings"
)

func productHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {

		http.Error(w, "Invalid Product URL", http.StatusBadRequest)

		return
	}

	productID := parts[2]

	fmt.Fprintln(w, "Product ID:", productID)
}

func main() {

	http.HandleFunc("/product/", productHandler)

	fmt.Println("Server Running on Port 8080")

	http.ListenAndServe(":8080", nil)
}