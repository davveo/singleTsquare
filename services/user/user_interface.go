package user

import (
	"github.com/davveo/singleTsquare/models"
)

type ServiceInterface interface {
	Create(uid, role uint, nickName, avatar, gender string) (*models.User, error)
	UpdateUser(user *models.User, role uint, nickName, avatar, gender string) error
	Close()
}
