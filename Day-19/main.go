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
	fmt.Println("POSTGRESQL Connected Successfully")
}