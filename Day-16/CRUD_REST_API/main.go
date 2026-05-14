package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


type Student struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

var students = []Student {

	{ID: 1, Name: "Dnyaneshwar", Age: 24},
	{ID: 2, Name: "Rahul", Age:23},
	
}

//GET ALL
func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(students)
}

//GET BY_ID

func getStudentByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	for _, student := range students {
		if strconv.Itoa(student.ID) == id {
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	http.Error(w, "Student Not Found", http.StatusNotFound)
}


//POST API - Add Student
func createStudent(w http.ResponseWriter, r * http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
		return
	}

	var student Student

	err :=  json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	students = append(students, student)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}


func updateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	var updatedStudent Student

	err := json.NewDecoder(r.Body).Decode(&updatedStudent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, student := range students {
		if strconv.Itoa(student.ID) == id {
			students[index] = updatedStudent

			json.NewEncoder(w).Encode(updatedStudent)
			return
		}
	}
	http.Error(w, "Student Not Found", http.StatusNotFound)
}


//Delete Student

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	for index, student := range students {
		if strconv.Itoa(student.ID) == id {
			students = append(
				students[:index],
				students[index+1:]...,
			)

			fmt.Fprintln(w, "Student Deleted Successfully")
			return
		}
	}
	http.Error(w, "Student Not Found", http.StatusNotFound)
}
func main() {
	http.HandleFunc("/students", getStudents)
	http.HandleFunc("/getstudent", getStudentByID)
	http.HandleFunc("/addstudent", createStudent)
	http.HandleFunc("/updatestudent", updateStudent)
	http.HandleFunc("/deleteStudent",deleteStudent)
	http.ListenAndServe(":8080",nil)

}