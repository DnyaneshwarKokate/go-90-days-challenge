package main

import (
	"encoding/json"
	"net/http"
	
)

type Student struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}


var students = []Student{
	{ID: 1, Name: "Dnyaneshwar", Age: 24},
	{ID: 2, Name: "Rahul", Age: 23},
	{ID: 3, Name: "Vaibhav", Age: 25},
}
//Get data
func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(students)
}
//Add Data
func createStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
		return
	}
	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	students = append(students, student)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func main(){
	http.HandleFunc("/students",getStudents)
	http.HandleFunc("/addStudent", createStudent)
	http.ListenAndServe(":8080", nil)
}