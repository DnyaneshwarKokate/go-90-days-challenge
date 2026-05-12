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
| Day 05 | Pointers | ✅ |
| Day 06 | Methods & Interfaces | ✅ |
| Day 07 | Mini Project | ✅ |
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

	Name   string
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
✅ Day 05 Completed  
✅ Day 06 Completed  
✅ Day 07 Completed  
🚀 Next: Packages & Modules in Go