package cmd

import (
	"github.com/equinor/radix-log-api/cmd/start"
	"github.com/urfave/cli/v2"
)

var CliCommands = cli.Commands{
	&start.Command,
}
