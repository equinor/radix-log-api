// Code generated by go-swagger; DO NOT EDIT.

package deployment

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

// GetDeploymentReader is a Reader for the GetDeployment structure.
type GetDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/deployments/{deploymentName}] getDeployment", response, response.Code())
	}
}

// NewGetDeploymentOK creates a GetDeploymentOK with default headers values
func NewGetDeploymentOK() *GetDeploymentOK {
	return &GetDeploymentOK{}
}

/*
GetDeploymentOK describes a response with status code 200, with default header values.

Successful get deployment
*/
type GetDeploymentOK struct {
	Payload *models.Deployment
}

// IsSuccess returns true when this get deployment o k response has a 2xx status code
func (o *GetDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployment o k response has a 3xx status code
func (o *GetDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment o k response has a 4xx status code
func (o *GetDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment o k response has a 5xx status code
func (o *GetDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment o k response a status code equal to that given
func (o *GetDeploymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get deployment o k response
func (o *GetDeploymentOK) Code() int {
	return 200
}

func (o *GetDeploymentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}][%d] getDeploymentOK %s", 200, payload)
}

func (o *GetDeploymentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}][%d] getDeploymentOK %s", 200, payload)
}

func (o *GetDeploymentOK) GetPayload() *models.Deployment {
	return o.Payload
}

func (o *GetDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentUnauthorized creates a GetDeploymentUnauthorized with default headers values
func NewGetDeploymentUnauthorized() *GetDeploymentUnauthorized {
	return &GetDeploymentUnauthorized{}
}

/*
GetDeploymentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeploymentUnauthorized struct {
}

// IsSuccess returns true when this get deployment unauthorized response has a 2xx status code
func (o *GetDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment unauthorized response has a 3xx status code
func (o *GetDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment unauthorized response has a 4xx status code
func (o *GetDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment unauthorized response has a 5xx status code
func (o *GetDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment unauthorized response a status code equal to that given
func (o *GetDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get deployment unauthorized response
func (o *GetDeploymentUnauthorized) Code() int {
	return 401
}

func (o *GetDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}][%d] getDeploymentUnauthorized", 401)
}

func (o *GetDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}][%d] getDeploymentUnauthorized", 401)
}

func (o *GetDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentNotFound creates a GetDeploymentNotFound with default headers values
func NewGetDeploymentNotFound() *GetDeploymentNotFound {
	return &GetDeploymentNotFound{}
}

/*
GetDeploymentNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetDeploymentNotFound struct {
}

// IsSuccess returns true when this get deployment not found response has a 2xx status code
func (o *GetDeploymentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment not found response has a 3xx status code
func (o *GetDeploymentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment not found response has a 4xx status code
func (o *GetDeploymentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment not found response has a 5xx status code
func (o *GetDeploymentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment not found response a status code equal to that given
func (o *GetDeploymentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get deployment not found response
func (o *GetDeploymentNotFound) Code() int {
	return 404
}

func (o *GetDeploymentNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}][%d] getDeploymentNotFound", 404)
}

func (o *GetDeploymentNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}][%d] getDeploymentNotFound", 404)
}

func (o *GetDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
