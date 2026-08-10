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

# 📅 Go 90 Days Challenge Roadmap

| Day | Topic | Status |
|---|---|---|
| Day 01 | Go Setup, Variables, Data Types | ✅ |
| Day 02 | Functions, Conditions, Loops | ✅ |
| Day 03 | Arrays, Slices, Maps | ✅ |
| Day 04 | Structs & Custom Types | ✅ |
| Day 05 | Pointers in Go | ✅ |
| Day 06 | Methods & Interfaces | ✅ |
| Day 07 | Student Management Mini Project | ✅ |
| Day 08 | Packages & Modules | ✅ |
| Day 09 | File Handling in Go | ✅ |
| Day 10 | Error Handling in Go | ✅ |
| Day 11 | Goroutines & Concurrency | ✅ |
| Day 12 | Channels in Go | ✅ |
| Day 13 | JSON Handling | ✅ |
| Day 14 | HTTP Package | ✅ |
| Day 15 | First REST API | ✅ |
| Day 16 | CRUD REST API | ✅ |
| Day 17 | Routing & URL Parameters | ✅ |
| Day 18 | Gin Framework | ✅ |
| Day 19 | PostgreSQL Integration | ✅ |
| Day 20 | GORM in Go | ✅ |
| Day 21 | JWT Authentication | ✅ |
| Day 22 | Middleware in Go | ✅ |
| Day 23 | Password Hashing | ✅ |
| Day 24 | Role-Based Authorization | ✅ |
| Day 25 | Environment Variables | ✅ |
| Day 26 | Clean Architecture | ✅ |
| Day 27 | Repository Pattern | ✅ |
| Day 28 | Dependency Injection | ✅ |
| Day 29 | Logging System | ✅ |
| Day 30 | Student Management REST API Project | ✅ |
| Day 31 | Advanced CRUD APIs | ✅ |
| Day 32 | Pagination & Filtering | ⏳ |
| Day 33 | File Upload API | ⏳ |
| Day 34 | Email Sending in Go | ⏳ |
| Day 35 | REST Client & External APIs | ⏳ |
| Day 36 | Redis Basics | ⏳ |
| Day 37 | Redis Caching | ⏳ |
| Day 38 | WebSockets in Go | ⏳ |
| Day 39 | Real-Time Chat Backend | ⏳ |
| Day 40 | Goroutine Worker Pools | ⏳ |
| Day 41 | Rate Limiting | ⏳ |
| Day 42 | Unit Testing in Go | ⏳ |
| Day 43 | Benchmark Testing | ⏳ |
| Day 44 | Docker Basics | ⏳ |
| Day 45 | Dockerizing Go API | ⏳ |
| Day 46 | Docker Compose | ⏳ |
| Day 47 | Kubernetes Basics | ⏳ |
| Day 48 | Deploy Go App on Kubernetes | ⏳ |
| Day 49 | CI/CD Basics | ⏳ |
| Day 50 | GitHub Actions CI/CD | ⏳ |
| Day 51 | Microservices Introduction | ⏳ |
| Day 52 | User Microservice | ⏳ |
| Day 53 | Product Microservice | ⏳ |
| Day 54 | API Gateway | ⏳ |
| Day 55 | Service Communication | ⏳ |
| Day 56 | gRPC Basics | ⏳ |
| Day 57 | Kafka/RabbitMQ Basics | ⏳ |
| Day 58 | Event-Driven Architecture | ⏳ |
| Day 59 | Monitoring & Logging | ⏳ |
| Day 60 | Production API Project | ⏳ |
| Day 61–70 | Advanced Backend Projects | ⏳ |
| Day 71–80 | Microservices Projects | ⏳ |
| Day 81–85 | System Design Basics | ⏳ |
| Day 86–88 | Interview Preparation | ⏳ |
| Day 89 | Resume & Portfolio Update | ⏳ |
| Day 90 | Final Production-Level Project | ⏳ |

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
---
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

# ✅ Day 05 — Pointers in Go

---

# 📖 Introduction to Pointers

Pointers are one of the most important concepts in Go.

Pointers are used for:
- Memory optimization
- Backend development
- Struct handling
- APIs
- Database operations

Pointers help us:
✅ Improve performance  
✅ Reduce memory usage  
✅ Modify original values directly

---

# 📖 What You Will Learn

- Pointers
- Memory Address
- Address Operator (&)
- Dereferencing (*)
- Pass by Value
- Pass by Pointer
- Struct Pointers
- new() Function
- Nil Pointers
- Interview Questions

---

# 📌 What is Pointer?

A pointer stores the memory address of another variable.

---

# 📌 Real Life Example

```go
age := 24
```

This variable is stored somewhere in memory.

A pointer stores:
👉 Address of `age`

---

# 📌 Pointer Symbols

| Symbol | Meaning |
|---|---|
| & | Address of variable |
| * | Value at address |

---

# 📌 Get Memory Address

```go
package main

import "fmt"

func main() {

	age := 24

	fmt.Println(age)

	fmt.Println(&age)
}
```

---

# 📌 Output

```text
24
0xc0000120a0
```

---

# 📌 What is This?

```text
0xc0000120a0
```

This is:
# ✅ Memory Address

---

# 📌 Pointer Variable

```go
package main

import "fmt"

func main() {

	age := 24

	var ptr *int

	ptr = &age

	fmt.Println(ptr)
}
```

---

# 📌 Explanation

| Code | Meaning |
|---|---|
| *int | Pointer to integer |
| &age | Address of age |
| ptr | Stores address |

---

# 📌 Dereferencing Pointer

Used to access actual value.

---

# ✅ Example

```go
package main

import "fmt"

func main() {

	age := 24

	ptr := &age

	fmt.Println(*ptr)
}
```

---

# 📌 Output

```text
24
```

---

# 📌 Understanding

| Expression | Meaning |
|---|---|
| ptr | Memory Address |
| *ptr | Actual Value |

---

# 📌 Modify Original Value Using Pointer

```go
package main

import "fmt"

func main() {

	age := 24

	ptr := &age

	*ptr = 30

	fmt.Println(age)
}
```

---

# 📌 Output

```text
30
```

---

# 📌 Why Value Changed?

Because pointer directly modified original variable.

Very important concept 🔥

---

# 📌 Pass by Value

Without pointer:
A copy of variable is sent.

---

# ❌ Example

```go
package main

import "fmt"

func update(age int) {

	age = 50
}

func main() {

	age := 24

	update(age)

	fmt.Println(age)
}
```

---

# 📌 Output

```text
24
```

---

# 📌 Why?

Because:
👉 Copy was modified, not original value.

---

# 📌 Pass by Pointer

Using pointer:
Original value gets modified.

---

# ✅ Example

```go
package main

import "fmt"

func update(age *int) {

	*age = 50
}

func main() {

	age := 24

	update(&age)

	fmt.Println(age)
}
```

---

# 📌 Output

```text
50
```

---

# 📌 Why Pointers are Important?

Pointers:
✅ Avoid unnecessary copying  
✅ Improve performance  
✅ Reduce memory usage  
✅ Modify original values

---

# 📌 Pointer with Struct

Very important for backend development.

---

# ✅ Example

```go
package main

import "fmt"

type Student struct {

	Name string
	Age  int
}

func updateStudent(student *Student) {

	student.Name = "Rahul"
	student.Age = 30
}

func main() {

	student := Student{

		Name: "Dnyaneshwar",
		Age: 24,
	}

	updateStudent(&student)

	fmt.Println(student)
}
```

---

# 📌 Why Struct Pointers Important?

In real projects:
- Large structs
- API models
- Database entities

Copying them repeatedly is expensive.

So Go developers mostly use:
# ✅ Struct Pointers

---

# 📌 new() Function

Used to create pointer.

---

# ✅ Example

```go
package main

import "fmt"

func main() {

	ptr := new(int)

	*ptr = 100

	fmt.Println(*ptr)
}
```

---

# 📌 Nil Pointer

Pointer with no assigned value.

```go
var ptr *int
```

Default value:
```text
nil
```

---

# 📌 Important Warning

Never dereference nil pointer.

❌ Wrong:

```go
fmt.Println(*ptr)
```

This causes runtime error.

---

# 💻 Day 05 Practice Program

```go
package main

import "fmt"

type Employee struct {

	Name   string
	Salary int
}

func updateSalary(emp *Employee) {

	emp.Salary = 70000
}

func main() {

	age := 24

	ptr := &age

	fmt.Println("Age:", age)
	fmt.Println("Address:", ptr)
	fmt.Println("Value using pointer:", *ptr)

	*ptr = 30

	fmt.Println("Updated Age:", age)

	employee := Employee{

		Name: "Dnyaneshwar",
		Salary: 50000,
	}

	updateSalary(&employee)

	fmt.Println(employee)
}
```

---

# 📘 Day 05 Interview Questions & Answers

---

## ❓ Q1: What is pointer in Go?

### ✅ Answer:
Pointer stores memory address of another variable.

---

## ❓ Q2: What does & operator do?

### ✅ Answer:
Returns memory address of variable.

---

## ❓ Q3: What does * operator do?

### ✅ Answer:
Used for:
- Dereferencing pointer
- Accessing actual value

---

## ❓ Q4: Why pointers are important?

### ✅ Answer:
Pointers:
- Reduce memory usage
- Improve performance
- Modify original data

---

## ❓ Q5: Difference between pass by value and pass by pointer?

| Pass by Value | Pass by Pointer |
|---|---|
| Copy sent | Address sent |
| Original unchanged | Original modified |

---

## ❓ Q6: What is nil pointer?

### ✅ Answer:
Pointer with no assigned memory address.

---

## ❓ Q7: Why pointers used with structs?

### ✅ Answer:
To avoid copying large struct data and improve performance.

---

# 📚 Day 05 Summary

Today I learned:
- Pointers
- Memory Address
- Dereferencing
- Pass by Value
- Pass by Pointer
- Struct Pointers
- new() Function
- Nil Pointers

I also practiced:
- Updating original values
- Struct updates using pointers
- Memory handling

---

# 🧠 Practice Tasks

✅ Create integer pointer  
✅ Print memory address  
✅ Modify value using pointer  
✅ Pass pointer to function  
✅ Create struct pointer  
✅ Update struct using pointer  
✅ Create Employee update function

---

# ✅ Day 06 — Methods & Interfaces

---

# 📖 Introduction to Methods & Interfaces

Methods and Interfaces are one of the most important concepts in Go backend development.

They are heavily used in:
- Backend APIs
- Clean Architecture
- Microservices
- Business Logic
- Abstraction
- Scalable Applications

If you know C#, then:
- Method in Go ≈ Class Method
- Interface in Go ≈ Interface in C#

---

# 📖 What You Will Learn

- Methods
- Method Receivers
- Value Receiver
- Pointer Receiver
- Interfaces
- Empty Interface
- Type Assertion
- Real Backend Examples
- Interview Questions

---

# 📌 What is Method?

A method is a function attached to a struct.

---

# 📌 Difference Between Function & Method

| Function | Method |
|---|---|
| Independent | Attached to struct |
| Called directly | Called using object |

---

# 📌 Simple Function Example

```go
func greet() {

	fmt.Println("Hello")
}
```

---

# 📌 Method Syntax

```go
func (receiver StructName) methodName() {

}
```

---

# 📌 Method Example

```go
package main

import "fmt"

type Student struct {

	Name string
	Age  int
}

func (s Student) display() {

	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
}

func main() {

	student := Student{

		Name: "Dnyaneshwar",
		Age: 24,
	}

	student.display()
}
```

---

# 📌 Explanation

| Code | Meaning |
|---|---|
| (s Student) | Receiver |
| display() | Method |
| s.Name | Access struct field |

---

# 📌 Why Methods are Important?

Methods help:
✅ Organize code  
✅ Add behavior to structs  
✅ Improve readability  
✅ Build scalable backend systems

---

# 📌 Value Receiver

Value receiver receives COPY of struct.

---

# ✅ Example

```go
package main

import "fmt"

type Employee struct {

	Name string
	Salary int
}

func (e Employee) updateSalary() {

	e.Salary = 70000
}

func main() {

	employee := Employee{

		Name: "Dnyaneshwar",
		Salary: 50000,
	}

	employee.updateSalary()

	fmt.Println(employee.Salary)
}
```

---

# 📌 Output

```text
50000
```

---

# 📌 Why Salary Not Changed?

Because:
👉 Value receiver receives COPY.

Original value remains unchanged.

---

# 📌 Pointer Receiver

Pointer receiver modifies original struct.

---

# ✅ Example

```go
package main

import "fmt"

type Employee struct {

	Name   string
	Salary int
}

func (e *Employee) updateSalary() {

	e.Salary = 70000
}

func main() {

	employee := Employee{
		Name: "Dnyaneshwar",
		Salary: 50000,
	}

	employee.updateSalary()

	fmt.Println(employee.Salary)
}
```

---

# 📌 Output

```text
70000
```

---

# 📌 Why Pointer Receivers Important?

Pointer receivers:
✅ Modify original value  
✅ Improve performance  
✅ Avoid copying large structs

Go developers mostly use:
# ✅ Pointer Receivers

---

# 📌 What is Interface?

Interface defines behavior.

Interface contains:
# ✅ Method Signatures

---

# 📌 Interface Syntax

```go
type Shape interface {

	area()
}
```

---

# 📌 Interface Example

```go
package main

import "fmt"

type Rectangle struct {

	Width  int
	Height int
}

func (r Rectangle) area() {

	fmt.Println("Area:", r.Width*r.Height)
}

type Shape interface {

	area()
}

func main() {

	rect := Rectangle{

		Width: 10,
		Height: 20,
	}

	var shape Shape

	shape = rect

	shape.area()
}
```

---

# 📌 Understanding

| Part | Meaning |
|---|---|
| interface | Defines behavior |
| area() | Method signature |
| Rectangle | Implements interface |

---

# 📌 Important Point

Go interfaces are:
# ✅ Implicit

No need for:
```text
implements keyword
```

Unlike Java or C#.

---

# 📌 Real Backend Example

Suppose:
We have payment system.

Different payment methods:
- UPI
- Card
- PayPal

All should implement:
```go
pay()
```

This is achieved using interfaces.

---

# ✅ Payment Interface Example

```go
package main

import "fmt"

type Payment interface {

	pay()
}

type UPI struct {}

func (u UPI) pay() {

	fmt.Println("Paid using UPI")
}

type Card struct {}

func (c Card) pay() {

	fmt.Println("Paid using Card")
}

func main() {

	var payment Payment

	payment = UPI{}
	payment.pay()

	payment = Card{}
	payment.pay()
}
```

---

# 📌 Empty Interface

```go
interface{}
```

Can store ANY datatype.

---

# ✅ Example

```go
package main

import "fmt"

func display(value interface{}) {

	fmt.Println(value)
}

func main() {

	display(10)
	display("Go")
	display(true)
}
```

---

# 📌 Why Empty Interface Important?

Used in:
- JSON handling
- Generic data
- Dynamic APIs

---

# 📌 Type Assertion

Used to get original datatype from interface.

---

# ✅ Example

```go
package main

import "fmt"

func main() {

	var value interface{} = "Go Lang"

	str := value.(string)

	fmt.Println(str)
}
```

---

# 💻 Day 06 Practice Program

```go
package main

import "fmt"

type Student struct {

	Name string
	Age  int
}

func (s Student) display() {

	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
}

type Employee struct {

	Name   string
	Salary int
}

func (e *Employee) updateSalary() {

	e.Salary = 70000
}

type Payment interface {

	pay()
}

type UPI struct {}

func (u UPI) pay() {

	fmt.Println("Payment using UPI")
}

func main() {

	student := Student{

		Name: "Dnyaneshwar",
		Age: 24,
	}

	student.display()

	employee := Employee{

		Name: "Rahul",
		Salary: 50000,
	}

	employee.updateSalary()

	fmt.Println(employee)

	var payment Payment

	payment = UPI{}

	payment.pay()
}
```

---

# 📘 Day 06 Interview Questions & Answers

---

## ❓ Q1: What is method in Go?

### ✅ Answer:
Method is a function attached to a struct.

---

## ❓ Q2: Difference between function and method?

| Function | Method |
|---|---|
| Independent | Attached to struct |

---

## ❓ Q3: What is receiver in Go?

### ✅ Answer:
Receiver connects method with struct.

---

## ❓ Q4: Difference between value receiver and pointer receiver?

| Value Receiver | Pointer Receiver |
|---|---|
| Copy passed | Original address passed |
| Cannot modify original | Can modify original |

---

## ❓ Q5: What is interface in Go?

### ✅ Answer:
Interface defines behavior using method signatures.

---

## ❓ Q6: Why interfaces are important?

### ✅ Answer:
Interfaces help achieve:
- Abstraction
- Loose coupling
- Clean architecture

---

## ❓ Q7: What is empty interface?

### ✅ Answer:
`interface{}` can store any datatype.

---

## ❓ Q8: What is type assertion?

### ✅ Answer:
Used to retrieve original datatype from interface.

---

# 📚 Day 06 Summary

Today I learned:
- Methods
- Receivers
- Pointer Receivers
- Interfaces
- Empty Interface
- Type Assertion

I also practiced:
- Struct methods
- Interface implementation
- Backend-style abstraction

---

# 🧠 Practice Tasks

✅ Create Student method  
✅ Create Employee pointer receiver  
✅ Create Payment interface  
✅ Create Card & UPI implementations  
✅ Use empty interface  
✅ Practice type assertion

---

# ✅ Day 07 — Mini Project in Go

---

# 📖 Introduction to Mini Project

Today we built our first real Go project.

This project combines:
- Structs
- Functions
- Slices
- Loops
- Conditions
- Switch Case
- CRUD Operations

This is very important because:
👉 Same concepts are used in real backend applications and APIs.

---

# 🎯 Project Name

# ✅ Student Management System

---

# 📖 What You Will Learn

- Real Project Structure
- CRUD Operations
- Struct Handling
- Slice Manipulation
- Menu-driven Programs
- Backend Thinking
- User Input Handling

---

# 📌 Features

✅ Add Student  
✅ View Students  
✅ Search Student  
✅ Update Student  
✅ Delete Student  
✅ Exit Program

---

# 📌 Project Structure

```text
Day-07/
 ├── main.go
 ├── README.md
```

---

# 📌 Create Project Folder

```bash
mkdir Day-07
cd Day-07
touch main.go
```

---

# 📌 Step 1 — Create Student Struct

Struct is used to store student information.

```go
type Student struct {

	ID   int
	Name string
	Age  int
	City string
}
```

---

# 📌 Explanation

| Field | Purpose |
|---|---|
| ID | Unique student ID |
| Name | Student name |
| Age | Student age |
| City | Student city |

---

# 📌 Step 2 — Create Slice

Slice stores multiple students dynamically.

```go
var students []Student
```

---

# 📌 Why Slice Used?

Because:
- Student count changes dynamically
- Slice can grow automatically

---

# 📌 Step 3 — Add Student Function

This function takes user input and stores student data.

```go
func addStudent() {

	var student Student

	fmt.Print("Enter ID: ")
	fmt.Scan(&student.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&student.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&student.Age)

	fmt.Print("Enter City: ")
	fmt.Scan(&student.City)

	students = append(students, student)

	fmt.Println("✅ Student Added Successfully")
}
```

---

# 📌 Concepts Used

| Concept | Used |
|---|---|
| Struct | Store student data |
| fmt.Scan | User input |
| append() | Add student dynamically |

---

# 📌 Step 4 — View Students Function

Used to display all students.

```go
func viewStudents() {

	if len(students) == 0 {

		fmt.Println("❌ No Students Found")
		return
	}

	fmt.Println("\n📚 Student List")

	for _, student := range students {

		fmt.Println("----------------------")
		fmt.Println("ID:", student.ID)
		fmt.Println("Name:", student.Name)
		fmt.Println("Age:", student.Age)
		fmt.Println("City:", student.City)
	}
}
```

---

# 📌 Concepts Used

| Concept | Used |
|---|---|
| len() | Check slice empty or not |
| range | Loop through students |
| Conditions | Validation |

---

# 📌 Step 5 — Search Student Function

Search student using ID.

```go
func searchStudent() {

	var id int

	fmt.Print("Enter Student ID: ")
	fmt.Scan(&id)

	for _, student := range students {

		if student.ID == id {

			fmt.Println("✅ Student Found")
			fmt.Println("Name:", student.Name)
			fmt.Println("Age:", student.Age)
			fmt.Println("City:", student.City)

			return
		}
	}

	fmt.Println("❌ Student Not Found")
}
```

---

# 📌 Search Logic

```text
Loop through all students
↓
Compare IDs
↓
If matched → show student
Else → not found
```

---

# 📌 Step 6 — Update Student Function

Update existing student data.

```go
func updateStudent() {

	var id int

	fmt.Print("Enter Student ID to Update: ")
	fmt.Scan(&id)

	for index, student := range students {

		if student.ID == id {

			fmt.Print("Enter New Name: ")
			fmt.Scan(&students[index].Name)

			fmt.Print("Enter New Age: ")
			fmt.Scan(&students[index].Age)

			fmt.Print("Enter New City: ")
			fmt.Scan(&students[index].City)

			fmt.Println("✅ Student Updated Successfully")

			return
		}
	}

	fmt.Println("❌ Student Not Found")
}
```

---

# 📌 Concepts Used

| Concept | Used |
|---|---|
| range | Get index |
| Slice update | Modify student |
| Conditions | Match ID |

---

# 📌 Step 7 — Delete Student Function

Delete student using ID.

```go
func deleteStudent() {

	var id int

	fmt.Print("Enter Student ID to Delete: ")
	fmt.Scan(&id)

	for index, student := range students {

		if student.ID == id {

			students = append(students[:index], students[index+1:]...)

			fmt.Println("✅ Student Deleted Successfully")

			return
		}
	}

	fmt.Println("❌ Student Not Found")
}
```

---

# 📌 Important Delete Logic

```go
students = append(students[:index], students[index+1:]...)
```

---

# 📌 What is Happening?

```text
Before Index + After Index
↓
Merged Together
↓
Current Student Removed
```

---

# 📌 Step 8 — Menu-driven Program

Used to continuously run program.

```go
for {

	fmt.Println("\n===== Student Management System =====")
	fmt.Println("1. Add Student")
	fmt.Println("2. View Students")
	fmt.Println("3. Search Student")
	fmt.Println("4. Update Student")
	fmt.Println("5. Delete Student")
	fmt.Println("6. Exit")
}
```

---

# 📌 Step 9 — Switch Case

Used to handle user choices.

```go
switch choice {

case 1:
	addStudent()

case 2:
	viewStudents()

case 3:
	searchStudent()

case 4:
	updateStudent()

case 5:
	deleteStudent()

case 6:
	fmt.Println("🚀 Exiting Program")
	return

default:
	fmt.Println("❌ Invalid Choice")
}
```

---

# 💻 Full Project Code — main.go

```go
package main

import "fmt"

// Struct
type Student struct {
	ID    int
	Name  string
	Age   int
	City  string
}

var students []Student

// Add Student
func addStudent() {

	var student Student

	fmt.Print("Enter ID: ")
	fmt.Scan(&student.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&student.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&student.Age)

	fmt.Print("Enter City: ")
	fmt.Scan(&student.City)

	students = append(students, student)

	fmt.Println("✅ Student Added Successfully")
}

// View Students
func viewStudents() {

	if len(students) == 0 {

		fmt.Println("❌ No Students Found")
		return
	}

	fmt.Println("\n📚 Student List")

	for _, student := range students {

		fmt.Println("----------------------")
		fmt.Println("ID:", student.ID)
		fmt.Println("Name:", student.Name)
		fmt.Println("Age:", student.Age)
		fmt.Println("City:", student.City)
	}
}

// Search Student
func searchStudent() {

	var id int

	fmt.Print("Enter Student ID: ")
	fmt.Scan(&id)

	for _, student := range students {

		if student.ID == id {

			fmt.Println("✅ Student Found")
			fmt.Println("Name:", student.Name)
			fmt.Println("Age:", student.Age)
			fmt.Println("City:", student.City)

			return
		}
	}

	fmt.Println("❌ Student Not Found")
}

// Delete Student
func deleteStudent() {

	var id int

	fmt.Print("Enter Student ID to Delete: ")
	fmt.Scan(&id)

	for index, student := range students {

		if student.ID == id {

			students = append(students[:index], students[index+1:]...)

			fmt.Println("✅ Student Deleted Successfully")

			return
		}
	}

	fmt.Println("❌ Student Not Found")
}

// Update Student
func updateStudent() {

	var id int

	fmt.Print("Enter Student ID to Update: ")
	fmt.Scan(&id)

	for index, student := range students {

		if student.ID == id {

			fmt.Print("Enter New Name: ")
			fmt.Scan(&students[index].Name)

			fmt.Print("Enter New Age: ")
			fmt.Scan(&students[index].Age)

			fmt.Print("Enter New City: ")
			fmt.Scan(&students[index].City)

			fmt.Println("✅ Student Updated Successfully")

			return
		}
	}

	fmt.Println("❌ Student Not Found")
}

func main() {

	for {

		fmt.Println("\n===== Student Management System =====")
		fmt.Println("1. Add Student")
		fmt.Println("2. View Students")
		fmt.Println("3. Search Student")
		fmt.Println("4. Update Student")
		fmt.Println("5. Delete Student")
		fmt.Println("6. Exit")

		var choice int

		fmt.Print("Enter Your Choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			addStudent()

		case 2:
			viewStudents()

		case 3:
			searchStudent()

		case 4:
			updateStudent()

		case 5:
			deleteStudent()

		case 6:
			fmt.Println("🚀 Exiting Program")
			return

		default:
			fmt.Println("❌ Invalid Choice")
		}
	}
}
```

---

# 📌 Run Project

```bash
go run main.go
```

---

# 📌 Expected Output

```text
===== Student Management System =====

1. Add Student
2. View Students
3. Search Student
4. Update Student
5. Delete Student
6. Exit
```

---

# 📌 Backend Understanding

This project is important because:
👉 Same logic is used in:
- REST APIs
- Database systems
- Backend applications

Later:
- Slice → Database
- Console → API Endpoints

---

# 📘 Day 07 Interview Questions & Answers

---

## ❓ Q1: Why slice used in this project?

### ✅ Answer:
Because data size changes dynamically.

---

## ❓ Q2: Why struct used?

### ✅ Answer:
Struct groups related student data together.

---

## ❓ Q3: Which CRUD operations implemented?

### ✅ Answer:
- Create
- Read
- Update
- Delete

---

## ❓ Q4: Why functions used?

### ✅ Answer:
Functions improve code reusability and readability.

---

## ❓ Q5: What is append() used for?

### ✅ Answer:
Used to add student into slice dynamically.

---

## ❓ Q6: Why range loop used?

### ✅ Answer:
Used to iterate through students slice.

---

## ❓ Q7: Why switch case used?

### ✅ Answer:
Used to handle menu choices efficiently.

---

# 📚 Day 07 Summary

Today I learned:
- Real Go project structure
- CRUD operations
- Struct handling
- Slice manipulation
- Backend logic

I also practiced:
- Add Student
- View Student
- Update Student
- Delete Student
- Search Student

---
# 🧠 Practice Tasks

✅ Add Email field  
✅ Add Marks field  
✅ Add Multiple Students  
✅ Add Exit Confirmation  
✅ Add Validation  
✅ Prevent Duplicate IDs
---

# ✅ Day 08 — Packages & Modules in Go

---

# 📖 Introduction to Packages & Modules

Packages and Modules are one of the most important concepts in Go backend development.

Real-world Go applications are divided into:
- Packages
- Modules
- Layers
- Services

Without packages:
❌ Code becomes messy  
❌ Difficult to maintain  
❌ Hard to scale

Using packages:
✅ Clean architecture  
✅ Better organization  
✅ Reusable code  
✅ Production-ready structure

---

# 📖 What You Will Learn

- Packages
- Main Package
- Custom Packages
- Modules
- go.mod
- Importing Packages
- Exported Functions
- Package Alias
- Backend Project Structure
- Interview Questions

---

# 📌 What is Package?

A package is a collection of Go files.

Packages help organize code into reusable components.

---

# 📌 Real Example

Suppose you are building backend project.

You can separate:
- User logic
- Product logic
- Authentication
- Database logic

into different packages.

---

# 📌 Package Syntax

```go
package main
```

---

# 📌 Important Rule

Every Go file must belong to a package.

---

# 📌 Main Package

```go
package main
```

This is a special package.

Program execution starts from:
```go
func main()
```

---

# 📌 Simple Example

```go
package main

import "fmt"

func main() {

	fmt.Println("Hello Go")
}
```

---

# 📌 Custom Package

We can create our own packages.

---

# 📌 Project Structure

```text
Day-08/
 ├── main.go
 ├── calculator/
 │    └── calculator.go
 ├── go.mod
```

---

# 📌 Step 1 — Create Project Folder

```bash
mkdir Day-08
cd Day-08
```

