package interfaces

import (
	"net/http"
	"apathy/models"
)

type IAccountService interface {
	CreateAccount(r *http.Request) (bool, error)
	LoginAccount(r *http.Request) (string, error)
	UpdateAccount(r *http.Request) (bool, error)
	FetchAccount(r *http.Request) (*models.Account, error)
	DeleteAccount(r *http.Request) (bool, error)
}