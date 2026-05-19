package main

import (
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Student struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`

}

func main() {
	connStr := "user=dnyaneshwarkokate password=postgres dbname=studentdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/students", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name, age from students")

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