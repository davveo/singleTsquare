package cmd

import (
	"github.com/davveo/singleTsquare/config"
	database "github.com/davveo/singleTsquare/initialize/db"
	"github.com/jinzhu/gorm"
)

func initConfig(configFile string) (*config.Config, *gorm.DB, error) {
	cfg := config.NewConfig(configFile)

	db, err := database.NewDataBase(cfg)
	if err != nil {
		return nil, nil, err
	}
	return cfg, db, nil
}
