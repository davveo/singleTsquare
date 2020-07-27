package account_platform

import (
	"errors"

	"github.com/davveo/singleTsquare/models"
	"github.com/jinzhu/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Close() {}

// 根据用户标示查找用户信息
func (s *Service) FindByIdentifyId(identifyId string) (*models.AccountPlatform, error) {
	accountPlatform := new(models.AccountPlatform)
	notFound := s.db.Where("identify_id = LOWER(?)", identifyId).
		Take(&accountPlatform).RecordNotFound()

	if notFound {
		return nil, errors.New("不存在该用户信息")
	}
	return accountPlatform, nil
}

// 根据用户唯一标示更新用户信息
func (s *Service) UpdateByIdentifyId(accesstoken,
	nickname, avatar string, accountPlatform *models.AccountPlatform) error {
	return s.updateAccountPlatform(s.db, accesstoken, nickname, avatar, accountPlatform)
}

// 创建用户信息
func (s *Service) Create(uid, platformType uint,
	identifyId, accesstoken, nickname, avatar string) (*models.AccountPlatform, error) {
	return s.createAccountPlatform(s.db, uid, platformType, identifyId, accesstoken, nickname, avatar)
}

func (s *Service) createAccountPlatform(
	db *gorm.DB, uid, platformType uint,
	identifyId, accesstoken, nickname, avatar string) (*models.AccountPlatform, error) {
	accountPlatform := &models.AccountPlatform{
		Uid:          uid,
		IdentifyId:   identifyId,
		Accesstoken:  accesstoken,
		Avatar:       avatar,
		NickName:     nickname,
		PlatformType: platformType,
	}
	if err := db.Create(accountPlatform).Error; err != nil {
		return nil, err
	}
	return accountPlatform, nil
}

func (s *Service) updateAccountPlatform(
	db *gorm.DB, accesstoken, nickname, avatar string,
	accountPlatform *models.AccountPlatform) error {
	accountPlatformUser := models.AccountPlatform{
		Accesstoken: accesstoken,
		NickName:    nickname,
		Avatar:      avatar,
	}
	return db.Model(&accountPlatform).Updates(&accountPlatformUser).Error
}
