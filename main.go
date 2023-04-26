package main

import (
	"os"

	"github.com/equinor/radix-log-api/cmd"
	_ "github.com/equinor/radix-log-api/docs"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// @BasePath /api/v1
// @Schemes http https
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description					Bearer is currently not supported by go-swag. Use "Bearer <JWT>" in value.
func main() {
	app := cli.NewApp()
	app.DefaultCommand = "start"
	app.Commands = cmd.CliCommands

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
