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
| Day 02 | Functions, Conditions, Loops | ✅ |
| Day 03 | Arrays, Slices, Maps | ✅ |
| Day 04 | Structs | ✅ |
| Day 05 | Pointers | ⏳ |
| Day 06 | Methods & Interfaces | ⏳ |
| Day 07 | Mini Project | ⏳ |
| Day 08 | Packages & Modules | ⏳ |
| Day 09 | File Handling | ⏳ |
| Day 10 | Error Handling | ⏳ |

---

# 📘 Daily Interview Questions & Answers

---

# ✅ Day 01 — Go Setup, Variables & Data Types

---

# 📖 Introduction to Go

Go (Golang) is an open-source programming language developed by Google.

It is designed for:
- High performance
- Scalability
- Backend development
- Cloud-native applications
- Microservices

---

# ⭐ Features of Go

✅ Simple Syntax  
✅ Fast Compilation  
✅ Built-in Concurrency  
✅ Garbage Collection  
✅ Cross-platform Support

---

# 📌 First Go Program

```go
package main

import "fmt"

func main() {

	fmt.Println("Hello Go Lang")
}
```

---

# 📌 Explanation

| Code | Meaning |
|---|---|
| package main | Entry package |
| import "fmt" | Import fmt package |
| func main() | Main function |
| fmt.Println() | Print output |

---

# 📌 Variables in Go

Variables are used to store data.

---

# ✅ Variable Example

```go
package main

import "fmt"

func main() {

	var name string = "Dnyaneshwar"
	var age int = 24

	fmt.Println(name)
	fmt.Println(age)
}
```

---

# 📌 Short Variable Declaration

```go
city := "Nashik"
```

Go automatically detects datatype.

---

# 📌 Data Types in Go

| Data Type | Example |
|---|---|
| int | 10 |
| string | "Go" |
| bool | true |
| float64 | 10.5 |

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

# 📘 Day 01 Interview Questions

---

## ❓ Q1: What is Go Language?

### ✅ Answer:
Go is an open-source programming language developed by Google.

---

## ❓ Q2: What are variables?

### ✅ Answer:
Variables are used to store data.

---

## ❓ Q3: Difference between var and := ?

| var | := |
|---|---|
| Explicit declaration | Short declaration |
| Can use globally | Inside function only |

---

## ❓ Q4: What is fmt package?

### ✅ Answer:
fmt package is used for input and output operations.

---

# 📚 Day 01 Summary

Today I learned:
- Go setup
- Variables
- Data types
- fmt package
- First Go program

---

# ✅ Day 02 — Functions, Conditions & Loops

---

# 📖 What You Will Learn

- Functions
- Parameters
- Return Values
- If-Else
- Switch
- Loops
- Variable Scope

---

# 📌 Functions in Go

Functions are reusable blocks of code.

---

# ✅ Simple Function

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

```go
package main

import "fmt"

func add(a int, b int) int {

	return a + b
}

func main() {

	result := add(10, 20)

	fmt.Println(result)
}
```

---

# 📌 Multiple Return Values

```go
package main

import "fmt"

func calculate(a int, b int) (int, int) {

	return a + b, a - b
}

func main() {

	sum, sub := calculate(20, 10)

	fmt.Println(sum)
	fmt.Println(sub)
}
```

---

# 📌 If-Else Condition

```go
package main

import "fmt"

func main() {

	age := 20

	if age >= 18 {

		fmt.Println("Eligible")
	} else {

		fmt.Println("Not Eligible")
	}
}
```

---

# 📌 Switch Case

```go
package main

import "fmt"

func main() {

	day := 2

	switch day {

	case 1:
		fmt.Println("Monday")

	case 2:
		fmt.Println("Tuesday")

	default:
		fmt.Println("Invalid")
	}
}
```

---

# 📌 Loops in Go

Go has only one loop:
# ✅ for loop

---

# ✅ Loop Example

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

# 📌 Break Example

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

# 📌 Continue Example

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

Scope defines where variable can be accessed.

---

# 💻 Day 02 Practice Program

