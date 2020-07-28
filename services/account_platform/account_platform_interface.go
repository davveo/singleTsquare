package account_platform

import "github.com/davveo/singleTsquare/models"

type ServiceInterface interface {
	FindByIdentifyId(identifyID string) (*models.AccountPlatform, error)
	UpdateByIdentifyId(accesstoken, nickname, avatar string, accountPlatform *models.AccountPlatform) error
	UpdateAccountId(accountID uint, accountPlatform *models.AccountPlatform) error
	Create(accountID, platformType uint, identifyID, accesstoken, nickname, avatar string) (*models.AccountPlatform, error)
	Close()
}
