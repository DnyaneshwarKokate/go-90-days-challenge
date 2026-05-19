package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	router := gin.Default()

	router.GET("/product/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"product_id": id,
		})
	})

	router.POST("/product", func(c *gin.Context)  {
		
		var product Product

		err := c.BindJSON(&product)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":"Product Created",
			"data": product,
		})
	})

	router.Run(":8080")
}
