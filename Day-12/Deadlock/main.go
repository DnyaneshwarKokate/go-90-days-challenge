package main

func main() {
	channel := make(chan int)

	channel <- 10
}