package db

import (
	"fmt"
	"time"

	"github.com/davveo/singleTsquare/config"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}
}

func NewDataBase(cfg *config.Config) (*gorm.DB, error) {
	// 这个地方可以考虑使用设计模式
	if cfg.Database.Type == "mysql" {
		args := fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.DatabaseName,
		)
		db, err := gorm.Open(cfg.Database.Type, args)
		if err != nil {
			return db, err
		}

		// Max idle connections
		db.DB().SetMaxIdleConns(cfg.Database.MaxIdleConns)

		// Max open connections
		db.DB().SetMaxOpenConns(cfg.Database.MaxOpenConns)

		// Database logging
		db.LogMode(cfg.IsDevelopment)

		//全局禁用表名复数
		db.SingularTable(true)

		return db, nil
	}
	return nil, fmt.Errorf("Database type %s not supported", cfg.Database.Type)
}
