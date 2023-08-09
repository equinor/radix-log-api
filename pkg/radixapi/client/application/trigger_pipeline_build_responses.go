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
	case 403:
		result := NewTriggerPipelineBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTriggerPipelineBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/pipelines/build] triggerPipelineBuild", response, response.Code())
	}
}

// NewTriggerPipelineBuildOK creates a TriggerPipelineBuildOK with default headers values
func NewTriggerPipelineBuildOK() *TriggerPipelineBuildOK {
	return &TriggerPipelineBuildOK{}
}

/*
TriggerPipelineBuildOK describes a response with status code 200, with default header values.

Successful trigger pipeline
*/
type TriggerPipelineBuildOK struct {
	Payload *models.JobSummary
}

// IsSuccess returns true when this trigger pipeline build o k response has a 2xx status code
func (o *TriggerPipelineBuildOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trigger pipeline build o k response has a 3xx status code
func (o *TriggerPipelineBuildOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline build o k response has a 4xx status code
func (o *TriggerPipelineBuildOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this trigger pipeline build o k response has a 5xx status code
func (o *TriggerPipelineBuildOK) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline build o k response a status code equal to that given
func (o *TriggerPipelineBuildOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the trigger pipeline build o k response
func (o *TriggerPipelineBuildOK) Code() int {
	return 200
}

func (o *TriggerPipelineBuildOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build][%d] triggerPipelineBuildOK  %+v", 200, o.Payload)
}

func (o *TriggerPipelineBuildOK) String() string {
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

// NewTriggerPipelineBuildForbidden creates a TriggerPipelineBuildForbidden with default headers values
func NewTriggerPipelineBuildForbidden() *TriggerPipelineBuildForbidden {
	return &TriggerPipelineBuildForbidden{}
}

/*
TriggerPipelineBuildForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TriggerPipelineBuildForbidden struct {
}

// IsSuccess returns true when this trigger pipeline build forbidden response has a 2xx status code
func (o *TriggerPipelineBuildForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger pipeline build forbidden response has a 3xx status code
func (o *TriggerPipelineBuildForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline build forbidden response has a 4xx status code
func (o *TriggerPipelineBuildForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger pipeline build forbidden response has a 5xx status code
func (o *TriggerPipelineBuildForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline build forbidden response a status code equal to that given
func (o *TriggerPipelineBuildForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the trigger pipeline build forbidden response
func (o *TriggerPipelineBuildForbidden) Code() int {
	return 403
}

func (o *TriggerPipelineBuildForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build][%d] triggerPipelineBuildForbidden ", 403)
}

func (o *TriggerPipelineBuildForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build][%d] triggerPipelineBuildForbidden ", 403)
}

func (o *TriggerPipelineBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerPipelineBuildNotFound creates a TriggerPipelineBuildNotFound with default headers values
func NewTriggerPipelineBuildNotFound() *TriggerPipelineBuildNotFound {
	return &TriggerPipelineBuildNotFound{}
}

/*
TriggerPipelineBuildNotFound describes a response with status code 404, with default header values.

Not found
*/
type TriggerPipelineBuildNotFound struct {
}

// IsSuccess returns true when this trigger pipeline build not found response has a 2xx status code
func (o *TriggerPipelineBuildNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger pipeline build not found response has a 3xx status code
func (o *TriggerPipelineBuildNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline build not found response has a 4xx status code
func (o *TriggerPipelineBuildNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger pipeline build not found response has a 5xx status code
func (o *TriggerPipelineBuildNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline build not found response a status code equal to that given
func (o *TriggerPipelineBuildNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the trigger pipeline build not found response
func (o *TriggerPipelineBuildNotFound) Code() int {
	return 404
}

func (o *TriggerPipelineBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build][%d] triggerPipelineBuildNotFound ", 404)
}

func (o *TriggerPipelineBuildNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build][%d] triggerPipelineBuildNotFound ", 404)
}

func (o *TriggerPipelineBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
