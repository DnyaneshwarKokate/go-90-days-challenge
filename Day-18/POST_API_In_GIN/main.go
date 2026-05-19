package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	router := gin.Default()
	router.POST("/student", func(c *gin.Context) {
		var student Student

		err := c.BindJSON(&student)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error(),
			})

			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "Student Created",
			"data":student,
		})
	})
	router.Run(":8080")
}