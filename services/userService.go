package services

import (
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

	ok := utils.ValidateUser(data.Email, data.Password)
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

	ok := utils.ValidateUser(data.Email, data.Password)
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
