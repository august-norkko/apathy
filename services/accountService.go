package services

import (
	"net/http"
	"encoding/json"
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

	ok := data.ValidateNewAccount(data.Email, data.Password)
	if !ok {
		return false, nil
	}

	hash, err := security.Generate([]byte(data.Password))
	if err != nil {
		return false, err
	}

	ok = service.CheckForExistingEmailInDatabase(r, data)
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

	ok := data.ValidateNewAccount(data.Email, data.Password)
	if !ok {
		return "Validation failed", nil
	}

	account, err := service.FetchAccountFromDatabase(r)
	if err != nil {
		return "", err
	}

	ok, err = security.Compare([]byte(account.Password), []byte(data.Password))
	if err != nil || !ok {
		return "", err
	}

	signedToken, err := security.GenerateToken(account.ID)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (service *AccountService) UpdateAccount(r *http.Request) (bool, error) {
	updatedAccount, err := decodeAccountModel(r)
	if err != nil {
		return false, err
	}

	ok, err := service.UpdateAccountInDatabase(r, updatedAccount)
	if !ok {
		return false, err
	}

	return true, nil
}

func (service *AccountService) FetchAccount(r *http.Request) (*models.Account, error) {
	account, err := service.FetchAccountFromDatabase(r)
	if err != nil {
		return &models.Account{}, err
	}

	return account, nil
}

func decodeAccountModel(r *http.Request) (*models.Account, error) {
	var data models.Account
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