```go
package main

import "fmt"

func greet() {

	fmt.Println("Welcome to go lang")
}

func greetWithName(name string) {

	fmt.Println("Welcome", name)
}

func add(a int, b int) int {

	return a + b
}

func checkAge(age int) {

	if age >= 18 {

		fmt.Println("Eligible")
	} else {

		fmt.Println("Not Eligible")
	}
}

func main() {

	greet()

	greetWithName("Dnyaneshwar")

	result := add(10, 20)

	fmt.Println(result)

	checkAge(24)

	for i := 1; i <= 5; i++ {

		fmt.Println(i)
	}
}
```

---

# 📘 Day 02 Interview Questions

---

## ❓ Q1: What is function?

### ✅ Answer:
Function is reusable block of code.

---

## ❓ Q2: Can Go return multiple values?

### ✅ Answer:
Yes.

---

## ❓ Q3: Which loop available in Go?

### ✅ Answer:
Only for loop.

---

## ❓ Q4: What is break statement?

### ✅ Answer:
Stops loop immediately.

---

## ❓ Q5: What is continue statement?

### ✅ Answer:
Skips current iteration.

---

# 📚 Day 02 Summary

Today I learned:
- Functions
- Parameters
- Return values
- Conditions
- Switch
- Loops

---

# ✅ Day 03 — Arrays, Slices & Maps

---

# 📖 Introduction to Arrays, Slices & Maps

In backend development, we often need to store and manage multiple values.

Examples:
- Student marks
- Product names
- User data
- Cities list
- Employee records

Go provides:
- Arrays → Fixed-size collection
- Slices → Dynamic collection
- Maps → Key-value storage

---

# 📌 Arrays in Go

An array is a fixed-size collection of elements of same datatype.

---

# 📌 Array Example

```go
package main

import "fmt"

func main() {

	numbers := [5]int{10,20,30,40,50}

	fmt.Println(numbers)
}
```

---

# 📌 Access Array Elements

```go
fmt.Println(numbers[2])
```

Output:
```text
30
```

---

# 📌 Loop with Array

```go
package main

import "fmt"

func main() {

	numbers := [5]int{10,20,30,40,50}

	for i := 0; i < len(numbers); i++ {

		fmt.Println(numbers[i])
	}
}
```

---

# 📌 Slices in Go

Slices are dynamic arrays.

---

# 📌 Slice Example

```go
package main

import "fmt"

func main() {

	cities := []string{"Pune","Mumbai"}

	cities = append(cities, "Nashik")

	fmt.Println(cities)
}
```

---

# 📌 len() and cap()

```go
package main

import "fmt"

func main() {

	numbers := []int{1,2,3,4,5}

	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}
```

---

# 📌 Maps in Go

Maps store key-value data.

---

# 📌 Map Example

```go
package main

import "fmt"

func main() {

	student := map[string]string{

		"name":"Dnyaneshwar",
		"city":"Nashik",
	}

	fmt.Println(student)
}
```

---

# 📌 Add Values in Map

```go
student["role"] = "Developer"
```

---

# 📌 Delete Values from Map

```go
delete(student, "city")
```

---

# 📌 Range Keyword

Used for iteration.

---

# 📌 Range Example

```go
package main

import "fmt"

func main() {

	numbers := []int{10,20,30,40,50}

	for index, value := range numbers {

		fmt.Println(index, value)
	}
}
```

---

# 💻 Day 03 Practice Program

```go
package main

import "fmt"

func main() {

	// Array
	marks := [5]int{80,75,90,85,70}

	fmt.Println(marks)

	// Slice
	cities := []string{"Pune","Mumbai"}

	cities = append(cities, "Nashik")

	fmt.Println(cities)

	// Map
	student := map[string]string{

		"name":"Dnyaneshwar",
		"city":"Nashik",
		"role":"Developer",
	}

	fmt.Println(student)

	// Range
	for index, value := range marks {

		fmt.Println(index, value)
	}
}
```

---

# 📘 Day 03 Interview Questions

---

## ❓ Q1: What is array?

### ✅ Answer:
Fixed-size collection of same datatype.

---

