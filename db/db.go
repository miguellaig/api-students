package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	CPF    int
	Email  string
	Age    int
	Active bool
}

func Init() {
	var err error
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

}

func AddStudent(DB *gorm.DB) {

	student := Student{
		Name:   "Bento",
		CPF:    12345,
		Email:  "bento@gmail.com",
		Age:    4,
		Active: true,
	}
	if result := DB.Create(&student); result.Error != nil {
		fmt.Println("Error to create student")
		return
	}
	fmt.Println("Create student!")
}
