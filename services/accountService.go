package services

import (
	"net/http"
	"encoding/json"
	"apathy/security"
	"apathy/interfaces"
	"apathy/models"
	"fmt"
)

type AccountService struct {
	interfaces.IAccountRepository
}

func (service *AccountService) Create(r *http.Request) (bool, error) {
	data, err := decodeAccountModel(r)
	if err != nil {
		return false, err
	}

	ok := data.ValidateNewAccount(data)
	if !ok {
		return false, nil
	}

	hash, err := security.Generate([]byte(data.Password))
	if err != nil {
		return false, err
	}

	fmt.Println(data)
	ok = service.CheckForEmailInUse(r, data.Email)
	if !ok {
		return false, nil
	}

	ok = service.CheckForUsernameInUse(r, data.Username)
	if !ok {
		return false, nil
	}

	ok, err = service.SaveAccount(r, hash, data)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}
	
	return true, nil
}

func (service *AccountService) Login(r *http.Request) (string, error) {
	data, err := decodeAccountModel(r)
	if err != nil {
		return "", err
	}

	account, err := service.FindByUsername(r, data)
	if err != nil {
		return "", err
	}

	ok, err := security.Compare([]byte(account.Password), []byte(data.Password))
	if !ok {
		return "", err
	}

	signedToken, err := security.GenerateToken(account.ID)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (service *AccountService) Update(r *http.Request) (bool, error) {
	updatedAccount, err := decodeAccountModel(r)
	if err != nil {
		return false, err
	}

	ok, err := service.UpdateAccount(r, updatedAccount)
	if !ok {
		return false, err
	}

	return true, nil
}

func (service *AccountService) Fetch(r *http.Request) (*models.Account, error) {
	account, err := service.FindById(r)
	if err != nil {
		return &models.Account{}, err
	}

	return account, nil
}

func (service *AccountService) Delete(r *http.Request) (bool, error) {
	account, err := decodeAccountModel(r)
	if err != nil {
		return false, err
	}

	ok, err := service.DeleteById(r, account)
	if !ok {
		return false, err
	}

	return true, nil
}

func decodeAccountModel(r *http.Request) (*models.Account, error) {
	var data models.Account
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

