package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Student struct {
	ID int
	Name string
	Age int

}

func main() {
	connStr :="user=dnyaneshwarkokate password=postgres dbname=studentdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err :=db.Query("Select id, name, age from students")
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