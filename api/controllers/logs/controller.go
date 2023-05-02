package logs

import (
	"net/http"

	"github.com/equinor/radix-common/utils/slice"
	"github.com/equinor/radix-log-api/api/controllers"
	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/models"
	"github.com/equinor/radix-log-api/api/params"
	"github.com/equinor/radix-log-api/pkg/constants"
	logservice "github.com/equinor/radix-log-api/services/logs"
	"github.com/gin-gonic/gin"
)

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
			Handler:               c.GetComponentInventory,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/environments/:envName/components/:componentName/logs",
			Handler:               c.GetComponentLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/environments/:envName/components/:componentName/replicas/:replicaName/logs",
			Handler:               c.GetComponentReplicaLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/environments/:envName/components/:componentName/replicas/:replicaName/containers/:containerId/logs",
			Handler:               c.GetComponentContainerLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
	}
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
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Router /applications/{appName}/environments/{envName}/components/{componentName} [get]
func (c *controller) GetComponentInventory(ctx *gin.Context) {
	var params struct {
		params.App
		params.Env
		params.Component
	}
	if err := ctx.BindUri(&params); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	queryParams, err := paramsFromContext[inventoryParams](ctx)
	if err != nil {
		ctx.Error(apierrors.NewBadRequestError(apierrors.WithCause(err)))
		ctx.Abort()
		return
	}

	pods, err := c.appLogsService.ComponentInventory(params.AppName, params.EnvName, params.ComponentName, queryParams.AsComponentPodInventoryOptions())
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	response := models.ComponentInventoryResponse{
		Replicas: slice.Map(pods, func(s logservice.Pod) models.Replica {
			return models.Replica{
				Name:              s.Name,
				CreationTimestamp: s.CreationTimestamp,
				Containers:        slice.Map(s.Containers, func(c logservice.Container) models.Container { return models.Container(c) })}
		}),
	}

	ctx.JSON(http.StatusOK, response)
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
// @Param rows query integer false "Number of rows to return in descending order by log time" example(100)
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Router /applications/{appName}/environments/{envName}/components/{componentName}/logs [get]
func (c *controller) GetComponentLog(ctx *gin.Context) {
	var uriParams struct {
		params.App
		params.Env
		params.Component
	}
	if err := ctx.BindUri(&uriParams); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	queryParams, err := paramsFromContext[logParams](ctx)
	if err != nil {
		ctx.Error(apierrors.NewBadRequestError(apierrors.WithCause(err)))
		ctx.Abort()
		return
	}

	logReader, err := c.appLogsService.ComponentLog(uriParams.AppName, uriParams.EnvName, uriParams.ComponentName, queryParams.AsLogOptions())
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.DataFromReader(200, -1, "text/plain; charset=utf-8", logReader, nil)
}

// GetComponentReplicaLog godoc
// @Summary Get log for a replica
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
// @Param replicaName path string true "Replica Name"
// @Param rows query integer false "Number of rows to return in descending order by log time" example(100)
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Router /applications/{appName}/environments/{envName}/components/{componentName}/replicas/{replicaName}/logs [get]
func (c *controller) GetComponentReplicaLog(ctx *gin.Context) {
	var params struct {
		params.App
		params.Env
		params.Component
		params.Replica
	}
	if err := ctx.BindUri(&params); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	queryParams, err := paramsFromContext[logParams](ctx)
	if err != nil {
		ctx.Error(apierrors.NewBadRequestError(apierrors.WithCause(err)))
		ctx.Abort()
		return
	}

	logReader, err := c.appLogsService.ComponentPodLog(params.AppName, params.EnvName, params.ComponentName, params.ReplicaName, queryParams.AsLogOptions())
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.DataFromReader(200, -1, "text/plain; charset=utf-8", logReader, nil)
}

// GetComponentContainerLog godoc
// @Summary Get log for a container
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
// @Param replicaName path string true "Replica Name"
// @Param containerId path string true "Container ID"
// @Param rows query integer false "Number of rows to return in descending order by log time" example(100)
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Router /applications/{appName}/environments/{envName}/components/{componentName}/replicas/{replicaName}/containers/{containerId}/logs [get]
func (c *controller) GetComponentContainerLog(ctx *gin.Context) {
	var params struct {
		params.App
		params.Env
		params.Component
		params.Replica
		params.Container
	}
	if err := ctx.BindUri(&params); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	queryParams, err := paramsFromContext[logParams](ctx)
	if err != nil {
		ctx.Error(apierrors.NewBadRequestError(apierrors.WithCause(err)))
		ctx.Abort()
		return
	}

	logReader, err := c.appLogsService.ComponentContainerLog(params.AppName, params.EnvName, params.ComponentName, params.ReplicaName, params.ContainerId, queryParams.AsLogOptions())
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.DataFromReader(200, -1, "text/plain; charset=utf-8", logReader, nil)
}
