package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/equinor/radix-log-api/cmd/start"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

const (
	logLevel  = "log-level"
	logPretty = "log-pretty"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    logLevel,
		Usage:   "Log level",
		Value:   zerolog.LevelInfoValue,
		EnvVars: []string{"LOG_API_LOG_LEVEL"},
	},
	&cli.BoolFlag{
		Name:    logPretty,
		Usage:   "Log in human readable format",
		Value:   false,
		EnvVars: []string{"LOG_API_LOG_PRETTY"},
	},
}

var commands = cli.Commands{
	&start.Command,
}

func Run() {
	app := cli.NewApp()
	app.DefaultCommand = "start"
	app.Commands = commands
	app.Flags = flags
	app.Before = initLogger

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func initLogger(ctx *cli.Context) error {
	level, err := zerolog.ParseLevel(ctx.String(logLevel))
	if err != nil {
		return fmt.Errorf("unknown log level `%s`", ctx.String(logLevel))
	}
	zerolog.SetGlobalLevel(level)
	logger := log.Logger
	if ctx.Bool(logPretty) {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly})
	}
	ctx.Context = logger.WithContext(ctx.Context)
	return nil
}
