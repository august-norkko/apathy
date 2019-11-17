package database

import (
	"log"
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"apathy/entity"
)

var database *gorm.DB

func Mysql() *gorm.DB {
	return database
}

func Initialize() {
	var uri string = "charset=utf8&parseTime=True&loc=Local"
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")
	uri = fmt.Sprintf("%s:%s@/%s?%s", username, password, name, uri)

	db, err := gorm.Open("mysql", uri)	
	if err != nil {
		log.Println("Connection to MySQL failed")
		log.Fatal(err)
	}
	log.Println("Connected to MySQL")
	
	db.AutoMigrate(&entity.User{})
	database = db
}
