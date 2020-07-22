package main

import (
	"os"

	"github.com/davveo/singleTsquare/cmd"
	"github.com/davveo/singleTsquare/utils/log"
	"github.com/urfave/cli"
)

var (
	cliApp     *cli.App
	configFile string
)

func init() {
	cliApp = cli.NewApp()
	cliApp.Name = "tsquare"
	cliApp.Version = "0.0.1"
	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "configFile",
			Value:       "config.yml",
			Destination: &configFile,
		},
	}
}

func main() {
	cliApp.Commands = []cli.Command{
		{
			Name:  "migrate",
			Usage: "run migrations",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "runserver",
			Usage: "run web server",
			Action: func(c *cli.Context) error {
				return cmd.RunServer(configFile)
			},
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		log.INFO.Fatal(err)
	}
}
