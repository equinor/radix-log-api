// Code generated by go-swagger; DO NOT EDIT.

package component

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

// EnvVarsReader is a Reader for the EnvVars structure.
type EnvVarsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnvVarsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnvVarsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewEnvVarsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/environments/{envName}/components/{componentName}/envvars] envVars", response, response.Code())
	}
}

// NewEnvVarsOK creates a EnvVarsOK with default headers values
func NewEnvVarsOK() *EnvVarsOK {
	return &EnvVarsOK{}
}

/*
EnvVarsOK describes a response with status code 200, with default header values.

environment variables
*/
type EnvVarsOK struct {
	Payload []*models.EnvVar
}

// IsSuccess returns true when this env vars o k response has a 2xx status code
func (o *EnvVarsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this env vars o k response has a 3xx status code
func (o *EnvVarsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this env vars o k response has a 4xx status code
func (o *EnvVarsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this env vars o k response has a 5xx status code
func (o *EnvVarsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this env vars o k response a status code equal to that given
func (o *EnvVarsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the env vars o k response
func (o *EnvVarsOK) Code() int {
	return 200
}

func (o *EnvVarsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/envvars][%d] envVarsOK %s", 200, payload)
}

func (o *EnvVarsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/envvars][%d] envVarsOK %s", 200, payload)
}

func (o *EnvVarsOK) GetPayload() []*models.EnvVar {
	return o.Payload
}

func (o *EnvVarsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnvVarsNotFound creates a EnvVarsNotFound with default headers values
func NewEnvVarsNotFound() *EnvVarsNotFound {
	return &EnvVarsNotFound{}
}

/*
EnvVarsNotFound describes a response with status code 404, with default header values.

Not found
*/
type EnvVarsNotFound struct {
}

// IsSuccess returns true when this env vars not found response has a 2xx status code
func (o *EnvVarsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this env vars not found response has a 3xx status code
func (o *EnvVarsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this env vars not found response has a 4xx status code
func (o *EnvVarsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this env vars not found response has a 5xx status code
func (o *EnvVarsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this env vars not found response a status code equal to that given
func (o *EnvVarsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the env vars not found response
func (o *EnvVarsNotFound) Code() int {
	return 404
}

func (o *EnvVarsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/envvars][%d] envVarsNotFound", 404)
}

func (o *EnvVarsNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/envvars][%d] envVarsNotFound", 404)
}

func (o *EnvVarsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
