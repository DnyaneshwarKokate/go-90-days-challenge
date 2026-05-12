package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("student.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File Created Successfully")
	file.Close()
}