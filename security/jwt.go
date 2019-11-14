package security

import (
	"time"
	"os"
	"apathy/entity"
	"github.com/dgrijalva/jwt-go"
	_ "log"
	"strings"
)

var secret string = os.Getenv("JWT_SECRET")

func GenerateToken(email string) (string, error) {
	expiration := time.Now().Add(10 * time.Minute)
	claim := &entity.Claim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(header string) (*jwt.Token, error) {
	tokenPointer := &entity.Token{}
	tokenPart := strings.Split(header, " ")[1] // don't want Bearer
	token, err := jwt.ParseWithClaims(tokenPart, tokenPointer, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}