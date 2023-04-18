package main

import (
	"os"

	"github.com/equinor/radix-log-api/internal/commands"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.DefaultCommand = "start"
	app.Commands = commands.CliCommands

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
