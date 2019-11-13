package database

import (
	"log"
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
	Email string
	Password string
}

func Initialize() {
	uri := "root:root@/golang?charset=utf8&parseTime=True&loc=Local" // dev
	db, err := gorm.Open("mysql", uri)	
	if err != nil {
		log.Println("Connection to MySQL failed")
		log.Fatal(err)
	} else {
		log.Println("Connected to MySQL")		
	}
	
	defer db.Close()
	database = db
	
	// Create User table.
	db.AutoMigrate(&User{})
}
