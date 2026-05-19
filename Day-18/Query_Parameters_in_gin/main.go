package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/search", func(c *gin.Context) {
		name := c.Query("name")

		c.JSON(http.StatusOK, gin.H{
			"name" : name,
		})
	})
	router.Run(":8080")
}