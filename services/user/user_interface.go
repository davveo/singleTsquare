package user

import (
	"github.com/davveo/singleTsquare/models"
)

type ServiceInterface interface {
	FindByUid(uid uint) (*models.User, error)
	Create(uid, role uint, nickName, avatar, gender string) (*models.User, error)
	UpdateUser(user *models.User, role uint, nickName, avatar, gender string) error
	Close()
}
