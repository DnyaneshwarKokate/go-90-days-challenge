package main

import(
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

	query := `Insert into students(name, age) values($1,$2)`

	_, err = db.Exec(query, "Shreyas Kokate", 18)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Student Inserted Successfully")

}