---

# 📌 Step 2 — Initialize Module

```bash
go mod init day08
```

---

# 📌 What is go.mod?

`go.mod` manages:
- Project dependencies
- Module name
- Go version

---

# 📌 Example go.mod

```go
module day08

go 1.24.0
```

---

# 📌 Step 3 — Create Package Folder

```bash
mkdir calculator
touch calculator/calculator.go
```

---

# 📌 calculator.go

```go
package calculator

func Add(a int, b int) int {

	return a + b
}

func Subtract(a int, b int) int {

	return a - b
}
```

---

# 📌 Important Point

Functions starting with CAPITAL letter are:
# ✅ Exported

Accessible outside package.

---

# 📌 Lowercase Functions

```go
func add()
```

❌ Not accessible outside package.

---

# 📌 Step 4 — main.go

```go
package main

import (
	"fmt"
	"day08/calculator"
)

func main() {

	result := calculator.Add(10, 20)

	fmt.Println("Addition:", result)

	sub := calculator.Subtract(30, 10)

	fmt.Println("Subtraction:", sub)
}
```

---

# 📌 Understanding Import

```go
import "day08/calculator"
```

Meaning:
👉 Import calculator package from module.

---

# 📌 Run Project

```bash
go run .
```

---

# 📌 Output

```text
Addition: 30
Subtraction: 20
```

---

# 📌 Why Packages are Important?

Packages help:
✅ Organize code  
✅ Reuse code  
✅ Separate business logic  
✅ Build scalable backend systems

---

# 📌 Real Backend Structure

Real Go backend projects look like:

```text
project/
 ├── cmd/
 ├── handlers/
 ├── services/
 ├── repositories/
 ├── models/
 ├── routes/
 ├── database/
 ├── go.mod
 └── main.go
```

---

# 📌 Package Alias

We can rename imported package.

---

# ✅ Example

```go
import calc "day08/calculator"
```

Usage:

```go
calc.Add(10,20)
```

---

# 📌 Blank Identifier Import

Used when package initialization needed only.

```go
import _ "package_name"
```

Used in:
- Database drivers
- Side effects

---

# 📌 Multiple Packages

One project can contain multiple packages.

Example:
- auth
- user
- product
- payment

---

# 📌 Important Rule

All files in same folder must belong to same package.

---

# ❌ Wrong

```text
folder/
 ├── package user
 ├── package product
```

---

# ✅ Correct

```text
user/
product/
```

---

# 📌 Backend Understanding

In real backend projects:
- Package = Layer
- Module = Project

Example:
- controllers package
- services package
- repositories package

Exactly like:
- ASP.NET Core Layers
- Clean Architecture

---

# 💻 Day 08 Practice Program

# 📌 calculator/calculator.go

```go
package calculator

func Add(a int, b int) int {

	return a + b
}

func Multiply(a int, b int) int {

	return a * b
}
```

---

# 📌 main.go

```go
package main

import (
	"fmt"
	"day08/calculator"
)

func main() {

	sum := calculator.Add(10, 20)

	multiply := calculator.Multiply(5, 5)

	fmt.Println("Addition:", sum)

	fmt.Println("Multiplication:", multiply)
}
```

---

# 📘 Day 08 Interview Questions & Answers

---

## ❓ Q1: What is package in Go?

### ✅ Answer:
Package is a collection of Go files used to organize code.

---

## ❓ Q2: What is main package?

### ✅ Answer:
Main package is special package where program execution starts.

---

## ❓ Q3: What is module in Go?

### ✅ Answer:
Module is collection of Go packages managed using go.mod.

---

## ❓ Q4: What is go.mod?

### ✅ Answer:
go.mod manages:
- Dependencies
- Module name
- Go version

---

## ❓ Q5: Why exported functions are important?

### ✅ Answer:
Exported functions can be accessed outside package.

---

## ❓ Q6: Difference between Add() and add()?

| Add() | add() |
|---|---|
| Exported | Unexported |
| Public | Private |

---

## ❓ Q7: Why packages are important?

### ✅ Answer:
Packages improve:
- Scalability
- Reusability
- Code organization

---

# 📚 Day 08 Summary

Today I learned:
- Packages
- Modules
- go.mod
- Importing packages
- Exported functions
- Backend project structure

I also practiced:
- Creating custom packages
- Importing packages
- Using modules
- Organizing backend code

---
# 🧠 Practice Tasks

✅ Create math package  
✅ Create auth package  
✅ Create user package  
✅ Add multiple functions  
✅ Use package alias  
✅ Create multi-package project
---

# ✅ Day 09 — File Handling in Go

---

# 📖 Introduction to File Handling

File handling is one of the most important concepts in backend development.

Real applications use file handling for:
- Logs
- Reports
- CSV Exports
- Upload Systems
- Config Files
- Data Storage

Using Go, we can:
✅ Create Files  
✅ Read Files  
✅ Write Files  
✅ Append Data  
✅ Delete Files  
✅ Rename Files

---

# 📖 What You Will Learn

- Create File
- Write File
- Read File
- Append Data
- Delete File
- Rename File
- File Information
- os Package
- defer Keyword
- Real Backend Examples
- Interview Questions

---

# 📌 What is File Handling?

File handling means:
👉 Creating and managing files using code.

---

# 📌 Important Packages

| Package | Purpose |
|---|---|
| os | File operations |
| fmt | Printing |
| bufio | Buffered reading |
| io/ioutil | Read/write helpers |

---

# 📌 Create File

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("student.txt")

	if err != nil {

		fmt.Println("Error:", err)
		return
	}

	fmt.Println("✅ File Created Successfully")

	file.Close()
}
```

---

# 📌 Write Data into File

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("student.txt")

	if err != nil {

		fmt.Println(err)
		return
	}

	defer file.Close()

	file.WriteString("Welcome to Go Lang")

	fmt.Println("✅ Data Written Successfully")
}
```

---

# 📌 Read File

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("student.txt")

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
```

---

# 📌 Append Data into File

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile(
		"student.txt",
		os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if err != nil {

		fmt.Println(err)
		return
	}

	defer file.Close()

	file.WriteString("\nLearning Go Lang")

	fmt.Println("✅ Data Appended")
}
```

---

# 📌 Delete File

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	err := os.Remove("student.txt")

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("✅ File Deleted")
}
```

---

# 📌 Rename File

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	err := os.Rename(
		"student.txt",
		"data.txt",
	)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("✅ File Renamed")
}
```

---

# 💻 Day 09 Practice Program

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("notes.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	file.WriteString("Learning File Handling in Go")

	data, err := os.ReadFile("notes.txt")

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	os.Rename("notes.txt", "golang.txt")

	fmt.Println("✅ File Renamed")
}
```

---

# 📘 Day 09 Interview Questions & Answers

---

## ❓ Q1: What is file handling?

### ✅ Answer:
File handling means creating, reading, updating, and deleting files using code.

---

## ❓ Q2: Which package is mainly used for file handling in Go?

### ✅ Answer:
`os` package.

---

## ❓ Q3: Why file.Close() important?

### ✅ Answer:
Used to release system resources and avoid memory leaks.

---

## ❓ Q4: What is defer in Go?

### ✅ Answer:
`defer` delays function execution until surrounding function ends.

---

## ❓ Q5: What does os.Create() do?

### ✅ Answer:
Creates new file.

---

# 📚 Day 09 Summary

Today I learned:
- File handling
- Create file
- Read file
- Write file
- Append file
- Delete file
- Rename file

---

# 🧠 Practice Tasks

✅ Create student.txt  
✅ Write student data  
✅ Read student file  
✅ Append new data  
✅ Rename file  
✅ Delete file
---

# ✅ Day 10 — Error Handling in Go

---

# 📖 Introduction to Error Handling

Error handling is one of the most important concepts in backend development.

In real applications:
- APIs fail
- Database connections fail
- Files may not exist
- User input can be invalid

Good developers:
✅ Handle errors properly  
✅ Prevent crashes  
✅ Return meaningful messages

That is why:
# ✅ Error handling is critical in Go

---

# 📖 What You Will Learn

- Errors in Go
- error Interface
- if err != nil
- Custom Errors
- errors.New()
- Panic
- Recover
- Logging Errors
- Best Practices
- Real Backend Examples
- Interview Questions

---

# 📌 What is Error?

Error means:
👉 Something went wrong during program execution.

Examples:
- File not found
- Database connection failed
- Invalid input
- API request failed

---

# 📌 Go Philosophy

Go does NOT use:
```text
try-catch
```

Instead Go uses:
# ✅ Explicit Error Handling

---

# 📌 Basic Error Handling Syntax

```go
result, err := someFunction()

if err != nil {

	fmt.Println(err)
	return
}
```

---

# 📌 Understanding

| Part | Meaning |
|---|---|
| err | Error object |
| nil | No error |
| if err != nil | Error exists |

---

# 📌 Simple Example

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("student.txt")

	if err != nil {

		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	fmt.Println("✅ File Opened Successfully")
}
```

---

# 📌 Output (If File Missing)

```text
Error: open student.txt: no such file or directory
```

---

# 📌 Why Error Handling Important?

Without error handling:
❌ Program crashes  
❌ Bad user experience  
❌ Difficult debugging

With error handling:
✅ Stable applications  
✅ Better debugging  
✅ Production-ready systems

---

# 📌 Custom Errors

Go allows creating custom errors.

---

# ✅ Example

```go
package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) error {

	if age < 18 {

		return errors.New("age must be 18 or above")
	}

	return nil
}

func main() {

	err := checkAge(16)

	if err != nil {

		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Eligible")
}
```

---

# 📌 errors.New()

Used to create custom error message.

---

# 📌 Return nil

```go
return nil
```

Means:
# ✅ No Error

---

# 📌 Multiple Return Values

Go functions commonly return:

```go
value, err
```

---

# ✅ Example

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	number, err := strconv.Atoi("100")

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(number)
}
```

---

# 📌 What is Happening?

```text
"100" → converted into integer
```

---

# 📌 Panic in Go

`panic` stops program immediately.

---

# ✅ Example

```go
package main

func main() {

	panic("Something went wrong")
}
```

---

# 📌 Output

```text
panic: Something went wrong
```

---

# 📌 When Panic Used?

Used for:
- Critical failures
- Impossible situations
- Serious system errors

Examples:
- Database corruption
- Server startup failure

---

# 📌 Recover in Go

`recover()` handles panic and prevents crash.

---

# ✅ Example

```go
package main

import "fmt"

func handlePanic() {

	recoverMessage := recover()

	if recoverMessage != nil {

		fmt.Println("Recovered:", recoverMessage)
	}
}

func main() {

	defer handlePanic()

	panic("Server crashed")
}
```

---

# 📌 Output

```text
Recovered: Server crashed
```

---

# 📌 defer + recover

Mostly used together.

---

# 📌 Real Backend Example

Suppose:
- Database connection fails
- API crashes

Instead of crashing server:
✅ Recover handles panic.

Very important in production systems.

---

# 📌 Logging Errors

Backend developers log errors for debugging.

---

# ✅ Example

```go
package main

import (
	"log"
)

func main() {

	log.Println("Server Started")
}
```

---

# 📌 Difference Between fmt and log

| fmt | log |
|---|---|
| Simple output | Logging |
| No timestamp | Includes timestamp |

---

# 📌 Best Practices

✅ Always check errors  
✅ Never ignore errors  
✅ Use meaningful messages  
✅ Return errors properly  
✅ Use panic only for critical issues

---

# ❌ Wrong Practice

```go
file, _ := os.Open("test.txt")
```

Ignoring error is bad practice.

---

# ✅ Correct Practice

```go
file, err := os.Open("test.txt")

if err != nil {

	fmt.Println(err)
	return
}
```

---

# 💻 Day 10 Practice Program

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

func checkNumber(number int) error {

	if number < 0 {

		return errors.New("negative number not allowed")
	}

	return nil
}

func main() {

	err := checkNumber(-10)

	if err != nil {

		fmt.Println("Error:", err)
		return
	}

	file, fileError := os.Open("data.txt")

	if fileError != nil {

		fmt.Println("File Error:", fileError)
		return
	}

	defer file.Close()

	fmt.Println("✅ Program Executed Successfully")
}
```

---

# 📘 Day 10 Interview Questions & Answers

---

## ❓ Q1: What is error handling?

### ✅ Answer:
Error handling is process of managing runtime problems in program.

---

## ❓ Q2: How Go handles errors?

### ✅ Answer:
Go uses explicit error handling with:

```go
if err != nil
```

---

## ❓ Q3: What is error interface?

### ✅ Answer:
Error interface represents runtime error in Go.

---

## ❓ Q4: What is errors.New()?

### ✅ Answer:
Used to create custom error.

---

## ❓ Q5: What is panic?

### ✅ Answer:
panic stops normal execution of program immediately.

---

## ❓ Q6: What is recover?

### ✅ Answer:
recover handles panic and prevents program crash.

---

## ❓ Q7: Difference between panic and error?

| error | panic |
|---|---|
| Recoverable | Critical failure |
| Normal flow | Stops execution |

---

## ❓ Q8: Why error handling important in backend?

### ✅ Answer:
Prevents crashes and improves application stability.

---

# 📚 Day 10 Summary

Today I learned:
- Error handling
- Custom errors
- errors.New()
- panic
- recover
- Logging
- Backend error management

I also practiced:
- File errors
- Input validation
- Panic recovery
- Production-style error handling

---

# 🧠 Practice Tasks

✅ Create divide function with error  
✅ Validate age input  
✅ Handle missing file  
✅ Create custom errors  
✅ Use panic & recover  
✅ Create log messages

---

# ✅ Day 11 — Goroutines & Concurrency in Go

---

# 📖 Introduction to Concurrency

Concurrency is one of the biggest reasons why Go is famous.

Go is widely used for:
- High-performance APIs
- Microservices
- Real-time systems
- Backend servers
- Cloud-native applications

Because Go provides:
# ✅ Goroutines

They are:
- Lightweight
- Fast
- Efficient

---

# 📖 What You Will Learn

- Concurrency
- Goroutines
- go Keyword
- Concurrent Execution
- Anonymous Goroutines
- sync.WaitGroup
- Sleep
- Real Backend Examples
- Interview Questions

---

# 📌 What is Concurrency?

Concurrency means:
👉 Multiple tasks running independently.

Examples:
- Sending email
- Processing payment
- Saving database
- Generating logs

All can run simultaneously.

---

# 📌 Why Go Famous for Concurrency?

Because Go provides:
# ✅ Goroutines

Goroutines are lightweight threads managed by Go runtime.

---

# 📌 What is Goroutine?

Goroutine is lightweight thread managed by Go runtime.

---

# 📌 Normal Function Example

```go
package main

import "fmt"

func printNumbers() {

	for i := 1; i <= 5; i++ {

		fmt.Println(i)
	}
}

func main() {

	printNumbers()

	fmt.Println("Main Function")
}
```

---

# 📌 Goroutine Syntax

```go
go functionName()
```

---

# 📌 First Goroutine Example

```go
package main

import (
	"fmt"
	"time"
)

func printNumbers() {

	for i := 1; i <= 5; i++ {

		fmt.Println(i)

		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	go printNumbers()

	time.Sleep(time.Second * 3)

	fmt.Println("Main Function")
}
```

---

# 📌 What Happened?

```text
printNumbers() runs concurrently
```

---

# 📌 Why time.Sleep Used?

Because:
👉 Main function exits quickly.

If main exits:
❌ Goroutines stop immediately.

---

# 📌 Important Rule

Main function must stay alive until goroutines finish.

---

# 📌 Multiple Goroutines

---

# ✅ Example

```go
package main

import (
	"fmt"
	"time"
)

func taskOne() {

	for i := 1; i <= 5; i++ {

		fmt.Println("Task One:", i)

		time.Sleep(time.Millisecond * 300)
	}
}

func taskTwo() {

	for i := 1; i <= 5; i++ {

		fmt.Println("Task Two:", i)

		time.Sleep(time.Millisecond * 300)
	}
}

func main() {

	go taskOne()

	go taskTwo()

	time.Sleep(time.Second * 3)
}
```

---

# 📌 Output

Output order may change because:
# ✅ Concurrent Execution

---

# 📌 Anonymous Goroutine

Goroutine without separate function.

---

# ✅ Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {

		fmt.Println("Anonymous Goroutine")
	}()

	time.Sleep(time.Second)
}
```

---

# 📌 What is sync.WaitGroup?

Used to wait for goroutines.

Instead of using:
```go
time.Sleep()
```

Professional developers use:
# ✅ WaitGroup

---

# 📌 WaitGroup Example

```go
package main

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 5; i++ {

		fmt.Println(i)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go printNumbers(&wg)

	wg.Wait()

	fmt.Println("Main Finished")
}
```

---

# 📌 Understanding WaitGroup

| Method | Purpose |
|---|---|
| Add() | Add goroutine count |
| Done() | Reduce count |
| Wait() | Wait until complete |

---

# 📌 defer wg.Done()

```go
defer wg.Done()
```

Means:
👉 Mark goroutine completed.

---

# 📌 Real Backend Usage

Concurrency used in:
- API requests
- Email sending
- Notifications
- Payment systems
- Parallel database operations

---

# 📌 Backend Example

Suppose:
User places order.

Backend can:
- Save order
- Send email
- Generate invoice
- Update inventory

simultaneously using goroutines.

Very powerful concept 🔥

---

# 📌 Goroutine vs Thread

| Goroutine | Thread |
|---|---|
| Lightweight | Heavy |
| Managed by Go | Managed by OS |
| Faster | Slower |

---

# 📌 Important Warning

Goroutines share memory.

Incorrect handling may cause:
# ❌ Race Conditions

We will learn this later.

---

# 💻 Day 11 Practice Program

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 5; i++ {

		fmt.Println("Numbers:", i)

		time.Sleep(time.Millisecond * 300)
	}
}

func printLetters(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 'A'; i <= 'E'; i++ {

		fmt.Printf("Letters: %c\n", i)

		time.Sleep(time.Millisecond * 300)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go printNumbers(&wg)

	go printLetters(&wg)

	wg.Wait()

	fmt.Println("✅ All Goroutines Completed")
}
```

---

# 📘 Day 11 Interview Questions & Answers

---

## ❓ Q1: What is concurrency?

### ✅ Answer:
Concurrency means executing multiple tasks independently.

---

## ❓ Q2: What is goroutine?

### ✅ Answer:
Goroutine is lightweight thread managed by Go runtime.

---

## ❓ Q3: How to create goroutine?

### ✅ Answer:

```go
go functionName()
```

---

## ❓ Q4: Why goroutines are faster?

### ✅ Answer:
Because they are lightweight and managed efficiently by Go runtime.

---

## ❓ Q5: What is sync.WaitGroup?

### ✅ Answer:
WaitGroup waits for goroutines to finish execution.

---

## ❓ Q6: Why wg.Done() used?

### ✅ Answer:
Marks goroutine as completed.

---

## ❓ Q7: Difference between goroutine and thread?

| Goroutine | Thread |
|---|---|
| Lightweight | Heavy |
| Faster | Slower |
| Managed by Go | Managed by OS |

---

# 📚 Day 11 Summary

Today I learned:
- Concurrency
- Goroutines
- go keyword
- Anonymous goroutines
- sync.WaitGroup
- Concurrent execution

I also practiced:
- Running multiple tasks
- Synchronization
- Backend concurrency logic

---

# 🧠 Practice Tasks

✅ Create multiple goroutines  
✅ Print numbers concurrently  
✅ Print alphabets concurrently  
✅ Use WaitGroup  
✅ Create anonymous goroutine  
✅ Simulate backend tasks

---

# ✅ Day 12 — Channels in Go

---

# 📖 Introduction to Channels

Channels are one of the most important concepts in Go concurrency.

Yesterday (Day 11):
You learned:
✅ Goroutines

Today:
You will learn:
# ✅ Communication Between Goroutines

That communication happens using:
# 🚀 Channels

---

# 📖 What You Will Learn

- Channels
- Channel Syntax
- Send & Receive Data
- Buffered Channels
- Unbuffered Channels
- Deadlock
- Channel Directions
- range & close()
- Real Backend Examples
- Interview Questions

---

# 📌 What is Channel?

Channel is used for:
# ✅ Communication between goroutines

Channels safely transfer data.

---

# 📌 Real Example

Suppose:
One goroutine:
👉 Processes payment

Another goroutine:
👉 Sends notification

They communicate using:
# ✅ Channels

---

# 📌 Why Channels Important?

Without channels:
❌ Difficult synchronization  
❌ Unsafe data sharing  
❌ Race conditions

With channels:
✅ Safe communication  
✅ Better concurrency  
✅ Cleaner code

---

# 📌 Channel Syntax

```go
channelName := make(chan datatype)
```

---

# 📌 Example

```go
messages := make(chan string)
```

---

# 📌 Understanding

| Part | Meaning |
|---|---|
| make() | Create channel |
| chan | Channel keyword |
| string | Data type |

---

# 📌 Send Data into Channel

```go
channel <- value
```

---

# 📌 Receive Data from Channel

```go
value := <-channel
```

---

# 📌 First Channel Example

```go
package main

import "fmt"

func main() {

	messageChannel := make(chan string)

	go func() {

		messageChannel <- "Hello from Goroutine"
	}()

	message := <-messageChannel

	fmt.Println(message)
}
```

---

# 📌 What Happened?

```text
Goroutine sent data
↓
Main function received data
```

---

# 📌 Important Rule

Channels are:
# ✅ Blocking

Meaning:
- Send waits for receiver
- Receive waits for sender

---

# 📌 Unbuffered Channel

Default channel type.

---

# ✅ Example

```go
channel := make(chan int)
```

---

# 📌 Deadlock Example

```go
package main

func main() {

	channel := make(chan int)

	channel <- 10
}
```

---

# 📌 Error

```text
fatal error: all goroutines are asleep - deadlock!
```

---

# 📌 Why Deadlock Happened?

Because:
❌ No receiver available.

---

# 📌 Fix Using Goroutine

```go
package main

import "fmt"

func main() {

	channel := make(chan int)

	go func() {

		channel <- 10
	}()

	value := <-channel

	fmt.Println("Received:", value)
}
```

---

# 📌 Buffered Channel

Buffered channel can store limited values without immediate receiver.

---

# ✅ Example

```go
package main

import "fmt"

func main() {

	channel := make(chan int, 2)

	channel <- 10
	channel <- 20

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
```

---

# 📌 Understanding

```go
make(chan int, 2)
```

Capacity:
```text
2 values
```

---

# 📌 Why Buffered Channels Important?

Useful when:
- Temporary storage needed
- Producer faster than consumer

---

# 📌 Channel with WaitGroup

Professional concurrency handling.

---

# ✅ Example

```go
package main

import (
	"fmt"
	"sync"
)

func sendData(channel chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	channel <- "Hello Go"
}

func main() {

	var wg sync.WaitGroup

	channel := make(chan string)

	wg.Add(1)

	go sendData(channel, &wg)

	message := <-channel

	fmt.Println(message)

	wg.Wait()
}
```

---

# 📌 Understanding WaitGroup

| Method | Purpose |
|---|---|
| Add() | Add goroutine count |
| Done() | Reduce count |
| Wait() | Wait until complete |

---

# 📌 Close Channel

Used to indicate:
# ✅ No more data will be sent

---

# ✅ Example

```go
close(channel)
```

---

# 📌 range with Channel

Used to continuously receive values.

---

# ✅ Example

```go
package main

import "fmt"

func main() {

	channel := make(chan int)

	go func() {

		for i := 1; i <= 5; i++ {

			channel <- i
		}

		close(channel)
	}()

	for value := range channel {

		fmt.Println(value)
	}
}
```

---

# 📌 What Happened?

```text
range receives values until channel closes
```

---

# 📌 Channel Directions

Restrict send/receive operations.

---

# 📌 Send Only Channel

```go
chan<- int
```

---

# 📌 Receive Only Channel

```go
<-chan int
```

---

# 📌 Example

```go
func send(channel chan<- int)
```

---

# 📌 Why Channel Directions Important?

Improves:
✅ Safety  
✅ Readability  
✅ Better architecture

---

# 📌 Real Backend Usage

Channels used in:
- Worker pools
- Background jobs
- API processing
- Notification systems
- Task queues

---

# 📌 Backend Example

Suppose:
User places order.

Different goroutines:
- Process payment
- Update inventory
- Send email

Channels coordinate communication safely.

Very powerful 🔥

---

# 📌 Goroutines + Channels

This combination is:
# 🚀 Core Power of Go

---

# 💻 Day 12 Practice Program

```go
package main

import (
	"fmt"
	"sync"
)

func sendNumbers(channel chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 5; i++ {

		channel <- i
	}

	close(channel)
}

func main() {

	var wg sync.WaitGroup

	channel := make(chan int)

	wg.Add(1)

	go sendNumbers(channel, &wg)

	for value := range channel {

		fmt.Println("Received:", value)
	}

	wg.Wait()

	fmt.Println("✅ Channel Processing Completed")
}
```

---

# 📘 Day 12 Interview Questions & Answers

---

## ❓ Q1: What is channel in Go?

### ✅ Answer:
Channel is used for communication between goroutines.

---

## ❓ Q2: Why channels important?

### ✅ Answer:
Channels provide safe communication and synchronization.

---

## ❓ Q3: What is buffered channel?

### ✅ Answer:
Buffered channel can store limited values before receiver reads them.

---

## ❓ Q4: What is deadlock?

### ✅ Answer:
Deadlock happens when goroutines wait forever.

---

## ❓ Q5: What does close(channel) do?

### ✅ Answer:
Indicates no more values will be sent.

---

## ❓ Q6: Difference between buffered and unbuffered channel?

| Buffered | Unbuffered |
|---|---|
| Stores values | No storage |
| Non-immediate receive | Immediate synchronization |

---

## ❓ Q7: What is range with channel?

### ✅ Answer:
Used to receive values continuously until channel closes.

---

# 📚 Day 12 Summary

Today I learned:
- Channels
- Buffered channels
- Unbuffered channels
- Deadlock
- close()
- range with channel
- Channel directions

I also practiced:
- Goroutine communication
- Concurrent processing
- Safe synchronization

---

# 🧠 Practice Tasks

✅ Create integer channel  
✅ Send & receive messages  
✅ Create buffered channel  
✅ Use range with channel  
✅ Close channel properly  
✅ Build producer-consumer example

---

# ✅ Day 13 — JSON Handling in Go

---

# 📖 Introduction to JSON Handling

JSON is one of the most important concepts in backend development.

JSON is used in:
- REST APIs
- Frontend communication
- Database responses
- Authentication systems
- External APIs

If you want to build APIs:
👉 JSON is mandatory.

---

# 📖 What You Will Learn

- JSON Basics
- Marshal
- Unmarshal
- Struct Tags
- JSON Arrays
- Nested JSON
- omitempty
- map with JSON
- Real API Examples
- Interview Questions

---

# 📌 What is JSON?

JSON means:
# ✅ JavaScript Object Notation

Used for:
👉 Data exchange between systems.

---

# 📌 JSON Example

```json
{
  "name": "Dnyaneshwar",
  "age": 24,
  "city": "Nashik"
}
```

---

# 📌 Why JSON Important?

Backend APIs send data in:
# ✅ JSON format

Frontend applications receive JSON responses.

---

# 📌 Go Package for JSON

```go
import "encoding/json"
```

---

# 📌 What is Marshal?

Convert:
# ✅ Go Struct → JSON

---

# 📌 Marshal Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

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
	jsonData, err := json.Marshal(student)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}
```

---

# 📌 Output

```json
{ 
	"Name":"Dnyaneshwar",
	"Age":24,
	"City":"Nashik"
}
```

---

# 📌 What Happened?

```text
Struct converted into JSON
```

---

# 📌 Why string(jsonData)?

Because:
👉 Marshal returns bytes.

Convert bytes → string.

---

# 📌 What is Unmarshal?

Convert:
# ✅ JSON → Go Struct

---

# 📌 Unmarshal Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
	City string
}

func main() {

	jsonData := `{
		"Name":"Dnyaneshwar",
		"Age":24,
		"City":"Nashik"
	}`

	var student Student

	err := json.Unmarshal([]byte(jsonData), &student)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(student.Name)
	fmt.Println(student.Age)
	fmt.Println(student.City)
}
```

---

# 📌 Understanding

| Function | Purpose |
|---|---|
| Marshal() | Struct → JSON |
| Unmarshal() | JSON → Struct |

---

# 📌 Struct Tags

Used to customize JSON field names.

---

# ✅ Example

```go
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}
```

---

# 📌 Why Struct Tags Important?

Because APIs mostly use:

```json
{
  "name":"Dnyaneshwar"
}
```

not:

```json
{
  "Name":"Dnyaneshwar"
}
```

---

