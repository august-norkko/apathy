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

func (repository *AccountRepository) SaveAccount(r *http.Request, hashedPassword []byte, account *models.Account) (bool, error) {
	account.Password = string(hashedPassword)
	err := database.Mysql().Create(account).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *AccountRepository) UpdateAccount(r *http.Request, data *models.Account) (bool, error) {
	id := r.Context().Value("id").(uint)
	account := &models.Account{}
	db := database.Mysql()

	err := db.Where("id = ?", id).First(&account).Error
	if err != nil {
		return false, err
	}

	account.Location = data.Location
	account.About = data.About
	err = db.Save(account).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *AccountRepository) CheckForEmailInUse(r *http.Request, email string) bool {
	db := database.Mysql()
	account := &models.Account{}
	err := db.Table(table).Where("email = ?", email).First(account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}

	if account.Email != "" {
		return false
	}

	return true
}

func (repository *AccountRepository) CheckForUsernameInUse(r *http.Request, username string) bool {
	db := database.Mysql()
	account := &models.Account{}
	err := db.Table(table).Where("username = ?", username).First(account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}

	if account.Username != "" {
		return false
	}

	return true
}

func (repository *AccountRepository) FindByUsername(r *http.Request, data *models.Account) (*models.Account, error) {
	account := &models.Account{}

	err := database.Mysql().Table(table).Where("username = ?", data.Username).First(account).Error
	if err != nil {
		return &models.Account{}, err
	}

	return account, nil
}

func (repository *AccountRepository) FindById(r *http.Request) (*models.Account, error) {
	account := &models.Account{}
	id := r.Context().Value("id").(uint)

	err := database.Mysql().Table(table).Where("id = ?", id).First(account).Error
	if err != nil {
		return &models.Account{}, err
	}

	return account, nil
}

func (repository *AccountRepository) DeleteById(r *http.Request, data *models.Account) (bool, error) {
	id := r.Context().Value("id").(uint)
	account := &models.Account{}
	db := database.Mysql()

	err := db.Where("id = ?", id).Delete(&account).Error
	if err != nil {
		return false, err
	}

	return true, nil
}