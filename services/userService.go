package services

import (
	"log"
	"regexp"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"apathy/utils"
	"apathy/database"
	"apathy/entity"
	"apathy/security"
)

type IUserService interface {
	CreateUser(r *http.Request) (int, string, error)
	LoginUser(r *http.Request) (string, error)
}

type Service struct {}

func UserService() *Service {
	return &Service{}
}

func validateUser(email, password string) string {
	match, _ := regexp.MatchString(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, email)
	if match == false {
		return "Email invalid (example@email.com)"
	}
	if len(password) < 5 {
		return "Password too short (min. 6 char)"
	}
	return ""
}

func (s *Service) CreateUser(r *http.Request) (int, string, error) {
	res, err := utils.Decode(r)
	if err != nil {
		return http.StatusBadRequest, "Unable to decode JSON payload", err
	}

	msg := validateUser(res.Email, res.Password)
	if len(msg) != 0 {
		return http.StatusBadRequest, msg, err
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

func (s *Service) LoginUser(r *http.Request) (string, error) {
	res, err := utils.Decode(r)
	if err != nil {
		return "", err
	}

	msg := validateUser(res.Email, res.Password)
	if len(msg) != 0 {
		return msg, err
	}

	db, user, password := database.Mysql(), &entity.User{}, []byte(res.Password)
	err = db.Table("users").Where("email = ?", res.Email).First(user).Error
	if err != nil {
		return "", err
	}

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
    if err != nil {
		return "", err
	}

	signedToken, err := security.GenerateToken(res.Email)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}