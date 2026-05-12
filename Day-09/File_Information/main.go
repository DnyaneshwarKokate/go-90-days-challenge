package main

import (
	"fmt"
	"os"
)
func main(){
	info, err := os.Stat("data.txt")

	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("File Name: ", info.Name())
	fmt.Println("File Size :", info.Size())
}