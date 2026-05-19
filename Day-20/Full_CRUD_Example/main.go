package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Student struct {
	ID uint `josn:"id" gorm:"primarykey"`
	Name string `josn:"name"`
	Age int `json:"age"`
}

func main() {
	dsn := "user=dnyaneshwarkokate password=postgres dbname=studentdb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Student{})

	router := gin.Default()

	//Get All Student

	router.GET("/students", func(c *gin.Context) {

		var students []Student

		db.Find(&students)

		c.JSON(http.StatusOK, students)
	})

	//Create Student


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

	//Updated Student

	router.PUT("/student/:id", func(c *gin.Context) {

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

	//DELETE STudent

	router.DELETE("/students/:id", func(c *gin.Context)  {


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
			"Message": "Student Deleted Successfully",
		})
	})

	router.Run(":8080")
}