package user

import (
	"errors"

	"github.com/davveo/singleTsquare/models"
	"github.com/jinzhu/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
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

func (s *Service) Create(uid, role uint, nickName, avatar, gender string) (*models.User, error) {
	return s.createUser(s.db, uid, role, nickName, avatar, gender)

}

func (s *Service) UpdateUser(user *models.User, role uint, nickName, avatar, gender string) error {
	return s.updateUser(s.db, user, role, nickName, avatar, gender)
}

func (s *Service) FindByUid(uid uint) (*models.User, error) {
	user := new(models.User)
	notFound := s.db.Where("uid = ?", uid).
		Take(&user).RecordNotFound()

	if notFound {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *Service) createUser(tx *gorm.DB,
	uid, role uint, nickName,
	avatar, gender string) (user *models.User, err error) {
	user = &models.User{
		Uid:      uid,
		NickName: nickName,
		Avatar:   avatar,
		Gender:   gender,
		Role:     role,
	}

	if err := tx.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) updateUser(tx *gorm.DB, user *models.User,
	role uint, nickName, avatar, gender string) error {
	updateUser := models.User{
		NickName: nickName,
		Avatar:   avatar,
		Gender:   gender,
		Role:     role,
	}
	return tx.Model(&user).Updates(&updateUser).Error

}
