package main
import "fmt"

type Student struct {
	ID int
	Name string
	Age int
	City string
}

var students []Student


//Add Student


func addStudent() {
	var student Student

	fmt.Print("Enter Id:")
	fmt.Scan(&student.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&student.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&student.Age)

	fmt.Print("Enter City: ")
	fmt.Scan(&student.City)
	students = append(students, student)

	fmt.Println("Student Added Successfully.")
}

//view Student

func viewStudents() {
	if len(students) == 0 {
		fmt.Println("No Students Found")
		return
	}

	fmt.Println("\n Student List")
	for _, student := range students {

		fmt.Println("_________")
		fmt.Println("Id:", student.ID)
		fmt.Println("Name:", student.Name)
		fmt.Println("Age:", student.Age)
		fmt.Println("City:", student.City)
	}
}

// Search Student

func searchStudent() {
	var id int

	fmt.Print("Enter Student Id: ")
	fmt.Scan(&id)

	for _, student := range students {
		if student.ID == id {
			fmt.Println("Student Found")
			fmt.Println("Name:", student.Name)
			fmt.Println("Age:",student.Age)
			fmt.Println("City:",student.City)
			return
		}
	}
	fmt.Println("Student Not Found")
}

//Delete Student

func deleteStudent() {
	var id int

	fmt.Print("Enter Student Id to Delete:")
	fmt.Scan(&id)

	for index, student := range students {
		if student.ID == id {
			students = append(students[:index], students[index+1:]...)
			fmt.Println("Student Deleted Successfully")
			return
		}
	}
	fmt.Println("Student Not Found")
}

func updateStudent() {
	var id int
	fmt.Print("Enter Student Id To Update:")
	fmt.Scan(&id)

	for index, student := range students {
		if student.ID == id {
			fmt.Print("Enter New Name: ")
			fmt.Scan(&students[index].Name)

			fmt.Print("Enter New Age: ")
			fmt.Scan(&students[index].Age)

			fmt.Print("Enter New City: ")
			fmt.Scan(&students[index].City)

			fmt.Println("Student Updated Successfully")
			return
		}
	}
	fmt.Println("Student Not Found")
}

func main() {
	for {
		fmt.Println("\n=======Student Management System=====")
		fmt.Println("1. Add Student")
		fmt.Println("2. View Student")
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
			fmt.Println("Exiting program")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}