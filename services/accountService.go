package services

import (
	"net/http"
	"encoding/json"
	"regexp"
	"apathy/security"
	"apathy/interfaces"
	"apathy/models"
)

type AccountService struct {
	interfaces.IAccountRepository
}

func (service *AccountService) CreateAccount(r *http.Request) (bool, error) {
	data, err := decodeAccountModel(r)
	if err != nil {
		return false, err
	}

	ok := validateAccount(data.Email, data.Password)
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

	ok, err = service.StoreAccountInDatabase(r, hash, data)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}
	
	return true, nil
}

func (service *AccountService) LoginAccount(r *http.Request) (string, error) {
	data, err := decodeAccountModel(r)
	if err != nil {
		return "", err
	}

	ok := validateAccount(data.Email, data.Password)
	if !ok {
		return "Validation failed", nil
	}

	user, err := service.FetchAccount(r, data.Email)
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

func decodeAccountModel(r *http.Request) (*models.Account, error) {
	var data models.Account
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func validateAccount(email, password string) bool {
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