# 📌 Marshal with Struct Tags

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {

	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {

	student := Student{

		Name: "Dnyaneshwar",
		Age: 24,
		City: "Nashik",
	}

	jsonData, _ := json.Marshal(student)

	fmt.Println(string(jsonData))
}
```

---

# 📌 Output

```json
{
	"name":"Dnyaneshwar",
	"age":24,
	"city":"Nashik"
}
```

---

# 📌 omitempty Tag

Removes empty fields from JSON.

---

# ✅ Example

```go
type Student struct {

	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}
```

---

# 📌 If Age Empty

Output:

```json
{
  "name":"Dnyaneshwar"
}
```

---

# 📌 JSON Array Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {

	Name string `json:"name"`
}

func main() {

	students := []Student{

		{Name: "Dnyaneshwar"},
		{Name: "Rahul"},
	}

	jsonData, _ := json.Marshal(students)

	fmt.Println(string(jsonData))
}
```

---

# 📌 Output

```json
[
  {"name":"Dnyaneshwar"},
  {"name":"Rahul"}
]
```

---

# 📌 Nested JSON

Struct inside another struct.

---

# ✅ Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {

	City  string `json:"city"`
	State string `json:"state"`
}

type Student struct {

	Name    string  `json:"name"`
	Address Address `json:"address"`
}

func main() {

	student := Student{

		Name: "Dnyaneshwar",

		Address: Address{

			City:  "Nashik",
			State: "Maharashtra",
		},
	}

	jsonData, _ := json.Marshal(student)

	fmt.Println(string(jsonData))
}
```

---

# 📌 map with JSON

---

# ✅ Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	data := map[string]interface{}{

		"name":"Dnyaneshwar",
		"age":24,
	}

	jsonData, _ := json.Marshal(data)

	fmt.Println(string(jsonData))
}
```

---

# 📌 Why interface{} Used?

Because:
👉 map values can store different datatypes.

---

# 📌 Real Backend Usage

JSON used in:
- REST APIs
- Authentication
- API responses
- Frontend communication
- External integrations

---

# 📌 Backend Example

Suppose:
Frontend sends:

```json
{
  "email":"test@gmail.com",
  "password":"123"
}
```

Backend:
👉 Converts JSON → Struct

Exactly using:
# ✅ Unmarshal

---

# 📌 Important Rule

Struct fields must start with CAPITAL letters.

Correct:

```go
Name string
```

Wrong:

```go
name string
```

Otherwise:
❌ JSON conversion fails.

---

# 💻 Day 13 Practice Program

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {

	Name   string `json:"name"`
	Role   string `json:"role"`
	Salary int    `json:"salary"`
}

func main() {

	employee := Employee{

		Name:   "Dnyaneshwar",
		Role:   "Backend Developer",
		Salary: 50000,
	}

	jsonData, err := json.Marshal(employee)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	var newEmployee Employee

	err = json.Unmarshal(jsonData, &newEmployee)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(newEmployee.Name)
	fmt.Println(newEmployee.Role)
}
```

---

# 📘 Day 13 Interview Questions & Answers

---

## ❓ Q1: What is JSON?

### ✅ Answer:
JSON is lightweight format used for data exchange.

---

## ❓ Q2: Which package used for JSON in Go?

### ✅ Answer:

```go
encoding/json
```

---

## ❓ Q3: What is Marshal?

### ✅ Answer:
Marshal converts Go object into JSON.

---

## ❓ Q4: What is Unmarshal?

### ✅ Answer:
Unmarshal converts JSON into Go object.

---

## ❓ Q5: Why struct tags important?

### ✅ Answer:
Struct tags customize JSON field names.

---

## ❓ Q6: What is omitempty?

### ✅ Answer:
Removes empty fields from JSON response.

---

## ❓ Q7: Why fields must start with capital letter?

### ✅ Answer:
Because exported fields are required for JSON conversion.

---

# 📚 Day 13 Summary

Today I learned:
- JSON handling
- Marshal
- Unmarshal
- Struct tags
- JSON arrays
- Nested JSON
- omitempty

I also practiced:
- API-style JSON handling
- Struct conversion
- Backend response formatting

---

# 🧠 Practice Tasks

✅ Create Student JSON  
✅ Convert struct → JSON  
✅ Convert JSON → struct  
✅ Use struct tags  
✅ Create nested JSON  
✅ Create JSON array

---

# ✅ Day 14 — HTTP Package in Go

---

# 📖 Introduction to HTTP Package

Today we started real backend development using Go.

The HTTP package is used to:
- Create web servers
- Handle routes
- Build APIs
- Send JSON responses
- Handle HTTP requests

This is the foundation for:
✅ REST APIs  
✅ Backend applications  
✅ Microservices

---

# 📖 What You Will Learn

- HTTP Package
- Web Server
- Routes
- Handler Functions
- Request & Response
- Query Parameters
- JSON Response
- POST Request Basics
- HTTP Methods
- Status Codes
- Real Backend Examples
- Interview Questions

---

# 📌 What is HTTP?

HTTP means:
# ✅ HyperText Transfer Protocol

Used for communication between:
- Client
- Server

Example:
- Browser → Backend API
- Mobile App → Server

---

# 📌 What is Web Server?

Web server receives requests and sends responses.

---

# 📌 Go HTTP Package

```go
import "net/http"
```

---

# 📌 First HTTP Server

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", homeHandler)

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome to Go Backend")
}
```

---

# 📌 Run Server

```bash
go run main.go
```

---

# 📌 Open Browser

```text
http://localhost:8080
```

---

# 📌 Output

```text
Welcome to Go Backend
```

---

# 📌 Understanding

| Function | Purpose |
|---|---|
| HandleFunc() | Create route |
| ListenAndServe() | Start server |
| ResponseWriter | Send response |
| Request | Receive request |

---

# 📌 What is Handler Function?

Function that handles HTTP request.

---

# 📌 Handler Syntax

```go
func handler(w http.ResponseWriter, r *http.Request)
```

---

# 📌 Parameters Meaning

| Parameter | Purpose |
|---|---|
| w | Send response |
| r | Receive request |

---

# 📌 Multiple Routes Example

```go
package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Home Page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "About Page")
}

func main() {

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on 8080")

	http.ListenAndServe(":8080", nil)
}
```

---

# 📌 Routes

| Route | Output |
|---|---|
| / | Home Page |
| /about | About Page |

---

# 📌 Query Parameters

Used to send data in URL.

---

# ✅ Example URL

```text
http://localhost:8080/user?name=Dnyaneshwar
```

---

# ✅ Query Parameter Example

```go
package main

import (
	"fmt"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	fmt.Fprintln(w, "Hello", name)
}

func main() {

	http.HandleFunc("/user", userHandler)

	http.ListenAndServe(":8080", nil)
}
```

---

# 📌 Output

```text
Hello Dnyaneshwar
```

---

# 📌 JSON Response

Backend APIs mostly send:
# ✅ JSON

---

# ✅ JSON Response Example

```go
package main

import (
	"encoding/json"
	"net/http"
)

type Student struct {

	Name string `json:"name"`
	Age  int    `json:"age"`
}

func studentHandler(w http.ResponseWriter, r *http.Request) {

	student := Student{

		Name: "Dnyaneshwar",
		Age: 24,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(student)
}

func main() {

	http.HandleFunc("/student", studentHandler)

	http.ListenAndServe(":8080", nil)
}
```

---

# 📌 Output

```json
{
  "name":"Dnyaneshwar",
  "age":24
}
```

---

# 📌 Why Header Important?

```go
w.Header().Set("Content-Type", "application/json")
```

Tells client:
# ✅ Response is JSON

---

# 📌 HTTP Methods

| Method | Purpose |
|---|---|
| GET | Retrieve data |
| POST | Create data |
| PUT | Update data |
| DELETE | Delete data |

---

# 📌 POST Request Example

```go
package main

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)

		return
	}

	fmt.Fprintln(w, "Login Successful")
}

func main() {

	http.HandleFunc("/login", loginHandler)

	http.ListenAndServe(":8080", nil)
}
```

---

# 📌 Test POST Request

Using curl:

```bash
curl -X POST http://localhost:8080/login
```

---

# 📌 Output

```text
Login Successful
```

---

# 📌 Status Codes

| Code | Meaning |
|---|---|
| 200 | OK |
| 201 | Created |
| 400 | Bad Request |
| 404 | Not Found |
| 500 | Server Error |

---

# 📌 Send Status Code

```go
w.WriteHeader(http.StatusCreated)
```

---

# 📌 Custom Error Response

```go
http.Error(w, "User Not Found", http.StatusNotFound)
```

---

# 📌 Real Backend Usage

HTTP package used in:
- REST APIs
- Authentication systems
- Microservices
- Backend servers
- Web applications

---

# 📌 Backend Example

Suppose:
Frontend sends request:

```text
GET /students
```

Backend:
- Receives request
- Fetches data
- Sends JSON response

Exactly same process used in real APIs 🔥

---

# 💻 Day 14 Practice Program

```go
package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {

	Name  string `json:"name"`
	Price int    `json:"price"`
}

func productHandler(w http.ResponseWriter, r *http.Request) {

	product := Product{

		Name:  "iPhone",
		Price: 120000,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(product)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Welcome to Go HTTP Server"))
}

func main() {

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/product", productHandler)

	http.ListenAndServe(":8080", nil)
}
```

---

# 📘 Day 14 Interview Questions & Answers

---

## ❓ Q1: Which package used for HTTP server in Go?

### ✅ Answer:

```go
net/http
```

---

## ❓ Q2: What is handler function?

### ✅ Answer:
Function that handles HTTP requests.

---

## ❓ Q3: What is ResponseWriter?

### ✅ Answer:
Used to send HTTP response to client.

---

## ❓ Q4: What is Request object?

### ✅ Answer:
Contains client request information.

---

## ❓ Q5: How to create route in Go?

### ✅ Answer:

```go
http.HandleFunc("/route", handler)
```

---

## ❓ Q6: Why Content-Type important?

### ✅ Answer:
Tells client response format.

---

## ❓ Q7: What is JSON response?

### ✅ Answer:
Response sent in JSON format.

---

## ❓ Q8: Difference between GET and POST?

| GET | POST |
|---|---|
| Retrieve data | Create data |
| Data in URL | Data in body |

---

# 📚 Day 14 Summary

Today I learned:
- HTTP package
- Web server
- Routes
- Handler functions
- Query parameters
- JSON responses
- HTTP methods
- Status codes

I also practiced:
- Creating server
- Handling routes
- Sending JSON
- API basics

---

# 🧠 Practice Tasks

✅ Create home route  
✅ Create about route  
✅ Create JSON API  
✅ Handle query parameters  
✅ Handle POST request  
✅ Send status codes
---

# ✅ Day 15 — First REST API in Go

---

# 📖 Introduction to REST API

Today we started real backend development using Go.

Until now we learned:
- Structs
- JSON
- HTTP Package
- Error Handling

Now we combine everything and build:
# ✅ First REST API

This is exactly how backend development works in companies.

---

# 📖 What You Will Learn

- REST API
- API Routes
- JSON Responses
- GET API
- POST API
- Request Body
- Decode JSON
- API Status Codes
- CRUD Basics
- Real Backend Flow
- Interview Questions

---

# 📌 What is REST API?

REST API means:
# ✅ Backend service communicating using HTTP + JSON

Used in:
- Web apps
- Mobile apps
- Frontend frameworks
- Microservices

---

# 📌 Real Example

Frontend sends:

```text
GET /students
```

Backend returns:

```json
[
  {
    "id":1,
    "name":"Dnyaneshwar"
  }
]
```

---

# 📌 REST API Uses

| Method | Purpose |
|---|---|
| GET | Fetch data |
| POST | Create data |
| PUT | Update data |
| DELETE | Delete data |

---

# 📌 Project Goal

Today we build:
# ✅ Student REST API

Features:
- Get Students
- Add Student

---

# 📌 Project Structure

```text
Day-15/
 ├── main.go
 ├── README.md
```

---

# 📌 Step 1 — Create Student Struct

```go
type Student struct {

	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

---

# 📌 Why Struct Tags Used?

Because API responses use:

```json
{
  "id":1
}
```

not:

```json
{
  "ID":1
}
```

---

# 📌 Step 2 — Create Dummy Data

```go
var students = []Student{

	{ID: 1, Name: "Dnyaneshwar", Age: 24},
	{ID: 2, Name: "Rahul", Age: 22},
}
```

---

# 📌 Why Dummy Data?

Currently:
❌ No database

Later:
✅ PostgreSQL  
✅ GORM

---

# 📌 Step 3 — GET API

Used to fetch students.

---

# ✅ Example

```go
func getStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(students)
}
```

---

# 📌 What Happened?

```text
students slice
↓
converted into JSON
↓
sent to client
```

---

# 📌 Step 4 — Register Route

```go
http.HandleFunc("/students", getStudents)
```

---

# 📌 Step 5 — Run Server

```go
http.ListenAndServe(":8080", nil)
```

---

# 📌 Open Browser

```text
http://localhost:8080/students
```

---

# 📌 Output

```json
[
  {
    "id":1,
    "name":"Dnyaneshwar",
    "age":24
  },
  {
    "id":2,
    "name":"Rahul",
    "age":22
  }
]
```

---

# 📌 POST API

Used to:
# ✅ Add Data

---

# 📌 What is Request Body?

Client sends JSON data inside request.

Example:

```json
{
  "id":3,
  "name":"Sai",
  "age":21
}
```

---

# 📌 Step 6 — Create POST API

---

# ✅ Example

```go
func createStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)

		return
	}

	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	students = append(students, student)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(student)
}
```

---

# 📌 Understanding

| Code | Purpose |
|---|---|
| NewDecoder() | Read JSON body |
| Decode() | Convert JSON → Struct |
| append() | Add student |
| StatusCreated | 201 status |

---

# 📌 Register POST Route

```go
http.HandleFunc("/add-student", createStudent)
```

---

# 📌 Test POST API

Using:
- Postman
- curl
- Thunder Client

---

# ✅ curl Example

```bash
curl -X POST http://localhost:8080/add-student \
-H "Content-Type: application/json" \
-d '{"id":3,"name":"Sai","age":21}'
```

---

# 📌 Response

```json
{
  "id":3,
  "name":"Sai",
  "age":21
}
```

---

# 📌 Full Project Code

# ✅ main.go

```go
package main

import (
	"encoding/json"
	"net/http"
)

type Student struct {

	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var students = []Student{

	{ID: 1, Name: "Dnyaneshwar", Age: 24},
	{ID: 2, Name: "Rahul", Age: 22},
}

// GET API
func getStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(students)
}

// POST API
func createStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)

		return
	}

	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	students = append(students, student)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(student)
}

func main() {

	http.HandleFunc("/students", getStudents)

	http.HandleFunc("/add-student", createStudent)

	http.ListenAndServe(":8080", nil)
}
```

---

# 📌 API Flow

```text
Client Request
↓
HTTP Route
↓
Handler Function
↓
Business Logic
↓
JSON Response
```

---

# 📌 Real Backend Understanding

This is exactly how:
- ASP.NET APIs
- Node.js APIs
- Java APIs
- Go APIs

work internally 🔥

---

# 📌 Current Limitation

Currently:
❌ Data stored in memory

After Day 20:
✅ PostgreSQL database  
✅ Permanent storage

---

# 💻 Backend Skills You Learned Today

✅ API development  
✅ JSON APIs  
✅ GET requests  
✅ POST requests  
✅ Request body handling  
✅ HTTP status codes  
✅ Route handling

---

# 📘 Day 15 Interview Questions & Answers

---

## ❓ Q1: What is REST API?

### ✅ Answer:
REST API is backend service using HTTP and JSON for communication.

---

## ❓ Q2: Difference between GET and POST?

| GET | POST |
|---|---|
| Fetch data | Create data |
| No body | Has body |

---

## ❓ Q3: What is request body?

### ✅ Answer:
Client sends data inside HTTP request.

---

## ❓ Q4: Why json.NewDecoder() used?

### ✅ Answer:
Used to convert JSON request into struct.

---

## ❓ Q5: Why json.NewEncoder() used?

### ✅ Answer:
Used to send JSON response.

---

## ❓ Q6: Why Content-Type important?

### ✅ Answer:
Indicates response format is JSON.

---

## ❓ Q7: What is StatusCreated?

### ✅ Answer:
HTTP 201 status indicating resource created successfully.

---

# 📚 Day 15 Summary

Today I learned:
- REST APIs
- GET API
- POST API
- JSON request handling
- JSON responses
- Request body
- API routes
- HTTP status codes

I also practiced:
- Building APIs
- Sending JSON
- Receiving JSON
- Backend API flow

---

# 🧠 Practice Tasks

✅ Create Product API  
✅ Create User API  
✅ Add GET route  
✅ Add POST route  
✅ Send JSON response  
✅ Decode request body

---

# ✅ Day 16 — CRUD REST API in Go

---

# 📖 Introduction to CRUD REST API

Today we moved from:
# ✅ Simple API
to:
# 🚀 Complete CRUD REST API

This is exactly what backend developers build in companies.

---

# 📖 What You Will Learn

- CRUD Operations
- GET API
- POST API
- PUT API
- DELETE API
- Query Parameters
- Update Data
- Delete Data
- API Design
- REST Standards
- Backend Logic
- Interview Questions

---

# 📌 What is CRUD?

CRUD means:

| Operation | Meaning |
|---|---|
| Create | Add data |
| Read | Fetch data |
| Update | Modify data |
| Delete | Remove data |

---

# 📌 Real Backend Example

Student Management System:
- Add Student
- Get Students
- Update Student
- Delete Student

This is:
# ✅ CRUD API

---

# 📌 Project Goal

Today we build:
# 🚀 Student CRUD REST API

Features:
✅ Get All Students  
✅ Get Student By ID  
✅ Add Student  
✅ Update Student  
✅ Delete Student

---

# 📌 Project Structure

```text
Day-16/
 ├── main.go
 ├── README.md
```

---

# 📌 Step 1 — Create Student Struct

```go
type Student struct {

	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

---

# 📌 Step 2 — Create Dummy Data

```go
var students = []Student{

	{ID: 1, Name: "Dnyaneshwar", Age: 24},
	{ID: 2, Name: "Rahul", Age: 22},
}
```

---

# 📌 GET All Students API

---

# ✅ Example

```go
func getStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(students)
}
```

---

# 📌 Route

```go
http.HandleFunc("/students", getStudents)
```

---

# 📌 Output

```json
[
  {
    "id":1,
    "name":"Dnyaneshwar",
    "age":24
  }
]
```

---

# 📌 GET Student By ID API

---

# 📌 Problem

Current API:

```text
/students
```

returns ALL students.

We need:

```text
/student?id=1
```

---

# ✅ Example

```go
func getStudentByID(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	for _, student := range students {

		if strconv.Itoa(student.ID) == id {

			json.NewEncoder(w).Encode(student)

			return
		}
	}

	http.Error(w, "Student Not Found", http.StatusNotFound)
}
```

---

# 📌 Why strconv.Itoa()?

Because:

```text
student.ID → integer
id → string
```

Need conversion.

---

# 📌 Route

```go
http.HandleFunc("/student", getStudentByID)
```

---

# 📌 Test

```text
http://localhost:8080/student?id=1
```

---

# 📌 POST API — Add Student

---

# ✅ Example

```go
func createStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)

		return
	}

	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	students = append(students, student)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(student)
}
```

---

# 📌 Route

```go
http.HandleFunc("/add-student", createStudent)
```

---

# 📌 PUT API — Update Student

---

# 📌 Goal

Update existing student.

---

# ✅ Example

```go
func updateStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {

		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)

		return
	}

	id := r.URL.Query().Get("id")

	var updatedStudent Student

	err := json.NewDecoder(r.Body).Decode(&updatedStudent)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	for index, student := range students {

		if strconv.Itoa(student.ID) == id {

			students[index] = updatedStudent

			json.NewEncoder(w).Encode(updatedStudent)

			return
		}
	}

	http.Error(w, "Student Not Found", http.StatusNotFound)
}
```

---

# 📌 Route

```go
http.HandleFunc("/update-student", updateStudent)
```

---

# 📌 Test PUT API

```text
PUT /update-student?id=1
```

---

# 📌 Request Body

```json
{
  "id":1,
  "name":"Updated Name",
  "age":25
}
```

---

# 📌 DELETE API

---

# 📌 Goal

Delete student by ID.

---

# ✅ Example

```go
func deleteStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {

		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)

		return
	}

	id := r.URL.Query().Get("id")

	for index, student := range students {

		if strconv.Itoa(student.ID) == id {

			students = append(
				students[:index],
				students[index+1:]...,
			)

			fmt.Fprintln(w, "Student Deleted Successfully")

			return
		}
	}

	http.Error(w, "Student Not Found", http.StatusNotFound)
}
```

---

# 📌 Delete Logic

```go
students[:index]
students[index+1:]
```

Removes current student.

---

# 📌 Route

```go
http.HandleFunc("/delete-student", deleteStudent)
```

---

# 📌 Full Project Code

# ✅ main.go

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Student struct {

	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var students = []Student{

	{ID: 1, Name: "Dnyaneshwar", Age: 24},
	{ID: 2, Name: "Rahul", Age: 22},
}

// GET ALL STUDENTS
func getStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(students)
}

// GET STUDENT BY ID
func getStudentByID(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	for _, student := range students {

		if strconv.Itoa(student.ID) == id {

			json.NewEncoder(w).Encode(student)

			return
		}
	}

	http.Error(w, "Student Not Found", http.StatusNotFound)
}

// CREATE STUDENT
func createStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)

		return
	}

	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	students = append(students, student)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(student)
}

// UPDATE STUDENT
func updateStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {

		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)

		return
	}

	id := r.URL.Query().Get("id")

	var updatedStudent Student

	err := json.NewDecoder(r.Body).Decode(&updatedStudent)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	for index, student := range students {

		if strconv.Itoa(student.ID) == id {

			students[index] = updatedStudent

			json.NewEncoder(w).Encode(updatedStudent)

			return
		}
	}

	http.Error(w, "Student Not Found", http.StatusNotFound)
}

// DELETE STUDENT
func deleteStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {

		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)

		return
	}

	id := r.URL.Query().Get("id")

	for index, student := range students {

		if strconv.Itoa(student.ID) == id {

			students = append(
				students[:index],
				students[index+1:]...,
			)

			fmt.Fprintln(w, "Student Deleted Successfully")

			return
		}
	}

	http.Error(w, "Student Not Found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/students", getStudents)

	http.HandleFunc("/student", getStudentByID)

	http.HandleFunc("/add-student", createStudent)

	http.HandleFunc("/update-student", updateStudent)

	http.HandleFunc("/delete-student", deleteStudent)

	fmt.Println("Server Running on Port 8080")

	http.ListenAndServe(":8080", nil)
}
```

---

# 📌 CRUD API Flow

```text
Client Request
↓
Route
↓
Handler
↓
Business Logic
↓
JSON Response
```

---

# 📌 Real Backend Understanding

This is exactly how:
- Company backend APIs
- ASP.NET APIs
- Node.js APIs
- Go APIs

work internally 🔥

---

# 📘 Day 16 Interview Questions & Answers

---

## ❓ Q1: What is CRUD?

### ✅ Answer:
CRUD means Create, Read, Update, Delete.

---

## ❓ Q2: Which HTTP method used for update?

### ✅ Answer:
PUT method.

---

## ❓ Q3: Which HTTP method used for delete?

### ✅ Answer:
DELETE method.

---

## ❓ Q4: Why json.NewDecoder() used?

### ✅ Answer:
Used to convert request JSON into struct.

---

## ❓ Q5: Why StatusCreated used?

### ✅ Answer:
Indicates successful resource creation.

---

## ❓ Q6: What is query parameter?

### ✅ Answer:
Data sent through URL.

Example:

```text
?id=1
```

---

## ❓ Q7: Why strconv.Itoa() used?

### ✅ Answer:
Converts integer into string.

---

# 📚 Day 16 Summary

Today I learned:
- CRUD APIs
- GET API
- POST API
- PUT API
- DELETE API
- Query parameters
- JSON request handling
- REST API design

I also practiced:
- Creating backend APIs
- Updating data
- Deleting data
- API routing

---

# 🧠 Practice Tasks

✅ Create Product CRUD API  
✅ Create Employee CRUD API  
✅ Add validation  
✅ Add custom error messages  
✅ Add status codes  
✅ Test APIs in Postman

---

# ✅ Day 17 — Routing & URL Parameters in Go

---

# 📖 Introduction to Routing & URL Parameters

Today we learned an important backend concept:
# 🚀 Routing & URL Parameters

Until now:
We created APIs using:

```text
/student?id=1
```

But real APIs use:

```text
/students/1
```

This is called:
# ✅ URL Parameters / Path Parameters

Today we learned:
- Dynamic routes
- URL parameters
- Clean API routing
- RESTful API structure

Used heavily in:
- REST APIs
- Microservices
- Production backend systems

---

# 📖 What You Will Learn

- Routing
- URL Parameters
- Dynamic Routes
- Path Parameters
- strings Package
- RESTful Routing
- Route Parsing
- Clean API Design
- Real Backend Examples
- Interview Questions

---

# 📌 What is Routing?

Routing means:
# ✅ Mapping URL to function

Example:

| Route | Handler |
|---|---|
| /students | getStudents |
| /products | getProducts |

---

# 📌 What are URL Parameters?

Dynamic values inside URL.

Example:

```text
/students/1
```

Here:

```text
1
```

is:
# ✅ URL Parameter

---

# 📌 Why URL Parameters Important?

Used for:
- Fetch by ID
- Update by ID
- Delete by ID

Examples:
- `/users/10`
- `/products/5`
- `/orders/100`

---

# 📌 Problem in net/http

Go standard `net/http` does NOT support dynamic routing directly.

So today we manually handled routes.

Later:
✅ Gin Framework  
✅ Professional routing

---

# 📌 Example Route

```text
/student/1
```

---

# 📌 Goal

Extract:

```text
1
```

from URL.

---

# 📌 Step 1 — Import strings Package

```go
import "strings"
```

---

# 📌 Why strings Package?

Used to split URL.

---

# 📌 Step 2 — Route Example

```text
http://localhost:8080/student/1
```

---

# 📌 Step 3 — Split URL

```go
parts := strings.Split(r.URL.Path, "/")
```

---

# 📌 Result

```text
["", "student", "1"]
```

---

# 📌 Access ID

```go
id := parts[2]
```

---

# 📌 First URL Parameter Example

```go
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func studentHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")

	id := parts[2]

	fmt.Fprintln(w, "Student ID:", id)
}

func main() {

	http.HandleFunc("/student/", studentHandler)

	http.ListenAndServe(":8080", nil)
}
```

---

# 📌 Test

```text
http://localhost:8080/student/10
```

---

# 📌 Output

```text
Student ID: 10
```

---

# 📌 Understanding

| Code | Purpose |
|---|---|
| URL.Path | Get URL |
| strings.Split() | Split route |
| parts[2] | Extract ID |

---

# 📌 Problem

If URL invalid:

```text
/student/
```

then:
❌ index out of range error

---

# 📌 Solution — Validation

---

# ✅ Example

```go
if len(parts) < 3 {

	http.Error(w, "Invalid URL", http.StatusBadRequest)

	return
}
```

---

# 📌 Updated Example

```go
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func studentHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {

		http.Error(w, "Invalid URL", http.StatusBadRequest)

		return
	}

	id := parts[2]

	fmt.Fprintln(w, "Student ID:", id)
}

func main() {

	http.HandleFunc("/student/", studentHandler)

	http.ListenAndServe(":8080", nil)
}
```

---

# 📌 Dynamic Student API

Now combine:
- Routing
- URL parameters
- JSON response

---

# ✅ Example

```go
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Student struct {

	ID   int    `json:"id"`
	Name string `json:"name"`
}

var students = []Student{

	{ID: 1, Name: "Dnyaneshwar"},
	{ID: 2, Name: "Rahul"},
}

func getStudent(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {

		http.Error(w, "Invalid URL", http.StatusBadRequest)

		return
	}

	id := parts[2]

	for _, student := range students {

		if strconv.Itoa(student.ID) == id {

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(student)

			return
		}
	}

	http.Error(w, "Student Not Found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/student/", getStudent)

	http.ListenAndServe(":8080", nil)
}
```

---

# 📌 Test

```text
http://localhost:8080/student/1
```

---

# 📌 Response

```json
{
  "id":1,
  "name":"Dnyaneshwar"
}
```

---

# 📌 RESTful Routing

Professional API structure.

---

# 📌 Good API Design

| Operation | Route |
|---|---|
| Get All | GET /students |
| Get By ID | GET /students/1 |
| Create | POST /students |
| Update | PUT /students/1 |
| Delete | DELETE /students/1 |

---

# 📌 Why RESTful Routing Important?

Improves:
✅ Clean architecture  
✅ API readability  
✅ Professional backend structure

---

# 📌 Current Limitation

Manual route parsing is difficult.

Tomorrow:
# 🚀 Gin Framework

will solve this professionally.

---

# 📌 Real Backend Usage

URL parameters used in:
- User APIs
- Product APIs
- Order APIs
- Authentication systems

Every production backend uses this.

---

# 💻 Day 17 Practice Program

```go
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func productHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {

		http.Error(w, "Invalid Product URL", http.StatusBadRequest)

		return
	}

	productID := parts[2]

	fmt.Fprintln(w, "Product ID:", productID)
}

