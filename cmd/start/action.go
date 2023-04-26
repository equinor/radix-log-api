package start

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"github.com/equinor/radix-log-api/api/controllers"
	logscontroller "github.com/equinor/radix-log-api/api/controllers/logs"
	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/equinor/radix-log-api/api/middleware/authz"
	"github.com/equinor/radix-log-api/api/router"
	"github.com/equinor/radix-log-api/pkg/authz/requirement"
	"github.com/equinor/radix-log-api/pkg/constants"
	"github.com/equinor/radix-log-api/pkg/jwt"
	"github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	"github.com/equinor/radix-log-api/server"
	logservice "github.com/equinor/radix-log-api/services/logs"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/urfave/cli/v2"
)

func action(ctx *cli.Context) error {
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
		logscontroller.New(logservice.New(logsClient, ctx.String(LogAnalyticsWorkspaceId))),
	}
	jwtValidator, err := jwt.NewValidator(ctx.String(AuthIssuerURL), ctx.String(AuthAudience))
	if err != nil {
		return nil, err
	}
	authn := []authn.AuthenticationProvider{
		authn.NewJwt(jwtValidator),
	}
	authz := buildAuthorization(ctx)
	return router.New(controllers, authn, authz)
}

func buildAuthorization(ctx *cli.Context) authz.Authorizer {
	client := buildRadixAPIApplicationClient(ctx.String(RadixAPIHost), ctx.String(RadixAPIPath), ctx.String(RadixAPIScheme))
	appOwnerRequirement := requirement.NewAppOwnerRequirement(client)
	auth := authz.NewAuthorizer(func(ab authz.AuthorizationBuilder) {
		ab.AddPolicy(constants.AuthorizationPolicyAppAdmin, func(pb authz.PolicyBuilder) {
			pb.RequireAuthenticatedUser().AddRequirement(appOwnerRequirement)
		})
	})
	return auth
}

func buildRadixAPIApplicationClient(host, path, scheme string) application.ClientService {
	return application.New(runtimeclient.New(host, path, []string{scheme}), strfmt.Default)
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
