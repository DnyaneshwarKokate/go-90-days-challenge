package main

import (
	"fmt"
	"sync"
)

func sendData(channel chan string, wg *sync.WaitGroup)  {
	defer wg.Done()

	channel <- "Hello Go"
}

func main() {
	var wg sync.WaitGroup

	channel := make(chan string)
	wg.Add(1)

	go sendData(channel, &wg)

	message := <- channel
	fmt.Println(message)
	wg.Wait()
}