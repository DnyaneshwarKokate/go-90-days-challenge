package main
import "fmt"

func handlePanic() {
	message := recover()

	if message != nil {
		fmt.Println("Recovered:", message)
	}
}

func main() {
	defer handlePanic()
	panic("Something went wrong")
}