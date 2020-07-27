package account_platform

import "github.com/davveo/singleTsquare/models"

type ServiceInterface interface {
	UpdateByIdentifyId(identifyId uint) error
	FindByIdentifyId(identifyId uint) (*models.AccountPlatform, error)
	Create(uid, identifyId, platformType uint, accesstoken, nickname, avatar string) (*models.AccountPlatform, error)
	Close()
}
