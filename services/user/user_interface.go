package user

import (
	"github.com/davveo/singleTsquare/models"
)

type ServiceInterface interface {
	Create(userName, password, phone, createIpAt string) (*models.AccountUser, error)
	FindUserByUsername(username string) (*models.AccountUser, error)
	FindUserByPhone(phone string) (*models.AccountUser, error)
	UserExistByUsername(username string) bool
	UserExistByPhone(phone string) bool
	Close()
}
