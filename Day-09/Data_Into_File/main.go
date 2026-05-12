package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("Student.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("Welcome to Go Lang")
	fmt.Println("Data Written Successfully.")
}
