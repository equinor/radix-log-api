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

// GetApplicationJobsReader is a Reader for the GetApplicationJobs structure.
type GetApplicationJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApplicationJobsOK creates a GetApplicationJobsOK with default headers values
func NewGetApplicationJobsOK() *GetApplicationJobsOK {
	return &GetApplicationJobsOK{}
}

/* GetApplicationJobsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetApplicationJobsOK struct {
	Payload []*models.JobSummary
}

func (o *GetApplicationJobsOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs][%d] getApplicationJobsOK  %+v", 200, o.Payload)
}
func (o *GetApplicationJobsOK) GetPayload() []*models.JobSummary {
	return o.Payload
}

func (o *GetApplicationJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationJobsUnauthorized creates a GetApplicationJobsUnauthorized with default headers values
func NewGetApplicationJobsUnauthorized() *GetApplicationJobsUnauthorized {
	return &GetApplicationJobsUnauthorized{}
}

/* GetApplicationJobsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetApplicationJobsUnauthorized struct {
}

func (o *GetApplicationJobsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs][%d] getApplicationJobsUnauthorized ", 401)
}

func (o *GetApplicationJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationJobsNotFound creates a GetApplicationJobsNotFound with default headers values
func NewGetApplicationJobsNotFound() *GetApplicationJobsNotFound {
	return &GetApplicationJobsNotFound{}
}

/* GetApplicationJobsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetApplicationJobsNotFound struct {
}

func (o *GetApplicationJobsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs][%d] getApplicationJobsNotFound ", 404)
}

func (o *GetApplicationJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
