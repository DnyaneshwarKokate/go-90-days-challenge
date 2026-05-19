package main

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secretkey := "mysecretkey"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":"Dnyanesh",
		"role":"admin",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretkey))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("JWT Token:")
	fmt.Println(tokenString)
}