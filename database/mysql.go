package database

import (
	"log"
	"fmt"
	"os"
	"apathy/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	host := "localhost"
	uri = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?%s", username, password, host, name, uri)

	db, err := gorm.Open("mysql", uri)	
	if err != nil {
		log.Println("Connection to MySQL failed")
		log.Fatal(err)
	}
	log.Println("Connected to MySQL")
	
	db.AutoMigrate(&entity.User{})
	database = db
}
