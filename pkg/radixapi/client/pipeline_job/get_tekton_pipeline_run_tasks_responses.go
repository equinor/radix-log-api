// Code generated by go-swagger; DO NOT EDIT.

package pipeline_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-log-api/pkg/radixapi/models"
)

// GetTektonPipelineRunTasksReader is a Reader for the GetTektonPipelineRunTasks structure.
type GetTektonPipelineRunTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTektonPipelineRunTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTektonPipelineRunTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTektonPipelineRunTasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTektonPipelineRunTasksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks] getTektonPipelineRunTasks", response, response.Code())
	}
}

// NewGetTektonPipelineRunTasksOK creates a GetTektonPipelineRunTasksOK with default headers values
func NewGetTektonPipelineRunTasksOK() *GetTektonPipelineRunTasksOK {
	return &GetTektonPipelineRunTasksOK{}
}

/*
GetTektonPipelineRunTasksOK describes a response with status code 200, with default header values.

List of Pipeline Run Tasks
*/
type GetTektonPipelineRunTasksOK struct {
	Payload []*models.PipelineRunTask
}

// IsSuccess returns true when this get tekton pipeline run tasks o k response has a 2xx status code
func (o *GetTektonPipelineRunTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tekton pipeline run tasks o k response has a 3xx status code
func (o *GetTektonPipelineRunTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline run tasks o k response has a 4xx status code
func (o *GetTektonPipelineRunTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tekton pipeline run tasks o k response has a 5xx status code
func (o *GetTektonPipelineRunTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline run tasks o k response a status code equal to that given
func (o *GetTektonPipelineRunTasksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tekton pipeline run tasks o k response
func (o *GetTektonPipelineRunTasksOK) Code() int {
	return 200
}

func (o *GetTektonPipelineRunTasksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks][%d] getTektonPipelineRunTasksOK %s", 200, payload)
}

func (o *GetTektonPipelineRunTasksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks][%d] getTektonPipelineRunTasksOK %s", 200, payload)
}

func (o *GetTektonPipelineRunTasksOK) GetPayload() []*models.PipelineRunTask {
	return o.Payload
}

func (o *GetTektonPipelineRunTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTektonPipelineRunTasksUnauthorized creates a GetTektonPipelineRunTasksUnauthorized with default headers values
func NewGetTektonPipelineRunTasksUnauthorized() *GetTektonPipelineRunTasksUnauthorized {
	return &GetTektonPipelineRunTasksUnauthorized{}
}

/*
GetTektonPipelineRunTasksUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTektonPipelineRunTasksUnauthorized struct {
}

// IsSuccess returns true when this get tekton pipeline run tasks unauthorized response has a 2xx status code
func (o *GetTektonPipelineRunTasksUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tekton pipeline run tasks unauthorized response has a 3xx status code
func (o *GetTektonPipelineRunTasksUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline run tasks unauthorized response has a 4xx status code
func (o *GetTektonPipelineRunTasksUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tekton pipeline run tasks unauthorized response has a 5xx status code
func (o *GetTektonPipelineRunTasksUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline run tasks unauthorized response a status code equal to that given
func (o *GetTektonPipelineRunTasksUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get tekton pipeline run tasks unauthorized response
func (o *GetTektonPipelineRunTasksUnauthorized) Code() int {
	return 401
}

func (o *GetTektonPipelineRunTasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks][%d] getTektonPipelineRunTasksUnauthorized", 401)
}

func (o *GetTektonPipelineRunTasksUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks][%d] getTektonPipelineRunTasksUnauthorized", 401)
}

func (o *GetTektonPipelineRunTasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTektonPipelineRunTasksNotFound creates a GetTektonPipelineRunTasksNotFound with default headers values
func NewGetTektonPipelineRunTasksNotFound() *GetTektonPipelineRunTasksNotFound {
	return &GetTektonPipelineRunTasksNotFound{}
}

/*
GetTektonPipelineRunTasksNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetTektonPipelineRunTasksNotFound struct {
}

// IsSuccess returns true when this get tekton pipeline run tasks not found response has a 2xx status code
func (o *GetTektonPipelineRunTasksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tekton pipeline run tasks not found response has a 3xx status code
func (o *GetTektonPipelineRunTasksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline run tasks not found response has a 4xx status code
func (o *GetTektonPipelineRunTasksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tekton pipeline run tasks not found response has a 5xx status code
func (o *GetTektonPipelineRunTasksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline run tasks not found response a status code equal to that given
func (o *GetTektonPipelineRunTasksNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get tekton pipeline run tasks not found response
func (o *GetTektonPipelineRunTasksNotFound) Code() int {
	return 404
}

func (o *GetTektonPipelineRunTasksNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks][%d] getTektonPipelineRunTasksNotFound", 404)
}

func (o *GetTektonPipelineRunTasksNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks][%d] getTektonPipelineRunTasksNotFound", 404)
}

func (o *GetTektonPipelineRunTasksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
