package services

import (
	"errors"
	"fmt"

	pass "github.com/davveo/singleTsquare/utils/password"

	"github.com/davveo/singleTsquare/utils"

	"github.com/davveo/singleTsquare/models"
	"github.com/jinzhu/gorm"
)

var (
	MinPasswordLength   = 6
	ErrUserNotFound     = errors.New("User not found")
	ErrUsernameTaken    = errors.New("Username taken")
	ErrPasswordTooShort = fmt.Errorf(
		"Password must be at least %d characters long",
		MinPasswordLength,
	)
)

type UserService struct {
	db *gorm.DB
}

func (s *UserService) Create(userName, password string) (*models.User, error) {
	return s.createUser(s.db, userName, password)

}

func (s *UserService) FindUserByUsername(username string) (*models.User, error) {
	// Usernames are case insensitive
	user := new(models.User)
	notFound := s.db.Where("username = LOWER(?)", username).
		First(user).RecordNotFound()

	// Not found
	if notFound {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *UserService) UserExists(username string) bool {
	_, err := s.FindUserByUsername(username)
	return err == nil
}

func (s *UserService) createUser(db *gorm.DB, userName, password string) (*models.User, error) {
	user := &models.User{
		Username: userName,
		Password: utils.StringOrNull(""),
	}
	if password != "" {
		if len(password) < MinPasswordLength {
			return nil, ErrPasswordTooShort
		}
		passwordHash, err := pass.HashPassword(password)
		if err != nil {
			return nil, err
		}
		user.Password = utils.StringOrNull(string(passwordHash))
	}

	if s.UserExists(user.Username) {
		return nil, ErrUsernameTaken
	}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
