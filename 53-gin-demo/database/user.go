package database

import (
	"demo/models"

	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) CreateUser(user *models.User) (*models.User, error) {
	u.db.AutoMigrate(models.User{})
	tx := u.db.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
