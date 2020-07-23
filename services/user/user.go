package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/davveo/singleTsquare/utils/randomstr"

	pass "github.com/davveo/singleTsquare/utils/password"

	"github.com/davveo/singleTsquare/utils"

	"github.com/davveo/singleTsquare/models"
	"github.com/jinzhu/gorm"
)

var (
	MinPasswordLength   = 6
	ErrUserNotFound     = errors.New("User not found")
	ErrUsernameTaken    = errors.New("Username taken")
	ErrUserPhoneTaken   = errors.New("Phone has taken")
	ErrPasswordTooShort = fmt.Errorf(
		"Password must be at least %d characters long",
		MinPasswordLength,
	)
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

func (s *Service) Create(userName, password, phone, createIpAt string) (*models.AccountUser, error) {
	return s.createUser(s.db, userName, password, phone, createIpAt)

}

func (s *Service) FindUserByUsername(username string) (*models.AccountUser, error) {
	// Usernames are case insensitive
	user := new(models.AccountUser)
	notFound := s.db.Where("username = LOWER(?)", username).
		First(user).RecordNotFound()

	// Not found
	if notFound {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *Service) FindUserByPhone(phone string) (*models.AccountUser, error) {
	user := new(models.AccountUser)
	notFound := s.db.Where("phone = LOWER(?)", phone).
		First(user).RecordNotFound()

	// Not found
	if notFound {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *Service) UserExistByUsername(username string) bool {
	_, err := s.FindUserByUsername(username)
	return err == nil
}

func (s *Service) UserExistByPhone(phone string) bool {
	_, err := s.FindUserByPhone(phone)
	return err == nil
}

func (s *Service) createUser(db *gorm.DB, userName, passWord, phone, createIpAt string) (*models.AccountUser, error) {
	accountUser := &models.AccountUser{
		Status:      1,
		Phone:       phone,
		UserName:    userName,
		CreateIpAt:  createIpAt,
		PassWord:    utils.StringOrNull(""),
		LastLoginAt: time.Now(),
	}

	if passWord != "" {
		if len(passWord) < MinPasswordLength {
			return nil, ErrPasswordTooShort
		}
		passwordHash, err := pass.HashPassword(passWord)
		if err != nil {
			return nil, err
		}
		accountUser.PassWord = utils.StringOrNull(string(passwordHash))
	}

	tx := db.Begin()
	if err := tx.Create(accountUser).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	NickName := fmt.Sprintf("tsquare_%s", randomstr.GenRandomString(8))
	member := &models.Member{
		Uid:      accountUser.ID,
		NickName: NickName,
		Avatar:   "./avatar.png",
	}
	if err := tx.Create(&member).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return accountUser, nil
}