## ❓ Q2: What is slice?

### ✅ Answer:
Dynamic version of array.

---

## ❓ Q3: Difference between array and slice?

| Array | Slice |
|---|---|
| Fixed size | Dynamic size |

---

## ❓ Q4: What is append()?

### ✅ Answer:
Used to add values in slice.

---

## ❓ Q5: What is map?

### ✅ Answer:
Map stores key-value data.

---

## ❓ Q6: What is range?

### ✅ Answer:
Used for iteration.

---

# 📚 Day 03 Summary

Today I learned:
- Arrays
- Slices
- Maps
- append()
- range
- len() & cap()

---
---

# ✅ Day 04 — Structs & Custom Types

---

# 📖 Introduction to Structs

Struct is one of the most important concepts in Go.

Structs are used to group related data together.

In backend development, structs are used for:
- API Models
- Database Entities
- JSON Handling
- Request & Response Models
- User Data
- Product Data

If you know C#, then:
👉 Struct in Go is similar to a Model/Class in ASP.NET Core.

---

# 📖 What You Will Learn

- Structs
- Custom Types
- Struct Initialization
- Access Struct Fields
- Nested Structs
- Anonymous Structs
- Struct with Functions
- Exported vs Unexported Fields
- Interview Questions

---

# 📌 What is Struct?

Struct is a custom datatype used to group multiple related values together.

---

# 📌 Real Life Example

Suppose we want to store:
- Student Name
- Age
- City

Without struct:

```go
name := "Dnyaneshwar"
age := 24
city := "Nashik"
```

This becomes difficult to manage.

Using struct:

```go
type Student struct {

	Name string
	Age  int
	City string
}
```

This is cleaner and professional.

---

# 📌 Struct Syntax

```go
type Student struct {

	Name string
	Age  int
	City string
}
```

---

# 📌 Explanation

| Part | Meaning |
|---|---|
| type | Create custom datatype |
| Student | Struct name |
| struct | Keyword |
| Name string | Field |

---

# 📌 Struct Example

```go
package main

import "fmt"

type Student struct {

	Name string
	Age  int
	City string
}

func main() {

	var student1 Student

	student1.Name = "Dnyaneshwar"
	student1.Age = 24
	student1.City = "Nashik"

	fmt.Println(student1)
}
```

---

# 📌 Output

```text
{Dnyaneshwar 24 Nashik}
```

---

# 📌 Access Struct Fields

```go
fmt.Println(student1.Name)
fmt.Println(student1.Age)
fmt.Println(student1.City)
```

---

# 📌 Short Struct Initialization

Instead of assigning values one by one:

```go
student := Student{

	Name: "Dnyaneshwar",
	Age: 24,
	City: "Nashik",
}
```

---

# ✅ Complete Example

```go
package main

import "fmt"

type Student struct {

	Name string
	Age  int
	City string
}

func main() {

	student := Student{

		Name: "Dnyaneshwar",
		Age: 24,
		City: "Nashik",
	}

	fmt.Println(student)
}
```

---

# 📌 Why Structs are Important?

Structs are used in almost every backend application.

Examples:
- User Models
- Product Models
- Employee Records
- Orders
- API Responses
- JSON Data

---

# 📌 Real Backend Example

Suppose we are creating Student API.

```go
type Student struct {

	Id    int
	Name  string
	Email string
	Age   int
}
```

This is similar to:
- Entity Class in ASP.NET Core
- Database Model

---

# 📌 Struct with Functions

Struct data can be passed into functions.

---

# ✅ Example

```go
package main

import "fmt"

type Employee struct {

	Name   string
	Salary int
}

func printEmployee(emp Employee) {

	fmt.Println("Employee Name:", emp.Name)
	fmt.Println("Salary:", emp.Salary)
}

func main() {

	employee := Employee{

		Name: "Dnyaneshwar",
		Salary: 50000,
	}

	printEmployee(employee)
}
```

---

# 📌 Anonymous Struct

Struct without name.

---

# ✅ Example

```go
package main

import "fmt"

func main() {

	employee := struct {

		Name string
		Age  int
	}{

		Name: "Dnyaneshwar",
		Age: 24,
	}

	fmt.Println(employee)
}
```

