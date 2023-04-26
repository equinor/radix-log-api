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

// UpdateApplicationAlertingConfigReader is a Reader for the UpdateApplicationAlertingConfig structure.
type UpdateApplicationAlertingConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateApplicationAlertingConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateApplicationAlertingConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateApplicationAlertingConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateApplicationAlertingConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateApplicationAlertingConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateApplicationAlertingConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateApplicationAlertingConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateApplicationAlertingConfigOK creates a UpdateApplicationAlertingConfigOK with default headers values
func NewUpdateApplicationAlertingConfigOK() *UpdateApplicationAlertingConfigOK {
	return &UpdateApplicationAlertingConfigOK{}
}

/* UpdateApplicationAlertingConfigOK describes a response with status code 200, with default header values.

Successful alerts config update
*/
type UpdateApplicationAlertingConfigOK struct {
	Payload *models.AlertingConfig
}

func (o *UpdateApplicationAlertingConfigOK) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigOK  %+v", 200, o.Payload)
}
func (o *UpdateApplicationAlertingConfigOK) GetPayload() *models.AlertingConfig {
	return o.Payload
}

func (o *UpdateApplicationAlertingConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateApplicationAlertingConfigBadRequest creates a UpdateApplicationAlertingConfigBadRequest with default headers values
func NewUpdateApplicationAlertingConfigBadRequest() *UpdateApplicationAlertingConfigBadRequest {
	return &UpdateApplicationAlertingConfigBadRequest{}
}

/* UpdateApplicationAlertingConfigBadRequest describes a response with status code 400, with default header values.

Invalid configuration
*/
type UpdateApplicationAlertingConfigBadRequest struct {
}

func (o *UpdateApplicationAlertingConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigBadRequest ", 400)
}

func (o *UpdateApplicationAlertingConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationAlertingConfigUnauthorized creates a UpdateApplicationAlertingConfigUnauthorized with default headers values
func NewUpdateApplicationAlertingConfigUnauthorized() *UpdateApplicationAlertingConfigUnauthorized {
	return &UpdateApplicationAlertingConfigUnauthorized{}
}

/* UpdateApplicationAlertingConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateApplicationAlertingConfigUnauthorized struct {
}

func (o *UpdateApplicationAlertingConfigUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigUnauthorized ", 401)
}

func (o *UpdateApplicationAlertingConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationAlertingConfigForbidden creates a UpdateApplicationAlertingConfigForbidden with default headers values
func NewUpdateApplicationAlertingConfigForbidden() *UpdateApplicationAlertingConfigForbidden {
	return &UpdateApplicationAlertingConfigForbidden{}
}

/* UpdateApplicationAlertingConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateApplicationAlertingConfigForbidden struct {
}

func (o *UpdateApplicationAlertingConfigForbidden) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigForbidden ", 403)
}

func (o *UpdateApplicationAlertingConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationAlertingConfigNotFound creates a UpdateApplicationAlertingConfigNotFound with default headers values
func NewUpdateApplicationAlertingConfigNotFound() *UpdateApplicationAlertingConfigNotFound {
	return &UpdateApplicationAlertingConfigNotFound{}
}

/* UpdateApplicationAlertingConfigNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateApplicationAlertingConfigNotFound struct {
}

func (o *UpdateApplicationAlertingConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigNotFound ", 404)
}

func (o *UpdateApplicationAlertingConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationAlertingConfigInternalServerError creates a UpdateApplicationAlertingConfigInternalServerError with default headers values
func NewUpdateApplicationAlertingConfigInternalServerError() *UpdateApplicationAlertingConfigInternalServerError {
	return &UpdateApplicationAlertingConfigInternalServerError{}
}

/* UpdateApplicationAlertingConfigInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateApplicationAlertingConfigInternalServerError struct {
}

func (o *UpdateApplicationAlertingConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigInternalServerError ", 500)
}

func (o *UpdateApplicationAlertingConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
