package start

import (
	"github.com/urfave/cli/v2"
)

var Command = cli.Command{
	Name:   "start",
	Usage:  "Starts the web server",
	Action: action,
	Flags:  flags,
}
