package main

import (
	"fmt"
)

func printNumbers(){
	for i := 1; i<= 5; i++ {
		fmt.Println(i)
	}
}

func main() {
	printNumbers()

	fmt.Println("Main Functions")
}