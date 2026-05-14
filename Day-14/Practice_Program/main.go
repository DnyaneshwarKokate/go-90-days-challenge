package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Name string `json:"name"`
	Price int `json:"price"`
}

func productHandler(w http.ResponseWriter, r *http.Request) {

	product := Product{
		Name: "IPhone",
		Price: 120000,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Go HTTP Server"))
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/product", productHandler)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Api Running Successfully!")

}