package entity

import (
	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Email	string	`json:"email"`
	jwt.StandardClaims
}