package database

import (
	"log"
	_ "fmt"
	_ "os"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
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
	uri := "root:root@/golang?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", uri)	
	if err != nil {
		log.Println("Connection to MySQL failed")
		log.Fatal(err)
	} else {
		log.Println("Connected to MySQL")		
	}
	
	db.AutoMigrate(&User{})
	database = db
}
