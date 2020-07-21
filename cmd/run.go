package cmd

import (
	"strconv"
	"time"

	"gopkg.in/tylerb/graceful.v1"

	"github.com/davveo/singleTsquare/router"
	"github.com/davveo/singleTsquare/services"
	log "github.com/davveo/singleTsquare/utils"
)

func RunServer(configFile string) error {
	cfg, db, err := initConfig(configFile)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := services.Init(cfg, db); err != nil {
		return err
	}
	defer services.Close()
	log.INFO.Println("Starting gravitee server on port ", cfg.ServerPort)
	graceful.Run(":"+strconv.Itoa(cfg.ServerPort), 5*time.Second, router.Routers())
	return nil
}
