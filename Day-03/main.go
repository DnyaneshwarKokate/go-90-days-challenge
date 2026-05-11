package main
import "fmt"

func main() {
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Array:", numbers)
	fmt.Println("Third Element:", numbers[2])
	fmt.Println("Length of Array:", len(numbers))

	cities := []string{"Pune", "Mumbai", "Delhi", "Bangalore"}
	cities = append(cities, "Chennai")
	fmt.Println("Cities:", cities)

	student := map[string]string{
		"name":  "Dnyaneshwar",
		"age":   "25",
		"city":  "Pune",
		"grade": "A",
	}
	fmt.Println("Student Details:", student)
}