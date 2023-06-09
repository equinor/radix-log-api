// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StopEnvironmentReader is a Reader for the StopEnvironment structure.
type StopEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStopEnvironmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopEnvironmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopEnvironmentOK creates a StopEnvironmentOK with default headers values
func NewStopEnvironmentOK() *StopEnvironmentOK {
	return &StopEnvironmentOK{}
}

/* StopEnvironmentOK describes a response with status code 200, with default header values.

Environment stopped ok
*/
type StopEnvironmentOK struct {
}

func (o *StopEnvironmentOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/stop][%d] stopEnvironmentOK ", 200)
}

func (o *StopEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopEnvironmentUnauthorized creates a StopEnvironmentUnauthorized with default headers values
func NewStopEnvironmentUnauthorized() *StopEnvironmentUnauthorized {
	return &StopEnvironmentUnauthorized{}
}

/* StopEnvironmentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StopEnvironmentUnauthorized struct {
}

func (o *StopEnvironmentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/stop][%d] stopEnvironmentUnauthorized ", 401)
}

func (o *StopEnvironmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopEnvironmentNotFound creates a StopEnvironmentNotFound with default headers values
func NewStopEnvironmentNotFound() *StopEnvironmentNotFound {
	return &StopEnvironmentNotFound{}
}

/* StopEnvironmentNotFound describes a response with status code 404, with default header values.

Not found
*/
type StopEnvironmentNotFound struct {
}

func (o *StopEnvironmentNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/stop][%d] stopEnvironmentNotFound ", 404)
}

func (o *StopEnvironmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
