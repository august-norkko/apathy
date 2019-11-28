package interfaces

import (
	"net/http"
	"apathy/models"
)

type IAccountRepository interface {
	StoreAccountInDatabase(r *http.Request, hashedPassword []byte, data *models.Account) (bool, error)
	UpdateAccountInDatabase(r *http.Request, data *models.Account) (bool, error)
	CheckForExistingEmailInDatabase(r *http.Request, data *models.Account) bool
	FetchAccountFromDatabase(r *http.Request) (*models.Account, error)
}