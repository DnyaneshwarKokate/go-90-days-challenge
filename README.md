# 🚀 Go 90 Days Challenge

<div align="center">

# Hi, I'm Dnyaneshwar Kokate 👋

### Full Stack .NET Developer | Future Go Backend Engineer

</div>

---

# 👨‍💻 About Me

I am a passionate software developer with experience in building scalable web applications using:

- ASP.NET Core
- C#
- Angular
- SQL Server
- Entity Framework Core
- Web API
- JWT Authentication

Currently, I am learning **Go (Golang)** from beginner to advanced level to become a high-performance backend engineer.

---

# 🎯 My Goal

✅ Become a strong Backend Developer  
✅ Master Go Lang & Microservices  
✅ Learn Docker & Kubernetes  
✅ Build scalable production-level APIs  
✅ Crack product-based company interviews  
✅ Improve problem-solving & system design skills  
✅ Upload code daily on GitHub consistently

---

# 🛠 Current Tech Stack

## Backend
- ASP.NET Core
- C#
- Web API
- Entity Framework Core
- Go (Learning)

## Frontend
- Angular
- React

## Database
- SQL Server
- PostgreSQL (Learning)

## Tools & DevOps
- Git & GitHub
- Docker (Learning)
- Kubernetes (Learning)
- Postman

---

# 📅 Go 90 Days Challenge Roadmap

| Day | Topic | Status |
|---|---|---|
| Day 01 | Go Setup, Variables, Data Types | ✅ |
| Day 02 | Functions, Conditions, Loops | ⏳ |
| Day 03 | Arrays, Slices, Maps | ⏳ |
| Day 04 | Structs | ⏳ |
| Day 05 | Pointers | ⏳ |
| Day 06 | Methods & Interfaces | ⏳ |
| Day 07 | Mini Project | ⏳ |
| Day 08 | Packages & Modules | ⏳ |
| Day 09 | File Handling | ⏳ |
| Day 10 | Error Handling | ⏳ |

---

# 📘 Daily Interview Questions & Answers

---

# ✅ Day 01 — Go Setup, Variables, Data Types

## ❓ Q1: What is Go Language?

### ✅ Answer:
Go (Golang) is an open-source programming language developed by Google.  
It is designed for building fast, scalable, and high-performance backend applications.

### ⭐ Features:
- Simple syntax
- Fast compilation
- Built-in concurrency
- Garbage collection
- Cross-platform support

---

## ❓ Q2: Why is Go popular for Backend Development?

### ✅ Answer:
Go is popular because:
- High performance
- Lightweight goroutines
- Fast APIs
- Easy concurrency handling
- Excellent for microservices

### ⭐ Companies using Go:
- Google
- Uber
- Netflix
- Docker
- Kubernetes

---

## ❓ Q3: What is the structure of a Go program?

### ✅ Answer:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
}
```

### Explanation:
- `package main` → Entry package
- `import` → Import packages
- `func main()` → Starting point of program
- `fmt.Println()` → Print output

---

## ❓ Q4: What are variables in Go?

### ✅ Answer:
Variables are used to store data.

### Example:

```go
var name string = "Dnyaneshwar"
var age int = 24
```

Short declaration:

```go
city := "Pune"
```

---

## ❓ Q5: What are data types in Go?

### ✅ Answer:

| Data Type | Example |
|---|---|
| int | 10 |
| string | "Go" |
| bool | true |
| float64 | 10.5 |

---

## ❓ Q6: Difference between var and := ?

### ✅ Answer:

| var | := |
|---|---|
| Explicit declaration | Short declaration |
| Used globally/local | Used only inside function |
| Requires datatype optional | Datatype auto detected |

### Example:

```go
var name string = "Go"

city := "Pune"
```

---

## ❓ Q7: What is fmt package?

### ✅ Answer:
`fmt` package is used for formatted input and output.

### Example:

```go
fmt.Println("Hello")
fmt.Printf("Age: %d", age)
```

---

# 💻 Day 01 Practice Program

```go
package main

import "fmt"

func main() {

	var name string = "Dnyaneshwar"
	var age int = 24
	var salary float64 = 45000.50
	var isDeveloper bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Developer:", isDeveloper)

	city := "Pune"
	fmt.Println("City:", city)
}
```

---
---

# ✅ Day 02 — Functions, Conditions & Loops

---

# 📖 What You Will Learn

- Functions
- Parameters
- Return Values
- Multiple Return Values
- If-Else Conditions
- Switch Case
- Loops
- Break & Continue
- Variable Scope

---

# 📌 Functions in Go

Functions are reusable blocks of code used to perform a specific task.

---

## ✅ Simple Function Example

```go
package main

import "fmt"

func greet() {
	fmt.Println("Welcome to Go Lang")
}

func main() {
	greet()
}
```

---

# 📌 Function with Parameters

Parameters allow passing values into functions.

## ✅ Example

```go
package main

import "fmt"

func greetWithName(name string) {
	fmt.Println("Welcome", name)
}

