package services

import (
	"errors"
	"fmt"
	"time"

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

type UserService struct {
	db *gorm.DB
}

func (s *UserService) Create(userName, password, phone string) (*models.AccountUser, error) {
	return s.createUser(s.db, userName, password, phone)

}

func (s *UserService) FindUserByUsername(username string) (*models.AccountUser, error) {
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

func (s *UserService) FindUserByPhone(phone string) (*models.AccountUser, error) {
	user := new(models.AccountUser)
	notFound := s.db.Where("phone = LOWER(?)", phone).
		First(user).RecordNotFound()

	// Not found
	if notFound {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *UserService) UserExistByUsername(username string) bool {
	_, err := s.FindUserByUsername(username)
	return err == nil
}

func (s *UserService) UserExistByPhone(phone string) bool {
	_, err := s.FindUserByPhone(phone)
	return err == nil
}

func (s *UserService) createUser(db *gorm.DB, userName, passWord, phone, createIpAt string) (*models.AccountUser, error) {
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

	member := &models.Member{Uid: accountUser.ID, NickName: "", Avatar: ""}
	if err := tx.Create(&member).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return accountUser, nil
}
