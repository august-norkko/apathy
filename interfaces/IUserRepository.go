package interfaces

import (
	"net/http"
	"apathy/models"
)

type IUserRepository interface {
	StoreUserInDatabase(r *http.Request, hashedPassword []byte, data *models.User) (bool, error)
	CheckForExistingEmail(r *http.Request, data *models.User) bool
}