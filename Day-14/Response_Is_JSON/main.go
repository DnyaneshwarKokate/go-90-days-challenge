package main

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r * http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, `{"message:"Login Successful""}`)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}