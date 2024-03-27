package start

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	HostName                = "host"
	PortNumber              = "port"
	AuthIssuer              = "auth-issuer"
	AuthAudience            = "auth-audience"
	LogAnalyticsWorkspaceId = "log-analytics-workspace-id"
	LogAnalyticsLogTable    = "log-analytics-log-table"
	RadixAPIHost            = "radix-api-host"
	RadixAPIPath            = "radix-api-path"
	RadixAPIScheme          = "radix-api-scheme"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    HostName,
		Usage:   "Host name/ip",
		Value:   "",
		EnvVars: []string{"LOG_API_HOST"},
	},
	&cli.IntFlag{
		Name:    PortNumber,
		Usage:   "Port number",
		Value:   8000,
		EnvVars: []string{"LOG_API_PORT"},
	},
	&cli.StringFlag{
		Name:     AuthIssuer,
		Usage:    "OIDC issuer",
		Required: true,
		EnvVars:  []string{"LOG_API_AUTH_ISSUER"},
	},
	&cli.StringFlag{
		Name:     AuthAudience,
		Usage:    "Audience",
		Required: true,
		EnvVars:  []string{"LOG_API_AUTH_AUDIENCE"},
	},
	&cli.StringFlag{
		Name:     LogAnalyticsLogTable,
		Usage:    "Log Analytics workspace table to query. ContainerLogV2, ContainerLog or Both. Defaults to ContainerLog",
		Required: true,
		EnvVars:  []string{"LOG_API_LOG_ANALYTICS_LOG_TABLE"},
	},
	&cli.StringFlag{
		Name:     LogAnalyticsWorkspaceId,
		Usage:    "Log Analytics workspace ID",
		Required: true,
		EnvVars:  []string{"LOG_API_LOG_ANALYTICS_WORKSPACE_ID"},
	},
	&cli.StringFlag{
		Name:     RadixAPIHost,
		Usage:    "Radix API host name",
		Value:    getRadixAPIHostFromEnv(),
		Required: false,
		EnvVars:  []string{"LOG_API_RADIX_API_HOST"},
	},
	&cli.StringFlag{
		Name:    RadixAPIPath,
		Usage:   "Radix API base path",
		Value:   "/api/v1",
		EnvVars: []string{"LOG_API_RADIX_API_PATH"},
	},
	&cli.StringFlag{
		Name:    RadixAPIScheme,
		Usage:   "Radix API scheme",
		Value:   "https",
		EnvVars: []string{"LOG_API_RADIX_API_SCHEME"},
	},
}

func getRadixAPIHostFromEnv() string {
	envName, clusterName, dnsZone := os.Getenv("RADIX_ENVIRONMENT"), os.Getenv("RADIX_CLUSTERNAME"), os.Getenv("RADIX_DNS_ZONE")
	if len(envName) == 0 || len(clusterName) == 0 || len(dnsZone) == 0 {
		return ""
	}
	return fmt.Sprintf("server-radix-api-%s.%s.%s", envName, clusterName, dnsZone)
}
