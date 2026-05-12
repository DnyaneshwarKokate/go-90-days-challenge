package main

import (
	"fmt"
	"os"
)

func main(){
	err := os.Remove("Student.txt")

	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("File Deleted")
	
}