package commands

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"github.com/equinor/radix-log-api/controllers"
	"github.com/equinor/radix-log-api/pkg/authn"
	"github.com/equinor/radix-log-api/pkg/jwt"
	"github.com/equinor/radix-log-api/router"
	"github.com/equinor/radix-log-api/server"
	"github.com/equinor/radix-log-api/services"
	"github.com/urfave/cli/v2"
)

const (
	HostName                = "host"
	PortNumber              = "port"
	AuthIssuerURL           = "auth-issuer-url"
	AuthAudience            = "auth-audience"
	LogAnalyticsWorkspaceId = "workspace-id"
)

var startFlags = []cli.Flag{
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
}

var Start = cli.Command{
	Name:   "start",
	Usage:  "Starts the web server",
	Action: startAction,
	Flags:  startFlags,
}

func startAction(ctx *cli.Context) error {
	cfg := initConfig(ctx)
	router, err := initRouter(ctx)
	if err != nil {
		return err
	}
	return server.Run(context.TODO(), router, cfg)
}

func initRouter(ctx *cli.Context) (http.Handler, error) {
	logsClient, err := initLogsAnalyticsClient()
	if err != nil {
		return nil, err
	}
	controllers := []controllers.Controller{
		controllers.NewAppLogs(services.NewAppLogs(logsClient, ctx.String(LogAnalyticsWorkspaceId))),
	}
	jwtValidator, err := jwt.NewValidator(ctx.String(AuthIssuerURL), ctx.String(AuthAudience))
	if err != nil {
		return nil, err
	}
	authn := []authn.Provider{
		authn.NewJwt(jwtValidator),
	}
	return router.New(controllers, authn)
}

func initLogsAnalyticsClient() (*azquery.LogsClient, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	return azquery.NewLogsClient(cred, nil)
}

func initConfig(ctx *cli.Context) *server.Config {
	cfg := &server.Config{
		Host: ctx.String(HostName),
		Port: ctx.Int(PortNumber),
	}
	return cfg
}
