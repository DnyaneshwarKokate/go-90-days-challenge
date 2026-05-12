package main

import (
	"fmt"
	"strconv"
)
func main(){
	number, err := strconv.Atoi("100")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(number)
}