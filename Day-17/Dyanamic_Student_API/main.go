package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Student struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var students = []Student {
	{ID: 1, Name: "Dnyaneshwar"},
	{ID: 2, Name: "Sai"},
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) <3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)

		return
	}

	id := parts[2]

	for _, student := range students {
		if strconv.Itoa(student.ID) == id {
			w.Header().Set("Content-Type","application/json")

			json.NewEncoder(w).Encode(student)
			return
		}
	}
	http.Error(w, "Student Not Found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/student/", getStudent)
	http.ListenAndServe(":8080", nil)
}