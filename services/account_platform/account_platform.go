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
func (s *Service) FindByIdentifyId(identifyId uint) (*models.AccountPlatform, error) {
	accountPlatform := new(models.AccountPlatform)
	notFound := s.db.Where("identify_id = LOWER(?)", identifyId).
		Take(&accountPlatform).RecordNotFound()

	if notFound {
		return nil, errors.New("不存在该用户信息")
	}
	return accountPlatform, nil
}

// 根据用户唯一标示更新用户信息
func (s *Service) UpdateByIdentifyId(identifyId uint) error {
	accountPlatform, err := s.FindByIdentifyId(identifyId)
	if err != nil {
		return errors.New("用户不存在, 更新失败")
	}
	return s.updateAccountPlatform(s.db, identifyId, accountPlatform)
}

// 创建用户信息
func (s *Service) Create(uid, identifyId, platformType uint,
	accesstoken, nickname, avatar string) (*models.AccountPlatform, error) {
	return s.createAccountPlatform(s.db, uid, identifyId, platformType, accesstoken, nickname, avatar)
}

func (s *Service) createAccountPlatform(
	db *gorm.DB, uid, identifyId, platformType uint,
	accesstoken, nickname, avatar string) (*models.AccountPlatform, error) {
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
	db *gorm.DB, identifyId uint,
	accountPlatform *models.AccountPlatform) error {
	accountPlatformUser := models.AccountPlatform{
		IdentifyId: identifyId,
	}
	return db.Model(&accountPlatform).Updates(&accountPlatformUser).Error
}
