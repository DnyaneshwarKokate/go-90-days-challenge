package main
import "fmt"

func main() {
	channel := make(chan int, 2)

	channel <- 10
	channel <- 20

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}