package logs

import (
	"io"
	"net/http"

	"github.com/equinor/radix-common/utils/slice"
	"github.com/equinor/radix-log-api/api/controllers"
	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/models"
	"github.com/equinor/radix-log-api/api/params"
	"github.com/equinor/radix-log-api/pkg/constants"
	logservice "github.com/equinor/radix-log-api/pkg/services/logs"
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
			Path:                  "/applications/:appName/environments/:envName/components/:componentName/log",
			Handler:               c.GetComponentLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/environments/:envName/components/:componentName/replicas/:replicaName/log",
			Handler:               c.GetComponentReplicaLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/environments/:envName/components/:componentName/replicas/:replicaName/containers/:containerId/log",
			Handler:               c.GetComponentContainerLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/environments/:envName/jobcomponents/:jobComponentName/jobs/:jobName",
			Handler:               c.GetJobInventory,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/environments/:envName/jobcomponents/:jobComponentName/jobs/:jobName/log",
			Handler:               c.GetJobLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/environments/:envName/jobcomponents/:jobComponentName/jobs/:jobName/replicas/:replicaName/log",
			Handler:               c.GetJobReplicaLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/environments/:envName/jobcomponents/:jobComponentName/jobs/:jobName/replicas/:replicaName/containers/:containerId/log",
			Handler:               c.GetJobContainerLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/pipelinejobs/:pipelineJobName",
			Handler:               c.GetPipelineJobInventory,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
		{
			Method:                http.MethodGet,
			Path:                  "/applications/:appName/pipelinejobs/:pipelineJobName/replicas/:replicaName/containers/:containerId/log",
			Handler:               c.GetPipelineJobContainerLog,
			AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
		},
	}
}

// GetComponentInventory godoc
// @Id getComponentInventory
// @Summary Get inventory (pods and containers) for a component
// @Tags Inventory
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.InventoryResponse
// @Failure 400 {object} errors.Status
// @Failure 401 {object} errors.Status
// @Failure 403 {object} errors.Status
// @Failure 500 {object} errors.Status
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
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.handleInventoryRequest(ctx, func(options *logservice.InventoryOptions) ([]logservice.Pod, error) {
		return c.appLogsService.ComponentInventory(ctx.Request.Context(), params.AppName, params.EnvName, params.ComponentName, options)
	})
}

// GetComponentLog godoc
// @Id getComponentLog
// @Summary Get log for a component
// @Tags Logs
// @Produce plain
// @Security ApiKeyAuth
// @Success 200 {string} string
// @Failure 400 {object} errors.Status
// @Failure 401 {object} errors.Status
// @Failure 403 {object} errors.Status
// @Failure 500 {object} errors.Status
// @Param appName path string true "Application Name"
// @Param envName path string true "Environment Name"
// @Param componentName path string true "Component Name"
// @Param tail query integer false "Number of rows to return from the tail of the log" example(100)
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Param file query boolean false "Response as attachment"
// @Router /applications/{appName}/environments/{envName}/components/{componentName}/log [get]
func (c *controller) GetComponentLog(ctx *gin.Context) {
	var params struct {
		params.App
		params.Env
		params.Component
	}
	if err := ctx.BindUri(&params); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.handleLogRequest(ctx, func(options *logservice.LogOptions) (io.Reader, error) {
		return c.appLogsService.ComponentLog(ctx.Request.Context(), params.AppName, params.EnvName, params.ComponentName, options)
	})
}

// GetComponentReplicaLog godoc
// @Id getComponentReplicaLog
// @Summary Get log for a component replica
// @Tags Logs
// @Produce plain
// @Security ApiKeyAuth
// @Success 200 {string} string
// @Failure 400 {object} errors.Status
// @Failure 401 {object} errors.Status
// @Failure 403 {object} errors.Status
// @Failure 500 {object} errors.Status
// @Param appName path string true "Application Name"
// @Param envName path string true "Environment Name"
// @Param componentName path string true "Component Name"
// @Param replicaName path string true "Replica Name"
// @Param tail query integer false "Number of rows to return from the tail of the log" example(100)
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Param file query boolean false "Response as attachment"
// @Router /applications/{appName}/environments/{envName}/components/{componentName}/replicas/{replicaName}/log [get]
func (c *controller) GetComponentReplicaLog(ctx *gin.Context) {
	var params struct {
		params.App
		params.Env
		params.Component
		params.Replica
	}
	if err := ctx.BindUri(&params); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.handleLogRequest(ctx, func(options *logservice.LogOptions) (io.Reader, error) {
		return c.appLogsService.ComponentPodLog(ctx.Request.Context(), params.AppName, params.EnvName, params.ComponentName, params.ReplicaName, options)
	})
}

