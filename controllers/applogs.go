package controllers

import (
	"net/http"

	"github.com/equinor/radix-log-api/internal/uriparams"
	"github.com/equinor/radix-log-api/services"
	"github.com/gin-gonic/gin"
)

type appURIParams struct {
	uriparams.App
	uriparams.Env
	uriparams.Component
}

func NewAppLogs(appLogsService services.AppLogs) Controller {
	return &appLogs{
		appLogsService: appLogsService,
	}
}

type appLogs struct {
	appLogsService services.AppLogs
}

func (c *appLogs) Endpoints() []Endpoint {
	return []Endpoint{
		{
			Method:  http.MethodGet,
			Path:    "/applications/:appName/environments/:envName/components/:componentName",
			Handler: c.GetComponentLog,
		},
	}
}

func (c *appLogs) GetComponentLog(ctx *gin.Context) {
	var params appURIParams
	if err := ctx.BindUri(&params); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	logReader, err := c.appLogsService.GetLogs(params.AppName, params.EnvName, params.ComponentName, nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.DataFromReader(200, -1, "text/plain; charset=utf-8", logReader, nil)
}
