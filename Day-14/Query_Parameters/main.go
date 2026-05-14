package main

import (
	"fmt"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")
	fmt.Fprintln(w, "Hello", name)
}

func main() {
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}