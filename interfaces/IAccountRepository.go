package interfaces

import (
	"net/http"
	"apathy/models"
)

type IAccountRepository interface {
	StoreAccountInDatabase(r *http.Request, hashedPassword []byte, data *models.Account) (bool, error)
	UpdateAccountInDatabase(r *http.Request, data *models.Account) (bool, error)
	FetchAccountFromDatabase(r *http.Request) (*models.Account, error)
	CheckForEmailInUse(r *http.Request, email string) bool
	CheckForUsernameInUse(r *http.Request, username string) bool
}