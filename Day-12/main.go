package main

import "fmt"

func main() {
	messageChannel := make(chan string)

	go func(){
		messageChannel <- "hello from Goroutine"
	}()

	message := <- messageChannel

	fmt.Println(message)

	
}