func main() {

	http.HandleFunc("/product/", productHandler)

	fmt.Println("Server Running on Port 8080")

	http.ListenAndServe(":8080", nil)
}
```

---

# 📘 Day 17 Interview Questions & Answers

---

## ❓ Q1: What is routing?

### ✅ Answer:
Routing maps URL to handler function.

---

## ❓ Q2: What are URL parameters?

### ✅ Answer:
Dynamic values passed inside URL.

Example:

```text
/students/1
```

---

## ❓ Q3: Why URL parameters important?

### ✅ Answer:
Used for fetching, updating, and deleting resources by ID.

---

## ❓ Q4: Which package used for URL splitting?

### ✅ Answer:

```go
strings
```

---

## ❓ Q5: What does strings.Split() do?

### ✅ Answer:
Splits string into parts.

---

## ❓ Q6: What is RESTful routing?

### ✅ Answer:
Professional API route structure using HTTP methods and clean URLs.

---

## ❓ Q7: Why validation important in URL parameters?

### ✅ Answer:
Prevents runtime errors and invalid requests.

---

# 📚 Day 17 Summary

Today I learned:
- Routing
- URL parameters
- Dynamic routes
- RESTful routing
- Route validation
- URL parsing

I also practiced:
- Extracting IDs
- Dynamic API routes
- JSON responses
- Backend route handling

---

# 🧠 Practice Tasks

✅ Create Product API with URL parameter  
✅ Create User API with URL parameter  
✅ Add route validation  
✅ Send JSON response  
✅ Handle invalid routes  
✅ Create RESTful routes

---

# ✅ Day 18 — Gin Framework in Go

---

# 📖 Introduction to Gin Framework

Today we started:
# 🚀 Professional Backend Development

Until now:
We used:

```text
net/http
```

But real Go backend developers mostly use:
# ✅ Gin Framework

Why?
Because Gin provides:
- Fast routing
- Middleware
- JSON handling
- Clean APIs
- Professional backend structure

Used by:
- Startups
- Product companies
- Microservices
- Production APIs

---

# 📖 What You Will Learn

- Gin Framework
- Gin Installation
- Routes in Gin
- GET API
- POST API
- JSON Response
- Request Body
- Path Parameters
- Query Parameters
- HTTP Status Codes
- Real Backend Examples
- Interview Questions

---

# 📌 What is Gin?

Gin is:
# ✅ Fast HTTP web framework for Go

Built on top of:

```text
net/http
```

---

# 📌 Why Gin Popular?

Because it is:
✅ Fast  
✅ Minimal  
✅ Clean  
✅ Easy for APIs  
✅ Production-ready

---

# 📌 Gin vs net/http

| Gin | net/http |
|---|---|
| Easy routing | Manual routing |
| Built-in JSON | More boilerplate |
| Middleware support | Manual handling |
| Cleaner APIs | More code |

---

# 📌 Step 1 — Initialize Go Module

Before installing Gin:

```bash
go mod init day18
```

---

# 📌 Why go.mod Important?

`go.mod` manages:
- Dependencies
- Package versions
- Project modules

Like:
- `package.json` in Node.js
- `.csproj` in .NET

---

# 📌 Install Gin

Run:

```bash
go get -u github.com/gin-gonic/gin
```

---

# 📌 Project Structure

```text
Day-18/
 ├── go.mod
 ├── go.sum
 ├── main.go
 ├── README.md
```

---

# 📌 Import Gin

```go
import "github.com/gin-gonic/gin"
```

---

# 📌 First Gin Server

```go
package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{

			"message": "Welcome to Gin Framework",
		})
	})

	router.Run(":8080")
}
```

---

# 📌 Run Server

```bash
go run main.go
```

---

# 📌 Open Browser

```text
http://localhost:8080
```

---

# 📌 Output

```json
{
  "message":"Welcome to Gin Framework"
}
```

---

# 📌 Understanding

| Code | Purpose |
|---|---|
| gin.Default() | Create router |
| router.GET() | Create GET route |
| c.JSON() | Send JSON response |
| router.Run() | Start server |

---

# 📌 What is gin.H?

Shortcut for:

```go
map[string]interface{}
```

Used for JSON response.

---

# 📌 GET API Example

```go
package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/students", func(c *gin.Context) {

		c.JSON(200, gin.H{

			"name": "Dnyaneshwar",
			"age": 24,
		})
	})

	router.Run(":8080")
}
```

---

# 📌 Path Parameters in Gin

Earlier:
Manual splitting required.

Now:
# ✅ Very Easy

---

# ✅ Example

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/student/:id", func(c *gin.Context) {

		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{

			"student_id": id,
		})
	})

	router.Run(":8080")
}
```

---

# 📌 Test

```text
http://localhost:8080/student/1
```

---

# 📌 Output

```json
{
  "student_id":"1"
}
```

---

# 📌 Query Parameters in Gin

---

# ✅ Example

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/search", func(c *gin.Context) {

		name := c.Query("name")

		c.JSON(http.StatusOK, gin.H{

			"name": name,
		})
	})

	router.Run(":8080")
}
```

---

# 📌 Test

```text
http://localhost:8080/search?name=Dnyaneshwar
```

---

# 📌 Output

```json
{
  "name":"Dnyaneshwar"
}
```

---

# 📌 POST API in Gin

---

# 📌 Create Struct

```go
type Product struct {

	Name  string `json:"name"`
	Price int    `json:"price"`
}
```

---

# ✅ POST API Example

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {

	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {

	router := gin.Default()

	router.GET("/product/:id", func(c *gin.Context) {

		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{

			"product_id": id,
		})
	})

	router.POST("/product", func(c *gin.Context) {

		var product Product

		err := c.BindJSON(&product)

		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{

				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{

			"message": "Product Created",
			"data": product,
		})
	})

	router.Run(":8080")
}
```

---

# 📌 Understanding

| Code | Purpose |
|---|---|
| BindJSON() | Convert request JSON → struct |
| c.JSON() | Send JSON response |
| StatusCreated | HTTP 201 |

---

# 📌 Test GET API

Open browser:

```text
http://localhost:8080/product/1
```

---

# 📌 GET Response

```json
{
  "product_id":"1"
}
```

---

# 📌 Test POST API

Using curl:

```bash
curl -X POST http://localhost:8080/product \
-H "Content-Type: application/json" \
-d '{"name":"iPhone","price":120000}'
```

---

# 📌 POST Response

```json
{
  "data": {
    "name": "iPhone",
    "price": 120000
  },
  "message": "Product Created"
}
```

---

# 📌 HTTP Status Codes in Gin

---

# ✅ Example

```go
c.JSON(http.StatusOK, gin.H{})
```

---

# 📌 Common Status Codes

| Code | Meaning |
|---|---|
| 200 | OK |
| 201 | Created |
| 400 | Bad Request |
| 404 | Not Found |
| 500 | Internal Server Error |

---

# 📌 Why Gin Important?

Gin makes:
✅ APIs cleaner  
✅ Development faster  
✅ Backend architecture better

---

# 📌 Real Backend Usage

Gin used in:
- REST APIs
- Authentication systems
- Microservices
- High-performance backends

---

# 📌 Company-Level Backend

Most Go backend interviews ask:
- Gin routing
- JSON handling
- CRUD APIs
- Middleware
- Authentication

Very important topic 🔥

---

# 💻 Day 18 Practice Program

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {

	Name string `json:"name"`
	Role string `json:"role"`
}

func main() {

	router := gin.Default()

	router.GET("/user/:id", func(c *gin.Context) {

		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{

			"user_id": id,
		})
	})

	router.POST("/user", func(c *gin.Context) {

		var user User

		err := c.BindJSON(&user)

		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{

				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{

			"message": "User Created",
			"data": user,
		})
	})

	router.Run(":8080")
}
```

---

# 📘 Day 18 Interview Questions & Answers

---

## ❓ Q1: What is Gin framework?

### ✅ Answer:
Gin is fast HTTP web framework for Go.

---

## ❓ Q2: Why Gin popular?

### ✅ Answer:
Because it provides fast routing, middleware, and clean API development.

---

## ❓ Q3: What is gin.Default()?

### ✅ Answer:
Creates Gin router with default middleware.

---

## ❓ Q4: What is gin.H?

### ✅ Answer:
Shortcut for:

```go
map[string]interface{}
```

---

## ❓ Q5: What does BindJSON() do?

### ✅ Answer:
Converts request JSON into struct.

---

## ❓ Q6: How to get path parameter in Gin?

### ✅ Answer:

```go
c.Param("id")
```

---

## ❓ Q7: How to get query parameter in Gin?

### ✅ Answer:

```go
c.Query("name")
```

---

## ❓ Q8: Why Gin better than net/http?

### ✅ Answer:
Gin provides easier routing, middleware, and cleaner APIs.

---

# 📚 Day 18 Summary

Today I learned:
- Gin framework
- Gin routing
- GET APIs
- POST APIs
- JSON responses
- Path parameters
- Query parameters
- Request body handling

I also practiced:
- Building APIs with Gin
- Dynamic routes
- JSON APIs
- Backend architecture

---

# 🧠 Practice Tasks

✅ Create Product API  
✅ Create User API  
✅ Add GET route  
✅ Add POST route  
✅ Use path parameters  
✅ Use query parameters

---

# ✅ Day 19 — PostgreSQL Integration in Go

---

# 📖 Introduction to PostgreSQL Integration

Today we moved from:
# ❌ Dummy Data
to:
# 🚀 Real Database Integration

This is where:
# ✅ Real Backend Development Starts

Until now:
Data stored in memory.

Problem:

```text
Server Restart
↓
All Data Lost
```

Solution:
# ✅ PostgreSQL Database

Today we learned:
- PostgreSQL connection
- Database queries
- Insert data
- Fetch data
- Real backend database flow

---

# 📖 What You Will Learn

- PostgreSQL
- Database Connection
- SQL Queries in Go
- Insert Data
- Fetch Data
- database/sql package
- pq driver
- API + Database Integration
- Real Backend Architecture
- Interview Questions

---

# 📌 What is PostgreSQL?

PostgreSQL is:
# ✅ Powerful Relational Database

Used in:
- Enterprise applications
- Product companies
- Microservices
- Scalable backend systems

---

# 📌 Why PostgreSQL Popular?

Because it is:
✅ Fast  
✅ Open-source  
✅ Secure  
✅ ACID compliant  
✅ Production-ready

---

# 📌 Real Backend Flow

```text
Client Request
↓
Go API
↓
PostgreSQL Database
↓
Response
```

---

# 🚀 PostgreSQL Installation on MacBook Air M2

---

# ✅ Step 1 — Check Homebrew

```bash
brew --version
```

---

# ✅ Step 2 — Install PostgreSQL

```bash
brew install postgresql
```

---

# ✅ Step 3 — Start PostgreSQL Service

```bash
brew services start postgresql
```

---

# ✅ Step 4 — Verify PostgreSQL

```bash
psql --version
```

Expected:

```text
psql (PostgreSQL) 18.x
```

---

# ✅ Step 5 — Open PostgreSQL Terminal

```bash
psql postgres
```

---

# ✅ Step 6 — Create Database

```sql
CREATE DATABASE studentdb;
```

---

# ✅ Step 7 — Connect Database

```sql
\c studentdb
```

---

# ✅ Step 8 — Create Table

```sql
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    age INT
);
```

---

# ✅ Step 9 — Insert Data

```sql
INSERT INTO students(name, age)
VALUES('Dnyaneshwar', 24);
```

---

# ✅ Step 10 — Fetch Data

```sql
SELECT * FROM students;
```

---

# ✅ Step 11 — Exit PostgreSQL

```sql
\q
```

---

# 📌 Important PostgreSQL Commands

| Command | Purpose |
|---|---|
| `\l` | Show databases |
| `\c dbname` | Connect database |
| `\dt` | Show tables |
| `SELECT * FROM table;` | Fetch data |
| `\q` | Exit PostgreSQL |

---

# 📌 Install PostgreSQL Driver

Run:

```bash
go get github.com/lib/pq
```

---

# 📌 Install Gin Framework

```bash
go get github.com/gin-gonic/gin
```

---

# 📌 Import Packages

```go
import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
```

---

# 📌 Why Blank Identifier `_` Used?

```go
_ "github.com/lib/pq"
```

Registers PostgreSQL driver internally.

---

# 📌 Connect PostgreSQL in Go

---

# ✅ Example

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=dnyaneshwarkokate dbname=studentdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {

		fmt.Println(err)
		return
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("PostgreSQL Connected Successfully")
}
```

---

# 📌 Run Program

```bash
go run main.go
```

---

# 📌 Output

```text
PostgreSQL Connected Successfully
```

---

# 📌 Understanding

| Code | Purpose |
|---|---|
| sql.Open() | Open DB connection |
| postgres | Driver name |
| connStr | Database credentials |
| db.Ping() | Verify DB connection |
| db.Close() | Close connection |

---

# 📌 Insert Data into Database

---

# ✅ Example

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=dnyaneshwarkokate dbname=studentdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {

		fmt.Println(err)
		return
	}

	defer db.Close()

	query := `
	INSERT INTO students(name, age)
	VALUES($1, $2)
	`

	_, err = db.Exec(query, "Dnyaneshwar", 24)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("Student Inserted Successfully")
}
```

---

# 📌 Why `$1`, `$2` Used?

Used for:
# ✅ Parameterized Queries

Prevents:
❌ SQL Injection

---

# 📌 Fetch Data from Database

---

# ✅ Example

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Student struct {

	ID   int
	Name string
	Age  int
}

func main() {

	connStr := "user=dnyaneshwarkokate dbname=studentdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {

		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, name, age FROM students")

	if err != nil {

		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {

		var student Student

		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Age,
		)

		if err != nil {

			fmt.Println(err)
			return
		}

		fmt.Println(student.ID, student.Name, student.Age)
	}
}
```

---

# 📌 Understanding

| Code | Purpose |
|---|---|
| db.Query() | Execute SELECT query |
| rows.Next() | Iterate rows |
| rows.Scan() | Read column values |

---

# 📌 API + PostgreSQL Integration

Now combine:
- Gin APIs
- PostgreSQL
- JSON Responses

This becomes:
# 🚀 Real Backend API

---

# 📌 Student API Example

---

# ✅ Example

```go
package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	connStr := "user=dnyaneshwarkokate dbname=studentdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/students", func(c *gin.Context) {

		rows, err := db.Query("SELECT id, name, age FROM students")

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		defer rows.Close()

		var students []Student

		for rows.Next() {

			var student Student

			rows.Scan(
				&student.ID,
				&student.Name,
				&student.Age,
			)

			students = append(students, student)
		}

		c.JSON(http.StatusOK, students)
	})

	router.Run(":8080")
}
```

---

# 📌 Test API

```text
http://localhost:8080/students
```

---

# 📌 JSON Response

```json
[
  {
    "id":1,
    "name":"Dnyaneshwar",
    "age":24
  }
]
```

---

# 📌 Why This Important?

Now:
# ✅ Data stored permanently

Server restart:
✔ Data still exists

---

# 📌 Real Backend Architecture

```text
Frontend
↓
Go Gin API
↓
PostgreSQL
↓
Persistent Data
```

---

# 📌 Common Database Methods

| Method | Purpose |
|---|---|
| Exec() | INSERT/UPDATE/DELETE |
| Query() | Multiple rows |
| QueryRow() | Single row |

---

# 📌 Real Backend Usage

PostgreSQL used in:
- E-commerce
- Banking
- Authentication systems
- Enterprise applications
- Microservices

---

# 📌 Very Important

Most Go interviews ask:
- Database connection
- SQL queries
- CRUD APIs
- PostgreSQL integration
- Gin + PostgreSQL

Very important topic 🔥

---

# 💻 Day 19 Practice Program

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=dnyaneshwarkokate dbname=studentdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {

		fmt.Println(err)
		return
	}

	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO students(name, age) VALUES($1, $2)",
		"Rahul",
		22,
	)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("Student Added Successfully")
}
```

---

# 📘 Day 19 Interview Questions & Answers

---

## ❓ Q1: What is PostgreSQL?

### ✅ Answer:
PostgreSQL is open-source relational database system.

---

## ❓ Q2: Which package used for database operations in Go?

### ✅ Answer:

```go
database/sql
```

---

## ❓ Q3: Why pq driver used?

### ✅ Answer:
Used for PostgreSQL database connectivity.

---

## ❓ Q4: What does sql.Open() do?

### ✅ Answer:
Creates database connection.

---

## ❓ Q5: Difference between Query() and Exec()?

| Query() | Exec() |
|---|---|
| SELECT queries | INSERT/UPDATE/DELETE |

---

## ❓ Q6: Why parameterized query important?

### ✅ Answer:
Prevents SQL Injection attacks.

---

## ❓ Q7: What does rows.Scan() do?

### ✅ Answer:
Reads database column values into variables.

---

## ❓ Q8: Why defer db.Close() used?

### ✅ Answer:
Closes database connection properly.

---

# 📚 Day 19 Summary

Today I learned:
- PostgreSQL integration
- Database connection
- Insert queries
- Select queries
- Query execution
- API + Database integration
- Persistent storage

I also practiced:
- Connecting PostgreSQL
- Running SQL queries
- Fetching database records
- Backend database flow

---

# 🧠 Practice Tasks

✅ Create Product Table  
✅ Insert Product Data  
✅ Fetch Product Data  
✅ Create User Table  
✅ Build API with PostgreSQL  
✅ Practice SQL Queries

---

# ✅ Day 20 — GORM ORM in Go

---

# 📖 Introduction to GORM ORM

Today we moved from:
# ❌ Raw SQL Queries
to:
# 🚀 Professional ORM Development

Until now:
We used:

```sql
SELECT * FROM students
```

But in real production backend:
Most Go developers use:
# ✅ GORM ORM

Why?
Because GORM provides:
- Cleaner code
- Faster development
- ORM features
- Auto migrations
- Model handling
- Easier CRUD

Used in:
- Startups
- Product companies
- Enterprise APIs
- Microservices

---

# 📖 What You Will Learn

- What is ORM
- GORM Introduction
- Install GORM
- PostgreSQL with GORM
- Models
- Auto Migration
- Create Data
- Fetch Data
- Update Data
- Delete Data
- GORM CRUD APIs
- Interview Questions

---

# 📌 What is ORM?

ORM means:
# ✅ Object Relational Mapping

Used to map:

```text
Go Struct
↓
Database Table
```

---

# 📌 Without ORM

You write:

```sql
SELECT * FROM students
```

---

# 📌 With GORM

You write:

```go
db.Find(&students)
```

Much cleaner ✅

---

# 📌 Why GORM Popular?

Because it provides:
✅ Cleaner code  
✅ Faster development  
✅ Less SQL writing  
✅ Easy CRUD  
✅ Auto migration  
✅ Production-ready architecture

---

# 📌 Install GORM

Run:

```bash
go get gorm.io/gorm
```

---

# 📌 Install PostgreSQL Driver

```bash
go get gorm.io/driver/postgres
```

---

# 📌 Project Structure

```text
Day-20/
 ├── go.mod
 ├── go.sum
 ├── main.go
 ├── README.md
```

---

# 📌 Import Packages

```go
import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)
```

---

# 📌 Connect PostgreSQL with GORM

---

# ✅ Example

```go
package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "user=dnyaneshwarkokate dbname=studentdb sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic(err)
	}

	fmt.Println("Database Connected Successfully")

	_ = db
}
```

---

# 📌 Run Program

```bash
go run main.go
```

---

# 📌 Output

```text
Database Connected Successfully
```

---

# 📌 What is Model?

Model means:
# ✅ Struct mapped with database table

---

# 📌 Student Model

```go
type Student struct {

	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

---

# 📌 Understanding

| Tag | Purpose |
|---|---|
| json | JSON response |
| gorm | Database configuration |

---

# 📌 What is AutoMigrate?

Automatically creates:
✅ Tables  
✅ Columns  
✅ Schema changes

---

# ✅ Example

```go
db.AutoMigrate(&Student{})
```

---

# 📌 Full Migration Example

```go
package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {

	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	dsn := "user=dnyaneshwarkokate dbname=studentdb sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic(err)
	}

	db.AutoMigrate(&Student{})

	fmt.Println("Migration Completed")
}
```

---

# 📌 What Happened?

GORM automatically created:
# ✅ students table

---

# 📌 Insert Data with GORM

---

# ✅ Example

```go
student := Student{

	Name: "Dnyaneshwar",
	Age: 24,
}

db.Create(&student)
```

---

# 📌 Fetch Data with GORM

---

# ✅ Example

```go
var students []Student

db.Find(&students)
```

---

# 📌 Fetch Single Record

```go
var student Student

db.First(&student, 1)
```

---

# 📌 Update Data

---

# ✅ Example

```go
db.Model(&student).Update("Age", 25)
```

---

# 📌 Delete Data

---

# ✅ Example

```go
db.Delete(&student, 1)
```

---

# 📌 Full CRUD Example

# ✅ main.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {

	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	dsn := "user=dnyaneshwarkokate dbname=studentdb sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic(err)
	}

	db.AutoMigrate(&Student{})

	router := gin.Default()

	// GET ALL STUDENTS
	router.GET("/students", func(c *gin.Context) {

		var students []Student

		db.Find(&students)

		c.JSON(http.StatusOK, students)
	})

	// CREATE STUDENT
	router.POST("/students", func(c *gin.Context) {

		var student Student

		err := c.BindJSON(&student)

		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{

				"error": err.Error(),
			})

			return
		}

		db.Create(&student)

		c.JSON(http.StatusCreated, student)
	})

	// UPDATE STUDENT
	router.PUT("/students/:id", func(c *gin.Context) {

		id := c.Param("id")

		var student Student

		db.First(&student, id)

		if student.ID == 0 {

			c.JSON(http.StatusNotFound, gin.H{

				"message": "Student Not Found",
			})

			return
		}

		var updatedStudent Student

		c.BindJSON(&updatedStudent)

		student.Name = updatedStudent.Name
		student.Age = updatedStudent.Age

		db.Save(&student)

		c.JSON(http.StatusOK, student)
	})

	// DELETE STUDENT
	router.DELETE("/students/:id", func(c *gin.Context) {

		id := c.Param("id")

		var student Student

		db.First(&student, id)

		if student.ID == 0 {

			c.JSON(http.StatusNotFound, gin.H{

				"message": "Student Not Found",
			})

			return
		}

		db.Delete(&student)

		c.JSON(http.StatusOK, gin.H{

			"message": "Student Deleted Successfully",
		})
	})

	router.Run(":8080")
}
```

---

# 📌 Test APIs

---

# ✅ GET

```text
GET /students
```

---

# ✅ POST

```text
POST /students
```

Body:

```json
{
  "name":"Dnyaneshwar",
  "age":24
}
```

---

# ✅ PUT

```text
PUT /students/1
```

---

# ✅ DELETE

```text
DELETE /students/1
```

---

# 📌 Real Backend Architecture

```text
Frontend
↓
Gin API
↓
GORM ORM
↓
PostgreSQL
```

---

# 📌 Why GORM Important?

GORM used in:
- Production APIs
- Enterprise backend
- Microservices
- SaaS applications

Very important for interviews 🔥

---

# 📌 Common GORM Methods

| Method | Purpose |
|---|---|
| Create() | Insert data |
| Find() | Fetch all |
| First() | Fetch single |
| Save() | Update |
| Delete() | Delete |

---

# 💻 Day 20 Practice Program

```go
package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {

	ID    uint   `gorm:"primaryKey"`
	Name  string
	Price int
}

func main() {

	dsn := "user=dnyaneshwarkokate dbname=studentdb sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic(err)
	}

	db.AutoMigrate(&Product{})

	product := Product{

		Name: "iPhone",
		Price: 120000,
	}

	db.Create(&product)

	fmt.Println("Product Added Successfully")
}
```

---

# 📘 Day 20 Interview Questions & Answers

---

## ❓ Q1: What is ORM?

### ✅ Answer:
ORM maps programming language objects with database tables.

---

## ❓ Q2: What is GORM?

### ✅ Answer:
GORM is ORM library for Go.

---

## ❓ Q3: Why GORM popular?

### ✅ Answer:
Because it simplifies database operations and reduces SQL code.

---

## ❓ Q4: What does AutoMigrate() do?

### ✅ Answer:
Automatically creates or updates database tables.

---

## ❓ Q5: Difference between Find() and First()?

| Find() | First() |
|---|---|
| Multiple records | Single record |

---

## ❓ Q6: What does Create() do?

### ✅ Answer:
Inserts data into database.

---

## ❓ Q7: Why GORM used in backend?

### ✅ Answer:
Makes CRUD operations cleaner and faster.

---

# 📚 Day 20 Summary

Today I learned:
- GORM ORM
- Database models
- Auto migration
- CRUD with GORM
- Gin + GORM integration
- PostgreSQL with GORM

I also practiced:
- Creating models
- Database migrations
- CRUD APIs
- ORM architecture

---

# 🧠 Practice Tasks

✅ Create Product CRUD API  
✅ Create Employee Model  
✅ Practice AutoMigrate  
✅ Build CRUD APIs with GORM  
✅ Test APIs in Postman

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

# ✅ Day 21 — JWT Authentication in Go

---

# 📖 Introduction to JWT Authentication

In modern backend web applications and REST APIs, stateless authentication is essential. Instead of maintaining sessions on the server, we use **JWT (JSON Web Token)**.

JWT allows clients (Frontend, Mobile Apps, Microservices) to authenticate once and send a signed token with every subsequent request via the `Authorization` header.

---

# 📖 What You Will Learn

- What is JWT (JSON Web Token)
- JWT Structure (Header, Payload, Signature)
- Installing `golang-jwt/jwt/v5`
- Creating & Signing JWT Tokens
- Verifying & Parsing JWT Tokens
- Building a Login & Protected Profile API with Gin
- Interview Questions & Answers

---

# 📌 What is JWT?

JWT stands for **JSON Web Token**. It is an open standard (RFC 7519) used to securely transmit information between parties as a JSON object.

Structure of a JWT token:
```text
Header.Payload.Signature
```

- **Header**: Contains the algorithm (e.g. HS256) and token type.
- **Payload**: Contains claims (data like `username`, `role`, `exp` expiration time).
- **Signature**: Used to verify that the sender is who it says it is and to ensure that the message wasn't changed along the way.

---

# 📌 Installing JWT Package

To use JWT in Go, install the official community-v5 package:

```bash
go get github.com/golang-jwt/jwt/v5
```

---

# 📌 Generating a JWT Token Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secretKey := "mysecretkey"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "Dnyanesh",
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("JWT Token:")
	fmt.Println(tokenString)
}
```

---

# 📌 Creating Login & Protected Profile API with Gin

```go
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	router := gin.Default()

	// LOGIN ENDPOINT (Generates JWT)
	router.POST("/login", func(c *gin.Context) {
		var login Login

		err := c.BindJSON(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if login.Username != "admin" || login.Password != "1234" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Credentials",
			})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": login.Username,
			"role":     "admin",
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte("mysecretkey"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	})

	// PROTECTED ROUTE (Requires & Verifies JWT)
	router.GET("/profile", func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token Required",
			})
			return
		}

		tokenString := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("mysecretkey"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or Expired Token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Claims",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Protected Route Accessed",
			"username": claims["username"],
			"role":     claims["role"],
		})
	})

	router.Run(":8080")
}
```

---

# 📘 Day 21 Interview Questions & Answers

---

## ❓ Q1: What is JWT and why is it used?

### ✅ Answer:
JWT (JSON Web Token) is an open standard for securely transmitting information between client and server as a JSON object. It is stateless, self-contained, and widely used for authentication and authorization in REST APIs.

---

## ❓ Q2: What are the three parts of a JWT?

### ✅ Answer:
1. **Header**: Contains token type (JWT) and signing algorithm (e.g. HS256).
2. **Payload**: Contains claims (user identity, roles, expiration).
3. **Signature**: Verifies token integrity using a secret key or private key.

---

## ❓ Q3: How is a JWT verified in Go?

### ✅ Answer:
Using `jwt.Parse()`, which takes the token string and a keyfunc callback that returns the secret signing key. If the token is valid and signature matches, claims can be safely extracted.

---

## ❓ Q4: Where should JWT tokens be stored on the client side?

### ✅ Answer:
Ideally in an `httpOnly` secure cookie to mitigate XSS (Cross-Site Scripting) risks, or in memory / Authorization headers (`Bearer <token>`).

---

# 📚 Day 21 Summary

Today I learned:
- JSON Web Token (JWT) concepts & security
- Creating signed tokens with `golang-jwt/jwt/v5`
- Extracting & verifying claims
- Building a JWT Login API & protecting routes in Gin

