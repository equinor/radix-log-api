// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-log-api/pkg/radixapi/models"
)

// GetEnvironmentAlertingConfigReader is a Reader for the GetEnvironmentAlertingConfig structure.
type GetEnvironmentAlertingConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvironmentAlertingConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnvironmentAlertingConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEnvironmentAlertingConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEnvironmentAlertingConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEnvironmentAlertingConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEnvironmentAlertingConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEnvironmentAlertingConfigOK creates a GetEnvironmentAlertingConfigOK with default headers values
func NewGetEnvironmentAlertingConfigOK() *GetEnvironmentAlertingConfigOK {
	return &GetEnvironmentAlertingConfigOK{}
}

/* GetEnvironmentAlertingConfigOK describes a response with status code 200, with default header values.

Successful get alerts config
*/
type GetEnvironmentAlertingConfigOK struct {
	Payload *models.AlertingConfig
}

func (o *GetEnvironmentAlertingConfigOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/alerting][%d] getEnvironmentAlertingConfigOK  %+v", 200, o.Payload)
}
func (o *GetEnvironmentAlertingConfigOK) GetPayload() *models.AlertingConfig {
	return o.Payload
}

func (o *GetEnvironmentAlertingConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnvironmentAlertingConfigUnauthorized creates a GetEnvironmentAlertingConfigUnauthorized with default headers values
func NewGetEnvironmentAlertingConfigUnauthorized() *GetEnvironmentAlertingConfigUnauthorized {
	return &GetEnvironmentAlertingConfigUnauthorized{}
}

/* GetEnvironmentAlertingConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetEnvironmentAlertingConfigUnauthorized struct {
}

func (o *GetEnvironmentAlertingConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/alerting][%d] getEnvironmentAlertingConfigUnauthorized ", 401)
}

func (o *GetEnvironmentAlertingConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEnvironmentAlertingConfigForbidden creates a GetEnvironmentAlertingConfigForbidden with default headers values
func NewGetEnvironmentAlertingConfigForbidden() *GetEnvironmentAlertingConfigForbidden {
	return &GetEnvironmentAlertingConfigForbidden{}
}

/* GetEnvironmentAlertingConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetEnvironmentAlertingConfigForbidden struct {
}

func (o *GetEnvironmentAlertingConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/alerting][%d] getEnvironmentAlertingConfigForbidden ", 403)
}

func (o *GetEnvironmentAlertingConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEnvironmentAlertingConfigNotFound creates a GetEnvironmentAlertingConfigNotFound with default headers values
func NewGetEnvironmentAlertingConfigNotFound() *GetEnvironmentAlertingConfigNotFound {
	return &GetEnvironmentAlertingConfigNotFound{}
}

/* GetEnvironmentAlertingConfigNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetEnvironmentAlertingConfigNotFound struct {
}

func (o *GetEnvironmentAlertingConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/alerting][%d] getEnvironmentAlertingConfigNotFound ", 404)
}

func (o *GetEnvironmentAlertingConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEnvironmentAlertingConfigInternalServerError creates a GetEnvironmentAlertingConfigInternalServerError with default headers values
func NewGetEnvironmentAlertingConfigInternalServerError() *GetEnvironmentAlertingConfigInternalServerError {
	return &GetEnvironmentAlertingConfigInternalServerError{}
}

/* GetEnvironmentAlertingConfigInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetEnvironmentAlertingConfigInternalServerError struct {
}

func (o *GetEnvironmentAlertingConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/alerting][%d] getEnvironmentAlertingConfigInternalServerError ", 500)
}

func (o *GetEnvironmentAlertingConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
