package main

import (
	"fmt"
	"os"
)

func main() {
	file , err := os.OpenFile(
		"Student.txt",
		os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("\nLearning Go Lang")
	fmt.Println("Data Appended")
}