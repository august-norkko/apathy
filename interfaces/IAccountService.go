package interfaces

import (
	"net/http"
	"apathy/models"
)

type IAccountService interface {
	Create(r *http.Request) (bool, error)
	Login(r *http.Request) (string, error)
	Update(r *http.Request) (bool, error)
	Fetch(r *http.Request) (*models.Account, error)
	Delete(r *http.Request) (bool, error)
}