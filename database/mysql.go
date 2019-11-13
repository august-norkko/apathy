package database

import (
	"log"
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

// access db from controllers
func Mysql() *gorm.DB {
	return database
}

type User struct {
	gorm.Model
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
}

func Initialize() {
	var uri string
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")
	uri = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, name)

	db, err := gorm.Open("mysql", uri)	
	if err != nil {
		log.Println("Connection to MySQL failed")
		log.Fatal(err)
	} else {
		log.Println("Connected to MySQL")		
	}
	
	database = db
	
	// Create User table.
	db.AutoMigrate(&User{})
}
