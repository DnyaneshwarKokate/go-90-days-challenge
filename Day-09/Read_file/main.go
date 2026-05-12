package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("Student.txt")

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}