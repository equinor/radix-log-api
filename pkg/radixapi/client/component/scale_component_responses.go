// Code generated by go-swagger; DO NOT EDIT.

package component

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ScaleComponentReader is a Reader for the ScaleComponent structure.
type ScaleComponentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScaleComponentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewScaleComponentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewScaleComponentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewScaleComponentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewScaleComponentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewScaleComponentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}] scaleComponent", response, response.Code())
	}
}

// NewScaleComponentNoContent creates a ScaleComponentNoContent with default headers values
func NewScaleComponentNoContent() *ScaleComponentNoContent {
	return &ScaleComponentNoContent{}
}

/*
ScaleComponentNoContent describes a response with status code 204, with default header values.

Success
*/
type ScaleComponentNoContent struct {
}

// IsSuccess returns true when this scale component no content response has a 2xx status code
func (o *ScaleComponentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this scale component no content response has a 3xx status code
func (o *ScaleComponentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scale component no content response has a 4xx status code
func (o *ScaleComponentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this scale component no content response has a 5xx status code
func (o *ScaleComponentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this scale component no content response a status code equal to that given
func (o *ScaleComponentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the scale component no content response
func (o *ScaleComponentNoContent) Code() int {
	return 204
}

func (o *ScaleComponentNoContent) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentNoContent", 204)
}

func (o *ScaleComponentNoContent) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentNoContent", 204)
}

func (o *ScaleComponentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScaleComponentBadRequest creates a ScaleComponentBadRequest with default headers values
func NewScaleComponentBadRequest() *ScaleComponentBadRequest {
	return &ScaleComponentBadRequest{}
}

/*
ScaleComponentBadRequest describes a response with status code 400, with default header values.

Invalid component
*/
type ScaleComponentBadRequest struct {
}

// IsSuccess returns true when this scale component bad request response has a 2xx status code
func (o *ScaleComponentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this scale component bad request response has a 3xx status code
func (o *ScaleComponentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scale component bad request response has a 4xx status code
func (o *ScaleComponentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this scale component bad request response has a 5xx status code
func (o *ScaleComponentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this scale component bad request response a status code equal to that given
func (o *ScaleComponentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the scale component bad request response
func (o *ScaleComponentBadRequest) Code() int {
	return 400
}

func (o *ScaleComponentBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentBadRequest", 400)
}

func (o *ScaleComponentBadRequest) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentBadRequest", 400)
}

func (o *ScaleComponentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScaleComponentUnauthorized creates a ScaleComponentUnauthorized with default headers values
func NewScaleComponentUnauthorized() *ScaleComponentUnauthorized {
	return &ScaleComponentUnauthorized{}
}

/*
ScaleComponentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ScaleComponentUnauthorized struct {
}

// IsSuccess returns true when this scale component unauthorized response has a 2xx status code
func (o *ScaleComponentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this scale component unauthorized response has a 3xx status code
func (o *ScaleComponentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scale component unauthorized response has a 4xx status code
func (o *ScaleComponentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this scale component unauthorized response has a 5xx status code
func (o *ScaleComponentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this scale component unauthorized response a status code equal to that given
func (o *ScaleComponentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the scale component unauthorized response
func (o *ScaleComponentUnauthorized) Code() int {
	return 401
}

func (o *ScaleComponentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentUnauthorized", 401)
}

func (o *ScaleComponentUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentUnauthorized", 401)
}

func (o *ScaleComponentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScaleComponentForbidden creates a ScaleComponentForbidden with default headers values
func NewScaleComponentForbidden() *ScaleComponentForbidden {
	return &ScaleComponentForbidden{}
}

/*
ScaleComponentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ScaleComponentForbidden struct {
}

// IsSuccess returns true when this scale component forbidden response has a 2xx status code
func (o *ScaleComponentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this scale component forbidden response has a 3xx status code
func (o *ScaleComponentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scale component forbidden response has a 4xx status code
func (o *ScaleComponentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this scale component forbidden response has a 5xx status code
func (o *ScaleComponentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this scale component forbidden response a status code equal to that given
func (o *ScaleComponentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the scale component forbidden response
func (o *ScaleComponentForbidden) Code() int {
	return 403
}

func (o *ScaleComponentForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentForbidden", 403)
}

func (o *ScaleComponentForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentForbidden", 403)
}

func (o *ScaleComponentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScaleComponentNotFound creates a ScaleComponentNotFound with default headers values
func NewScaleComponentNotFound() *ScaleComponentNotFound {
	return &ScaleComponentNotFound{}
}

/*
ScaleComponentNotFound describes a response with status code 404, with default header values.

Not found
*/
type ScaleComponentNotFound struct {
}

// IsSuccess returns true when this scale component not found response has a 2xx status code
func (o *ScaleComponentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this scale component not found response has a 3xx status code
func (o *ScaleComponentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scale component not found response has a 4xx status code
func (o *ScaleComponentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this scale component not found response has a 5xx status code
func (o *ScaleComponentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this scale component not found response a status code equal to that given
func (o *ScaleComponentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the scale component not found response
func (o *ScaleComponentNotFound) Code() int {
	return 404
}

func (o *ScaleComponentNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentNotFound", 404)
}

func (o *ScaleComponentNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentNotFound", 404)
}

func (o *ScaleComponentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
