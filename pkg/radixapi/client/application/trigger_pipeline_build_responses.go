// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-log-api/pkg/radixapi/models"
)

// TriggerPipelineBuildReader is a Reader for the TriggerPipelineBuild structure.
type TriggerPipelineBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerPipelineBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTriggerPipelineBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewTriggerPipelineBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTriggerPipelineBuildOK creates a TriggerPipelineBuildOK with default headers values
func NewTriggerPipelineBuildOK() *TriggerPipelineBuildOK {
	return &TriggerPipelineBuildOK{}
}

/* TriggerPipelineBuildOK describes a response with status code 200, with default header values.

Successful trigger pipeline
*/
type TriggerPipelineBuildOK struct {
	Payload *models.JobSummary
}

func (o *TriggerPipelineBuildOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build][%d] triggerPipelineBuildOK  %+v", 200, o.Payload)
}
func (o *TriggerPipelineBuildOK) GetPayload() *models.JobSummary {
	return o.Payload
}

func (o *TriggerPipelineBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggerPipelineBuildNotFound creates a TriggerPipelineBuildNotFound with default headers values
func NewTriggerPipelineBuildNotFound() *TriggerPipelineBuildNotFound {
	return &TriggerPipelineBuildNotFound{}
}

/* TriggerPipelineBuildNotFound describes a response with status code 404, with default header values.

Not found
*/
type TriggerPipelineBuildNotFound struct {
}

func (o *TriggerPipelineBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build][%d] triggerPipelineBuildNotFound ", 404)
}

func (o *TriggerPipelineBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
