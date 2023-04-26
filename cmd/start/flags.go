package start

import "github.com/urfave/cli/v2"

const (
	HostName                = "host"
	PortNumber              = "port"
	AuthIssuerURL           = "auth-issuer-url"
	AuthAudience            = "auth-audience"
	LogAnalyticsWorkspaceId = "workspace-id"
	RadixAPIHost            = "radix-api-host"
	RadixAPIPath            = "radix-api-path"
	RadixAPIScheme          = "radix-api-scheme"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    HostName,
		Usage:   "Host name",
		Value:   "",
		EnvVars: []string{"RADIX_LOG_API_HOST"},
	},
	&cli.IntFlag{
		Name:    PortNumber,
		Usage:   "Port number",
		Value:   8080,
		EnvVars: []string{"RADIX_LOG_API_PORT"},
	},
	&cli.StringFlag{
		Name:     AuthIssuerURL,
		Usage:    "OIDC issuer URL",
		Required: true,
		EnvVars:  []string{"RADIX_LOG_API_ISSUER"},
	},
	&cli.StringFlag{
		Name:     AuthAudience,
		Usage:    "Audience",
		Required: true,
		EnvVars:  []string{"RADIX_LOG_API_AUDIENCE"},
	},
	&cli.StringFlag{
		Name:     LogAnalyticsWorkspaceId,
		Usage:    "Log Analytics workspace ID",
		Required: true,
		EnvVars:  []string{"RADIX_LOG_API_WORKSPACE_ID"},
	},
	&cli.StringFlag{
		Name:     RadixAPIHost,
		Usage:    "Radix API host name",
		Value:    "",
		Required: true,
		EnvVars:  []string{"RADIX_LOG_API_RADIX_API_HOST"},
	},
	&cli.StringFlag{
		Name:    RadixAPIPath,
		Usage:   "Radix API base path",
		Value:   "/api/v1",
		EnvVars: []string{"RADIX_LOG_API_RADIX_API_PATH"},
	},
	&cli.StringFlag{
		Name:    RadixAPIScheme,
		Usage:   "Radix API sceme",
		Value:   "https",
		EnvVars: []string{"RADIX_LOG_API_RADIX_API_SCHEME"},
	},
}
