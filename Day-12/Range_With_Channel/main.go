package main

import "fmt"

func main() {
	channel := make(chan int)

	go func ()  {
		for i := 1; i <= 5; i++ {
			channel <- i 
		}
		close(channel)
	}()

	for value := range channel {
		fmt.Println(value)
	}
}