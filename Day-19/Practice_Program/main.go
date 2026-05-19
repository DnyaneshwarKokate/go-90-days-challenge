package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=dnyaneshwarkokate password=postgres dbname=studentdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	_, err = db.Exec(
		"INSERT INTo students(name, age) values($1, $2)",
		 "Ritesh Gadakh", 23,)

	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println("Student Added Successfully")
}