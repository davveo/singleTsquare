package services

import (
	"github.com/davveo/singleTsquare/config"
	"github.com/jinzhu/gorm"
)

func Init(cfg *config.Config, db *gorm.DB) error {
	return nil
}

func Close() {

}
