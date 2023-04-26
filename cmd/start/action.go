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
	"github.com/equinor/radix-log-api/pkg/constants"
	"github.com/equinor/radix-log-api/pkg/jwt"
	"github.com/equinor/radix-log-api/server"
	logservice "github.com/equinor/radix-log-api/services/logs"
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
	authz := buildAuthorization()
	return router.New(controllers, authn, authz)
}

func buildAuthorization() authz.Authorizer {
	auth := authz.NewAuthorizer(func(ab authz.AuthorizationBuilder) {
		ab.AddPolicy(constants.AuthorizationPolicyAuthenticated, func(pb authz.PolicyBuilder) {
			pb.RequireAuthenticatedUser()
		})
		ab.AddPolicy(constants.AuthorizationPolicyAppAdmin, func(pb authz.PolicyBuilder) {
			pb.RequireAuthenticatedUser()
			pb.AddRequirement(authz.RequirementFunc(func(ctx *authz.AuthorizationContext) error {
				// fmt.Println(ctx.User().Token())
				// return apierrors.NewForbiddenError()
				return nil
			}))
		})
	})
	return auth
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