---

# 📌 Nested Struct

Struct inside another struct.

---

# ✅ Example

```go
package main

import "fmt"

type Address struct {

	City string
	State string
}

type Employee struct {

	Name string
	Address Address
}

func main() {

	employee := Employee{

		Name: "Dnyaneshwar",

		Address: Address{

			City: "Nashik",
			State: "Maharashtra",
		},
	}

	fmt.Println(employee)
}
```

---

# 📌 Exported vs Unexported Fields

Important for APIs & JSON.

---

# ✅ Exported Field

Starts with Capital Letter:

```go
Name string
```

Accessible outside package.

---

# ❌ Unexported Field

Starts with Small Letter:

```go
name string
```

Accessible only inside same package.

---

# 📌 Important for JSON APIs

Correct:

```go
Name string
```

Wrong:

```go
name string
```

---

# 📌 Why?

JSON conversion only works properly with exported fields.

---

# 📌 Struct Comparison

Structs can be compared if fields are comparable.

```go
student1 == student2
```

---

# 💻 Day 04 Practice Program

```go
package main

import "fmt"

type Student struct {

	Name string
	Age  int
	City string
}

type Employee struct {

	Name   string
	Role   string
	Salary int
}

func printStudent(student Student) {

	fmt.Println("Student Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("City:", student.City)
}

func main() {

	student := Student{

		Name: "Dnyaneshwar",
		Age: 24,
		City: "Nashik",
	}

	employee := Employee{

		Name: "Rahul",
		Role: "Backend Developer",
		Salary: 60000,
	}

	fmt.Println(student)

	fmt.Println(employee)

	printStudent(student)
}
```

---

# 📘 Day 04 Interview Questions & Answers

---

## ❓ Q1: What is struct in Go?

### ✅ Answer:
Struct is a custom datatype used to group related data together.

---

## ❓ Q2: Why structs are important?

### ✅ Answer:
Structs are used for:
- API Models
- Database Entities
- JSON Handling
- Backend Development

---

## ❓ Q3: Difference between struct and array?

| Struct | Array |
|---|---|
| Different datatypes | Same datatype |
| Represents object | Represents collection |

---

## ❓ Q4: What is nested struct?

### ✅ Answer:
Struct inside another struct.

---

## ❓ Q5: Why exported fields are important?

### ✅ Answer:
Exported fields are required for:
- JSON conversion
- Package access
- API responses

---

## ❓ Q6: Difference between Name and name in Go?

| Name | name |
|---|---|
| Exported | Unexported |
| Public | Private |

---

## ❓ Q7: What is anonymous struct?

### ✅ Answer:
Struct without a name is called anonymous struct.

---

# 📚 Day 04 Summary

Today I learned:
- Structs
- Custom Types
- Nested Structs
- Anonymous Structs
- Struct Functions
- Exported vs Unexported Fields

I also practiced:
- Creating models
- Accessing fields
- Passing structs into functions
- Backend-style struct design

---

# 🧠 Practice Tasks

✅ Create Student struct  
✅ Create Product struct  
✅ Create Employee struct  
✅ Create nested Address struct  
✅ Print struct fields  
✅ Pass struct into function  
✅ Create User API model structure

---

# 🔥 My Learning Rules

✅ Code Daily  
✅ Push Daily on GitHub  
✅ Learn by Building Projects  
✅ Practice Interview Questions  
✅ Focus on Backend Engineering  
✅ Build Production-Level Projects

---

# 📫 Connect With Me

📧 Email: dnyaneshwarkokatevip@gmail.com  
💻 GitHub: https://github.com/Dnyanesh0902  
🌐 Portfolio: https://dnyanesh.miracledevelopers.in  
🔗 LinkedIn: https://www.linkedin.com/in/dnyaneshwar-kokate-04a12b258/

---

# ⭐ Challenge Progress

✅ Day 01 Completed  
✅ Day 02 Completed  
✅ Day 03 Completed  
✅ Day 04 Completed  
🚀 Next: Pointers in Go