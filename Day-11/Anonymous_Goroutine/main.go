package main 

import (
	"fmt"
	"time"
)

func main() {
	go func ()  {
		fmt.Println("Anonymous Goroutine")
	}()
	time.Sleep(time.Second)
}