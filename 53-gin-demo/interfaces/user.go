package interfaces

import (
	"demo/models"
)

type IUser interface {
	CreateUser(user *models.User) (*models.User, error)
}
