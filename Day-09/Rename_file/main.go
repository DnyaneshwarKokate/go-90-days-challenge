package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Rename(
		"Student.txt",
		"data.txt",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File Renamed")
}
