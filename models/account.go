package models

import (
	"github.com/jinzhu/gorm"
	"regexp"
	"fmt"
)

type Account struct {
	gorm.Model
	Username	string  `json:"username"`
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
	Location	string  `json:"location"`
	About		string  `json:"about"`
}

func (account *Account) ValidateNewAccount(data *Account) bool {
	ok, _ := regexp.MatchString(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, data.Email)
	if !ok {
		return false
	}

	ok, _ = regexp.MatchString(`(^[A-Za-z0-9]+(?:[ _-][A-Za-z0-9]+)*$)`, data.Username)
	if !ok {
		return false
	}
	
	if len(data.Password) < 6 {
		return false
	}

	return true
}
