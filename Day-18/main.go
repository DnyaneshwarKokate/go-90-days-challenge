package main

import "github.com/gin-gonic/gin"


func main() {
	router := gin.Default()

	router.GET("/student/:id", func(c *gin.Context)  {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"student_id":id,
		})
	})
	router.Run(":8080")
}