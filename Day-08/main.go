package main 

import (
	"fmt"
	"day08/calculator"
)

func main() {
	result := calculator.Add(10,20)
	fmt.Println("Addition: ", result)

	sub  := calculator.Subtract(30, 10)
	fmt.Println("Subtraction:", sub)
}