// Code generated by go-swagger; DO NOT EDIT.

package buildstatus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetBuildStatusReader is a Reader for the GetBuildStatus structure.
type GetBuildStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetBuildStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/environments/{envName}/buildstatus] getBuildStatus", response, response.Code())
	}
}

// NewGetBuildStatusOK creates a GetBuildStatusOK with default headers values
func NewGetBuildStatusOK() *GetBuildStatusOK {
	return &GetBuildStatusOK{}
}

/*
GetBuildStatusOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetBuildStatusOK struct {
}

// IsSuccess returns true when this get build status o k response has a 2xx status code
func (o *GetBuildStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get build status o k response has a 3xx status code
func (o *GetBuildStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get build status o k response has a 4xx status code
func (o *GetBuildStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get build status o k response has a 5xx status code
func (o *GetBuildStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get build status o k response a status code equal to that given
func (o *GetBuildStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get build status o k response
func (o *GetBuildStatusOK) Code() int {
	return 200
}

func (o *GetBuildStatusOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/buildstatus][%d] getBuildStatusOK", 200)
}

func (o *GetBuildStatusOK) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/buildstatus][%d] getBuildStatusOK", 200)
}

func (o *GetBuildStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBuildStatusInternalServerError creates a GetBuildStatusInternalServerError with default headers values
func NewGetBuildStatusInternalServerError() *GetBuildStatusInternalServerError {
	return &GetBuildStatusInternalServerError{}
}

/*
GetBuildStatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetBuildStatusInternalServerError struct {
}

// IsSuccess returns true when this get build status internal server error response has a 2xx status code
func (o *GetBuildStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get build status internal server error response has a 3xx status code
func (o *GetBuildStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get build status internal server error response has a 4xx status code
func (o *GetBuildStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get build status internal server error response has a 5xx status code
func (o *GetBuildStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get build status internal server error response a status code equal to that given
func (o *GetBuildStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get build status internal server error response
func (o *GetBuildStatusInternalServerError) Code() int {
	return 500
}

func (o *GetBuildStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/buildstatus][%d] getBuildStatusInternalServerError", 500)
}

func (o *GetBuildStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/buildstatus][%d] getBuildStatusInternalServerError", 500)
}

func (o *GetBuildStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
