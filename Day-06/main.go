// package main
// import "fmt"
// type Student struct{
// 	Name string
// 	Age int
// }

// func (s Student) display(){
// 	fmt.Println("Name", s.Name)
// 	fmt.Println("Age", s.Age)
// }


// func main(){
// 	student := Student{
// 		Name : "Dnyaneshwar",
// 		Age : 24,
// 	}
// 	student.display()
// 	}

//value receiver
// package main
// import "fmt"

// type Employee struct {
// 	Name string
// 	Salary int
// }

// func (e Employee) updateSalary()  {
// 	e.Salary = 70000
// }

// func main() {

// 	employee := Employee{
// 		Name: "Dnyaneshwar",
// 		Salary: 50000,
// 	}
// 	employee.updateSalary()
// 	fmt.Println(employee.Salary)
// }

//pointer receiver

// package main

// import "fmt"
// type Employee struct {
// 	Name string
// 	Salary int
// }

// func (e *Employee) updateSalary()  {
// 	e.Salary = 70000
// }

// func main() {
// 	employee := Employee{
// 		Name: "Dnyaneshwar",
// 		Salary: 50000,
// 	}
// 	employee.updateSalary()
// 	fmt.Println(employee.Salary)
// }
//Real Example
// package main

// import "fmt"

// type Rectangle struct {
// 	Width int
// 	Height int
// }

// func (r Rectangle) area() {
// 	fmt.Println("Area:", r.Width*r.Height)
// }

// type Shape interface {
// 	area()
// }

// func main() {
// 	rect := Rectangle{
// 		Width:  10,
// 		Height: 30,
// 	}
// 	var shape Shape
// 	shape = rect
// 	shape.area()
// }

package main

import "fmt"

type Payment interface {
	Pay()
}
type UPI struct {}

func (u UPI) Pay() {
	fmt.Println("Paid using UPI")
}

type Card struct {}

func (c Card) Pay() {
	fmt.Println("Paid using Card")
}


func main() {
	var payment Payment
	payment = UPI{}
	payment.Pay()

	payment = Card{}
	payment.Pay()
}