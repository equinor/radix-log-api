package router

import (
	"net/http"

	commongin "github.com/equinor/radix-common/pkg/gin"
	"github.com/equinor/radix-log-api/api/controllers"
	authnmiddleware "github.com/equinor/radix-log-api/api/middleware/authn"
	authzmiddleware "github.com/equinor/radix-log-api/api/middleware/authz"
	errmiddleware "github.com/equinor/radix-log-api/api/middleware/error"
	"github.com/equinor/radix-log-api/pkg/authz/requirement"
	"github.com/equinor/radix-log-api/pkg/constants"
	applicationclient "github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(jwtValidator authnmiddleware.JwtValidator, applicationClient applicationclient.ClientService, controllers ...controllers.Controller) (http.Handler, error) {
	engine := gin.New()
	engine.RemoveExtraSlash = true
	engine.Use(commongin.SetZerologLogger(commongin.ZerologLoggerWithRequestId))
	engine.Use(commongin.ZerologRequestLogger(), gzip.Gzip(gzip.DefaultCompression), gin.Recovery())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	authz, err := buildAuthorizer(applicationClient)
	if err != nil {
		return nil, err
	}
	authn := authnmiddleware.New(authnmiddleware.NewJwtProvider(jwtValidator))
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
	auth := authzmiddleware.NewAuthorizer(func(ab authzmiddleware.AuthorizationConfiguration) {
		ab.AddPolicy(constants.AuthorizationPolicyAppAdmin, func(pb authzmiddleware.PolicyConfiguration) {
			pb.RequireAuthenticatedUser().AddRequirement(appOwnerRequirement)
		})
	})
	return auth, nil
}
