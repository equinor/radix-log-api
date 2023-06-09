// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StopJobReader is a Reader for the StopJob structure.
type StopJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStopJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStopJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopJobNoContent creates a StopJobNoContent with default headers values
func NewStopJobNoContent() *StopJobNoContent {
	return &StopJobNoContent{}
}

/* StopJobNoContent describes a response with status code 204, with default header values.

Success
*/
type StopJobNoContent struct {
}

func (o *StopJobNoContent) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/stop][%d] stopJobNoContent ", 204)
}

func (o *StopJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopJobBadRequest creates a StopJobBadRequest with default headers values
func NewStopJobBadRequest() *StopJobBadRequest {
	return &StopJobBadRequest{}
}

/* StopJobBadRequest describes a response with status code 400, with default header values.

Invalid job
*/
type StopJobBadRequest struct {
}

func (o *StopJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/stop][%d] stopJobBadRequest ", 400)
}

func (o *StopJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopJobUnauthorized creates a StopJobUnauthorized with default headers values
func NewStopJobUnauthorized() *StopJobUnauthorized {
	return &StopJobUnauthorized{}
}

/* StopJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StopJobUnauthorized struct {
}

func (o *StopJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/stop][%d] stopJobUnauthorized ", 401)
}

func (o *StopJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopJobForbidden creates a StopJobForbidden with default headers values
func NewStopJobForbidden() *StopJobForbidden {
	return &StopJobForbidden{}
}

/* StopJobForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StopJobForbidden struct {
}

func (o *StopJobForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/stop][%d] stopJobForbidden ", 403)
}

func (o *StopJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopJobNotFound creates a StopJobNotFound with default headers values
func NewStopJobNotFound() *StopJobNotFound {
	return &StopJobNotFound{}
}

/* StopJobNotFound describes a response with status code 404, with default header values.

Not found
*/
type StopJobNotFound struct {
}

func (o *StopJobNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/stop][%d] stopJobNotFound ", 404)
}

func (o *StopJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
