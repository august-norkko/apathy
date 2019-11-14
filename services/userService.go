package services

import (
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"apathy/utils"
	"apathy/database"
	"apathy/entity"
	"apathy/security"
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

	signedToken, err := security.GenerateToken(res.Email)
	if err != nil {
		return http.StatusBadRequest, "Unable to generate JWT token", err
	}

	return http.StatusOK, signedToken, nil
}