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
	student := Student {
		Name: "Dnyaneshwar",
		Age: 24,
		City: "Nashik",
	}
	jsonData, err := json.Marshal(student)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}