package account

import "github.com/davveo/singleTsquare/models"

type ServiceInterface interface {
	ExistByPhone(phone string) bool
	ExistByMail(email string) bool
	ExistByUserName(username string) bool
	FindByPhone(phone string) (*models.Account, error)
	FindByEmail(email string) (*models.Account, error)
	FindByLoginId(loginId string) (*models.Account, error)
	FindByName(username string) (*models.Account, error)
	UpdateAccount(lastLoginIpAt string, account *models.Account) error
	Create(userName, password, phone, email, createIpAt string) (*models.Account, error)
	Close()
}
