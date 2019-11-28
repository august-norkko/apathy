package security

import (
	"time"
	"os"
	"apathy/models"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

var secret string = os.Getenv("JWT_SECRET")

const (
	EXPIRESIN = 60 * time.Minute
)

func GenerateToken(id uint) (string, error) {
	expiration := time.Now().Add(EXPIRESIN)
	claim := &models.Token{
		Id: id,
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

func ParseToken(header string) (*jwt.Token, *models.Token, error) {
	tokenModel := &models.Token{}
	headerPart := strings.Split(header, " ")[1]

	token, err := jwt.ParseWithClaims(headerPart, tokenModel, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	return token, tokenModel, nil
}