// GetComponentContainerLog godoc
// @Id getComponentContainerLog
// @Summary Get log for a component container
// @Tags Logs
// @Produce plain
// @Security ApiKeyAuth
// @Success 200 {string} string
// @Failure 400 {object} errors.Status
// @Failure 401 {object} errors.Status
// @Failure 403 {object} errors.Status
// @Failure 500 {object} errors.Status
// @Param appName path string true "Application Name"
// @Param envName path string true "Environment Name"
// @Param componentName path string true "Component Name"
// @Param replicaName path string true "Replica Name"
// @Param containerId path string true "Container ID"
// @Param tail query integer false "Number of rows to return from the tail of the log" example(100)
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Param file query boolean false "Response as attachment"
// @Router /applications/{appName}/environments/{envName}/components/{componentName}/replicas/{replicaName}/containers/{containerId}/log [get]
func (c *controller) GetComponentContainerLog(ctx *gin.Context) {
	var params struct {
		params.App
		params.Env
		params.Component
		params.Replica
		params.Container
	}
	if err := ctx.BindUri(&params); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.handleLogRequest(ctx, func(options *logservice.LogOptions) (io.Reader, error) {
		return c.appLogsService.ComponentContainerLog(ctx.Request.Context(), params.AppName, params.EnvName, params.ComponentName, params.ReplicaName, params.ContainerId, options)
	})
}

// GetJobInventory godoc
// @Id getJobInventory
// @Summary Get inventory (pods and containers) for a job
// @Tags Inventory
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.InventoryResponse
// @Failure 400 {object} errors.Status
// @Failure 401 {object} errors.Status
// @Failure 403 {object} errors.Status
// @Failure 500 {object} errors.Status
// @Param appName path string true "Application Name"
// @Param envName path string true "Environment Name"
// @Param jobComponentName path string true "Job Component Name"
// @Param jobName path string true "Job Name"
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Router /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName} [get]
func (c *controller) GetJobInventory(ctx *gin.Context) {
	var params struct {
		params.App
		params.Env
		params.JobComponent
		params.Job
	}
	if err := ctx.BindUri(&params); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.handleInventoryRequest(ctx, func(options *logservice.InventoryOptions) ([]logservice.Pod, error) {
		return c.appLogsService.JobInventory(ctx.Request.Context(), params.AppName, params.EnvName, params.JobComponentName, params.JobName, options)
	})
}

// GetJobLog godoc
// @Id getJobLog
// @Summary Get log for a job
// @Tags Logs
// @Produce plain
// @Security ApiKeyAuth
// @Success 200 {string} string
// @Failure 400 {object} errors.Status
// @Failure 401 {object} errors.Status
// @Failure 403 {object} errors.Status
// @Failure 500 {object} errors.Status
// @Param appName path string true "Application Name"
// @Param envName path string true "Environment Name"
// @Param jobComponentName path string true "Job Component Name"
// @Param jobName path string true "Job Name"
// @Param tail query integer false "Number of rows to return from the tail of the log" example(100)
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Param file query boolean false "Response as attachment"
// @Router /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/log [get]
func (c *controller) GetJobLog(ctx *gin.Context) {
	var params struct {
		params.App
		params.Env
		params.JobComponent
		params.Job
	}
	if err := ctx.BindUri(&params); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.handleLogRequest(ctx, func(options *logservice.LogOptions) (io.Reader, error) {
		return c.appLogsService.JobLog(ctx.Request.Context(), params.AppName, params.EnvName, params.JobComponentName, params.JobName, options)
	})
}

// GetJobReplicaLog godoc
// @Id getJobReplicaLog
// @Summary Get log for a job replica
// @Tags Logs
// @Produce plain
// @Security ApiKeyAuth
// @Success 200 {string} string
// @Failure 400 {object} errors.Status
// @Failure 401 {object} errors.Status
// @Failure 403 {object} errors.Status
// @Failure 500 {object} errors.Status
// @Param appName path string true "Application Name"
// @Param envName path string true "Environment Name"
// @Param jobComponentName path string true "Job Component Name"
// @Param jobName path string true "Job Name"
// @Param replicaName path string true "Replica Name"
// @Param tail query integer false "Number of rows to return from the tail of the log" example(100)
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Param file query boolean false "Response as attachment"
// @Router /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/replicas/{replicaName}/log [get]
func (c *controller) GetJobReplicaLog(ctx *gin.Context) {
	var params struct {
		params.App
		params.Env
		params.JobComponent
		params.Job
		params.Replica
	}
	if err := ctx.BindUri(&params); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.handleLogRequest(ctx, func(options *logservice.LogOptions) (io.Reader, error) {
		return c.appLogsService.JobPodLog(ctx.Request.Context(), params.AppName, params.EnvName, params.JobComponentName, params.JobName, params.ReplicaName, options)
	})
}

