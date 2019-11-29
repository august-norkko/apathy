package interfaces

import (
	"net/http"
	"apathy/models"
)

type IAccountRepository interface {
	SaveAccount(r *http.Request, hashedPassword []byte, data *models.Account) (bool, error)
	UpdateAccount(r *http.Request, data *models.Account) (bool, error)

	CheckForEmailInUse(r *http.Request, email string) bool
	CheckForUsernameInUse(r *http.Request, username string) bool

	FindByUsername(r *http.Request, data *models.Account) (*models.Account, error)
	FindById(r *http.Request) (*models.Account, error)
	DeleteById(r *http.Request, data *models.Account) (bool, error)
}