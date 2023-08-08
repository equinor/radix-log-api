package start

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"github.com/equinor/radix-log-api/api/controllers"
	logscontroller "github.com/equinor/radix-log-api/api/controllers/logs"
	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/equinor/radix-log-api/api/router"
	"github.com/equinor/radix-log-api/api/server"
	"github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	logservice "github.com/equinor/radix-log-api/pkg/services/logs"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/urfave/cli/v2"
)

func action(cliCtx *cli.Context) error {
	router, err := buildRouter(cliCtx)
	if err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(cliCtx.Context, os.Interrupt, os.Kill)
	defer cancel()
	return server.Run(ctx, router, buildServerConfig(cliCtx))
}

func buildRouter(cliCtx *cli.Context) (http.Handler, error) {
	jwtValidator, err := buildJwtValidator(cliCtx)
	if err != nil {
		return nil, err
	}
	applicationClient, err := buildApplicationClient(cliCtx)
	if err != nil {
		return nil, err
	}
	controllers, err := buildControllers(cliCtx)
	if err != nil {
		return nil, err
	}
	return router.New(jwtValidator, applicationClient, controllers...)
}

func buildControllers(cliCtx *cli.Context) ([]controllers.Controller, error) {
	logService, err := buildLogService(cliCtx)
	if err != nil {
		return nil, err
	}
	contollers := []controllers.Controller{
		logscontroller.New(logService),
	}
	return contollers, nil
}

func buildLogService(cliCtx *cli.Context) (logservice.Service, error) {
	logsClient, err := buildLogsAnalyticsClient()
	if err != nil {
		return nil, err
	}
	logService := logservice.New(logsClient, cliCtx.String(LogAnalyticsWorkspaceId))
	return logService, nil
}

func buildJwtValidator(cliCtx *cli.Context) (authn.JwtValidator, error) {
	return authn.NewValidator(cliCtx.String(AuthIssuer), cliCtx.String(AuthAudience))
}

func buildApplicationClient(cliCtx *cli.Context) (application.ClientService, error) {
	radixApiHost := cliCtx.String(RadixAPIHost)
	if len(radixApiHost) == 0 {
		return nil, fmt.Errorf("required argument %s is not set", RadixAPIHost)
	}
	client := application.New(runtimeclient.New(radixApiHost, cliCtx.String(RadixAPIPath), []string{cliCtx.String(RadixAPIScheme)}), strfmt.Default)
	return client, nil
}

func buildLogsAnalyticsClient() (*azquery.LogsClient, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	return azquery.NewLogsClient(cred, nil)
}

func buildServerConfig(cliCtx *cli.Context) *server.Config {
	cfg := &server.Config{
		Host: cliCtx.String(HostName),
		Port: cliCtx.Int(PortNumber),
	}
	return cfg
}
