package services

import (
	"log"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"apathy/utils"
	"apathy/database"
	"apathy/entity"
)

const (
	secret = "secret"
)

type IUserService interface {
	CreateUser(email, password string)
}

type Service struct {}

func UserService() *Service {
	return &Service{}
}

func (s *Service) CreateUser(r *http.Request) (int, string, error) {
	res, err := utils.Decode(r)
	if err != nil {
		return http.StatusBadRequest, "Unable to decode JSON", err
	}

	password := []byte(res.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
		return http.StatusBadRequest, "Unable to generate hash", err
	}

	db := database.Mysql()
	err = db.Create(&entity.User{ Email: res.Email, Password: string(hashedPassword) }).Error
	if err != nil {
		return http.StatusBadRequest, "Unable to create user", err
	}

	log.Println("Saved user ", res.Email)
	return http.StatusOK, "User created successfully", nil
}

func (s *Service) LoginUser(r *http.Request) (int, string, error) {
	res, err := utils.Decode(r)
	if err != nil {
		return http.StatusBadRequest, "Unable to decode JSON", err
	}

	db, user, password := database.Mysql(), &entity.User{}, []byte(res.Password)
	err = db.Table("users").Where("email = ?", res.Email).First(user).Error
	if err != nil {
		return http.StatusBadRequest, "Email not found", err
	}

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
    if err != nil {
		return http.StatusBadRequest, "Email or password incorrect", err
	}

	expiration := time.Now().Add(10 * time.Minute)
	claim := &entity.Claim{
		Email: res.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return http.StatusBadRequest, "Unable to sign token", err
	}

	return http.StatusOK, signedToken, nil
}