---

# ✅ Day 22 — Middleware in Go

---

# 📖 Introduction to Middleware

In web applications and microservices, **Middleware** acts as a bridge between an incoming HTTP request and the final route handler.

Middleware functions intercept, process, authorize, log, or transform requests before passing control to the next handler using `c.Next()`, or abort processing using `c.Abort()`.

Common use cases:
- Logging & Latency tracking
- Authentication & Authorization
- CORS (Cross-Origin Resource Sharing)
- Rate Limiting & Header Injection
- Panic Recovery

---

# 📖 What You Will Learn

- What is Middleware in Go (net/http & Gin)
- Global Middleware vs Route-Specific Middleware
- Using `c.Next()` and `c.Abort()`
- Building Custom Logger Middleware
- Building Custom JWT Authentication Middleware
- Grouping Protected Routes with Middleware
- Day 22 Interview Questions & Answers

---

# 📌 How Middleware Works in Gin

In Gin, middleware functions have the signature `func(*gin.Context)` (or `gin.HandlerFunc`).

```text
Incoming Request -> [ Middleware 1 ] -> [ Middleware 2 ] -> [ Handler ]
                                                                 |
Response        <- [ Middleware 1 ] <- [ Middleware 2 ] <--------+
```

Key control methods:
- `c.Next()`: Passes control to the next handler in the chain.
- `c.Abort()`: Stops executing subsequent handlers (e.g. on unauthorized access).
- `c.Set("key", value)`: Passes contextual data down the execution chain.

---

# 📌 Custom Logger & Request ID Middleware Example

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Custom Logger Middleware
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		fmt.Printf("[LOG] Incoming Request: %s %s\n", c.Request.Method, c.Request.URL.Path)

		c.Next() // Pass to next handler

		latency := time.Since(startTime)
		status := c.Writer.Status()

		fmt.Printf("[LOG] Completed Request: %s %s | Status: %d | Time: %v\n",
			c.Request.Method, c.Request.URL.Path, status, latency)
	}
}

func main() {
	router := gin.New()
	router.Use(LoggerMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.Run(":8080")
}
```

---

# 📌 JWT Authentication Middleware with Route Groups

```go
package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("mysecretkey")

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired JWT token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("username", claims["username"])
			c.Set("role", claims["role"])
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	// Protected API Group
	api := router.Group("/api")
	api.Use(JWTAuthMiddleware())
	{
		api.GET("/profile", func(c *gin.Context) {
			user, _ := c.Get("username")
			c.JSON(http.StatusOK, gin.H{"message": "Access granted", "user": user})
		})
	}

	router.Run(":8080")
}
```

---

# 📘 Day 22 Interview Questions & Answers

---

## ❓ Q1: What is Middleware in web development?

### ✅ Answer:
Middleware is a software layer/function that intercepts incoming HTTP requests before they reach the main request handler. It is used for tasks like logging, authentication, request validation, CORS, and response transformation.

---

## ❓ Q2: What is the difference between `c.Next()` and `c.Abort()` in Gin?

### ✅ Answer:
- `c.Next()`: Suspends execution of current middleware, executes downstream handlers/middlewares, and then resumes execution after `c.Next()` completes.
- `c.Abort()`: Prevents pending handlers in the chain from being called. It is used when a middleware decides to reject a request (e.g. 401 Unauthorized).

---

## ❓ Q3: How do you pass data from a Middleware to a Request Handler in Gin?

### ✅ Answer:
Using `c.Set("key", value)` in the middleware, and retrieving it inside downstream handlers using `value, exists := c.Get("key")`.

---

## ❓ Q4: How do you apply middleware to specific routes or groups in Gin?

### ✅ Answer:
By using route groups:
```go
api := router.Group("/api")
api.Use(AuthMiddleware())
```
Or directly passing the middleware function to individual routes:
```go
router.GET("/admin", AuthMiddleware(), AdminHandler)
```

---

# 📚 Day 22 Summary

Today I learned:
- Middleware fundamentals in Go & Gin
- Creating custom Logger and Request ID middlewares
- Writing reusable JWT Authentication Middleware
- Using `c.Next()`, `c.Abort()`, and `c.Set()`
- Grouping protected API endpoints with `router.Group().Use()`

---

# ⭐ Challenge Progress
✅ Day 01 Completed  
✅ Day 02 Completed  
✅ Day 03 Completed  
✅ Day 04 Completed  
✅ Day 05 Completed  
✅ Day 06 Completed  
✅ Day 07 Completed  
✅ Day 08 Completed  
✅ Day 09 Completed   
✅ Day 10 Completed  
✅ Day 11 Completed  
✅ Day 12 Completed  
✅ Day 13 Completed  
✅ Day 14 Completed   
✅ Day 15 Completed  
✅ Day 16 Completed  
✅ Day 17 Completed  
✅ Day 18 Completed  
✅ Day 19 Completed  
✅ Day 20 Completed  
✅ Day 21 Completed  
✅ Day 22 Completed  
✅ Day 23 Completed  
🚀 Next: Role-Based Authorization in Go

---

# ✅ Day 23 — Password Hashing in Go

---

# 📖 Introduction to Password Hashing

Storing plain text passwords in a database is a critical security vulnerability. If a database leak occurs, user accounts across multiple services become compromised.

To secure user authentication, passwords must be securely hashed using a cryptographic key derivation function such as **Bcrypt** (`golang.org/x/crypto/bcrypt`).

### 🔒 Key Security Concepts:
- **Hashing vs Encryption**: Encryption is two-way (reversible with a key); Hashing is one-way (irreversible).
- **Salting**: Automatically adding random data to the password before hashing to prevent Rainbow Table attacks.
- **Work Factor (Cost)**: Adjusts computational complexity to slow down brute-force attacks.

---

# 📖 What You Will Learn

- Why plain text passwords should never be stored
- Difference between Hashing & Encryption
- Using `golang.org/x/crypto/bcrypt` in Go
- Generating password hashes with `bcrypt.GenerateFromPassword()`
- Verifying passwords using `bcrypt.CompareHashAndPassword()`
- Building a Secure User Registration & Login REST API with Gin
- Day 23 Interview Questions & Answers

---

# 📌 Password Hashing Helper Functions in Go

```go
package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash from plain text password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares plain text password with hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "SuperSecret123"

	// Generate Hash
	hash, _ := HashPassword(password)
	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	// Compare Correct Password
	match := CheckPasswordHash("SuperSecret123", hash)
	fmt.Println("Correct Password Match:", match) // true

	// Compare Wrong Password
	wrongMatch := CheckPasswordHash("WrongPass", hash)
	fmt.Println("Wrong Password Match:  ", wrongMatch) // false
}
```

---

# 📌 User Registration & Login API with Gin and Bcrypt

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]User)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	router := gin.Default()

	// Register Route
	router.POST("/register", func(c *gin.Context) {
		var input struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := HashPassword(input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		users[input.Username] = User{
			Username: input.Username,
			Password: hashedPassword,
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
	})

	// Login Route
	router.POST("/login", func(c *gin.Context) {
		var input struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, exists := users[input.Username]
		if !exists || !CheckPasswordHash(input.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
	})

	router.Run(":8080")
}
```

---

# 📘 Day 23 Interview Questions & Answers

---

## ❓ Q1: Why should you use Bcrypt instead of MD5 or SHA256 for passwords?

### ✅ Answer:
MD5 and SHA256 are fast general-purpose cryptographic hash functions. Because they are fast, attackers can run billions of guesses per second using GPUs or precomputed Rainbow Tables. **Bcrypt** is an adaptive key derivation algorithm specifically designed for passwords; it includes automatic salting and configurable work factor (cost) to slow down brute-force and dictionary attacks.

---

## ❓ Q2: What is "Salt" in password hashing?

### ✅ Answer:
A salt is a unique, randomly generated byte sequence appended to a password before hashing. It ensures that two users with identical passwords end up with completely different hash strings, preventing Rainbow Table attacks. In Bcrypt, the salt is generated automatically and embedded directly within the output hash string.

---

## ❓ Q3: What is the Cost Factor in Bcrypt?

### ✅ Answer:
The Cost Factor determines how many iteration rounds ($2^{cost}$) the hashing algorithm performs. Increasing the cost exponentially increases the time required to hash a password, allowing developers to scale password verification times as hardware speed increases over time.

---

## ❓ Q4: How do you verify a password against a stored Bcrypt hash in Go?

### ✅ Answer:
Using `bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))`. It returns `nil` if the password matches the hash, or an `error` if they do not match.

---

# 📚 Day 23 Summary

Today I learned:
- Importance of secure password hashing in backend development
- Differences between Encryption, Hashing, and Salting
- Using `golang.org/x/crypto/bcrypt` in Go
- Password hashing with `bcrypt.GenerateFromPassword()`
- Password verification with `bcrypt.CompareHashAndPassword()`
- Integrating password hashing into Gin User Registration & Login endpoints

---

# ⭐ Challenge Progress
✅ Day 01 Completed  
✅ Day 02 Completed  
✅ Day 03 Completed  
✅ Day 04 Completed  
✅ Day 05 Completed  
✅ Day 06 Completed  
✅ Day 07 Completed  
✅ Day 08 Completed  
✅ Day 09 Completed   
✅ Day 10 Completed  
✅ Day 11 Completed  
✅ Day 12 Completed  
✅ Day 13 Completed  
✅ Day 14 Completed   
✅ Day 15 Completed  
✅ Day 16 Completed  
✅ Day 17 Completed  
✅ Day 18 Completed  
✅ Day 19 Completed  
✅ Day 20 Completed  
✅ Day 21 Completed  
✅ Day 22 Completed  
✅ Day 23 Completed  
✅ Day 24 Completed  
🚀 Next: Environment Variables in Go

---

# ✅ Day 24 — Role-Based Authorization in Go

---

# 📖 Introduction to Role-Based Authorization (RBAC)

Role-Based Access Control (RBAC) is an authorization paradigm where system permissions are grouped into **Roles** (e.g., `admin`, `manager`, `user`), and users are assigned one or more roles. Instead of checking individual user IDs for permissions, application endpoints restrict access based on the authenticated user's role.

### 🔑 Key Authorization Concepts:
- **Authentication vs Authorization**: Authentication answers *"Who are you?"* (e.g. login with password/JWT). Authorization answers *"What are you allowed to do?"* (e.g. RBAC middleware).
- **JWT Role Claims**: Embedding user roles directly within signed JWT payload claims allows backend APIs to perform stateless authorization checks without querying the database on every request.
- **Middleware-Based Guards**: Protecting API routes by chaining JWT authentication middleware and Role authorization middleware in web frameworks like **Gin**.

---

# 📖 What You Will Learn

- Concept of Role-Based Access Control (RBAC) in API design
- Embedding role claims into JSON Web Tokens (JWT) using `golang-jwt/jwt/v5`
- Creating custom Gin authorization middleware `AuthorizeRole(allowedRoles ...string)`
- Protecting route groups by user roles (`/admin`, `/manager`, `/user`)
- Enforcing HTTP 401 Unauthorized vs HTTP 403 Forbidden responses
- Day 24 Technical Interview Questions & Answers

---

# 📌 Standalone Role-Based Authorization Example in Go

```go
package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleUser    = "user"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecretKey = []byte("super-secret-rbac-key")

func GenerateJWT(username, role string) (string, error) {
	claims := &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func CheckPermission(userRole string, allowedRoles []string) bool {
	for _, r := range allowedRoles {
		if userRole == r {
			return true
		}
	}
	return false
}

func main() {
	token, _ := GenerateJWT("alice", RoleAdmin)
	claims, _ := ValidateJWT(token)

	if CheckPermission(claims.Role, []string{RoleAdmin}) {
		fmt.Println("Access granted to Admin Settings!")
	}
}
```

---

# 📌 Role-Based Authorization API with Gin and JWT Middleware

```go
package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("rbac-secret")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid header format"})
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(parts[1], claims, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func AuthorizeRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found in context"})
			c.Abort()
			return
		}

		roleStr := userRole.(string)
		for _, role := range allowedRoles {
			if roleStr == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: insufficient privileges"})
		c.Abort()
	}
}

func main() {
	router := gin.Default()

	// Public Route
	router.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Public access"})
	})

	// User Routes (user, manager, admin)
	userRoutes := router.Group("/user")
	userRoutes.Use(AuthMiddleware(), AuthorizeRole("user", "manager", "admin"))
	{
		userRoutes.GET("/profile", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "User profile data"})
		})
	}

	// Admin Routes (admin only)
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(AuthMiddleware(), AuthorizeRole("admin"))
	{
		adminRoutes.GET("/settings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Admin settings data"})
		})
	}

	router.Run(":8080")
}
```

---

# 📘 Day 24 Interview Questions & Answers

---

## ❓ Q1: What is the difference between Authentication and Authorization?

### ✅ Answer:
- **Authentication (AuthN)** is the process of verifying *who a user is* (e.g. checking credentials like username/password or validating a JWT signature).
- **Authorization (AuthZ)** is the process of verifying *what permissions an authenticated user has* (e.g. checking if a user with role `user` can access `/admin/settings`).

---

## ❓ Q2: What is the difference between 401 Unauthorized and 403 Forbidden?

### ✅ Answer:
- **`401 Unauthorized`**: Indicates that the client has not authenticated (missing or invalid credentials/token).
- **`403 Forbidden`**: Indicates that the client is authenticated, but does *not* have sufficient permissions to access the requested resource.

---

## ❓ Q3: What is RBAC vs ABAC?

### ✅ Answer:
- **RBAC (Role-Based Access Control)** assigns permissions to predefined roles (e.g., `admin`, `user`). Access is checked against the user's assigned role.
- **ABAC (Attribute-Based Access Control)** dynamically determines access based on attributes of the user, resource, environment, and action (e.g., allow user access to document X only if department matches and time is between 9 AM - 5 PM).

---

## ❓ Q4: Why embed roles inside JWT claims in microservices?

### ✅ Answer:
Embedding roles in JWT claims enables **stateless authorization**. Downstream microservices can verify the JWT signature and extract the user's role directly from the token without needing to query a central database or authentication server for every incoming request.

---

# 📚 Day 24 Summary

Today I learned:
- Difference between Authentication and Authorization
- Concepts of Role-Based Access Control (RBAC) in Go APIs
- Embedding and parsing custom role claims in JWT (`golang-jwt/jwt/v5`)
- Building custom role authorization middleware (`AuthorizeRole`) in Gin
- Protecting endpoints using Gin route groups and multi-role guards
- Proper HTTP status codes (`401 Unauthorized` vs `403 Forbidden`)

---

# ⭐ Challenge Progress
✅ Day 01 Completed  
✅ Day 02 Completed  
✅ Day 03 Completed  
✅ Day 04 Completed  
✅ Day 05 Completed  
✅ Day 06 Completed  
✅ Day 07 Completed  
✅ Day 08 Completed  
✅ Day 09 Completed   
✅ Day 10 Completed  
✅ Day 11 Completed  
✅ Day 12 Completed  
✅ Day 13 Completed  
✅ Day 14 Completed   
✅ Day 15 Completed  
✅ Day 16 Completed  
✅ Day 17 Completed  
✅ Day 18 Completed  
✅ Day 19 Completed  
✅ Day 20 Completed  
✅ Day 21 Completed  
✅ Day 22 Completed  
✅ Day 23 Completed  
✅ Day 24 Completed  
✅ Day 25 Completed  
🚀 Next: Clean Architecture in Go

---

# ✅ Day 25 — Environment Variables in Go

---

# 📖 Introduction to Environment Variables in Go

Environment variables are key-value pairs configured outside application code at the operating system or deployment environment level (e.g. Docker, Kubernetes, systemd). Following the **12-Factor App methodology**, external configuration allows the exact same application binary to run seamlessly across local development, staging, and production environments without changing source code.

### 🔒 Key Configuration Principles:
- **Never Hardcode Secrets**: Database passwords, API credentials, and JWT signing keys must never be committed to source control.
- **`os` Package**: Standard Go library (`os.Getenv`, `os.LookupEnv`, `os.Setenv`) reads variables directly from the host system process environment.
- **`.env` Files & `godotenv`**: During local development, `.env` files manage local environment variables. `github.com/joho/godotenv` loads `.env` key-value pairs into the process environment.
- **Configuration Struct & Fallbacks**: Parsing environment variables into a strongly-typed `Config` struct with sensible fallback defaults ensures safe application startup.

---

# 📖 What You Will Learn

- Reading environment variables using standard library `os.Getenv()`
- Distinguishing between empty and missing variables using `os.LookupEnv()`
- Programmatically setting environment variables with `os.Setenv()`
- Loading `.env` files into environment variables using `github.com/joho/godotenv`
- Building a centralized, strongly-typed `Config` struct with fallback defaults
- Exposing dynamic configuration and health endpoints in a Gin REST API
- Day 25 Technical Interview Questions & Answers

---

# 📌 Standalone Environment Variables Example in Go (`os` package)

```go
package main

import (
	"fmt"
	"os"
)

func getEnvWithDefault(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}

func main() {
	// Set environment variables programmatically
	os.Setenv("APP_NAME", "Go90DaysChallenge")
	os.Setenv("PORT", "8080")

	// Read environment variables
	appName := os.Getenv("APP_NAME")
	fmt.Println("APP_NAME:", appName)

	// Check variable existence with LookupEnv
	if dbHost, exists := os.LookupEnv("DB_HOST"); exists {
		fmt.Println("DB_HOST:", dbHost)
	} else {
		fmt.Println("DB_HOST is not set!")
	}

	// Read with fallback default value
	dbPort := getEnvWithDefault("DB_PORT", "5432")
	fmt.Println("DB_PORT (with fallback):", dbPort)
}
```

---

# 📌 Environment Configured REST API with Gin and `godotenv`

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	AppName     string
	Environment string
	DBHost      string
	DBPort      string
}

var AppConfig Config

func getEnv(key, defaultValue string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return defaultValue
}

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️  Warning: .env file not found, using system environment variables")
	}

	AppConfig = Config{
		Port:        getEnv("PORT", "8085"),
		AppName:     getEnv("APP_NAME", "Default_Go_API"),
		Environment: getEnv("ENVIRONMENT", "development"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
	}
}

func main() {
	LoadConfig()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      "UP",
			"app_name":    AppConfig.AppName,
			"environment": AppConfig.Environment,
		})
	})

	router.Run(":" + AppConfig.Port)
}
```

---

# 📘 Day 25 Interview Questions & Answers

---

## ❓ Q1: What is the difference between `os.Getenv()` and `os.LookupEnv()` in Go?

### ✅ Answer:
- **`os.Getenv(key)`** returns the string value of the environment variable. If the variable is not set, it returns an empty string `""`. However, it cannot differentiate between an unset variable and a variable explicitly set to an empty string.
- **`os.LookupEnv(key)`** returns two values: `(value string, exists bool)`. The boolean `exists` indicates whether the environment variable is present in the process environment, even if set to `""`.

---

## ❓ Q2: Why is `.env.example` committed to Git, while `.env` is added to `.gitignore`?

### ✅ Answer:
- **`.env`** contains real credentials (API keys, DB passwords, secret tokens) and local machine settings. It must be listed in `.gitignore` to prevent sensitive secrets from leaking to source repositories.
- **`.env.example`** contains non-sensitive key names and sample placeholder values. It acts as a template and documentation for developers setting up the application environment.

---

## ❓ Q3: What is the 12-Factor App principle regarding configuration?

### ✅ Answer:
The 12-Factor App methodology states that an application's **config should be strictly separated from code**. Configuration that varies between deployments (staging, production, dev) should be stored in environment variables, allowing identical build artifacts to run in any environment.

---

## ❓ Q4: What popular packages exist in Go for configuration management?

### ✅ Answer:
- **`github.com/joho/godotenv`**: Lightweight package for loading `.env` files into process environment variables.
- **`github.com/spf13/viper`**: Complete configuration solution supporting JSON, TOML, YAML, environment variables, flags, and remote config systems (Etcd, Consul).
- **`github.com/kelseyhightower/envconfig`**: Decodes environment variables directly into Go structs using struct tags.

---

# 📚 Day 25 Summary

Today I learned:
- Importance of externalizing configuration following 12-Factor App methodology
- Reading environment variables with `os.Getenv()` and `os.LookupEnv()` in standard library `os`
- Loading local `.env` files into Go environment using `github.com/joho/godotenv`
- Designing a centralized, strongly-typed `Config` struct with fallback default values
- Building a Gin REST API that dynamically configures port, database connection settings, and environment state
- Security best practices: keeping secrets out of Git using `.env.example` templates

---

# ⭐ Challenge Progress
✅ Day 01 Completed  
✅ Day 02 Completed  
✅ Day 03 Completed  
✅ Day 04 Completed  
✅ Day 05 Completed  
✅ Day 06 Completed  
✅ Day 07 Completed  
✅ Day 08 Completed  
✅ Day 09 Completed   
✅ Day 10 Completed  
✅ Day 11 Completed  
✅ Day 12 Completed  
✅ Day 13 Completed  
✅ Day 14 Completed   
✅ Day 15 Completed  
✅ Day 16 Completed  
✅ Day 17 Completed  
✅ Day 18 Completed  
✅ Day 19 Completed  
✅ Day 20 Completed  
✅ Day 21 Completed  
✅ Day 22 Completed  
✅ Day 23 Completed  
✅ Day 24 Completed  
✅ Day 25 Completed  
✅ Day 26 Completed  
🚀 Next: Repository Pattern in Go

---

# ✅ Day 26 — Clean Architecture in Go

---

# 📖 Introduction to Clean Architecture in Go

Clean Architecture (advocated by Robert C. Martin / Uncle Bob) is a software design pattern aimed at creating applications with high maintainability, testability, and independence from external frameworks, databases, or UI libraries.

### 🏗️ Layer Responsibilities in Go:

1. **Domain Layer (`domain/`)**: The innermost layer. Contains pure enterprise entities, data structures, and contract interfaces (`UserRepository`, `UserUseCase`). Has zero external dependencies.
2. **Use Case Layer (`usecase/`)**: Implements business rules and workflows. Depends only on domain interfaces and entities.
3. **Repository Layer (`repository/`)**: Handles database persistence (PostgreSQL, GORM, or In-Memory). Implements `domain.UserRepository`.
4. **Delivery/Handler Layer (`handler/`)**: Handles web transport details (Gin, HTTP routes, JSON parsing, DTO binding). Implements HTTP handlers.

```text
       +---------------------------------------------+
       |         Delivery / Handler Layer            |
       |       (Gin REST API / Controllers)          |
       +---------------------------------------------+
                              |
                              v
       +---------------------------------------------+
       |          Use Case / Service Layer           |
       |            (Business Logic)                 |
       +---------------------------------------------+
                              |
                              v
       +---------------------------------------------+
       |           Domain / Entity Layer             |
       |     (Structs & Interface Contracts)         |
       +---------------------------------------------+
                              ^
                              |
       +---------------------------------------------+
       |          Repository / Data Layer            |
       |       (PostgreSQL / GORM / In-Memory)       |
       +---------------------------------------------+
```

---

# 📖 What You Will Learn

- Principles of Clean Architecture and Dependency Inversion in Go
- Organizing Go packages by layer (`domain`, `repository`, `usecase`, `handler`)
- Defining interface contracts in the domain layer
- Implementing dependency injection across layers
- Decoupling business logic from web frameworks (Gin) and data storage
- Day 26 Technical Interview Questions & Answers

---

# 📌 Code Walkthrough: Clean Architecture Implementation

### 1. Domain Layer (`domain/user.go`)
```go
package domain

import "errors"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user with this email already exists")
)

type UserRepository interface {
	Create(user *User) (*User, error)
	GetByID(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]User, error)
}

type UserUseCase interface {
	RegisterUser(name, email, role string) (*User, error)
	GetUserByID(id int) (*User, error)
	ListUsers() ([]User, error)
}
```

### 2. Repository Layer (`repository/user_repository.go`)
```go
package repository

import (
	"sync"
	"day-26/domain"
)

type inMemoryUserRepository struct {
	mu     sync.RWMutex
	users  map[int]domain.User
	nextID int
}

func NewInMemoryUserRepository() domain.UserRepository {
	return &inMemoryUserRepository{
		users:  make(map[int]domain.User),
		nextID: 1,
	}
}

func (r *inMemoryUserRepository) Create(user *domain.User) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user.ID = r.nextID
	r.users[user.ID] = *user
	r.nextID++
	return user, nil
}

func (r *inMemoryUserRepository) GetByID(id int) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id]
	if !ok {
		return nil, domain.ErrUserNotFound
	}
	return &u, nil
}

func (r *inMemoryUserRepository) GetByEmail(email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, u := range r.users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, nil
}

func (r *r *inMemoryUserRepository) GetAll() ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]domain.User, 0, len(r.users))
	for _, u := range r.users {
		list = append(list, u)
	}
	return list, nil
}
```

### 3. Use Case Layer (`usecase/user_usecase.go`)
```go
package usecase

import "day-26/domain"

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{userRepo: repo}
}

func (u *userUseCase) RegisterUser(name, email, role string) (*domain.User, error) {
	existing, _ := u.userRepo.GetByEmail(email)
	if existing != nil {
		return nil, domain.ErrUserAlreadyExists
	}
	newUser := &domain.User{Name: name, Email: email, Role: role}
	return u.userRepo.Create(newUser)
}

func (u *userUseCase) GetUserByID(id int) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUseCase) ListUsers() ([]domain.User, error) {
	return u.userRepo.GetAll()
}
```

### 4. Main Entry Point (`main.go`)
```go
package main

import (
	"day-26/handler"
	"day-26/repository"
	"day-26/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewInMemoryUserRepository()
	useCase := usecase.NewUserUseCase(repo)
	userHandler := handler.NewUserHandler(useCase)

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.RegisterUser)
		api.GET("/users", userHandler.ListUsers)
		api.GET("/users/:id", userHandler.GetUserByID)
	}

	router.Run(":8086")
}
```

---

# 📘 Day 26 Interview Questions & Answers

---

## ❓ Q1: What is the primary benefit of Clean Architecture in backend development?

### ✅ Answer:
The primary benefit is **Separation of Concerns and Testability**. By isolating core business logic (Use Cases) from external frameworks (Gin) and storage mechanisms (PostgreSQL, Redis), developers can write unit tests for business logic without mocking HTTP request objects or spinning up real database instances.

---

## ❓ Q2: What is the Dependency Inversion Principle (DIP) and how is it used in Go Clean Architecture?

### ✅ Answer:
Dependency Inversion states that high-level modules (Use Cases) should not depend on low-level modules (Database Repositories); both should depend on abstractions (Go Interfaces). In Go, interface contracts are defined in the `domain` layer. Use Cases accept interface dependencies rather than concrete struct implementations.

---

## ❓ Q3: How do you handle database transactions in Clean Architecture?

### ✅ Answer:
Database transactions can be managed by defining a `TxManager` or `UnitOfWork` interface in the `domain` layer. The Use Case initiates the transaction interface, passing the transaction context down to repository methods, ensuring atomic operations while keeping SQL/ORM logic out of the Use Case.

---

## ❓ Q4: Why shouldn't HTTP Handlers call Repository methods directly?

### ✅ Answer:
Bypassing Use Cases violates layer separation. HTTP Handlers should only handle transport concerns (JSON serialization, HTTP status codes) and delegate domain rules to Use Cases. Directly calling Repositories duplicates business validation logic across multiple handlers and makes testing difficult.

---

# 📚 Day 26 Summary

Today I learned:
- Core principles of Clean Architecture and layered design in Go
- Decoupling domain entities, business logic (Use Cases), persistence (Repositories), and HTTP handlers
- Using Go interfaces to satisfy Dependency Inversion
- Assembling dependencies via constructor injection in `main.go`
- Unit testing advantages provided by interface-based architecture

---

# ⭐ Challenge Progress
✅ Day 01 Completed  
✅ Day 02 Completed  
✅ Day 03 Completed  
✅ Day 04 Completed  
✅ Day 05 Completed  
✅ Day 06 Completed  
✅ Day 07 Completed  
✅ Day 08 Completed  
✅ Day 09 Completed   
✅ Day 10 Completed  
✅ Day 11 Completed  
✅ Day 12 Completed  
✅ Day 13 Completed  
✅ Day 14 Completed   
✅ Day 15 Completed  
✅ Day 16 Completed  
✅ Day 17 Completed  
✅ Day 18 Completed  
✅ Day 19 Completed  
✅ Day 20 Completed  
✅ Day 21 Completed  
✅ Day 22 Completed  
✅ Day 23 Completed  
✅ Day 24 Completed  
✅ Day 25 Completed  
✅ Day 26 Completed  
✅ Day 27 Completed  
🚀 Next: Dependency Injection in Go

---

# ✅ Day 27 — Repository Pattern in Go

---

# 📖 Introduction to Repository Pattern in Go

The **Repository Pattern** is a software design pattern that abstracts data access logic and mediates between the domain/business layer and the persistence layer (SQL, NoSQL, ORM, or In-Memory).

By decoupling persistence mechanisms from domain models and business logic:
- Business logic (`usecase`) operates purely on **Go Interfaces** (`domain.ProductRepository`).
- Swapping database drivers (e.g., switching from In-Memory to PostgreSQL / GORM or MongoDB) requires **zero changes** to business logic or HTTP handlers.
- Writing fast, isolated unit tests becomes effortless because repositories can be mocked or swapped with in-memory implementations.

---

# 🏗️ Architecture & Component Flow

```text
  +-------------------------------------------------------------+
  |                   HTTP / Delivery Layer                     |
  |                (handler/product_handler.go)                 |
  +-------------------------------------------------------------+
                               |
                               v
  +-------------------------------------------------------------+
  |                  Business Use Case Layer                    |
  |               (usecase/product_usecase.go)                  |
  +-------------------------------------------------------------+
                               |
                               | (Depends on Interface)
                               v
  +-------------------------------------------------------------+
  |                   Domain Interface Contract                 |
  |                 (domain.ProductRepository)                  |
  +-------------------------------------------------------------+
                 /                             \
                /                               \
   (Implements)                                 (Implements)
              v                                   v
