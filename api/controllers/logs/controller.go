package logs

import (
	"net/http"

	"github.com/equinor/radix-log-api/api/controllers"
	"github.com/equinor/radix-log-api/api/params"
	"github.com/equinor/radix-log-api/pkg/constants"
	logservice "github.com/equinor/radix-log-api/services/logs"
	"github.com/gin-gonic/gin"
)

type appURIParams struct {
	params.App
	params.Env
	params.Component
}

func New(appLogsService logservice.Service) controllers.Controller {
	return &controller{
		appLogsService: appLogsService,
	}
}

type controller struct {
	appLogsService logservice.Service
}

func (c *controller) Endpoints() []controllers.Endpoint {
	return []controllers.Endpoint{
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/environments/:envName/components/:componentName",
			Handler:               c.GetComponentLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
	}
}

// GetComponentLog godoc
// @Summary Get log for a component
// @Tags GetComponentLog
// @Produce plain
// @Success 200 {string} ALogRecord
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Param appName path string true "Application Name"
// @Param envName path string true "Environment Name"
// @Param componentName path string true "Component Name"
// @Router /applications/{appName}/environments/{envName}/components/{componentName} [get]
func (c *controller) GetComponentLog(ctx *gin.Context) {
	var params appURIParams
	if err := ctx.BindUri(&params); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	logReader, err := c.appLogsService.Component(params.AppName, params.EnvName, params.ComponentName, nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.DataFromReader(200, -1, "text/plain; charset=utf-8", logReader, nil)
}
