package models

import (
	"github.com/jinzhu/gorm"
	"regexp"
)

type Account struct {
	gorm.Model
	Username	string  `json:"username"`
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
	Location	string  `json:"location"`
	About		string  `json:"about"`
}

func (account *Account) ValidateNewAccount(email, password string) bool {
	ok, err := regexp.MatchString(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, email)
	if err != nil {
		return false
	}

	if !ok {
		return false
	}

	if len(password) < 6 {
		return false
	}

	return true
}
