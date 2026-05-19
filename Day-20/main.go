package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	ID uint `json:"id" gorm:"primarykey"`
	Name string `json:"name:"`
	Age int `json:"age"`
}

func main() {
	dsn := "user=dnyaneshwarkokate dbname=studentdb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Student{})
	fmt.Println("Migration Completed")
	
}