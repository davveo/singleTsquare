package account_platform

import "github.com/davveo/singleTsquare/models"

type ServiceInterface interface {
	FindByIdentifyId(identifyId string) (*models.AccountPlatform, error)
	UpdateByIdentifyId(accesstoken, nickname, avatar string, accountPlatform *models.AccountPlatform) error
	Create(uid, platformType uint, identifyId, accesstoken, nickname, avatar string) (*models.AccountPlatform, error)
	Close()
}
