package main

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	student := Student {
		Name: "Dnyaneshwar",
		Age: 24,
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(student)
}

func main() {
	http.HandleFunc("/student", studentHandler)

	http.ListenAndServe(":8080", nil)
}