+-----------------------------+   +-----------------------------+
|    In-Memory Repository     |   |       GORM Repository       |
| (repository/memory_repo.go) |   |  (repository/gorm_repo.go)  |
+-----------------------------+   +-----------------------------+
```

---

# 📖 What You Will Learn

- Designing clean repository interfaces in `domain/` package
- Implementing thread-safe In-Memory storage repositories with `sync.RWMutex`
- Implementing relational database repositories using `gorm.io/gorm`
- Unit testing repository implementations independently using Go `testing` framework and SQLite in-memory DB
- Wiring repositories, use cases, and Gin handlers via Constructor Dependency Injection
- Day 27 Technical Interview Questions & Answers

---

# 📌 Code Walkthrough: Repository Pattern Implementation

### 1. Domain Model & Interface Contract (`domain/product.go`)

```go
package domain

import (
	"errors"
	"time"
)

type Product struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	SKU       string    `json:"sku" gorm:"uniqueIndex;not null"`
	Price     float64   `json:"price" gorm:"not null"`
	Stock     int       `json:"stock" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	ErrProductNotFound = errors.New("product not found")
	ErrDuplicateSKU    = errors.New("product with this SKU already exists")
	ErrInvalidProduct  = errors.New("product details are invalid")
)

type ProductRepository interface {
	Create(product *Product) (*Product, error)
	GetByID(id int) (*Product, error)
	GetBySKU(sku string) (*Product, error)
	GetAll() ([]Product, error)
	Update(id int, product *Product) (*Product, error)
	Delete(id int) error
}

type ProductUseCase interface {
	CreateProduct(name, sku string, price float64, stock int) (*Product, error)
	GetProductByID(id int) (*Product, error)
	GetProductBySKU(sku string) (*Product, error)
	ListProducts() ([]Product, error)
	UpdateProduct(id int, name string, price float64, stock int) (*Product, error)
	DeleteProduct(id int) error
}
```

---

### 2. Thread-Safe In-Memory Repository (`repository/memory_repository.go`)

```go
package repository

import (
	"sync"
	"time"

	"day-27/domain"
)

type memoryProductRepository struct {
	mu       sync.RWMutex
	products map[int]domain.Product
	nextID   int
}

func NewMemoryProductRepository() domain.ProductRepository {
	return &memoryProductRepository{
		products: make(map[int]domain.Product),
		nextID:   1,
	}
}

func (r *memoryProductRepository) Create(p *domain.Product) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, item := range r.products {
		if item.SKU == p.SKU {
			return nil, domain.ErrDuplicateSKU
		}
	}

	p.ID = r.nextID
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
	r.products[p.ID] = *p
	r.nextID++

	return p, nil
}

func (r *memoryProductRepository) GetByID(id int) (*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, exists := r.products[id]
	if !exists {
		return nil, domain.ErrProductNotFound
	}
	return &p, nil
}

func (r *memoryProductRepository) GetBySKU(sku string) (*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, p := range r.products {
		if p.SKU == sku {
			return &p, nil
		}
	}
	return nil, domain.ErrProductNotFound
}

func (r *memoryProductRepository) GetAll() ([]domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]domain.Product, 0, len(r.products))
	for _, p := range r.products {
		list = append(list, p)
	}
	return list, nil
}

func (r *memoryProductRepository) Update(id int, p *domain.Product) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, exists := r.products[id]
	if !exists {
		return nil, domain.ErrProductNotFound
	}

	if p.Name != "" {
		existing.Name = p.Name
	}
	if p.Price > 0 {
		existing.Price = p.Price
	}
	if p.Stock >= 0 {
		existing.Stock = p.Stock
	}
	existing.UpdatedAt = time.Now()

	r.products[id] = existing
	return &existing, nil
}

func (r *memoryProductRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[id]; !exists {
		return domain.ErrProductNotFound
	}

	delete(r.products, id)
	return nil
}
```

---

### 3. GORM DB Repository (`repository/gorm_repository.go`)

```go
package repository

import (
	"errors"

	"day-27/domain"
	"gorm.io/gorm"
)

type gormProductRepository struct {
	db *gorm.DB
}

func NewGORMProductRepository(db *gorm.DB) domain.ProductRepository {
	return &gormProductRepository{db: db}
}

func (r *gormProductRepository) Create(p *domain.Product) (*domain.Product, error) {
	var existing domain.Product
	err := r.db.Where("sku = ?", p.SKU).First(&existing).Error
	if err == nil {
		return nil, domain.ErrDuplicateSKU
	}

	if err := r.db.Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (r *gormProductRepository) GetByID(id int) (*domain.Product, error) {
	var p domain.Product
	if err := r.db.First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (r *gormProductRepository) GetBySKU(sku string) (*domain.Product, error) {
	var p domain.Product
	if err := r.db.Where("sku = ?", sku).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (r *gormProductRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *gormProductRepository) Update(id int, p *domain.Product) (*domain.Product, error) {
	existing, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	updates := map[string]interface{}{}
	if p.Name != "" {
		updates["name"] = p.Name
	}
	if p.Price > 0 {
		updates["price"] = p.Price
	}
	if p.Stock >= 0 {
		updates["stock"] = p.Stock
	}

	if err := r.db.Model(existing).Updates(updates).Error; err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *gormProductRepository) Delete(id int) error {
	result := r.db.Delete(&domain.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrProductNotFound
	}
	return nil
}
```

---

### 4. Unit Testing Repositories (`repository/repository_test.go`)

```go
package repository_test

import (
	"testing"

	"day-27/domain"
	"day-27/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func runRepositoryTests(t *testing.T, repo domain.ProductRepository) {
	p1 := &domain.Product{Name: "Mechanical Keyboard", SKU: "KB-101", Price: 99.99, Stock: 15}

	created, err := repo.Create(p1)
	if err != nil || created.ID == 0 {
		t.Fatalf("failed to create product: %v", err)
	}

	pDup := &domain.Product{Name: "Another Keyboard", SKU: "KB-101", Price: 80.00, Stock: 5}
	if _, err := repo.Create(pDup); err != domain.ErrDuplicateSKU {
		t.Errorf("expected ErrDuplicateSKU, got %v", err)
	}

	fetched, err := repo.GetByID(created.ID)
	if err != nil || fetched.Name != p1.Name {
		t.Errorf("failed to fetch created product correctly")
	}

	updated, err := repo.Update(created.ID, &domain.Product{Price: 109.99, Stock: 20})
	if err != nil || updated.Price != 109.99 {
		t.Errorf("failed to update product price")
	}

	if err := repo.Delete(created.ID); err != nil {
		t.Errorf("failed to delete product: %v", err)
	}
}

func TestMemoryProductRepository(t *testing.T) {
	runRepositoryTests(t, repository.NewMemoryProductRepository())
}

func TestGORMProductRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite memory db: %v", err)
	}
	db.AutoMigrate(&domain.Product{})
	runRepositoryTests(t, repository.NewGORMProductRepository(db))
}
```

---

### 5. Application Entry Point (`main.go`)

```go
package main

import (
	"fmt"
	"log"

	"day-27/domain"
	"day-27/handler"
	"day-27/repository"
	"day-27/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 Day 27: Repository Pattern in Go")

	// Plug in desired storage repository (Memory or GORM)
	var productRepo domain.ProductRepository = repository.NewMemoryProductRepository()

	productUseCase := usecase.NewProductUseCase(productRepo)
	productHandler := handler.NewProductHandler(productUseCase)

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products", productHandler.ListProducts)
		api.GET("/products/:id", productHandler.GetProductByID)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	log.Println("⚡ Server listening on http://localhost:8087")
	router.Run(":8087")
}
```

---

# 📘 Day 27 Interview Questions & Answers

---

## ❓ Q1: What is the main purpose of the Repository Pattern in Go?

### ✅ Answer:
The Repository Pattern abstracts data persistence details behind an interface contract. It decouples domain models and business logic (`usecases`) from database-specific drivers or ORMs (SQL, GORM, Mongo, Redis). This promotes clean separation of concerns, simplifies swapping storage backends, and enables fast unit testing with mock or in-memory repositories.

---

## ❓ Q2: Where should the Repository Interface contract be defined in a Clean Architecture Go project?

### ✅ Answer:
The Repository Interface (e.g., `type ProductRepository interface`) must be defined in the **Domain Layer** (`domain/` package). According to the **Dependency Inversion Principle (DIP)**, high-level modules should depend on abstractions defined within the core domain, not on concrete implementations in the infrastructure or persistence layers.

---

## ❓ Q3: How does the Repository Pattern improve Unit Testing?

### ✅ Answer:
Because business logic depends on an interface rather than concrete SQL/ORM structs, unit tests can supply a lightweight In-Memory repository or a generated mock (e.g., `mockery` or `gomock`). Tests execute in milliseconds without needing to start a real database server or execute slow network/file I/O operations.

---

## ❓ Q4: What is the difference between an Active Record Pattern and a Repository Pattern?

### ✅ Answer:
- **Active Record**: Database columns and tables map directly to entity methods where entity structs perform their own SQL/database operations (e.g., `user.Save()`, `user.Delete()`).
- **Repository**: Entities are pure data structs without persistence methods. A separate Repository component handles all queries and storage operations on entities (`repo.Create(user)`). Repository pattern provides much cleaner separation of concerns and testability.

---

# 📚 Day 27 Summary

Today I learned:
- How to implement the **Repository Pattern** in Go to abstract persistence operations.
- Defining repository interface contracts in `domain/` package.
- Creating multiple repository implementations (Thread-Safe In-Memory map & GORM DB adapter).
- Writing table-driven repository unit tests with Go's `testing` package and SQLite in-memory engine.
- Integrating repositories with Use Case business logic and Gin HTTP handler endpoints.

---

# ⭐ Challenge Progress
✅ Day 01 Completed  
✅ Day 02 Completed  
✅ Day 03 Completed  
✅ Day 04 Completed  
✅ Day 05 Completed  
✅ Day 06 Completed  
✅ Day 07 Completed  
✅ Day 08 Completed  
✅ Day 09 Completed   
✅ Day 10 Completed  
✅ Day 11 Completed  
✅ Day 12 Completed  
✅ Day 13 Completed  
✅ Day 14 Completed   
✅ Day 15 Completed  
✅ Day 16 Completed  
✅ Day 17 Completed  
✅ Day 18 Completed  
✅ Day 19 Completed  
✅ Day 20 Completed  
✅ Day 21 Completed  
✅ Day 22 Completed  
✅ Day 23 Completed  
✅ Day 24 Completed  
✅ Day 25 Completed  
✅ Day 26 Completed  
✅ Day 27 Completed  
✅ Day 28 Completed  
🚀 Next: Logging System in Go

---

# ✅ Day 28 — Dependency Injection in Go

---

# 📖 Introduction to Dependency Injection in Go

**Dependency Injection (DI)** is a fundamental software design technique in object-oriented and functional architecture that implements **Inversion of Control (IoC)**. Instead of a struct or component creating its own dependencies (such as initializing database clients, HTTP clients, logger instances, or notification services internally), dependencies are **supplied ("injected") from the outside**, usually through constructor functions.

In Go, Dependency Injection directly upholds the **Dependency Inversion Principle (DIP)** from SOLID principles:
1. **High-level modules** (business use cases & HTTP handlers) should not depend on low-level modules (database drivers, SQL queries, SMTP servers). Both should depend on **abstractions** (Go interfaces).
2. **Abstractions** should not depend on details. Details (concrete implementations) should depend on abstractions.

---

# 🏗️ Architecture & Component Flow

```text
  +---------------------------------------------------------------+
  |                     Main Entrypoint / Container               |
  |                       (di/container.go & main.go)             |
  +---------------------------------------------------------------+
           |                        |                       |
  Constructs & Injects    Constructs & Injects   Constructs & Injects
           |                        |                       |
           v                        v                       v
  +--------------------+   +--------------------+  +------------------+
  |  Console Logger    |   | Memory Order Repo  |  |  Email Notifier  |
  +--------------------+   +--------------------+  +------------------+
           |                        |                       |
  Satisfies interface      Satisfies interface     Satisfies interface
           |                        |                       |
           v                        v                       v
  +---------------------------------------------------------------+
  |                    Domain Interface Contracts                 |
  |          (domain.Logger, OrderRepository, NotificationService)|
  +---------------------------------------------------------------+
                                    ^
                                    | Injected via Constructor
  +---------------------------------------------------------------+
  |                     Business Use Case Layer                   |
  |                   (usecase/order_usecase.go)                  |
  +---------------------------------------------------------------+
                                    ^
                                    | Injected via Constructor
  +---------------------------------------------------------------+
  |                    HTTP Handler / Delivery                    |
  |                   (handler/order_handler.go)                  |
  +---------------------------------------------------------------+
```

---

# 🎯 Why Dependency Injection Matters in Go

- **Effortless Unit Testing & Mocking**: Business logic can be tested in isolation in milliseconds without connecting to external databases or live SMTP servers.
- **Decoupled Architecture**: Components can be swapped (e.g., swapping `MemoryOrderRepository` with `PostgresOrderRepository` or `EmailService` with `MockEmailService`) without altering business logic.
- **Explicit Component Contracts**: Dependencies are explicitly declared in constructor parameters (`NewOrderUseCase(...)`), making codebase dependencies clear and readable.
- **Single Responsibility Principle**: Individual components focus solely on their core logic rather than managing the lifecycles of their dependencies.

---

# 💡 Key DI Patterns in Go

### 1. Constructor Injection (Idiomatic Go)
In Go, constructor functions (`New...`) accepting interface types are the standard and most explicit form of dependency injection:

```go
func NewOrderUseCase(
    repo domain.OrderRepository,
    notifier domain.NotificationService,
    logger domain.Logger,
) *OrderUseCase {
    return &OrderUseCase{
        repo:     repo,
        notifier: notifier,
        logger:   logger,
    }
}
```

### 2. Functional Options Pattern for Dynamic DI
When dependencies or configurations are optional, Go developers use the **Functional Options Pattern**:

```go
type Option func(*OrderUseCase)

func WithLogger(logger domain.Logger) Option {
    return func(u *OrderUseCase) {
        if logger != nil {
            u.logger = logger
        }
    }
}

func NewOrderUseCase(repo domain.OrderRepository, opts ...Option) *OrderUseCase {
    uc := &OrderUseCase{repo: repo}
    for _, opt := range opts {
        opt(uc)
    }
    return uc
}
```

### 3. Application DI Container
A centralized container handles instantiating components in topological order (Infrastructure -> Repositories -> Services -> Handlers):

```go
type Container struct {
    Logger       domain.Logger
    OrderRepo    domain.OrderRepository
    Notification domain.NotificationService
    OrderUseCase *usecase.OrderUseCase
    OrderHandler *handler.OrderHandler
}

func NewContainer(cfg Config) *Container { ... }
```

---

# ⚖️ Manual DI vs. Framework DI (Google Wire vs. Uber FX)

| Feature | Manual DI (Idiomatic) | Google Wire | Uber FX |
| :--- | :--- | :--- | :--- |
| **Mechanism** | Handcrafted Constructors & Containers | Code Generation (`wire` CLI) | Runtime Reflection |
| **Safety** | Compile-time safe | Compile-time safe | Runtime errors on missing dependencies |
| **Reflection Overhead** | None | None | Uses `reflect` package |
| **Debugging** | Extremely easy & visible stack traces | Generates readable Go code | Harder to debug runtime graph failures |
| **Best Used For** | Small to Medium Microservices | Medium to Large Codebases | Large Enterprise Applications |

---

# 🛠️ Implementation Walkthrough — Day 28 Project

### 1. Domain Interfaces (`domain/order.go`)

