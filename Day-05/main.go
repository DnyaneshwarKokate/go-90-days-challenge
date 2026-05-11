package main

import "fmt"
func main() {
	age := 25
	// fmt.Println(age)
	// fmt.Println(&age)

	// var ptr *int
	ptr := &age
	fmt.Println(*ptr)

}