package main

import (
	"fmt"
	"errors"
)
func checkAge(age int) error{
	if age < 18 {
		return  errors.New("Age must be 18 or above")
	}
	return  nil
}

func main(){
	err := checkAge(18)

	if err != nil{
		fmt.Println("Error",err)
		return
	}
	fmt.Println("Eligible")
}