```go
package domain

import (
	"errors"
	"time"
)

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrInvalidOrderAmount = errors.New("order amount must be greater than zero")
	ErrEmptyCustomerEmail = errors.New("customer email cannot be empty")
)

type Order struct {
	ID            string    `json:"id"`
	CustomerEmail string    `json:"customer_email"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreateOrderInput struct {
	CustomerEmail string  `json:"customer_email" binding:"required,email"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
}

type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type NotificationService interface {
	SendOrderConfirmation(order *Order) error
}

type OrderRepository interface {
	Save(order *Order) error
	FindByID(id string) (*Order, error)
	FindAll() ([]*Order, error)
}
```

---

### 2. Concrete Logger Implementation (`logger/logger.go`)

```go
package logger

import (
	"fmt"
	"log"
	"sync"

	"day-28/domain"
)

type ConsoleLogger struct {
	prefix string
}

func NewConsoleLogger(prefix string) domain.Logger {
	return &ConsoleLogger{prefix: prefix}
}

func (l *ConsoleLogger) Info(msg string, args ...interface{}) {
	log.Printf("[INFO] [%s] %s", l.prefix, fmt.Sprintf(msg, args...))
}

func (l *ConsoleLogger) Error(msg string, args ...interface{}) {
	log.Printf("[ERROR] [%s] %s", l.prefix, fmt.Sprintf(msg, args...))
}

type MockLogger struct {
	mu   sync.Mutex
	Logs []string
}

func NewMockLogger() *MockLogger {
	return &MockLogger{Logs: make([]string, 0)}
}

func (m *MockLogger) Info(msg string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Logs = append(m.Logs, fmt.Sprintf("[INFO] "+msg, args...))
}

func (m *MockLogger) Error(msg string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Logs = append(m.Logs, fmt.Sprintf("[ERROR] "+msg, args...))
}
```

---

### 3. Notification Service Implementation (`notification/email_service.go`)

```go
package notification

import (
	"fmt"
	"sync"

	"day-28/domain"
)

type EmailNotificationService struct {
	logger domain.Logger
}

func NewEmailNotificationService(logger domain.Logger) domain.NotificationService {
	return &EmailNotificationService{logger: logger}
}

func (s *EmailNotificationService) SendOrderConfirmation(order *domain.Order) error {
	s.logger.Info("📧 Sending email confirmation for Order ID %s to %s ($%.2f)",
		order.ID, order.CustomerEmail, order.Amount)
	return nil
}

type MockNotificationService struct {
	mu         sync.Mutex
	SentOrders []*domain.Order
	ShouldFail bool
}

func NewMockNotificationService() *MockNotificationService {
	return &MockNotificationService{SentOrders: make([]*domain.Order, 0)}
}

func (m *MockNotificationService) SendOrderConfirmation(order *domain.Order) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.ShouldFail {
		return fmt.Errorf("simulated notification failure")
	}
	m.SentOrders = append(m.SentOrders, order)
	return nil
}
```

---

### 4. Repository Implementation (`repository/order_repository.go`)

```go
package repository

import (
	"sync"

	"day-28/domain"
)

type memoryOrderRepository struct {
	mu     sync.RWMutex
	orders map[string]*domain.Order
}

func NewMemoryOrderRepository() domain.OrderRepository {
	return &memoryOrderRepository{
		orders: make(map[string]*domain.Order),
	}
}

func (r *memoryOrderRepository) Save(order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID] = order
	return nil
}

func (r *memoryOrderRepository) FindByID(id string) (*domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	order, exists := r.orders[id]
	if !exists {
		return nil, domain.ErrOrderNotFound
	}
	return order, nil
}

func (r *memoryOrderRepository) FindAll() ([]*domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*domain.Order, 0, len(r.orders))
	for _, order := range r.orders {
		result = append(result, order)
	}
	return result, nil
}
```

---

### 5. Business Use Case with DI (`usecase/order_usecase.go`)

```go
package usecase

import (
	"fmt"
	"time"

	"day-28/domain"
)

type OrderUseCase struct {
	repo     domain.OrderRepository
	notifier domain.NotificationService
	logger   domain.Logger
}

type Option func(*OrderUseCase)

func WithLogger(logger domain.Logger) Option {
	return func(u *OrderUseCase) {
		if logger != nil {
			u.logger = logger
		}
	}
}

func NewOrderUseCase(
	repo domain.OrderRepository,
	notifier domain.NotificationService,
	logger domain.Logger,
	opts ...Option,
) *OrderUseCase {
	uc := &OrderUseCase{
		repo:     repo,
		notifier: notifier,
		logger:   logger,
	}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

func (u *OrderUseCase) CreateOrder(input domain.CreateOrderInput) (*domain.Order, error) {
	if input.Amount <= 0 {
		u.logger.Error("Validation failed: order amount must be positive (received %.2f)", input.Amount)
		return nil, domain.ErrInvalidOrderAmount
	}
	if input.CustomerEmail == "" {
		u.logger.Error("Validation failed: customer email is empty")
		return nil, domain.ErrEmptyCustomerEmail
	}

	order := &domain.Order{
		ID:            fmt.Sprintf("ORD-%d", time.Now().UnixNano()),
		CustomerEmail: input.CustomerEmail,
		Amount:        input.Amount,
		Status:        "COMPLETED",
		CreatedAt:     time.Now(),
	}

	if err := u.repo.Save(order); err != nil {
		u.logger.Error("Failed to save order %s: %v", order.ID, err)
		return nil, fmt.Errorf("repository save error: %w", err)
	}

	u.logger.Info("Order %s saved successfully", order.ID)

	if u.notifier != nil {
		if err := u.notifier.SendOrderConfirmation(order); err != nil {
			u.logger.Error("Failed to send order confirmation for %s: %v", order.ID, err)
		}
	}

	return order, nil
}
```

---

### 6. Application Container (`di/container.go`)

```go
package di

import (
	"day-28/domain"
	"day-28/handler"
	"day-28/logger"
	"day-28/notification"
	"day-28/repository"
	"day-28/usecase"

	"github.com/gin-gonic/gin"
)

type Container struct {
	Logger       domain.Logger
	OrderRepo    domain.OrderRepository
	Notification domain.NotificationService
	OrderUseCase *usecase.OrderUseCase
	OrderHandler *handler.OrderHandler
}

type Config struct {
	Env string
}

func NewContainer(cfg Config) *Container {
	var appLogger domain.Logger
	var notifier domain.NotificationService

	if cfg.Env == "test" {
		appLogger = logger.NewMockLogger()
		notifier = notification.NewMockNotificationService()
	} else {
		appLogger = logger.NewConsoleLogger("APP")
		notifier = notification.NewEmailNotificationService(appLogger)
	}

	orderRepo := repository.NewMemoryOrderRepository()
	orderUseCase := usecase.NewOrderUseCase(orderRepo, notifier, appLogger)
	orderHandler := handler.NewOrderHandler(orderUseCase, appLogger)

	return &Container{
		Logger:       appLogger,
		OrderRepo:    orderRepo,
		Notification: notifier,
		OrderUseCase: orderUseCase,
		OrderHandler: orderHandler,
	}
}

func (c *Container) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("/orders", c.OrderHandler.CreateOrder)
		api.GET("/orders", c.OrderHandler.ListOrders)
		api.GET("/orders/:id", c.OrderHandler.GetOrder)
	}
}
```

---

### 7. Unit Testing with Mock Injection (`usecase/order_usecase_test.go`)

```go
package usecase_test

import (
	"testing"

	"day-28/domain"
	"day-28/logger"
	"day-28/notification"
	"day-28/repository"
	"day-28/usecase"
)

func TestCreateOrder_Success(t *testing.T) {
	mockRepo := repository.NewMemoryOrderRepository()
	mockNotifier := notification.NewMockNotificationService()
	mockLogger := logger.NewMockLogger()

	orderUC := usecase.NewOrderUseCase(mockRepo, mockNotifier, mockLogger)

	input := domain.CreateOrderInput{
		CustomerEmail: "alice@example.com",
		Amount:        150.75,
	}

	order, err := orderUC.CreateOrder(input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if order.ID == "" {
		t.Errorf("expected generated Order ID, got empty string")
	}

	if len(mockNotifier.SentOrders) != 1 {
		t.Errorf("expected 1 sent notification, got %d", len(mockNotifier.SentOrders))
	}
}
```

---

# 📘 Day 28 Interview Questions & Answers

---

## ❓ Q1: What is Dependency Injection (DI) and how does it relate to SOLID principles in Go?

### ✅ Answer:
Dependency Injection is a design pattern implementing **Inversion of Control (IoC)**, where a struct receives its dependencies from an external caller rather than constructing them internally.
It directly relates to the **Dependency Inversion Principle (DIP)**: High-level modules (usecases/handlers) depend on abstractions (Go interfaces) rather than concrete implementations (SQL DB, HTTP clients, SMTP servers).

---

## ❓ Q2: What are the main advantages of using manual Constructor Injection in Go over magic DI frameworks?

### ✅ Answer:
1. **Compile-Time Safety**: Type errors or missing dependencies are caught by the Go compiler at compile time, eliminating runtime panics.
2. **Zero Performance Overhead**: No runtime reflection (`reflect` package) or indirect lookup tables are needed.
3. **Transparency & Debuggability**: Navigation to constructors (`NewService(...)`) is clear in IDEs and stack traces are straightforward without framework-generated stack noise.

---

## ❓ Q3: How does Dependency Injection improve Unit Testability?

### ✅ Answer:
When components depend on interfaces rather than concrete structs, unit test suites can pass lightweight **Mock**, **Fake**, or **Stub** implementations into constructors. This isolates business logic testing from external networks, databases, or third-party APIs, allowing thousands of unit tests to execute in milliseconds without external side effects.

---

## ❓ Q4: How does the Functional Options Pattern work with Dependency Injection in Go?

### ✅ Answer:
The Functional Options Pattern uses functions matching `type Option func(*Struct)` to configure optional or secondary dependencies on a struct during initialization. It avoids long parameter lists in constructor functions (`NewService(...)`) and preserves backward compatibility when adding new optional dependencies.

---

## ❓ Q5: What is the difference between Google Wire and Uber FX in the Go ecosystem?

### ✅ Answer:
- **Google Wire**: Uses static analysis and code generation (`wire` CLI) to generate compile-time standard Go constructor calls. It has zero runtime performance impact and guarantees compile-time type safety.
- **Uber FX**: Uses runtime reflection (`reflect`) to automatically discover, construct, and wire dependencies into an application container graph with lifecycle hooks (`OnStart`, `OnStop`).

---

# 📚 Day 28 Summary

Today I learned:
- The core principles of **Dependency Injection (DI)** and **Dependency Inversion Principle (DIP)** in Go.
- Implementing **Constructor Injection** using small, focused Go interfaces.
- Designing an **Application DI Container** for centralized dependency wireup.
- Utilizing the **Functional Options Pattern** for dynamic dependency configuration.
- Writing fast, isolated unit tests by injecting mock dependencies.
- Evaluating **Manual DI vs Compile-Time DI (Google Wire) vs Runtime DI (Uber FX)**.

---

# ⭐ Challenge Progress
✅ Day 01 Completed  
✅ Day 02 Completed  
✅ Day 03 Completed  
✅ Day 04 Completed  
✅ Day 05 Completed  
✅ Day 06 Completed  
✅ Day 07 Completed  
✅ Day 08 Completed  
✅ Day 09 Completed   
✅ Day 10 Completed  
✅ Day 11 Completed  
✅ Day 12 Completed  
✅ Day 13 Completed  
✅ Day 14 Completed   
✅ Day 15 Completed  
✅ Day 16 Completed  
✅ Day 17 Completed  
✅ Day 18 Completed  
✅ Day 19 Completed  
✅ Day 20 Completed  
✅ Day 21 Completed  
✅ Day 22 Completed  
✅ Day 23 Completed  
✅ Day 24 Completed  
✅ Day 25 Completed  
✅ Day 26 Completed  
✅ Day 27 Completed  
✅ Day 28 Completed  
✅ Day 29 Completed  
🚀 Next: Student Management REST API Project in Go

---

# ✅ Day 29 — Logging System in Go

---

# 📖 Introduction to Production Logging in Go

**Logging** is a critical operational foundation for modern backend services and microservices. In production, logs provide visibility into system behavior, aid in debugging errors, monitor health metrics, and enable security auditing.

### Why Unstructured Logging Fails in Production
Traditional unstructured logging (using `fmt.Println` or Go's standard `log.Printf`) outputs plain text strings without standardized keys:
```text
2026/08/10 19:30:00 Failed to save user dnyaneshwar@example.com due to duplicate email
```
Parsing millions of such unstructured log lines in log aggregators (Elasticsearch / ELK, Grafana Loki, Datadog, Splunk) requires complex regular expressions and is error-prone.

### Structured Logging & Log Levels
**Structured Logging** outputs logs in standardized formats (such as JSON or logfmt) with key-value pairs:
```json
{
  "timestamp": "2026-08-10T19:30:00.123Z",
  "level": "WARN",
  "message": "User registration rejected: Email already registered",
  "request_id": "req-trace-103",
  "email": "dnyaneshwar@example.com"
}
```

Common **Log Levels**:
- **`DEBUG`**: Detailed diagnostic information for developers during local debugging.
- **`INFO`**: Normal operational events (e.g., service startup, successful order creation).
- **`WARN`**: Unexpected non-fatal conditions (e.g., invalid user input, duplicate email attempt).
- **`ERROR`**: Actionable errors or failures (e.g., DB connection loss, repository write failures).
- **`FATAL`**: Critical errors causing application termination (`os.Exit(1)`).

---

# 🏗️ Architecture & Contextual Log Flow

```text
  +---------------------------------------------------------------+
  |                     Incoming HTTP Request                     |
  |                  (Header: X-Request-ID / UUID)                |
  +---------------------------------------------------------------+
                                  |
                                  v
  +---------------------------------------------------------------+
  |                 RequestIDMiddleware (Gin)                     |
  |    Generates/Extracts Request ID -> Attaches to context.Ctx   |
  +---------------------------------------------------------------+
                                  |
                                  v
  +---------------------------------------------------------------+
  |             StructuredLoggerMiddleware (Gin)                  |
  |     Intercepts status, method, path, IP, latency, errors     |
  +---------------------------------------------------------------+
                                  |
                                  v
  +---------------------------------------------------------------+
  |                     HTTP Handler Layer                        |
  |             (Passes c.Request.Context() down)                 |
  +---------------------------------------------------------------+
                                  |
                                  v
  +---------------------------------------------------------------+
  |                     Business UseCase Layer                    |
  |           Logger.Info(ctx, "Processing user...", ...)         |
  +---------------------------------------------------------------+
                                  |
                                  v
  +---------------------------------------------------------------+
  |                    Repository / DB Layer                      |
  |           Logger.Debug(ctx, "Executing query...", ...)        |
  +---------------------------------------------------------------+
                                  |
                                  v
  +---------------------------------------------------------------+
  |                     Zap / Slog Core Logger                    |
  |    Extracts Request ID from Context -> Formats JSON Record    |
  +---------------------------------------------------------------+
                   |                             |
                   v                             v
  +-------------------------------+ +-----------------------------+
  |    Stdout (Console Output)    | |   app.log (Persistent File) |
  +-------------------------------+ +-----------------------------+
```

---

# ⚡ Go Logging Ecosystem Comparison

| Logger | Mechanism / Features | Allocation Performance | Structured JSON | Go Stdlib Inclusion |
| :--- | :--- | :--- | :--- | :--- |
| **Standard `log`** | Basic string formatting (`Printf`) | Moderate | ❌ Manual string building | Built-in |
| **Standard `log/slog`** | Structured logging, Handlers, `slog.Attr` | Extremely High | ✅ Built-in JSON & Text Handlers | Built-in (Go 1.21+) |
| **`go.uber.org/zap`** | Ultra high-performance, zero allocation fields | Maximum | ✅ JSON & Console Encoders | External Library |
| **`zerolog`** | Fast JSON logger with fluent API | Maximum | ✅ JSON | External Library |

---

# 🛠️ Implementation Walkthrough — Day 29 Project

### 1. Module Definition (`go.mod`)

```go
module day-29

go 1.22.0

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/google/uuid v1.6.0
	go.uber.org/zap v1.27.0
)
```

---

### 2. Domain Models & Logger Interface (`domain/user.go`)

```go
package domain

import (
	"context"
	"errors"
	"time"
)

type contextKey string

const (
	RequestIDKey contextKey = "X-Request-ID"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("user with this email already exists")
	ErrInvalidEmail       = errors.New("invalid email address format")
	ErrEmptyName          = errors.New("user name cannot be empty")
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Role  string `json:"role"`
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindAll(ctx context.Context) ([]*User, error)
}

type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}
```

---

### 3. Zap & Slog Structured Logger Implementation (`logger/logger.go`)

```go
package logger

import (
	"context"
	"log/slog"
	"os"

	"day-29/domain"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(env string, logFilePath string) (*ZapLogger, error) {
	var encoderConfig zapcore.EncoderConfig

	if env == "production" {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var core zapcore.Core

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	if env == "production" {
		consoleEncoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	consoleSyncer := zapcore.AddSync(os.Stdout)

	var syncers []zapcore.WriteSyncer
	syncers = append(syncers, consoleSyncer)

	if logFilePath != "" {
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
			fileSyncer := zapcore.AddSync(file)
			fileCore := zapcore.NewCore(jsonEncoder, fileSyncer, zap.InfoLevel)

			consoleCore := zapcore.NewCore(consoleEncoder, consoleSyncer, zap.DebugLevel)
			core = zapcore.NewTee(consoleCore, fileCore)
		} else {
			core = zapcore.NewCore(consoleEncoder, consoleSyncer, zap.DebugLevel)
		}
	} else {
		core = zapcore.NewCore(consoleEncoder, consoleSyncer, zap.DebugLevel)
	}

	zapLog := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &ZapLogger{logger: zapLog}, nil
}

func (z *ZapLogger) extractFields(ctx context.Context, keysAndValues ...interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(keysAndValues)/2+1)

	if reqID, ok := ctx.Value(domain.RequestIDKey).(string); ok && reqID != "" {
		fields = append(fields, zap.String("request_id", reqID))
	}

	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			key, ok := keysAndValues[i].(string)
			if !ok {
				key = "invalid_key"
			}
			fields = append(fields, zap.Any(key, keysAndValues[i+1]))
		}
	}

	return fields
}

func (z *ZapLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	fields := z.extractFields(ctx, keysAndValues...)
	z.logger.Info(msg, fields...)
}

func (z *ZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	fields := z.extractFields(ctx, keysAndValues...)
	z.logger.Warn(msg, fields...)
}

func (z *ZapLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	fields := z.extractFields(ctx, keysAndValues...)
	z.logger.Error(msg, fields...)
}

func (z *ZapLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	fields := z.extractFields(ctx, keysAndValues...)
	z.logger.Debug(msg, fields...)
}

func (z *ZapLogger) Sync() error {
	return z.logger.Sync()
}

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(env string) *SlogLogger {
	var handler slog.Handler
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}

	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return &SlogLogger{logger: slog.New(handler)}
}

func (s *SlogLogger) extractAttrs(ctx context.Context, keysAndValues ...interface{}) []interface{} {
	attrs := make([]interface{}, 0, len(keysAndValues)+2)
	if reqID, ok := ctx.Value(domain.RequestIDKey).(string); ok && reqID != "" {
		attrs = append(attrs, "request_id", reqID)
	}
	attrs = append(attrs, keysAndValues...)
	return attrs
}

func (s *SlogLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.logger.InfoContext(ctx, msg, s.extractAttrs(ctx, keysAndValues...)...)
}

func (s *SlogLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.logger.WarnContext(ctx, msg, s.extractAttrs(ctx, keysAndValues...)...)
}

func (s *SlogLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.logger.ErrorContext(ctx, msg, s.extractAttrs(ctx, keysAndValues...)...)
}

func (s *SlogLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.logger.DebugContext(ctx, msg, s.extractAttrs(ctx, keysAndValues...)...)
}
```

---

### 4. Gin Request ID & HTTP Access Middleware (`middleware/logger_middleware.go`)

```go
package middleware

import (
	"context"
	"time"

	"day-29/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}

		c.Header("X-Request-ID", reqID)

		ctx := context.WithValue(c.Request.Context(), domain.RequestIDKey, reqID)
		c.Request = c.Request.WithContext(ctx)
		c.Set(string(domain.RequestIDKey), reqID)

		c.Next()
	}
}

func StructuredLoggerMiddleware(log domain.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		rawQuery := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		fullPath := path
		if rawQuery != "" {
			fullPath = path + "?" + rawQuery
		}

		keysAndValues := []interface{}{
			"status", statusCode,
			"method", method,
			"path", fullPath,
			"ip", clientIP,
			"latency_ms", latency.Milliseconds(),
			"latency_human", latency.String(),
			"user_agent", c.Request.UserAgent(),
		}

		if errorMessage != "" {
			keysAndValues = append(keysAndValues, "error", errorMessage)
		}

		ctx := c.Request.Context()

		if statusCode >= 500 {
			log.Error(ctx, "HTTP Request Failed (Server Error)", keysAndValues...)
		} else if statusCode >= 400 {
			log.Warn(ctx, "HTTP Request Warning (Client Error)", keysAndValues...)
		} else {
			log.Info(ctx, "HTTP Request Completed", keysAndValues...)
		}
	}
}
```

---

### 5. Repository Layer with Logging (`repository/user_repository.go`)

```go
package repository

import (
	"context"
	"sync"
	"time"

	"day-29/domain"
)

type memoryUserRepository struct {
	mu     sync.RWMutex
	users  map[string]*domain.User
	emails map[string]string
	logger domain.Logger
}

func NewMemoryUserRepository(logger domain.Logger) domain.UserRepository {
	return &memoryUserRepository{
		users:  make(map[string]*domain.User),
		emails: make(map[string]string),
		logger: logger,
	}
}

func (r *memoryUserRepository) Save(ctx context.Context, user *domain.User) error {
	start := time.Now()
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug(ctx, "Executing DB query: Save User", "user_id", user.ID, "email", user.Email)

	if existingID, exists := r.emails[user.Email]; exists && existingID != user.ID {
		r.logger.Warn(ctx, "DB Conflict: Email already exists", "email", user.Email, "existing_user_id", existingID)
		return domain.ErrEmailAlreadyExists
	}

	r.users[user.ID] = user
	r.emails[user.Email] = user.ID

	r.logger.Info(ctx, "DB Write successful", "user_id", user.ID, "duration_ms", time.Since(start).Milliseconds())
	return nil
}

func (r *memoryUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	start := time.Now()
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "Executing DB query: FindByID", "user_id", id)

	user, exists := r.users[id]
	if !exists {
		r.logger.Warn(ctx, "DB Record not found", "user_id", id, "duration_ms", time.Since(start).Milliseconds())
		return nil, domain.ErrUserNotFound
	}

	r.logger.Debug(ctx, "DB Record retrieved", "user_id", id, "duration_ms", time.Since(start).Milliseconds())
	return user, nil
}

func (r *memoryUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "Executing DB query: FindByEmail", "email", email)

	id, exists := r.emails[email]
	if !exists {
		return nil, domain.ErrUserNotFound
	}

	return r.users[id], nil
}

func (r *memoryUserRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	start := time.Now()
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "Executing DB query: FindAll")

	result := make([]*domain.User, 0, len(r.users))
	for _, user := range r.users {
		result = append(result, user)
	}

	r.logger.Info(ctx, "DB Records fetched", "count", len(result), "duration_ms", time.Since(start).Milliseconds())
	return result, nil
}
```

---

### 6. Business UseCase Layer (`usecase/user_usecase.go`)

```go
package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"day-29/domain"

	"github.com/google/uuid"
)

type UserUseCase struct {
	repo   domain.UserRepository
	logger domain.Logger
}

func NewUserUseCase(repo domain.UserRepository, logger domain.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, input domain.CreateUserInput) (*domain.User, error) {
	u.logger.Info(ctx, "Processing user registration request", "name", input.Name, "email", input.Email)

	if strings.TrimSpace(input.Name) == "" {
		u.logger.Warn(ctx, "Validation failed: Empty user name")
		return nil, domain.ErrEmptyName
	}

	if !strings.Contains(input.Email, "@") {
		u.logger.Warn(ctx, "Validation failed: Invalid email format", "email", input.Email)
		return nil, domain.ErrInvalidEmail
	}

	existingUser, err := u.repo.FindByEmail(ctx, input.Email)
	if err == nil && existingUser != nil {
		u.logger.Warn(ctx, "User registration rejected: Email already registered", "email", input.Email)
		return nil, domain.ErrEmailAlreadyExists
	}

	role := input.Role
	if role == "" {
		role = "USER"
	}

	now := time.Now()
	user := &domain.User{
		ID:        fmt.Sprintf("usr_%s", uuid.New().String()[:8]),
		Name:      strings.TrimSpace(input.Name),
		Email:     strings.ToLower(strings.TrimSpace(input.Email)),
		Role:      strings.ToUpper(role),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.repo.Save(ctx, user); err != nil {
		u.logger.Error(ctx, "Failed to persist new user to repository", "user_id", user.ID, "error", err.Error())
		return nil, err
	}

	u.logger.Info(ctx, "User successfully registered", "user_id", user.ID, "email", user.Email, "role", user.Role)
	return user, nil
}

func (u *UserUseCase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	u.logger.Debug(ctx, "Fetching user profile by ID", "user_id", id)

	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if err == domain.ErrUserNotFound {
			u.logger.Warn(ctx, "User profile requested but not found", "user_id", id)
		} else {
			u.logger.Error(ctx, "Unexpected error retrieving user", "user_id", id, "error", err.Error())
		}
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) ListUsers(ctx context.Context) ([]*domain.User, error) {
	u.logger.Info(ctx, "Listing all active users")
	return u.repo.FindAll(ctx)
}
```

---

### 7. Main Application Entry Point (`main.go`)

```go
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"day-29/domain"
	"day-29/handler"
	"day-29/logger"
	"day-29/middleware"
	"day-29/repository"
	"day-29/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("🚀 Day 29: Production-Grade Logging System in Go")
	fmt.Println("==================================================")

	logFilePath := "app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()
	zapLog.Info(ctx, "Logger initialized successfully", "log_file", logFilePath, "environment", "development")

	userRepo := repository.NewMemoryUserRepository(zapLog)
	userUC := usecase.NewUserUseCase(userRepo, zapLog)
	userHandler := handler.NewUserHandler(userUC, zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.RegisterUser)
		api.GET("/users", userHandler.ListUsers)
		api.GET("/users/:id", userHandler.GetUserByID)
	}

	createUser(router, "Dnyaneshwar Kokate", "dnyaneshwar@example.com", "ADMIN", "req-trace-101")
	createUser(router, "Duplicate User", "dnyaneshwar@example.com", "USER", "req-trace-103")
	listUsers(router, "req-trace-104")
}

func createUser(router *gin.Engine, name, email, role, requestID string) {
	payload := domain.CreateUserInput{Name: name, Email: email, Role: role}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if requestID != "" {
		req.Header.Set("X-Request-ID", requestID)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
}

func listUsers(router *gin.Engine, requestID string) {
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	if requestID != "" {
		req.Header.Set("X-Request-ID", requestID)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
}
```

---

# 📘 Day 29 Interview Questions & Answers

---

## ❓ Q1: Why is structured JSON logging preferred over standard `fmt.Printf` or `log.Printf` in production backend systems?

### ✅ Answer:
Standard `log.Printf` outputs plain text strings that are difficult to parse and index in log management tools (e.g., Elasticsearch, Grafana Loki, Datadog). Structured JSON logging outputs key-value pairs (`{"level":"INFO","user_id":"123","request_id":"abc"}`), allowing log aggregation systems to automatically index fields. This enables fast searching, metric aggregation, filtering by user ID or request ID, and real-time dashboard visualization.

---

## ❓ Q2: What is Request Tracing / Correlation ID and how is it propagated in Go microservices?

### ✅ Answer:
A Correlation ID (`X-Request-ID`) is a unique string assigned to an incoming HTTP request at the API gateway or ingress middleware.
In Go, it is attached to the request's `context.Context` (`context.WithValue(ctx, RequestIDKey, reqID)`). As the request traverses handlers, use cases, repository queries, and downstream HTTP/gRPC client calls, the context is passed along, allowing the logger to attach `request_id` to every log entry emitted during that request lifecycle.

---

## ❓ Q3: How does `go.uber.org/zap` achieve near-zero memory allocations compared to traditional loggers?

### ✅ Answer:
`Zap` avoids interface boxing (`interface{}`) and runtime reflection by using strongly typed fields (`zap.String()`, `zap.Int()`, `zap.Duration()`). It also uses an internal buffer pool (`zapcore.BufferPool`) to construct log messages without allocating memory on the heap for every log call.

---

## ❓ Q4: What is `log/slog` introduced in Go 1.21, and how does it compare to `zap` and `zerolog`?

### ✅ Answer:
`log/slog` is the standard library structured logging package added in Go 1.21. It defines standard `Logger`, `Record`, and `Handler` interfaces (`slog.TextHandler`, `slog.JSONHandler`). While `zap` and `zerolog` offer slightly higher raw throughput in extreme benchmarks, `log/slog` provides a standardized interface built into the standard library without external third-party dependencies.

---

## ❓ Q5: How do you handle multi-output logging (writing to both Stdout and a persistent log file) in Zap?

### ✅ Answer:
In Zap, multi-output logging is implemented using `zapcore.NewTee(...)` to combine multiple `zapcore.Core` instances. One core writes formatted console output (or JSON) to `os.Stdout` (for container log collectors like Docker/K8s), while a second core writes JSON output to an appended file writer (`os.OpenFile("app.log", ...)`).

---

# 📚 Day 29 Summary

Today I learned:
- The necessity of **Structured JSON Logging** and standard **Log Levels** (`DEBUG`, `INFO`, `WARN`, `ERROR`).
- Integrating **Uber Zap Logger** and Go 1.21+ **`log/slog`** with Clean Architecture.
- Implementing **Request Tracing (`X-Request-ID`)** via Gin middleware and `context.Context` propagation.
- Building custom **Structured HTTP Access Logging Middleware** to measure latency, status codes, and IP addresses.
- Multi-output logging configuration using `zapcore.NewTee` for simultaneous Console and File (`app.log`) persistence.
- Writing unit tests verifying request ID headers and HTTP logging behavior.

---

# ⭐ Challenge Progress
✅ Day 01 Completed  
✅ Day 02 Completed  
✅ Day 03 Completed  
✅ Day 04 Completed  
✅ Day 05 Completed  
✅ Day 06 Completed  
✅ Day 07 Completed  
✅ Day 08 Completed  
✅ Day 09 Completed   
✅ Day 10 Completed  
✅ Day 11 Completed  
✅ Day 12 Completed  
✅ Day 13 Completed  
✅ Day 14 Completed   
✅ Day 15 Completed  
✅ Day 16 Completed  
✅ Day 17 Completed  
✅ Day 18 Completed  
✅ Day 19 Completed  
✅ Day 20 Completed  
✅ Day 21 Completed  
✅ Day 22 Completed  
✅ Day 23 Completed  
✅ Day 24 Completed  
✅ Day 25 Completed  
✅ Day 26 Completed  
✅ Day 27 Completed  
✅ Day 28 Completed  
✅ Day 29 Completed  
✅ Day 30 Completed  
🚀 Next: Advanced CRUD APIs in Go

---

# ✅ Day 30 — Student Management REST API Project

---

# 📖 Overview — Milestone REST API Project

Day 30 marks the **First Major Milestone REST API Project** in the 90-day Go Challenge. This production-grade application consolidates all architecture and infrastructure patterns developed over Days 15 to 29 into a complete, modular, tested Student Management Backend Microservice.

### Key Integrated Technologies & Features
- **Clean Architecture & SOLID Principles**: Decoupled layers (`domain`, `repository`, `usecase`, `handler`, `middleware`, `logger`, `config`, `di`).
- **Dependency Injection (DI)**: Manual constructor injection and a central application container (`di/container.go`).
- **JWT Authentication & RBAC Authorization**: Role-based endpoints for `ADMIN`, `TEACHER`, and `STUDENT` roles using `golang-jwt/jwt/v5`.
- **Security & Hashing**: Bcrypt password hashing (`golang.org/x/crypto/bcrypt`).
- **Structured Logging & Tracing**: High-performance Uber Zap (`go.uber.org/zap`) with custom HTTP access logging and `X-Request-ID` correlation context tracking.
- **RESTful API Operations**: Full CRUD lifecycle for Student resources with status and department filtering.
- **Automated Testing**: Comprehensive HTTP integration test suite (`tests/student_api_test.go`).

---

# 🏗️ Architecture & Component Flow

```text
  +---------------------------------------------------------------------------------+
  |                            Incoming HTTP Client / Postman                       |
  +---------------------------------------------------------------------------------+
                                           |
                    Header: Authorization: Bearer <JWT> & X-Request-ID
                                           v
  +---------------------------------------------------------------------------------+
  |                       RequestIDMiddleware (Gin Middleware)                      |
  |         Extracts/Generates UUID -> Injects into c.Request.Context()             |
  +---------------------------------------------------------------------------------+
                                           |
                                           v
  +---------------------------------------------------------------------------------+
  |                   StructuredLoggerMiddleware (Gin Middleware)                   |
  |           Audits HTTP Method, Path, Status Code, Client IP, Latency             |
  +---------------------------------------------------------------------------------+
                                           |
                                           v
  +---------------------------------------------------------------------------------+
  |                     JWTAuthMiddleware (Authentication Layer)                    |
  |         Parses & Validates HMAC Token -> Sets UserID, Role in Context           |
  +---------------------------------------------------------------------------------+
                                           |
                                           v
  +---------------------------------------------------------------------------------+
  |                  RequireRoleMiddleware (Authorization / RBAC)                   |
  |           Verifies Role Permissions (ADMIN / TEACHER / STUDENT)                 |
  +---------------------------------------------------------------------------------+
                                           |
                                           v
  +---------------------------------------------------------------------------------+
  |                       HTTP Handlers (handler/*_handler.go)                      |
  |               Parses JSON Payloads & Returns JSON Responses                     |
  +---------------------------------------------------------------------------------+
                                           |
                                           v
  +---------------------------------------------------------------------------------+
  |                       Business Use Cases (usecase/*_usecase.go)                 |
  |         Contains Password Hashing, Validation, Business Workflows               |
  +---------------------------------------------------------------------------------+
                                           |
                                           v
  +---------------------------------------------------------------------------------+
  |                     Repositories (repository/*_repository.go)                   |
  |         Thread-Safe In-Memory / PostgreSQL Database Persistence                 |
  +---------------------------------------------------------------------------------+
```

---

# 🌐 API Endpoint Matrix & Permission Matrix

| Method | Endpoint | Description | Allowed Roles | Auth Required |
| :--- | :--- | :--- | :--- | :--- |
| `POST` | `/api/v1/auth/register` | Register a new User Account | Public | ❌ |
| `POST` | `/api/v1/auth/login` | Authenticate User & Issue JWT | Public | ❌ |
| `GET` | `/api/v1/students` | List All Students (with filters) | `ADMIN`, `TEACHER`, `STUDENT` | ✅ Bearer JWT |
| `GET` | `/api/v1/students/:id` | Fetch Student Profile by ID | `ADMIN`, `TEACHER`, `STUDENT` | ✅ Bearer JWT |
| `POST` | `/api/v1/students` | Register New Student Profile | `ADMIN`, `TEACHER` | ✅ Bearer JWT |
| `PUT` | `/api/v1/students/:id` | Update Student Details / GPA / Status | `ADMIN`, `TEACHER` | ✅ Bearer JWT |
| `DELETE` | `/api/v1/students/:id` | Permanently Delete Student Profile | `ADMIN` Only | ✅ Bearer JWT |

---

# 🛠️ Implementation Walkthrough — Day 30 Project

### 1. Configuration Loader (`config/config.go`)

```go
package config

import (
	"os"
	"time"
)

type Config struct {
	Port         string
	Environment  string
	JWTSecret    string
	JWTExpiresIn time.Duration
	LogFilePath  string
}

func LoadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "super-secret-student-mgmt-key-90-days-challenge"
	}
	logPath := os.Getenv("LOG_FILE_PATH")
	if logPath == "" {
		logPath = "student_app.log"
	}

	return Config{
		Port:         port,
		Environment:  env,
		JWTSecret:    secret,
		JWTExpiresIn: 24 * time.Hour,
		LogFilePath:  logPath,
	}
}
```

---

### 2. Domain Entities & Interfaces (`domain/student.go`)

```go
package domain

import (
	"context"
	"errors"
	"time"
)

type ContextKey string

const (
	RequestIDKey ContextKey = "X-Request-ID"
	UserIDKey    ContextKey = "user_id"
	UserRoleKey  ContextKey = "user_role"
	UserEmailKey ContextKey = "user_email"
)

const (
	RoleAdmin   = "ADMIN"
	RoleTeacher = "TEACHER"
	RoleStudent = "STUDENT"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrStudentNotFound    = errors.New("student record not found")
	ErrEmailExists        = errors.New("email address is already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUnauthorized       = errors.New("authentication token is missing or invalid")
	ErrForbidden          = errors.New("access forbidden: insufficient permissions")
	ErrInvalidInput       = errors.New("invalid request payload")
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Student struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	GPA        float64   `json:"gpa"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	User      *User     `json:"user"`
}

type CreateStudentInput struct {
	FullName   string  `json:"full_name" binding:"required"`
	Email      string  `json:"email" binding:"required,email"`
	Department string  `json:"department" binding:"required"`
	GPA        float64 `json:"gpa" binding:"min=0,max=4"`
}

type UpdateStudentInput struct {
	FullName   string  `json:"full_name"`
	Department string  `json:"department"`
	GPA        float64 `json:"gpa" binding:"min=0,max=4"`
	Status     string  `json:"status"`
}

type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}

type UserRepository interface {
	SaveUser(ctx context.Context, user *User) error
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	FindUserByID(ctx context.Context, id string) (*User, error)
}

type StudentRepository interface {
	Save(ctx context.Context, student *Student) error
	FindByID(ctx context.Context, id string) (*Student, error)
	FindByEmail(ctx context.Context, email string) (*Student, error)
	FindAll(ctx context.Context, department string, status string) ([]*Student, error)
	Update(ctx context.Context, student *Student) error
	Delete(ctx context.Context, id string) error
}
```

---

### 3. JWT Authentication & Role Middlewares (`middleware/auth_middleware.go`)

```go
package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"day-30/domain"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(jwtSecret string, logger domain.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warn(c.Request.Context(), "Authentication failed: missing Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": domain.ErrUnauthorized.Error()})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			logger.Warn(c.Request.Context(), "Authentication failed: malformed Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be Bearer token"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			logger.Warn(c.Request.Context(), "Authentication failed: invalid or expired token", "error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired authentication token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logger.Warn(c.Request.Context(), "Authentication failed: invalid token claims payload")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID, _ := claims["sub"].(string)
		role, _ := claims["role"].(string)
		email, _ := claims["email"].(string)

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, domain.UserIDKey, userID)
		ctx = context.WithValue(ctx, domain.UserRoleKey, role)
		ctx = context.WithValue(ctx, domain.UserEmailKey, email)

		c.Request = c.Request.WithContext(ctx)

		c.Set(string(domain.UserIDKey), userID)
		c.Set(string(domain.UserRoleKey), role)

		c.Next()
	}
}

func RequireRoleMiddleware(logger domain.Logger, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get(string(domain.UserRoleKey))
		if !exists {
			logger.Warn(c.Request.Context(), "Authorization rejected: Role missing from context")
			c.JSON(http.StatusForbidden, gin.H{"error": domain.ErrForbidden.Error()})
			c.Abort()
			return
		}

		userRole, _ := roleVal.(string)
		roleAllowed := false
		for _, allowed := range allowedRoles {
			if strings.EqualFold(userRole, allowed) {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			logger.Warn(c.Request.Context(), "Authorization rejected: Access denied for role", "user_role", userRole, "required_roles", allowedRoles)
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("Role '%s' is not authorized to perform this operation", userRole)})
			c.Abort()
			return
		}

		c.Next()
	}
}
```

---

### 4. Dependency Injection Container (`di/container.go`)

```go
package di

import (
	"day-30/config"
	"day-30/domain"
	"day-30/handler"
	"day-30/logger"
	"day-30/middleware"
	"day-30/repository"
	"day-30/usecase"

	"github.com/gin-gonic/gin"
)

type Container struct {
	Config         config.Config
	Logger         domain.Logger
	UserRepo       domain.UserRepository
	StudentRepo    domain.StudentRepository
	AuthUseCase    *usecase.AuthUseCase
	StudentUseCase *usecase.StudentUseCase
	AuthHandler    *handler.AuthHandler
	StudentHandler *handler.StudentHandler
}

