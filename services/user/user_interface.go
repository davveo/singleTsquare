package user

import (
	"github.com/davveo/singleTsquare/models"
)

type ServiceInterface interface {
	FindByAccountID(accountID uint) (*models.User, error)
	Create(accountID, role uint, nickName, avatar, gender string) (*models.User, error)
	UpdateUser(user *models.User, role uint, nickName, avatar, gender string) error
	Close()
}
