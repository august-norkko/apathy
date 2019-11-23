package interfaces

import (
	"net/http"
	"apathy/models"
)

type IUserRepository interface {
	StoreUserInDatabase(r *http.Request, hashedPassword []byte, data *models.User) (bool, error)
	FetchUser(r *http.Request, email string) (*models.User, error)
	CheckForExistingEmail(r *http.Request, data *models.User) bool
}