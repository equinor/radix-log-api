package router

import (
	"net/http"

	"github.com/equinor/radix-log-api/api/controllers"
	logscontroller "github.com/equinor/radix-log-api/api/controllers/logs"
	authnmiddleware "github.com/equinor/radix-log-api/api/middleware/authn"
	authzmiddleware "github.com/equinor/radix-log-api/api/middleware/authz"
	errmiddleware "github.com/equinor/radix-log-api/api/middleware/error"
	"github.com/equinor/radix-log-api/pkg/authz/requirement"
	"github.com/equinor/radix-log-api/pkg/constants"
	applicationclient "github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	logservice "github.com/equinor/radix-log-api/services/logs"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(logService logservice.Service, jwtValidator authnmiddleware.JwtValidator, applicationClient applicationclient.ClientService) (http.Handler, error) {
	authz, err := buildAuthorizer(applicationClient)
	if err != nil {
		return nil, err
	}
	authn := authnmiddleware.New(authnmiddleware.NewJwt(jwtValidator))
	controllers := buildControllers(logService)

	engine := gin.New()
	engine.RemoveExtraSlash = true
	engine.Use(gin.Logger(), gzip.Gzip(gzip.DefaultCompression), gin.Recovery())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	g := engine.Group("/api/v1", cors.Default(), errmiddleware.ErrorHandler, authn)
	{
		mapControllers(controllers, g, authz)
	}

	return engine, nil
}

func mapControllers(controllers []controllers.Controller, router gin.IRoutes, authz authzmiddleware.Authorizer) {
	for _, controller := range controllers {
		for _, endpoint := range controller.Endpoints() {
			router.Handle(endpoint.Method, endpoint.Path, authz.Authorize(endpoint.AuthorizationPolicies...), endpoint.Handler)
		}
	}
}

func buildAuthorizer(applicationClient applicationclient.ClientService) (authzmiddleware.Authorizer, error) {
	appOwnerRequirement := requirement.NewAppOwnerRequirement(applicationClient)
	auth := authzmiddleware.NewAuthorizer(func(ab authzmiddleware.AuthorizationBuilder) {
		ab.AddPolicy(constants.AuthorizationPolicyAppAdmin, func(pb authzmiddleware.PolicyBuilder) {
			pb.RequireAuthenticatedUser().AddRequirement(appOwnerRequirement)
		})
	})
	return auth, nil
}

func buildControllers(logService logservice.Service) []controllers.Controller {
	return []controllers.Controller{
		logscontroller.New(logService),
	}
}
