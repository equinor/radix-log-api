// Code generated by go-swagger; DO NOT EDIT.

package pipeline_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-log-api/pkg/radixapi/models"
)

// GetTektonPipelineRunsReader is a Reader for the GetTektonPipelineRuns structure.
type GetTektonPipelineRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTektonPipelineRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTektonPipelineRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTektonPipelineRunsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTektonPipelineRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTektonPipelineRunsOK creates a GetTektonPipelineRunsOK with default headers values
func NewGetTektonPipelineRunsOK() *GetTektonPipelineRunsOK {
	return &GetTektonPipelineRunsOK{}
}

/* GetTektonPipelineRunsOK describes a response with status code 200, with default header values.

List of PipelineRun-s
*/
type GetTektonPipelineRunsOK struct {
	Payload []*models.PipelineRun
}

func (o *GetTektonPipelineRunsOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns][%d] getTektonPipelineRunsOK  %+v", 200, o.Payload)
}
func (o *GetTektonPipelineRunsOK) GetPayload() []*models.PipelineRun {
	return o.Payload
}

func (o *GetTektonPipelineRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTektonPipelineRunsUnauthorized creates a GetTektonPipelineRunsUnauthorized with default headers values
func NewGetTektonPipelineRunsUnauthorized() *GetTektonPipelineRunsUnauthorized {
	return &GetTektonPipelineRunsUnauthorized{}
}

/* GetTektonPipelineRunsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTektonPipelineRunsUnauthorized struct {
}

func (o *GetTektonPipelineRunsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns][%d] getTektonPipelineRunsUnauthorized ", 401)
}

func (o *GetTektonPipelineRunsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTektonPipelineRunsNotFound creates a GetTektonPipelineRunsNotFound with default headers values
func NewGetTektonPipelineRunsNotFound() *GetTektonPipelineRunsNotFound {
	return &GetTektonPipelineRunsNotFound{}
}

/* GetTektonPipelineRunsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetTektonPipelineRunsNotFound struct {
}

func (o *GetTektonPipelineRunsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns][%d] getTektonPipelineRunsNotFound ", 404)
}

func (o *GetTektonPipelineRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
