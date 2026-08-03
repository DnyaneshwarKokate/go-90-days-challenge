package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context){
		var login Login

		err := c.BindJSON(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if login.Username != "admin" || login.Password != "1234" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"Invalid Credentials",
			})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": login.Username,
			"role": "admin",
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString , err := token.SignedString([]byte("mysecretkey"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	})

	router.GET("/profile", func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token Required",
			})
			return
		}

		tokenString := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("mysecretkey"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or Expired Token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Claims",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Protected Route Accessed",
			"username": claims["username"],
			"role":     claims["role"],
		})
	})

	router.Run(":8080")
}