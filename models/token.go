package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Id	uint
	jwt.StandardClaims
}