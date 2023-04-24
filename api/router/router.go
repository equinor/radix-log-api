package router

import (
	"net/http"

	"github.com/equinor/radix-log-api/api/controllers"
	authnmiddleware "github.com/equinor/radix-log-api/api/middleware/authn"
	authzmiddleware "github.com/equinor/radix-log-api/api/middleware/authz"
	errmiddleware "github.com/equinor/radix-log-api/api/middleware/error"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	LogLevel string
}

func New(controllers []controllers.Controller, authn []authnmiddleware.AuthenticationProvider, authz authzmiddleware.Authorizer) (http.Handler, error) {
	engine := gin.New()
	engine.RemoveExtraSlash = true

	engine.Use(gin.Logger(), gin.Recovery(), gzip.Gzip(gzip.DefaultCompression))

	g := engine.Group("/api/v1", cors.Default(), errmiddleware.ErrorHandler, authnmiddleware.New(authn...))
	{
		mapControllers(controllers, g, authz)
	}
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return engine, nil
}

func mapControllers(controllers []controllers.Controller, router gin.IRoutes, authz authzmiddleware.Authorizer) {
	for _, controller := range controllers {
		for _, endpoint := range controller.Endpoints() {
			router.Handle(endpoint.Method, endpoint.Path, authz.Authorize(endpoint.AuthorizationPolicies...), endpoint.Handler)
		}
	}
}
