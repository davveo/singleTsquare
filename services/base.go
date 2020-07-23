package services

import (
	"reflect"

	"github.com/davveo/singleTsquare/services/healthcheck"
	"github.com/davveo/singleTsquare/services/user"

	"github.com/davveo/singleTsquare/config"
	"github.com/jinzhu/gorm"
)

var (
	HealthService healthcheck.ServiceInterface
	UserService   user.ServiceInterface
)

// UseHealthService sets the health service
func UseHealthService(h healthcheck.ServiceInterface) {
	HealthService = h
}

// UseOauthService sets the oAuth service
func UseOauthService(u user.ServiceInterface) {
	UserService = u
}

func Init(cfg *config.Config, db *gorm.DB) error {
	if nil == reflect.TypeOf(HealthService) {
		HealthService = healthcheck.NewService(db)
	}

	if nil == reflect.TypeOf(UserService) {
		UserService = user.NewService(db)
	}

	return nil
}

func Close() {
	HealthService.Close()
	UserService.Close()
}
