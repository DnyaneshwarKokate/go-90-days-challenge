package main

import (
	"fmt"
	"time"
)

func printNumber() {
	for i := 1; i <=5; i++{
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printNumber()

	time.Sleep(time.Second * 3)
	fmt.Println("Main Function")
}