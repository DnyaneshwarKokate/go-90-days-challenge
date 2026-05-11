package main

import "fmt"

// Simple Function Example
func greet() {
	fmt.Println("Welcome to go lang")
}

// Function with Parameters
func greetWithName(name string) {
	fmt.Println("Welcome", name, "to go lang")
}

// Function with return type
func add(a int, b int) int {
	return a + b
}

// Multiple Return Values
func calculate(a int, b int) (int, int) {
	return a + b, a - b
}

// if-else condition
func checkAge(age int) {

	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}
}

//Else If
func checkGrade(score int)  {
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
}

// Switch Statement
func checkDay(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}

// Loops in Go
// for loop
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func main() {

	greet()

	greetWithName("Dnyaneshwar")

	result := add(10, 20)

	fmt.Println("The sum is:", result)

	sum, sub := calculate(10, 5)

	fmt.Println("Sum:", sum)
	fmt.Println("Subtraction:", sub)

	age := 25

	checkAge(age)
	checkGrade(85)
	checkDay(5)
	printNumbers()
	fmt.Println("Age:", age)
}