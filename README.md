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

These are heavily used in:
- REST APIs
- Database operations
- JSON handling
- Backend systems

---

# 📖 What You Will Learn

- Arrays
- Slices
- Maps
- Range Keyword
- Slice Operations
- Append Function
- Length & Capacity
- Iteration
- Interview Questions

---

# 📌 Arrays in Go

An array is a fixed-size collection of elements of the same data type.

### Important Points:
✅ Fixed Size  
✅ Same Data Type  
✅ Index Starts from 0

---

# 📌 Real Life Example of Array

Suppose we want to store marks of 5 students.

Without array:

```go
mark1 := 80
mark2 := 75
mark3 := 90
```

This is difficult to manage.

Using array:

```go
marks := [5]int{80,75,90,85,70}
```

This is cleaner and easier.

---

## ✅ Array Syntax

```go
var numbers [5]int
```

### Explanation:
- `var` → variable declaration
- `numbers` → array name
- `[5]` → array size
- `int` → datatype

---

## ✅ Array Example

```go
package main

import "fmt"

func main() {

	var numbers [5]int

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println(numbers)
}
```

---

# 📌 Understanding Index

| Index | Value |
|---|---|
| 0 | 10 |
| 1 | 20 |
| 2 | 30 |
| 3 | 40 |
| 4 | 50 |

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

# 📌 Short Array Declaration

```go
package main

import "fmt"

func main() {

	numbers := [5]int{1,2,3,4,5}

	fmt.Println(numbers)
}
```

---

# 📌 Array Length

```go
package main

import "fmt"

func main() {

	numbers := [5]int{1,2,3,4,5}

	fmt.Println("Length:", len(numbers))
}
```

---

# 📌 Loops with Arrays

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

# 📌 Problem with Arrays

Arrays are fixed-size.

❌ Cannot add dynamic values  
❌ Not flexible

That’s why Go developers mostly use:
# ✅ SLICES

---

# 📌 Slices in Go

Slices are dynamic and flexible versions of arrays.

Unlike arrays:
✅ Dynamic size  
✅ More powerful  
✅ Used most in real projects

---

# 📌 Why Slices are Important?

Slices are used in real projects because data size changes dynamically.

Examples:
- Products in E-Commerce
- Users in Applications
- Orders in Database
- API Responses

Arrays are fixed-size, but slices can grow dynamically.

---

# ✅ Slice Syntax

```go
numbers := []int{1,2,3}
```

---

# ✅ Slice Example

```go
package main

import "fmt"

func main() {

	cities := []string{"Pune","Mumbai","Nashik"}

	fmt.Println(cities)
}
```

---

# 📌 Append in Slice

Used to add elements dynamically.

```go
package main

import "fmt"

func main() {

	numbers := []int{1,2,3}

	numbers = append(numbers, 4)
	numbers = append(numbers, 5)

	fmt.Println(numbers)
}
```

---

# 📌 Slice Length & Capacity

```go
package main

import "fmt"

func main() {

	numbers := []int{1,2,3,4,5}

	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}
```

---

# 📌 Difference Between len() and cap()

| len() | cap() |
|---|---|
| Current elements count | Total allocated memory |

---

# 📌 Slice from Array

```go
package main

import "fmt"

func main() {

	arr := [5]int{10,20,30,40,50}

	slice := arr[1:4]

	fmt.Println(slice)
}
```

Output:

```text
[20 30 40]
```

---

# 📌 Maps in Go

Maps store data in key-value format.

---

# 📌 Real Life Example of Map

Maps are useful for storing related information.

Example:

```text
name -> Dnyaneshwar
city -> Nashik
role -> Backend Developer
```

Maps are commonly used in:
- JSON Data
- API Responses
- Database Records
- User Information

---

# ✅ Map Syntax

```go
student := map[string]string{

	"name":"Dnyaneshwar",
	"city":"Nashik",
}
```

---

# ✅ Map Example

```go
package main

import "fmt"

func main() {

	student := map[string]string{

		"name":"Dnyaneshwar",
		"city":"Nashik",
		"role":"Developer",
	}

	fmt.Println(student)
}
```

---

# 📌 Access Map Values

```go
fmt.Println(student["name"])
```

Output:

```text
Dnyaneshwar
```

---

# 📌 Add Values to Map

```go
student["salary"] = "50000"
```

---

# 📌 Delete from Map

```go
delete(student, "role")
```

---

# 📌 Range Keyword

Used for iteration.

Works with:
- Arrays
- Slices
- Maps

---

# ✅ Range with Array

```go
package main

import "fmt"

func main() {

	numbers := [5]int{10,20,30,40,50}

	for index, value := range numbers {

		fmt.Println(index, value)
	}
}
```

---

# ✅ Range with Slice

```go
package main

import "fmt"

func main() {

	cities := []string{"Pune","Mumbai","Nashik"}

	for index, value := range cities {

		fmt.Println(index, value)
	}
}
```

---

# ✅ Range with Map

```go
package main

import "fmt"

func main() {

	student := map[string]string{

		"name":"Dnyaneshwar",
		"city":"Nashik",
	}

	for key, value := range student {

		fmt.Println(key, value)
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

	fmt.Println("Marks:", marks)

	// Slice
	cities := []string{"Pune","Mumbai"}

	cities = append(cities, "Nashik")

	fmt.Println("Cities:", cities)

	// Map
	student := map[string]string{

		"name":"Dnyaneshwar",
		"city":"Nashik",
		"role":"Developer",
	}

	fmt.Println(student)

	// Range Loop
	for index, value := range marks {

		fmt.Println(index, value)
	}
}
```

---

# 📘 Day 03 Interview Questions & Answers

---

## ❓ Q1: What is an array in Go?

### ✅ Answer:
An array is a fixed-size collection of elements of the same data type.

---

## ❓ Q2: What is a slice in Go?

### ✅ Answer:
A slice is a dynamic and flexible version of an array.

---

## ❓ Q3: Difference between array and slice?

| Array | Slice |
|---|---|
| Fixed size | Dynamic size |
| Less flexible | More flexible |
| Memory fixed | Dynamic memory |

---

## ❓ Q4: What is append() in Go?

### ✅ Answer:
`append()` is used to add values dynamically into slices.

---

## ❓ Q5: What is map in Go?

### ✅ Answer:
Map stores data in key-value format.

---

## ❓ Q6: What is range keyword?

### ✅ Answer:
`range` is used for iteration over arrays, slices, and maps.

---

## ❓ Q7: Difference between len() and cap()?

| len() | cap() |
|---|---|
| Current elements count | Total allocated capacity |

---

# 📚 Day 03 Summary

Today I learned:
- Arrays
- Slices
- Maps
- Range keyword
- append()
- len() & cap()

I also practiced:
- Array iteration
- Slice operations
- Map handling
- Range loops

---

# 🧠 Practice Tasks

✅ Create student marks array  
✅ Create city slice  
✅ Add values using append()  
✅ Create employee map  
✅ Print all values using range  
✅ Delete map values  
✅ Create product price map

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

# 🚀 GitHub Commands

```bash
git add .
git commit -m "Day 3 completed arrays slices maps"
git push
```

---

# ⭐ Challenge Progress

✅ Day 03 Completed  
🚀 Next: Structs & Custom Types