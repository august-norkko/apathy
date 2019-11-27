package models

import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
}