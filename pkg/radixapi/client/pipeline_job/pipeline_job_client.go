// Code generated by go-swagger; DO NOT EDIT.

package pipeline_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pipeline job API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pipeline job API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetApplicationJob(params *GetApplicationJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationJobOK, error)

	GetApplicationJobs(params *GetApplicationJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationJobsOK, error)

	GetPipelineJobStepLogs(params *GetPipelineJobStepLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPipelineJobStepLogsOK, error)

	GetTektonPipelineRun(params *GetTektonPipelineRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunOK, error)

	GetTektonPipelineRunTask(params *GetTektonPipelineRunTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunTaskOK, error)

	GetTektonPipelineRunTaskStepLogs(params *GetTektonPipelineRunTaskStepLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunTaskStepLogsOK, error)

	GetTektonPipelineRunTaskSteps(params *GetTektonPipelineRunTaskStepsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunTaskStepsOK, error)

	GetTektonPipelineRunTasks(params *GetTektonPipelineRunTasksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunTasksOK, error)

	GetTektonPipelineRuns(params *GetTektonPipelineRunsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunsOK, error)

	StopApplicationJob(params *StopApplicationJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopApplicationJobNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetApplicationJob gets the detail of a given pipeline job for a given application
*/
func (a *Client) GetApplicationJob(params *GetApplicationJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApplicationJob",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/jobs/{jobName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetApplicationJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApplicationJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetApplicationJobs gets the summary of jobs for a given application
*/
func (a *Client) GetApplicationJobs(params *GetApplicationJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApplicationJobs",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetApplicationJobsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApplicationJobs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPipelineJobStepLogs gets logs of a pipeline job step
*/
func (a *Client) GetPipelineJobStepLogs(params *GetPipelineJobStepLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPipelineJobStepLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelineJobStepLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPipelineJobStepLogs",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/jobs/{jobName}/logs/{stepName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPipelineJobStepLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPipelineJobStepLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPipelineJobStepLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTektonPipelineRun gets a pipeline run for a pipeline job
*/
func (a *Client) GetTektonPipelineRun(params *GetTektonPipelineRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTektonPipelineRunParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTektonPipelineRun",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTektonPipelineRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTektonPipelineRunOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTektonPipelineRun: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTektonPipelineRunTask gets list of pipeline run task of a pipeline job
*/
func (a *Client) GetTektonPipelineRunTask(params *GetTektonPipelineRunTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTektonPipelineRunTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTektonPipelineRunTask",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTektonPipelineRunTaskReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTektonPipelineRunTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTektonPipelineRunTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTektonPipelineRunTaskStepLogs gets logs of pipeline runs for a pipeline job
*/
func (a *Client) GetTektonPipelineRunTaskStepLogs(params *GetTektonPipelineRunTaskStepLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunTaskStepLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTektonPipelineRunTaskStepLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTektonPipelineRunTaskStepLogs",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}/logs/{stepName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTektonPipelineRunTaskStepLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTektonPipelineRunTaskStepLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTektonPipelineRunTaskStepLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTektonPipelineRunTaskSteps gets list of steps for a pipeline run task of a pipeline job
*/
func (a *Client) GetTektonPipelineRunTaskSteps(params *GetTektonPipelineRunTaskStepsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunTaskStepsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTektonPipelineRunTaskStepsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTektonPipelineRunTaskSteps",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}/steps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTektonPipelineRunTaskStepsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTektonPipelineRunTaskStepsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTektonPipelineRunTaskSteps: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTektonPipelineRunTasks gets list of pipeline run tasks of a pipeline job
*/
func (a *Client) GetTektonPipelineRunTasks(params *GetTektonPipelineRunTasksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTektonPipelineRunTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTektonPipelineRunTasks",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTektonPipelineRunTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTektonPipelineRunTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTektonPipelineRunTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTektonPipelineRuns gets list of pipeline runs for a pipeline job
*/
func (a *Client) GetTektonPipelineRuns(params *GetTektonPipelineRunsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTektonPipelineRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTektonPipelineRunsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTektonPipelineRuns",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/jobs/{jobName}/pipelineruns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTektonPipelineRunsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTektonPipelineRunsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTektonPipelineRuns: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopApplicationJob stops job
*/
func (a *Client) StopApplicationJob(params *StopApplicationJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopApplicationJobNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopApplicationJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopApplicationJob",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/jobs/{jobName}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopApplicationJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopApplicationJobNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopApplicationJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
