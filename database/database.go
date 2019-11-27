package database

import (
	"log"
	"fmt"
	"os"
	"apathy/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

func Mysql() *gorm.DB {
	return database
}

func Initialize() {
	db, err := gorm.Open("mysql", constructUri())	
	if err != nil {
		log.Println("Connection to MySQL failed")
		log.Fatal(err)
	}
	db.LogMode(true)
	database = db
	migrations()
}

func constructUri() string {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	name	 := os.Getenv("MYSQL_DATABASE")
	host 	 := os.Getenv("MYSQL_HOST")
	uri 	 := "charset=utf8&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?%s", username, password, host, name, uri)
}

func migrations() {
	database.AutoMigrate(&models.User{})
}