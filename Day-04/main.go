package main
import "fmt"

type Student struct {
	Name string
	Age  int
	City string
}
func main() {
	var student1 Student
	student1.Name = "Alice"
	student1.Age = 20
	student1.City = "New York"

	fmt.Println(student1)
}