// GetJobContainerLog godoc
// @Id getJobContainerLog
// @Summary Get log for a job container
// @Tags Logs
// @Produce plain
// @Security ApiKeyAuth
// @Success 200 {string} string
// @Failure 400 {object} errors.Status
// @Failure 401 {object} errors.Status
// @Failure 403 {object} errors.Status
// @Failure 500 {object} errors.Status
// @Param appName path string true "Application Name"
// @Param envName path string true "Environment Name"
// @Param jobComponentName path string true "Job Component Name"
// @Param jobName path string true "Job Name"
// @Param replicaName path string true "Replica Name"
// @Param containerId path string true "Container ID"
// @Param tail query integer false "Number of rows to return from the tail of the log" example(100)
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Param file query boolean false "Response as attachment"
// @Router /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/replicas/{replicaName}/containers/{containerId}/log [get]
func (c *controller) GetJobContainerLog(ctx *gin.Context) {
	var params struct {
		params.App
		params.Env
		params.JobComponent
		params.Job
		params.Replica
		params.Container
	}
	if err := ctx.BindUri(&params); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.handleLogRequest(ctx, func(options *logservice.LogOptions) (io.Reader, error) {
		return c.appLogsService.JobContainerLog(ctx.Request.Context(), params.AppName, params.EnvName, params.JobComponentName, params.JobName, params.ReplicaName, params.ContainerId, options)
	})
}

// GetPipelineJobInventory godoc
// @Id getPipelineJobInventory
// @Summary Get inventory (pods and containers) for a pipeline job
// @Tags Inventory
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.InventoryResponse
// @Failure 400 {object} errors.Status
// @Failure 401 {object} errors.Status
// @Failure 403 {object} errors.Status
// @Failure 500 {object} errors.Status
// @Param appName path string true "Application Name"
// @Param pipelineJobName path string true "Pipeline Job Name"
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Router /applications/{appName}/pipelinejobs/{pipelineJobName} [get]
func (c *controller) GetPipelineJobInventory(ctx *gin.Context) {
	var params struct {
		params.App
		params.PipelineJob
	}
	if err := ctx.BindUri(&params); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.handleInventoryRequest(ctx, func(options *logservice.InventoryOptions) ([]logservice.Pod, error) {
		return c.appLogsService.PipelineJobInventory(ctx.Request.Context(), params.AppName, params.PipelineJobName, options)
	})
}

// GetPipelineJobContainerLog godoc
// @Id getPipelineJobContainerLog
// @Summary Get log for a pipeline job container
// @Tags Logs
// @Produce plain
// @Security ApiKeyAuth
// @Success 200 {string} string
// @Failure 400 {object} errors.Status
// @Failure 401 {object} errors.Status
// @Failure 403 {object} errors.Status
// @Failure 500 {object} errors.Status
// @Param appName path string true "Application Name"
// @Param pipelineJobName path string true "Pipeline Job Name"
// @Param replicaName path string true "Replica Name"
// @Param containerId path string true "Container ID"
// @Param tail query integer false "Number of rows to return from the tail of the log" example(100)
// @Param start query string false "Start time" format(date-time) example(2023-05-01T08:15:00+02:00)
// @Param end query string false "End time" format(date-time) example(2023-05-02T12:00:00Z)
// @Param file query boolean false "Response as attachment"
// @Router /applications/{appName}/pipelinejobs/{pipelineJobName}/replicas/{replicaName}/containers/{containerId}/log [get]
func (c *controller) GetPipelineJobContainerLog(ctx *gin.Context) {
	var params struct {
		params.App
		params.PipelineJob
		params.Replica
		params.Container
	}
	if err := ctx.BindUri(&params); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.handleLogRequest(ctx, func(options *logservice.LogOptions) (io.Reader, error) {
		return c.appLogsService.PipelineJobContainerLog(ctx.Request.Context(), params.AppName, params.PipelineJobName, params.ReplicaName, params.ContainerId, options)
	})
}

func (c *controller) handleInventoryRequest(ctx *gin.Context, inventorySource func(options *logservice.InventoryOptions) ([]logservice.Pod, error)) {
	queryParams, err := paramsFromContext[inventoryParams](ctx)
	if err != nil {
		_ = ctx.Error(apierrors.NewBadRequestError(apierrors.WithCause(err)))
		ctx.Abort()
		return
	}

	options := queryParams.AsInventoryOptions()
	pods, err := inventorySource(&options)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}

	response := inventoryRepsonseFromPods(pods)
	ctx.JSON(http.StatusOK, response)
}

func (c *controller) handleLogRequest(ctx *gin.Context, logSource func(options *logservice.LogOptions) (io.Reader, error)) {
	queryParams, err := paramsFromContext[logParams](ctx)
	if err != nil {
		_ = ctx.Error(apierrors.NewBadRequestError(apierrors.WithCause(err)))
		ctx.Abort()
		return
	}

	logOptions := queryParams.AsLogOptions()
	logReader, err := logSource(&logOptions)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}

	extraHeaders := make(map[string]string)
	if queryParams.File {
		extraHeaders["Content-Disposition"] = `attachment; filename="log.txt"`
	}

	ctx.DataFromReader(200, -1, "text/plain; charset=utf-8", logReader, extraHeaders)
}

func inventoryRepsonseFromPods(pods []logservice.Pod) models.InventoryResponse {
	return models.InventoryResponse{
		Replicas: slice.Map(pods, func(s logservice.Pod) models.Replica {
			return models.Replica{
				Name:              s.Name,
				CreationTimestamp: s.CreationTimestamp,
				LastKnown:         s.LastKnown,
				Containers:        slice.Map(s.Containers, func(c logservice.Container) models.Container { return models.Container(c) })}
		}),
	}
}
