package services

import (
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"apathy/response"
	"apathy/security"
	"apathy/interfaces"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) CreateUser(r *http.Request) (bool, error) {
	data, err := response.Decode(r)
	if err != nil {
		return false, err
	}

	ok := response.ValidateUser(data.Email, data.Password)
	if !ok {
		return false, nil
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
	data, err := response.Decode(r)
	if err != nil {
		return "", err
	}

	ok := response.ValidateUser(data.Email, data.Password)
	if !ok {
		return "Validation failed", nil
	}

	user, err := service.FetchUser(r, data.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return "", err
	}

	signedToken, err := security.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
