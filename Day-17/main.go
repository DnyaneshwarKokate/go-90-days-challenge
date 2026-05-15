package main

import (
	"fmt"
	"net/http"
	"strings"
)

func studentHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path,"/")

	if len(parts) <3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id := parts[2]

	fmt.Fprintln(w, "Student ID:", id)

}

func main() {
	http.HandleFunc("/student/", studentHandler)
	http.ListenAndServe(":8080", nil)

}