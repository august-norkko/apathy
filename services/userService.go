package services

import (
	"net/http"
	"encoding/json"
	"regexp"
	"apathy/security"
	"apathy/interfaces"
	"apathy/models"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) CreateUser(r *http.Request) (bool, error) {
	data, err := decode(r)
	if err != nil {
		return false, err
	}

	ok := validateUser(data.Email, data.Password)
	if !ok {
		return false, nil
	}

	hash, err := security.Generate([]byte(data.Password))
	if err != nil {
		return false, err
	}

	ok = service.CheckForExistingEmail(r, data)
	if !ok {
		return false, nil
	}

	ok, err = service.StoreUserInDatabase(r, hash, data)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}
	
	return true, nil
}

func (service *UserService) LoginUser(r *http.Request) (string, error) {
	data, err := decode(r)
	if err != nil {
		return "", err
	}

	ok := validateUser(data.Email, data.Password)
	if !ok {
		return "Validation failed", nil
	}

	user, err := service.FetchUser(r, data.Email)
	if err != nil {
		return "", err
	}

	ok, err = security.Compare([]byte(user.Password), []byte(data.Password))
	if err != nil || !ok {
		return "", err
	}

	signedToken, err := security.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func decode(r *http.Request) (*models.User, error) {
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
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