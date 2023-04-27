package logs

import (
	"net/http"

	"github.com/equinor/radix-common/utils/slice"
	"github.com/equinor/radix-log-api/api/controllers"
	"github.com/equinor/radix-log-api/api/models"
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
			Path:                  "/applications/:appName/environments/:envName/components/:componentName/logs",
			Handler:               c.GetComponentLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/environments/:envName/components/:componentName",
			Handler:               c.GetComponentInventory,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
	}
}

// GetComponentLog godoc
// @Summary Get log for a component
// @Tags Logs
// @Produce plain
// @Security ApiKeyAuth
// @Success 200 {string} ALogRecord
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Param appName path string true "Application Name"
// @Param envName path string true "Environment Name"
// @Param componentName path string true "Component Name"
// @Router /applications/{appName}/environments/{envName}/components/{componentName}/logs [get]
func (c *controller) GetComponentLog(ctx *gin.Context) {
	var params appURIParams
	if err := ctx.BindUri(&params); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	logReader, err := c.appLogsService.ComponentLog(params.AppName, params.EnvName, params.ComponentName, nil)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.DataFromReader(200, -1, "text/plain; charset=utf-8", logReader, nil)
}

// GetComponentInventory godoc
// @Summary Get inventory (pods and their containers) for a component
// @Tags Inventory
// @Produce json
// @Security ApiKeyAuth
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
func (c *controller) GetComponentInventory(ctx *gin.Context) {
	var params appURIParams
	if err := ctx.BindUri(&params); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	pods, err := c.appLogsService.ComponentPodInventory(params.AppName, params.EnvName, params.ComponentName, nil)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	response := models.ComponentInventoryResponse{
		Pods: slice.Map(pods, func(s logservice.Pod) models.Pod {
			return models.Pod{
				Name:              s.Name,
				CreationTimestamp: s.CreationTimestamp,
				Containers:        slice.Map(s.Containers, func(c logservice.Container) models.Container { return models.Container(c) })}
		}),
	}

	ctx.JSON(http.StatusOK, response)
}
