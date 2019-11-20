package services

import (
	"regexp"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"apathy/utils"
	"apathy/database"
	"apathy/models"
	"apathy/security"
	"apathy/interfaces"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) CreateUser(r *http.Request) (bool, error) {
	data, err := utils.Decode(r)
	if err != nil {
		return false, err
	}

	ok := validateUser(data.Email, data.Password)
	if !ok {
		return false, err
	}

	password := []byte(data.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
		return false, err
	}

	ok = service.CheckForExistingEmail(r, data)
	if !ok {
		return false, nil
	}

	ok, err = service.StoreUserInDatabase(r, hashedPassword, data)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}
	
	return true, nil
}

func (service *UserService) LoginUser(r *http.Request) (string, error) {
	data, err := utils.Decode(r)
	if err != nil {
		return "", err
	}

	ok := validateUser(data.Email, data.Password)
	if !ok {
		return "Validation failed", err
	}

	db := database.Mysql()
	user := &models.User{}
	password := []byte(data.Password)
	err = db.Table("users").Where("email = ?", data.Email).First(user).Error
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	if err != nil {
		return "", err
	}

	signedToken, err := security.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func validateUser(email, password string) bool {
	ok, err := regexp.MatchString(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, email)
	if err != nil {
		return false
	}

	if !ok {
		return false
	}

	if len(password) < 5 {
		return false
	}

	return true
}