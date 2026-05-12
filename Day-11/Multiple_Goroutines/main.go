package main

import (
	"fmt"
	"time"
)


func taskOne() {
	for i := 1; i <= 5; i++{
		fmt.Println("Task One:", i)

		time.Sleep(time.Millisecond * 300)
	}
}

func taskTwo() {
	for i := 1; i <= 5; i++{
		fmt.Println("Task Two:", i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	go taskOne()
	go taskTwo()
	time.Sleep(time.Second * 3)
}