func NewContainer(cfg config.Config) (*Container, error) {
	appLogger, err := logger.NewZapLogger(cfg.Environment, cfg.LogFilePath)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewMemoryUserRepository(appLogger)
	studentRepo := repository.NewMemoryStudentRepository(appLogger)

	authUseCase := usecase.NewAuthUseCase(userRepo, cfg, appLogger)
	studentUseCase := usecase.NewStudentUseCase(studentRepo, appLogger)

	authHandler := handler.NewAuthHandler(authUseCase, appLogger)
	studentHandler := handler.NewStudentHandler(studentUseCase, appLogger)

	return &Container{
		Config:         cfg,
		Logger:         appLogger,
		UserRepo:       userRepo,
		StudentRepo:    studentRepo,
		AuthUseCase:    authUseCase,
		StudentUseCase: studentUseCase,
		AuthHandler:    authHandler,
		StudentHandler: studentHandler,
	}, nil
}

func (c *Container) SetupRouter() *gin.Engine {
	if c.Config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(c.Logger))

	v1 := router.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", c.AuthHandler.Register)
			authGroup.POST("/login", c.AuthHandler.Login)
		}

		studentGroup := v1.Group("/students")
		studentGroup.Use(middleware.JWTAuthMiddleware(c.Config.JWTSecret, c.Logger))
		{
			studentGroup.GET("", middleware.RequireRoleMiddleware(c.Logger, domain.RoleAdmin, domain.RoleTeacher, domain.RoleStudent), c.StudentHandler.ListStudents)
			studentGroup.GET("/:id", middleware.RequireRoleMiddleware(c.Logger, domain.RoleAdmin, domain.RoleTeacher, domain.RoleStudent), c.StudentHandler.GetStudentByID)
			studentGroup.POST("", middleware.RequireRoleMiddleware(c.Logger, domain.RoleAdmin, domain.RoleTeacher), c.StudentHandler.CreateStudent)
			studentGroup.PUT("/:id", middleware.RequireRoleMiddleware(c.Logger, domain.RoleAdmin, domain.RoleTeacher), c.StudentHandler.UpdateStudent)
			studentGroup.DELETE("/:id", middleware.RequireRoleMiddleware(c.Logger, domain.RoleAdmin), c.StudentHandler.DeleteStudent)
		}
	}

	return router
}
```

---

### 5. Automated Integration Test Suite (`tests/student_api_test.go`)

```go
package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"day-30/config"
	"day-30/di"
	"day-30/domain"
)

func setupTestApp() (*di.Container, http.Handler) {
	cfg := config.LoadConfig()
	cfg.Environment = "test"

	container, _ := di.NewContainer(cfg)
	router := container.SetupRouter()
	return container, router
}

func getAuthToken(t *testing.T, handler http.Handler, email, password, role string) string {
	regPayload := domain.RegisterInput{Email: email, Password: password, Role: role}
	regBody, _ := json.Marshal(regPayload)
	req1, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(regBody))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req1)

	loginPayload := domain.LoginInput{Email: email, Password: password}
	loginBody, _ := json.Marshal(loginPayload)
	req2, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(loginBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req2)

	var res map[string]interface{}
	json.Unmarshal(w2.Body.Bytes(), &res)
	dataMap := res["data"].(map[string]interface{})
	return dataMap["token"].(string)
}

func TestStudentAPI_FullLifecycle(t *testing.T) {
	_, handler := setupTestApp()

	adminToken := getAuthToken(t, handler, "test.admin@school.com", "AdminPassword123", "ADMIN")
	studentToken := getAuthToken(t, handler, "test.student@school.com", "StudentPassword123", "STUDENT")

	createPayload := domain.CreateStudentInput{
		FullName:   "Alice Johnson",
		Email:      "alice.johnson@school.com",
		Department: "PHYSICS",
		GPA:        3.88,
	}
	createBody, _ := json.Marshal(createPayload)

	reqCreate, _ := http.NewRequest("POST", "/api/v1/students", bytes.NewBuffer(createBody))
	reqCreate.Header.Set("Content-Type", "application/json")
	reqCreate.Header.Set("Authorization", "Bearer "+adminToken)
	wCreate := httptest.NewRecorder()
	handler.ServeHTTP(wCreate, reqCreate)

	if wCreate.Code != http.StatusCreated {
		t.Fatalf("expected status 201 Created, got %d. Body: %s", wCreate.Code, wCreate.Body.String())
	}

	var createRes map[string]interface{}
	json.Unmarshal(wCreate.Body.Bytes(), &createRes)
	studentData := createRes["data"].(map[string]interface{})
	studentID := studentData["id"].(string)

	reqGet, _ := http.NewRequest("GET", "/api/v1/students/"+studentID, nil)
	reqGet.Header.Set("Authorization", "Bearer "+studentToken)
	wGet := httptest.NewRecorder()
	handler.ServeHTTP(wGet, reqGet)

	if wGet.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", wGet.Code)
	}

	reqForbiddenDelete, _ := http.NewRequest("DELETE", "/api/v1/students/"+studentID, nil)
	reqForbiddenDelete.Header.Set("Authorization", "Bearer "+studentToken)
	wForbiddenDelete := httptest.NewRecorder()
	handler.ServeHTTP(wForbiddenDelete, reqForbiddenDelete)

	if wForbiddenDelete.Code != http.StatusForbidden {
		t.Errorf("expected status 403 Forbidden for STUDENT role deletion, got %d", wForbiddenDelete.Code)
	}

	reqDelete, _ := http.NewRequest("DELETE", "/api/v1/students/"+studentID, nil)
	reqDelete.Header.Set("Authorization", "Bearer "+adminToken)
	wDelete := httptest.NewRecorder()
	handler.ServeHTTP(wDelete, reqDelete)

	if wDelete.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK for ADMIN role deletion, got %d", wDelete.Code)
	}
}
```

---

# 📘 Day 30 Interview Questions & Answers

---

## ❓ Q1: How does Clean Architecture decouple business logic from framework and database dependencies in a production REST API?

### ✅ Answer:
Clean Architecture enforces strict **Dependency Inversion**:
1. **Domain Layer**: Defines core business entities and abstract interfaces (`UserRepository`, `StudentRepository`, `Logger`) with zero external imports.
2. **Use Case Layer**: Contains business rules (e.g. validating input, hashing passwords with Bcrypt, issuing JWTs) relying solely on domain interfaces.
3. **Handler / Infrastructure Layer**: Implements web routing (Gin framework), persistence (in-memory map / GORM PostgreSQL), and logging (Uber Zap).
Because high-level business use cases depend on interfaces rather than concrete structs, database drivers or web frameworks can be replaced without modifying core business logic.

---

## ❓ Q2: How does JWT Authentication and Role-Based Authorization (RBAC) work in Gin middleware?

### ✅ Answer:
1. **Authentication Middleware (`JWTAuthMiddleware`)**: Intercepts `Authorization: Bearer <token>`, validates HMAC-SHA256 signature using secret key, extracts claims (`sub`, `role`, `email`), and attaches them to `c.Request.Context()`.
2. **Authorization Middleware (`RequireRoleMiddleware`)**: Reads the user's role from request context and compares it against allowed roles (e.g. `ADMIN`, `TEACHER`, `STUDENT`). If unauthorized, it immediately halts execution (`c.Abort()`) and returns `403 Forbidden`.

---

## ❓ Q3: Why is Request Correlation Tracing (`X-Request-ID`) vital in production REST APIs and Microservices?

### ✅ Answer:
In a multi-threaded web server handling thousands of concurrent requests, log lines from different requests interleave in stdout. A **Correlation ID** (`X-Request-ID`) generated at the ingress middleware assigns a unique UUID to each request. Propagating this ID through `context.Context` enables loggers to tag every log record across middleware, use case, and repository layers, allowing engineers to trace a single request's full lifecycle in log aggregators (ELK, Grafana Loki).

---

## ❓ Q4: What are the security benefits of using Bcrypt for password storage in Go REST APIs?

### ✅ Answer:
Bcrypt is a salted, adaptive cryptographic key derivation function. It includes:
1. **Automatic Salting**: Unique random salts prevent rainbow table lookup attacks.
2. **Configurable Cost Factor**: Increasing the cost factor increases computation time per password hash, making brute-force dictionary attacks computationally infeasible while keeping API authentication fast for legitimate users.

---

## ❓ Q5: How does a Dependency Injection Container (`di/container.go`) simplify application setup and testing?

### ✅ Answer:
A Central DI Container instantiates components in topological order (Config -> Logger -> Repositories -> Use Cases -> Handlers -> Router). It centralizes route definitions and dependency wiring, eliminating initialization boilerplate in `main.go`. In unit or integration test environments, mock dependencies can be injected into the container without changing production application code.

---

# 📚 Day 30 Summary

Today I completed the **Day 30 Student Management REST API Milestone Project**:
- Consolidated Clean Architecture, Dependency Injection, Repository Pattern, JWT Authentication, and RBAC into a unified backend service.
- Configured Uber Zap structured logging with contextual `X-Request-ID` correlation tracing.
- Implemented full REST API CRUD endpoints for Student resources with status and department filtering.
- Implemented password hashing with Bcrypt and JWT token authentication.
- Wrote end-to-end integration tests (`httptest`) asserting authentication, student lifecycle, and role permission enforcement.

---

# ⭐ Challenge Progress
✅ Day 01 Completed  
✅ Day 02 Completed  
✅ Day 03 Completed  
✅ Day 04 Completed  
✅ Day 05 Completed  
✅ Day 06 Completed  
✅ Day 07 Completed  
✅ Day 08 Completed  
✅ Day 09 Completed   
✅ Day 10 Completed  
✅ Day 11 Completed  
✅ Day 12 Completed  
✅ Day 13 Completed  
✅ Day 14 Completed   
✅ Day 15 Completed  
✅ Day 16 Completed  
✅ Day 17 Completed  
✅ Day 18 Completed  
✅ Day 19 Completed  
✅ Day 20 Completed  
✅ Day 21 Completed  
✅ Day 22 Completed  
✅ Day 23 Completed  
✅ Day 24 Completed  
✅ Day 25 Completed  
✅ Day 26 Completed  
✅ Day 27 Completed  
✅ Day 28 Completed  
✅ Day 29 Completed  
✅ Day 30 Completed  
✅ Day 31 Completed  
🚀 Next: Pagination & Filtering in Go

---

# ✅ Day 31 — Advanced CRUD APIs in Go

---

# 📖 Introduction to Advanced CRUD Patterns in Go

In production backend engineering, basic CRUD operations (`CREATE`, `READ`, `UPDATE`, `DELETE`) are insufficient for high-volume enterprise systems. **Advanced CRUD APIs** handle complex scenarios such as partial resource updates, bulk processing, record soft deletion, and concurrency control.

### Key Concepts & Design Patterns

1. **Partial Updates (`PATCH` vs `PUT`)**:
   - `PUT`: Requires replacing the entire resource payload. If a field is omitted, it will be wiped or set to zero values.
   - `PATCH`: Modifies only the explicitly supplied fields (represented using pointer types in Go structs, e.g. `*string`, `*float64`), leaving all other resource attributes untouched.

2. **Bulk / Batch Operations (`POST /bulk`, `POST /bulk-delete`)**:
   - Instead of triggering hundreds of separate HTTP requests to insert or delete records individually, bulk endpoints process slices of items (`[]Item`) in a single database transaction/batch call.

3. **Soft Deletion & Restoration**:
   - Deleting database rows permanently (`HARD DELETE`) can cause accidental data loss and break foreign key integrity.
   - `SOFT DELETE` sets an `is_deleted = true` flag and populates a `deleted_at` timestamp. Standard queries filter out deleted rows, while a `POST /:id/restore` endpoint allows restoring soft-deleted records.

4. **Optimistic Concurrency Control (OCC)**:
   - When multiple users attempt to update the same record simultaneously, a **Lost Update** anomaly can occur.
   - OCC assigns a `Version` (or `ETag`) counter to each record. Updating a record requires passing the expected version (via `If-Match` header or JSON body). If `CurrentVersion != ExpectedVersion`, the API rejects the write with a `409 Conflict` status code.

---

# 🏗️ Architecture & Component Flow

```text
  +---------------------------------------------------------------------------------+
  |                       Client / API Consumer (Postman / HTTP)                    |
  +---------------------------------------------------------------------------------+
                                           |
                   Header: X-Request-ID & Optional If-Match: <version>
                                           v
  +---------------------------------------------------------------------------------+
  |                 RequestIDMiddleware & StructuredLoggerMiddleware                |
  |             Attaches X-Request-ID -> Context & Audits HTTP Execution            |
  +---------------------------------------------------------------------------------+
                                           |
                                           v
  +---------------------------------------------------------------------------------+
  |                       HTTP Handlers (handler/product_handler.go)                |
  |     Binds JSON (Create / Bulk / Patch) -> Extracts If-Match Concurrency Header    |
  +---------------------------------------------------------------------------------+
                                           |
                                           v
  +---------------------------------------------------------------------------------+
  |                     Business UseCase Layer (usecase/product_usecase.go)         |
  |     Enforces Non-Nil Field Merging, SKU Uniqueness, Soft Delete & Restoration   |
  +---------------------------------------------------------------------------------+
                                           |
                                           v
  +---------------------------------------------------------------------------------+
  |                    Repository Layer (repository/product_repository.go)           |
  |   Verifies Version Match (Returns 409 on Mismatch), Mutates State, Increments Ver|
  +---------------------------------------------------------------------------------+
```

---

# ⚡ Endpoint Matrix

| Method | Endpoint | Description | Status Code | Notes |
| :--- | :--- | :--- | :--- | :--- |
| `POST` | `/api/v1/products` | Create Single Product | `201 Created` | Sets initial `Version = 1` |
| `POST` | `/api/v1/products/bulk` | Bulk Insert Products | `201 Created` | Atomically inserts multiple products |
| `GET` | `/api/v1/products` | List Products | `200 OK` | Excludes soft deleted (`?include_deleted=true`) |
| `GET` | `/api/v1/products/:id` | Fetch Product by ID | `200 OK` / `404 Not Found` | Returns ETag header with `Version` |
| `PATCH` | `/api/v1/products/:id` | Partial Update Product | `200 OK` / `409 Conflict` | Validates `If-Match` or `expected_version` |
| `POST` | `/api/v1/products/:id/restore` | Restore Soft-Deleted Product | `200 OK` | Resets `is_deleted = false` |
| `DELETE` | `/api/v1/products/:id` | Soft Delete Product | `200 OK` | Sets `is_deleted = true`, populates `deleted_at` |
| `POST` | `/api/v1/products/bulk-delete` | Bulk Soft Delete | `200 OK` | Marks multiple IDs as soft-deleted |

---

# 🛠️ Implementation Walkthrough — Day 31 Project

### 1. Domain Entities & Concurrency Errors (`domain/product.go`)

```go
package domain

import (
	"context"
	"errors"
	"time"
)

type ContextKey string

const (
	RequestIDKey ContextKey = "X-Request-ID"
)

var (
	ErrProductNotFound     = errors.New("product record not found")
	ErrSKUExists           = errors.New("product SKU already exists")
	ErrAlreadyDeleted      = errors.New("product is already soft deleted")
	ErrNotDeleted          = errors.New("product is not deleted")
	ErrConcurrencyConflict = errors.New("concurrency conflict: record was modified by another request")
	ErrInvalidInput        = errors.New("invalid request input parameters")
	ErrEmptyBulkRequest    = errors.New("bulk payload cannot be empty")
)

type Product struct {
	ID        string     `json:"id"`
	SKU       string     `json:"sku"`
	Name      string     `json:"name"`
	Category  string     `json:"category"`
	Price     float64    `json:"price"`
	Stock     int        `json:"stock"`
	Version   int        `json:"version"` // Optimistic Concurrency Control counter
	IsDeleted bool       `json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CreateProductInput struct {
	SKU      string  `json:"sku" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Stock    int     `json:"stock" binding:"gte=0"`
}

type BulkCreateInput struct {
	Items []CreateProductInput `json:"items" binding:"required,dive"`
}

type PatchProductInput struct {
	Name            *string  `json:"name,omitempty"`
	Category        *string  `json:"category,omitempty"`
	Price           *float64 `json:"price,omitempty"`
	Stock           *int     `json:"stock,omitempty"`
	ExpectedVersion *int     `json:"expected_version,omitempty"`
}

type BulkDeleteInput struct {
	IDs []string `json:"ids" binding:"required"`
}

type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}

type ProductRepository interface {
	Save(ctx context.Context, product *Product) error
	BulkSave(ctx context.Context, products []*Product) ([]*Product, error)
	FindByID(ctx context.Context, id string, includeDeleted bool) (*Product, error)
	FindBySKU(ctx context.Context, sku string) (*Product, error)
	FindAll(ctx context.Context, includeDeleted bool) ([]*Product, error)
	PatchUpdate(ctx context.Context, id string, input PatchProductInput) (*Product, error)
	SoftDelete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) (*Product, error)
	BulkSoftDelete(ctx context.Context, ids []string) (int, error)
}
```

---

### 2. Repository with Optimistic Concurrency & Soft Delete (`repository/product_repository.go`)

```go
package repository

import (
	"context"
	"strings"
	"sync"
	"time"

	"day-31/domain"
)

type memoryProductRepository struct {
	mu       sync.RWMutex
	products map[string]*domain.Product
	skus     map[string]string
	logger   domain.Logger
}

func NewMemoryProductRepository(logger domain.Logger) domain.ProductRepository {
	return &memoryProductRepository{
		products: make(map[string]*domain.Product),
		skus:     make(map[string]string),
		logger:   logger,
	}
}

func (r *memoryProductRepository) PatchUpdate(ctx context.Context, id string, input domain.PatchProductInput) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	product, exists := r.products[id]
	if !exists || product.IsDeleted {
		return nil, domain.ErrProductNotFound
	}

	if input.ExpectedVersion != nil {
		if product.Version != *input.ExpectedVersion {
			r.logger.Warn(ctx, "DB Concurrency Conflict detected", "product_id", id, "current_version", product.Version, "expected_version", *input.ExpectedVersion)
			return nil, domain.ErrConcurrencyConflict
		}
	}

	if input.Name != nil && strings.TrimSpace(*input.Name) != "" {
		product.Name = strings.TrimSpace(*input.Name)
	}
	if input.Category != nil && strings.TrimSpace(*input.Category) != "" {
		product.Category = strings.ToUpper(strings.TrimSpace(*input.Category))
	}
	if input.Price != nil && *input.Price > 0 {
		product.Price = *input.Price
	}
	if input.Stock != nil && *input.Stock >= 0 {
		product.Stock = *input.Stock
	}

	product.Version++
	product.UpdatedAt = time.Now()

	r.products[id] = product
	r.logger.Info(ctx, "DB Product patched & version incremented", "product_id", id, "new_version", product.Version)
	return product, nil
}

func (r *memoryProductRepository) SoftDelete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	product, exists := r.products[id]
	if !exists {
		return domain.ErrProductNotFound
	}
	if product.IsDeleted {
		return domain.ErrAlreadyDeleted
	}

	now := time.Now()
	product.IsDeleted = true
	product.DeletedAt = &now
	product.UpdatedAt = now

	r.products[id] = product
	return nil
}

func (r *memoryProductRepository) Restore(ctx context.Context, id string) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	product, exists := r.products[id]
	if !exists {
		return nil, domain.ErrProductNotFound
	}
	if !product.IsDeleted {
		return nil, domain.ErrNotDeleted
	}

	product.IsDeleted = false
	product.DeletedAt = nil
	product.UpdatedAt = time.Now()

	r.products[id] = product
	return product, nil
}
```

---

### 3. Business Use Case Layer (`usecase/product_usecase.go`)

```go
package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"day-31/domain"

	"github.com/google/uuid"
)

type ProductUseCase struct {
	repo   domain.ProductRepository
	logger domain.Logger
}

func NewProductUseCase(repo domain.ProductRepository, logger domain.Logger) *ProductUseCase {
	return &ProductUseCase{repo: repo, logger: logger}
}

func (u *ProductUseCase) CreateProduct(ctx context.Context, input domain.CreateProductInput) (*domain.Product, error) {
	sku := strings.ToUpper(strings.TrimSpace(input.SKU))
	existing, err := u.repo.FindBySKU(ctx, sku)
	if err == nil && existing != nil {
		return nil, domain.ErrSKUExists
	}

	now := time.Now()
	product := &domain.Product{
		ID:        fmt.Sprintf("prd_%s", uuid.New().String()[:8]),
		SKU:       sku,
		Name:      strings.TrimSpace(input.Name),
		Category:  strings.ToUpper(strings.TrimSpace(input.Category)),
		Price:     input.Price,
		Stock:     input.Stock,
		Version:   1,
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.repo.Save(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}

func (u *ProductUseCase) PatchUpdateProduct(ctx context.Context, id string, input domain.PatchProductInput) (*domain.Product, error) {
	return u.repo.PatchUpdate(ctx, id, input)
}
```

---

### 4. Gin HTTP Handler (`handler/product_handler.go`)

```go
package handler

import (
	"net/http"
	"strconv"

	"day-31/domain"
	"day-31/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase *usecase.ProductUseCase
	logger         domain.Logger
}

func NewProductHandler(productUseCase *usecase.ProductUseCase, logger domain.Logger) *ProductHandler {
	return &ProductHandler{productUseCase: productUseCase, logger: logger}
}

func (h *ProductHandler) PatchUpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var input domain.PatchProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	if ifMatch := c.GetHeader("If-Match"); ifMatch != "" {
		if ver, err := strconv.Atoi(ifMatch); err == nil {
			input.ExpectedVersion = &ver
		}
	}

	product, err := h.productUseCase.PatchUpdateProduct(c.Request.Context(), id, input)
	if err != nil {
		switch err {
		case domain.ErrConcurrencyConflict:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case domain.ErrProductNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.Header("ETag", strconv.Itoa(product.Version))
	c.JSON(http.StatusOK, gin.H{"message": "Product patched successfully", "data": product})
}
```

---

### 5. Integration Test Suite (`handler/product_handler_test.go`)

```go
package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"day-31/domain"
	"day-31/handler"
	"day-31/logger"
	"day-31/middleware"
	"day-31/repository"
	"day-31/usecase"

	"github.com/gin-gonic/gin"
)

func TestAdvancedCRUD_Lifecycle(t *testing.T) {
	gin.SetMode(gin.TestMode)
	zapLog, _ := logger.NewZapLogger("test", "")
	productRepo := repository.NewMemoryProductRepository(zapLog)
	productUC := usecase.NewProductUseCase(productRepo, zapLog)
	productHandler := handler.NewProductHandler(productUC, zapLog)

	router := gin.New()
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	v1 := router.Group("/api/v1/products")
	{
		v1.POST("", productHandler.CreateProduct)
		v1.PATCH("/:id", productHandler.PatchUpdateProduct)
		v1.DELETE("/:id", productHandler.SoftDeleteProduct)
		v1.POST("/:id/restore", productHandler.RestoreProduct)
	}

	createInput := domain.CreateProductInput{
		SKU:      "TEST-SKU-01",
		Name:     "Test Keyboard",
		Category: "PERIPHERALS",
		Price:    89.99,
		Stock:    20,
	}
	body, _ := json.Marshal(createInput)
	req1, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	var res1 map[string]interface{}
	json.Unmarshal(w1.Body.Bytes(), &res1)
	productData := res1["data"].(map[string]interface{})
	productID := productData["id"].(string)

	// Test Optimistic Concurrency Conflict (Stale Version)
	staleVersion := 0
	newPrice := 79.99
	patchInputStale := domain.PatchProductInput{Price: &newPrice, ExpectedVersion: &staleVersion}
	staleBody, _ := json.Marshal(patchInputStale)
	req3, _ := http.NewRequest("PATCH", "/api/v1/products/"+productID, bytes.NewBuffer(staleBody))
	req3.Header.Set("Content-Type", "application/json")
	w3 := httptest.NewRecorder()
	router.ServeHTTP(w3, req3)

	if w3.Code != http.StatusConflict {
		t.Errorf("expected status 403 Conflict on stale version, got %d", w3.Code)
	}
}
```

---

# 📘 Day 31 Interview Questions & Answers

---

## ❓ Q1: What is the technical difference between `PUT` and `PATCH` in HTTP REST APIs, and how is `PATCH` implemented in Go?

### ✅ Answer:
- **`PUT`**: Is an idempotent replacement operation. The request payload must represent the full state of the resource. Any omitted fields are reset to their default zero values (`""`, `0`, `nil`).
- **`PATCH`**: Applies a partial update. Only the fields present in the request body are modified.
In Go, `PATCH` is implemented by defining DTO fields as **pointers** (e.g. `*string`, `*float64`). If a field is omitted in the JSON request, its pointer value is `nil` and the repository leaves that struct field untouched. If a field is provided, its pointer is non-nil and the value is updated.

---

## ❓ Q2: What is Optimistic Concurrency Control (OCC) and how does it prevent "Lost Updates" in REST APIs?

### ✅ Answer:
Optimistic Concurrency Control assumes database write conflicts are rare and avoids expensive pessimistic locks (e.g. `SELECT FOR UPDATE`).
Each record maintains a `Version` counter (or `ETag`). When a client reads a resource, it receives `Version = N`. When submitting a `PATCH` or `PUT` request, the client supplies `ExpectedVersion = N` (or `If-Match: N`). The database query updates the record only `WHERE id = ID AND version = N`, incrementing `version = N + 1`. If another request updated the record first (`version = N + 1`), zero rows are updated, and the API returns `409 Conflict`.

---

## ❓ Q3: What is Soft Deletion and what are its performance and architectural trade-offs?

### ✅ Answer:
**Soft Deletion** retains the database row while marking `is_deleted = true` and setting `deleted_at = NOW()`.
- **Pros**: Prevents accidental data loss, maintains audit history, preserves foreign key integrity, and allows simple restoration (`POST /restore`).
- **Cons**: Database tables grow continuously over time, requiring all queries to include `WHERE is_deleted = false`. Unique constraints (e.g. unique `sku`) must be composite indexes (`sku, deleted_at`) to allow re-registering an SKU after soft deletion.

---

## ❓ Q4: How do Batch/Bulk API endpoints improve performance over individual CRUD requests?

### ✅ Answer:
Individual REST requests incur significant overhead: HTTP network round-trips, TLS handshakes, middleware execution, and separate database transactions (`BEGIN` ... `COMMIT`).
Bulk API endpoints (e.g. `POST /api/v1/products/bulk`) accept an array of resources and execute multi-row batch inserts (`INSERT INTO products VALUES (...), (...)`) within a single database transaction. This reduces network latency and database log flush operations by orders of magnitude.

---

## ❓ Q5: How do you handle conditional requests using HTTP `ETag` and `If-Match` headers in Go REST handlers?

### ✅ Answer:
In the `GET` handler, the server returns the record's current version formatted as an `ETag` header (`ETag: "2"`).
During subsequent `PATCH` or `PUT` requests, the client sends `If-Match: "2"`. The Gin handler extracts `c.GetHeader("If-Match")`, converts it to an integer `expected_version`, and passes it to the use case. If the database record version does not match, the handler responds with `409 Conflict` or `412 Precondition Failed`.

---

# 📚 Day 31 Summary

Today I completed **Day 31 Advanced CRUD APIs in Go**:
- Implemented **Partial Updates (`PATCH`)** using pointer DTO fields to update only supplied JSON properties.
- Built **Bulk Operations (`POST /bulk`, `POST /bulk-delete`)** for batch creation and deletion.
- Implemented **Soft Deletion & Restoration** (`is_deleted`, `deleted_at`, `POST /:id/restore`).
- Implemented **Optimistic Concurrency Control (OCC)** using version counters and `If-Match` header validation returning `409 Conflict`.
- Integrated Uber Zap structured logging and request ID correlation tracing.
- Wrote integration unit tests verifying PATCH updates, OCC conflicts, and soft delete/restore lifecycles.

---

# ⭐ Challenge Progress
✅ Day 01 Completed  
✅ Day 02 Completed  
✅ Day 03 Completed  
✅ Day 04 Completed  
✅ Day 05 Completed  
✅ Day 06 Completed  
✅ Day 07 Completed  
✅ Day 08 Completed  
✅ Day 09 Completed   
✅ Day 10 Completed  
✅ Day 11 Completed  
✅ Day 12 Completed  
✅ Day 13 Completed  
✅ Day 14 Completed   
✅ Day 15 Completed  
✅ Day 16 Completed  
✅ Day 17 Completed  
✅ Day 18 Completed  
✅ Day 19 Completed  
✅ Day 20 Completed  
✅ Day 21 Completed  
✅ Day 22 Completed  
✅ Day 23 Completed  
✅ Day 24 Completed  
✅ Day 25 Completed  
✅ Day 26 Completed  
✅ Day 27 Completed  
✅ Day 28 Completed  
✅ Day 29 Completed  
✅ Day 30 Completed  
✅ Day 31 Completed  
🚀 Next: Pagination & Filtering in Go