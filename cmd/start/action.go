package start

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"github.com/equinor/radix-log-api/api/controllers"
	logscontroller "github.com/equinor/radix-log-api/api/controllers/logs"
	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/equinor/radix-log-api/api/router"
	"github.com/equinor/radix-log-api/api/server"
	"github.com/equinor/radix-log-api/pkg/jwt"
	"github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	logservice "github.com/equinor/radix-log-api/services/logs"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/urfave/cli/v2"
)

func action(ctx *cli.Context) error {
	router, err := buildRouter(ctx)
	if err != nil {
		return err
	}
	return server.Run(context.TODO(), router, buildServerConfig(ctx))
}

func buildRouter(ctx *cli.Context) (http.Handler, error) {

	jwtValidator, err := buildJwtValidator(ctx)
	if err != nil {
		return nil, err
	}
	applicationClient, err := buildApplicationClient(ctx)
	if err != nil {
		return nil, err
	}
	controllers, err := buildControllers(ctx)
	if err != nil {
		return nil, err
	}
	return router.New(jwtValidator, applicationClient, controllers...)
}

func buildControllers(ctx *cli.Context) ([]controllers.Controller, error) {
	logService, err := buildLogService(ctx)
	if err != nil {
		return nil, err
	}
	contollers := []controllers.Controller{
		logscontroller.New(logService),
	}
	return contollers, nil
}

func buildLogService(ctx *cli.Context) (logservice.Service, error) {
	logsClient, err := buildLogsAnalyticsClient()
	if err != nil {
		return nil, err
	}
	logService := logservice.New(logsClient, ctx.String(LogAnalyticsWorkspaceId))
	return logService, nil
}

func buildJwtValidator(ctx *cli.Context) (authn.JwtValidator, error) {
	return jwt.NewValidator(ctx.String(AuthIssuer), ctx.String(AuthAudience))
}

func buildApplicationClient(ctx *cli.Context) (application.ClientService, error) {
	radixApiHost := ctx.String(RadixAPIHost)
	if len(radixApiHost) == 0 {
		return nil, fmt.Errorf("required argument %s is not set", RadixAPIHost)
	}
	client := application.New(runtimeclient.New(radixApiHost, ctx.String(RadixAPIPath), []string{ctx.String(RadixAPIScheme)}), strfmt.Default)
	return client, nil
}

func buildLogsAnalyticsClient() (*azquery.LogsClient, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	return azquery.NewLogsClient(cred, nil)
}

func buildServerConfig(ctx *cli.Context) *server.Config {
	cfg := &server.Config{
		Host: ctx.String(HostName),
		Port: ctx.Int(PortNumber),
	}
	return cfg
}
