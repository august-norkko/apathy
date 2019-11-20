package repositories

import (
	"net/http"
	"apathy/interfaces"
	"apathy/models"
	"apathy/database"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	UserRepository interfaces.IUserRepository
}

func (repository *UserRepository) CheckForExistingEmail(r *http.Request, data *models.User) bool {
	db := database.Mysql()
	user := &models.User{}
	err := db.Table("users").Where("email = ?", data.Email).First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}

	if user.Email != "" {
		return false
	}

	return true
}

func (repository *UserRepository) StoreUserInDatabase(r *http.Request, hashedPassword []byte, data *models.User) (bool, error) {
	db := database.Mysql()
	user := &models.User{ Email: data.Email, Password: string(hashedPassword) }
	err := db.Create(user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

