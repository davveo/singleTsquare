package account

import (
	"errors"
	"fmt"
	"time"

	"github.com/davveo/singleTsquare/services/user"

	"github.com/davveo/singleTsquare/models"
	"github.com/davveo/singleTsquare/utils"
	pass "github.com/davveo/singleTsquare/utils/password"
	"github.com/davveo/singleTsquare/utils/randomstr"
	"github.com/jinzhu/gorm"
)

var (
	MinPasswordLength   = 6
	DEFAULT_AVATAR_PATH = "./avatar.png"
	ErrUserNotFound     = errors.New("user not found")
	ErrUsernameTaken    = errors.New("username taken")
	ErrUserPhoneTaken   = errors.New("phone has taken")
	ErrPasswordTooShort = fmt.Errorf("password must be at least %d characters long", MinPasswordLength)
)

type Service struct {
	db          *gorm.DB
	userService user.ServiceInterface
}

func NewService(db *gorm.DB, userService user.ServiceInterface) *Service {
	return &Service{
		db:          db,
		userService: userService,
	}
}

func (s *Service) Close() {}

func (s *Service) Create(userName, password, phone, email, createIpAt string) (*models.Account, error) {
	return s.createAccount(s.db, userName, password, phone, email, createIpAt)
}

func (s *Service) FindByName(username string) (*models.Account, error) {
	// TODO 用户名进行脱敏处理
	account := new(models.Account)
	notFound := s.db.Where("username = LOWER(?)", username).
		Take(&account).RecordNotFound()

	if notFound {
		return nil, ErrUserNotFound
	}

	return account, nil
}

func (s *Service) FindByPhone(phone string) (*models.Account, error) {
	account := new(models.Account)
	notFound := s.db.Where(
		"phone = LOWER(?)", phone).Take(&account).RecordNotFound()

	if notFound {
		return nil, ErrUserNotFound
	}

	return account, nil
}

func (s *Service) FindByEmail(email string) (*models.Account, error) {
	account := new(models.Account)
	notFound := s.db.Where(
		"email = LOWER(?)", email).Take(&account).RecordNotFound()

	if notFound {
		return nil, ErrUserNotFound
	}

	return account, nil
}

func (s *Service) ExistByUserName(username string) bool {
	_, err := s.FindByName(username)
	return err == nil
}

func (s *Service) ExistByPhone(phone string) bool {
	_, err := s.FindByPhone(phone)
	return err == nil
}

func (s *Service) ExistByMail(email string) bool {
	_, err := s.FindByEmail(email)
	return err == nil
}

func (s *Service) UpdateAccountIp(lastLoginIpAt string, account *models.Account) error {
	return s.updateAccountLoginIp(s.db, lastLoginIpAt, account)
}

func (s *Service) UpdateAccountPassword(password string, account *models.Account) error {
	return s.updateAccountPassword(s.db, password, account)
}

func (s *Service) FindByLoginId(loginId string) (*models.Account, error) {
	// login_id= email或者phone或者username
	var account *models.Account
	notFound := s.db.Where("email = LOWER(?)", loginId).
		Or("username = LOWER(?)", loginId).
		Or("phone = ?", loginId).
		Take(&account).RecordNotFound()
	if notFound {
		return nil, ErrUserNotFound
	}
	return account, nil
}

func (s *Service) createAccount(
	db *gorm.DB, userName,
	passWord, phone, email,
	createIpAt string) (*models.Account, error) {
	account := &models.Account{
		Status:        1,
		Phone:         phone,
		Email:         email,
		UserName:      userName,
		CreateIpAt:    createIpAt,
		LastLoginIpAt: createIpAt,
		PassWord:      utils.StringOrNull(""),
		LastLoginAt:   time.Now(),
	}

	if passWord != "" {
		if len(passWord) < MinPasswordLength {
			return nil, ErrPasswordTooShort
		}
		passwordHash, err := pass.HashPassword(passWord)
		if err != nil {
			return nil, err
		}
		account.PassWord = utils.StringOrNull(string(passwordHash))
	}

	tx := db.Begin()
	if err := tx.Create(account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	NickName := fmt.Sprintf("tsquare_%s", randomstr.GenRandomString(8))

	// TODO 默认用户头像路径
	if _, err := s.userService.Create(
		account.ID, 0, NickName,
		DEFAULT_AVATAR_PATH, "unknow"); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return account, nil
}

func (s *Service) updateAccountLoginIp(
	db *gorm.DB,
	lastLoginIpAt string,
	account *models.Account) error {
	accountUser := models.Account{
		LastLoginAt:   time.Now(),
		LastLoginIpAt: lastLoginIpAt,
		LoginTimes:    account.LoginTimes + 1,
	}
	return db.Model(&account).Updates(&accountUser).Error
}

func (s *Service) updateAccountPassword(db *gorm.DB,
	password string,
	account *models.Account) error {
	accountUser := models.Account{
		PassWord:   utils.StringOrNull(""),
		LoginTimes: account.LoginTimes + 1,
	}

	if password != "" {
		if len(password) < MinPasswordLength {
			return ErrPasswordTooShort
		}
		passwordHash, err := pass.HashPassword(password)
		if err != nil {
			return err
		}
		account.PassWord = utils.StringOrNull(string(passwordHash))
	}

	return db.Model(&account).Updates(&accountUser).Error
}
