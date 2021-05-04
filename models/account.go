package models

import (
	"errors"

	"gorm.io/gorm"
)

type Account struct {
	Id        int    `json:"id"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	Name      string `json:"name" form:"name"`
	Create_at string `json:"-"`
	Update_at string `json:"-"`
}

// // Repository generic repo.
// type Repository interface {
// 	GetAll() []interface{}
// 	GetOne(id string) interface{}
// 	Save(item interface{}) interface{}
// 	Delete(id string) interface{}
// }

// AccountRepository repo for account
type AccountsRepository struct {
	DB *gorm.DB
}

func (Account) TableName() string {
	return "wt_account"
}

func (ar AccountsRepository) GetAll() []Account {
	var account []Account
	ar.DB.Find(&account)
	return account
}

// GetOne fetch book by id from database
func (ar AccountsRepository) GetOne(id string) (Account, error) {
	var account Account
	ar.DB.Find(&account, id)
	if account.Username == "" {
		return Account{}, errors.New("can't find this account")
	}
	return account, nil
}

// Save add new book to database
func (ar AccountsRepository) Save(item Account) Account {
	var account Account = item
	ar.DB.Create(&account)
	return account
}

// Delete a book by id from database
func (ar AccountsRepository) Delete(id string) error {
	var account Account
	ar.DB.First(&account, id)
	if account.Username == "" {
		return errors.New("This account not found")
	}
	ar.DB.Delete(&account)
	return nil
}
