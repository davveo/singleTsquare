package account

import "github.com/davveo/singleTsquare/models"

type ServiceInterface interface {
	ExistByPhone(phone string) bool
	ExistByUserName(username string) bool
	FindByPhone(phone string) (*models.Account, error)
	FindByName(username string) (*models.Account, error)
	UpdateAccount(lastLoginIpAt string, account *models.Account) error
	Create(userName, password, phone, createIpAt string) (*models.Account, error)
	Close()
}
