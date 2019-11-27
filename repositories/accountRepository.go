package repositories

import (
	"net/http"
	"apathy/interfaces"
	"apathy/models"
	"apathy/database"
	"github.com/jinzhu/gorm"
)

const (
	table = "accounts"
)

type AccountRepository struct {
	AccountRepository interfaces.IAccountRepository
}

func (repository *AccountRepository) CheckForExistingEmail(r *http.Request, data *models.Account) bool {
	db := database.Mysql()
	user := &models.Account{}
	err := db.Table(table).Where("email = ?", data.Email).First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}

	if user.Email != "" {
		return false
	}

	return true
}

func (repository *AccountRepository) StoreAccountInDatabase(r *http.Request, hashedPassword []byte, data *models.Account) (bool, error) {
	db := database.Mysql()
	user := &models.Account{ Email: data.Email, Password: string(hashedPassword) }
	err := db.Create(user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *AccountRepository) FetchAccount(r *http.Request, email string) (*models.Account, error) {
	db := database.Mysql()
	user := &models.Account{}

	err := db.Table(table).Where("email = ?", email).First(user).Error
	if err != nil {
		return &models.Account{}, err
	}

	return user, nil
}
