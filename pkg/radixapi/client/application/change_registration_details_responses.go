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

// ChangeRegistrationDetailsReader is a Reader for the ChangeRegistrationDetails structure.
type ChangeRegistrationDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeRegistrationDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeRegistrationDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewChangeRegistrationDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewChangeRegistrationDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewChangeRegistrationDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewChangeRegistrationDetailsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewChangeRegistrationDetailsOK creates a ChangeRegistrationDetailsOK with default headers values
func NewChangeRegistrationDetailsOK() *ChangeRegistrationDetailsOK {
	return &ChangeRegistrationDetailsOK{}
}

/* ChangeRegistrationDetailsOK describes a response with status code 200, with default header values.

Change registration operation result
*/
type ChangeRegistrationDetailsOK struct {
	Payload *models.ApplicationRegistrationUpsertResponse
}

func (o *ChangeRegistrationDetailsOK) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsOK  %+v", 200, o.Payload)
}
func (o *ChangeRegistrationDetailsOK) GetPayload() *models.ApplicationRegistrationUpsertResponse {
	return o.Payload
}

func (o *ChangeRegistrationDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationRegistrationUpsertResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeRegistrationDetailsBadRequest creates a ChangeRegistrationDetailsBadRequest with default headers values
func NewChangeRegistrationDetailsBadRequest() *ChangeRegistrationDetailsBadRequest {
	return &ChangeRegistrationDetailsBadRequest{}
}

/* ChangeRegistrationDetailsBadRequest describes a response with status code 400, with default header values.

Invalid application
*/
type ChangeRegistrationDetailsBadRequest struct {
}

func (o *ChangeRegistrationDetailsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsBadRequest ", 400)
}

func (o *ChangeRegistrationDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeRegistrationDetailsUnauthorized creates a ChangeRegistrationDetailsUnauthorized with default headers values
func NewChangeRegistrationDetailsUnauthorized() *ChangeRegistrationDetailsUnauthorized {
	return &ChangeRegistrationDetailsUnauthorized{}
}

/* ChangeRegistrationDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ChangeRegistrationDetailsUnauthorized struct {
}

func (o *ChangeRegistrationDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsUnauthorized ", 401)
}

func (o *ChangeRegistrationDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeRegistrationDetailsNotFound creates a ChangeRegistrationDetailsNotFound with default headers values
func NewChangeRegistrationDetailsNotFound() *ChangeRegistrationDetailsNotFound {
	return &ChangeRegistrationDetailsNotFound{}
}

/* ChangeRegistrationDetailsNotFound describes a response with status code 404, with default header values.

Not found
*/
type ChangeRegistrationDetailsNotFound struct {
}

func (o *ChangeRegistrationDetailsNotFound) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsNotFound ", 404)
}

func (o *ChangeRegistrationDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeRegistrationDetailsConflict creates a ChangeRegistrationDetailsConflict with default headers values
func NewChangeRegistrationDetailsConflict() *ChangeRegistrationDetailsConflict {
	return &ChangeRegistrationDetailsConflict{}
}

/* ChangeRegistrationDetailsConflict describes a response with status code 409, with default header values.

Conflict
*/
type ChangeRegistrationDetailsConflict struct {
}

func (o *ChangeRegistrationDetailsConflict) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsConflict ", 409)
}

func (o *ChangeRegistrationDetailsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
