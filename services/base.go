package services

import (
	"reflect"

	"github.com/davveo/singleTsquare/services/account"
	"github.com/davveo/singleTsquare/services/account_platform"
	"github.com/davveo/singleTsquare/services/healthcheck"
	"github.com/davveo/singleTsquare/services/user"

	"github.com/davveo/singleTsquare/config"
	"github.com/jinzhu/gorm"
)

var (
	HealthService          healthcheck.ServiceInterface
	UserService            user.ServiceInterface
	AccountService         account.ServiceInterface
	AccountPlatformService account_platform.ServiceInterface
)

func UseHealthService(h healthcheck.ServiceInterface) {
	HealthService = h
}

func UseUserService(u user.ServiceInterface) {
	UserService = u
}

func UseAccountService(u user.ServiceInterface) {
	UserService = u
}

func UseAccountPlatformService(u account_platform.ServiceInterface) {
	AccountPlatformService = u
}

func Init(cfg *config.Config, db *gorm.DB) error {
	if nil == reflect.TypeOf(HealthService) {
		HealthService = healthcheck.NewService(db)
	}

	if nil == reflect.TypeOf(UserService) {
		UserService = user.NewService(db)
	}

	if nil == reflect.TypeOf(AccountService) {
		AccountService = account.NewService(db, UserService)
	}

	if nil == reflect.TypeOf(AccountPlatformService) {
		AccountPlatformService = account_platform.NewService(db)
	}

	return nil
}

func Close() {
	HealthService.Close()
	UserService.Close()
	AccountService.Close()
	AccountPlatformService.Close()
}
