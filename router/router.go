package router

import (
	"net/http"

	"github.com/equinor/radix-log-api/controllers"
	"github.com/equinor/radix-log-api/middleware"
	"github.com/equinor/radix-log-api/pkg/authn"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Config struct {
	LogLevel string
}

func New(controllers []controllers.Controller, authn []authn.Provider) (http.Handler, error) {
	engine := gin.New()
	engine.RemoveExtraSlash = true

	engine.Use(gin.Logger(), gin.Recovery(), gzip.Gzip(gzip.DefaultCompression))

	g := engine.Group("/api/v1", cors.Default(), middleware.Authentication(authn...))
	{
		mapControllers(controllers, g)
	}

	return engine, nil
}

func mapControllers(controllers []controllers.Controller, router gin.IRoutes) {
	for _, controller := range controllers {
		for _, endpoint := range controller.Endpoints() {
			router.Handle(endpoint.Method, endpoint.Path, endpoint.Handler)
		}
	}
}