func main() {
	greetWithName("Dnyaneshwar")
}
```

---

# 📌 Function with Return Type

Functions can return values.

## ✅ Example

```go
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {

	result := add(10, 20)

	fmt.Println("Addition:", result)
}
```

---

# 📌 Multiple Return Values

Go supports multiple return values.

## ✅ Example

```go
package main

import "fmt"

func calculate(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {

	sum, sub := calculate(20, 10)

	fmt.Println("Sum:", sum)
	fmt.Println("Subtraction:", sub)
}
```

---

# 📌 If-Else Condition

Used for decision making.

## ✅ Example

```go
package main

import "fmt"

func main() {

	age := 20

	if age >= 18 {
		fmt.Println("Eligible for voting")
	} else {
		fmt.Println("Not eligible")
	}
}
```

---

# 📌 Else If Example

```go
package main

import "fmt"

func main() {

	marks := 75

	if marks >= 90 {
		fmt.Println("Grade A")
	} else if marks >= 70 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}
}
```

---

# 📌 Switch Statement

Switch is cleaner than multiple if-else conditions.

## ✅ Example

```go
package main

import "fmt"

func main() {

	day := 3

	switch day {

	case 1:
		fmt.Println("Monday")

	case 2:
		fmt.Println("Tuesday")

	case 3:
		fmt.Println("Wednesday")

	default:
		fmt.Println("Invalid Day")
	}
}
```

---

# 📌 Loops in Go

Go has only one loop:
## ✅ for loop

---

# ✅ Simple Loop

```go
package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}
```

---

# ✅ While Loop Style

```go
package main

import "fmt"

func main() {

	i := 1

	for i <= 5 {
		fmt.Println(i)
		i++
	}
}
```

---

# 📌 Break Statement

```go
package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		if i == 5 {
			break
		}

		fmt.Println(i)
	}
}
```

---

# 📌 Continue Statement

```go
package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {

		if i == 3 {
			continue
		}

		fmt.Println(i)
	}
}
```

---

# 📌 Variable Scope

Scope defines where a variable can be accessed.

## ✅ Local Scope Example

```go
package main

import "fmt"

func test() {
	name := "Go"
	fmt.Println(name)
}

func main() {
	test()
}
```

---

# 💻 Day 02 Practice Program

```go
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

	fmt.Println("Age:", age)

	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
	}
}
```

---

# 📘 Day 02 Interview Questions & Answers

---

## ❓ Q1: What is a function in Go?

### ✅ Answer:
A function is a reusable block of code used to perform a specific task.

### Example:

```go
func greet() {
	fmt.Println("Hello")
}
```

---

## ❓ Q2: How to pass parameters in Go function?

### ✅ Answer:

```go
func add(a int, b int) int {
	return a + b
}
```

Parameters:
- `a int`
- `b int`

---

## ❓ Q3: Can Go return multiple values?

### ✅ Answer:
Yes, Go supports multiple return values.

### Example:

```go
func calculate(a int, b int) (int, int) {
	return a+b, a-b
}
```

---

## ❓ Q4: Difference between if-else and switch?

| if-else | switch |
|---|---|
| Multiple conditions | Cleaner for fixed cases |
| Complex logic | Better readability |

---

## ❓ Q5: Which loop is available in Go?

### ✅ Answer:
Go has only one loop:
```go
for
```

It can behave like:
- for loop
- while loop
- infinite loop

---

## ❓ Q6: What is break statement?

### ✅ Answer:
`break` stops loop execution immediately.

---

## ❓ Q7: What is continue statement?

### ✅ Answer:
`continue` skips current iteration and moves to next iteration.

---

## ❓ Q8: What is variable scope in Go?

### ✅ Answer:
Scope defines where a variable can be accessed.

### Types:
- Local Scope
- Global Scope

### Example:

```go
var company = "Google"

func main() {
	name := "Dnyaneshwar"
}
```

---

# 🧠 Mini Practice Tasks

✅ Create multiplication function  
✅ Create calculator using switch  
✅ Print numbers 1–100  
✅ Check even/odd using if-else  
✅ Create voting eligibility checker

---

# 🔥 My Learning Rules

✅ Code Daily  
✅ Push Daily on GitHub  
✅ Learn by Building Projects  
✅ Practice Interview Questions  
✅ Focus on Backend Engineering  
✅ Build Production-Level Projects

---

# 📌 Future Projects

- Todo API
- Student Management API
- JWT Authentication API
- E-Commerce Backend
- URL Shortener
- Chat Application
- Microservices Architecture

---

# 📫 Connect With Me

📧 Email: dnyaneshwarkokatevip@gmail.com  
💻 GitHub: https://github.com/Dnyanesh0902  
🌐 Portfolio: https://dnyanesh.miracledevelopers.in  
🔗 LinkedIn: https://www.linkedin.com/in/dnyaneshwar-kokate-04a12b258/

---

# ⭐ Challenge Progress

✅ Day 02 Completed  
🚀 Next: Arrays, Slices & Maps