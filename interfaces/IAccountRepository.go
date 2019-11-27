package interfaces

import (
	"net/http"
	"apathy/models"
)

type IAccountRepository interface {
	StoreAccountInDatabase(r *http.Request, hashedPassword []byte, data *models.Account) (bool, error)
	FetchAccount(r *http.Request, email string) (*models.Account, error)
	CheckForExistingEmail(r *http.Request, data *models.Account) bool
}