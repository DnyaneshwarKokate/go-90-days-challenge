package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age int
	City string
}

func main() {
	jsonData := `{
	"Name":"Dnyaneshwar",
	"Age":24,
	"City":"Nashik"
	}`

	var student Student

	err := json.Unmarshal([]byte(jsonData), &student)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(student.Name)
	fmt.Println(student.Age)
	fmt.Println(student.City)
}