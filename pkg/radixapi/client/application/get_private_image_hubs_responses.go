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

// GetPrivateImageHubsReader is a Reader for the GetPrivateImageHubs structure.
type GetPrivateImageHubsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateImageHubsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateImageHubsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPrivateImageHubsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPrivateImageHubsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateImageHubsOK creates a GetPrivateImageHubsOK with default headers values
func NewGetPrivateImageHubsOK() *GetPrivateImageHubsOK {
	return &GetPrivateImageHubsOK{}
}

/* GetPrivateImageHubsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetPrivateImageHubsOK struct {
	Payload []*models.ImageHubSecret
}

func (o *GetPrivateImageHubsOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/privateimagehubs][%d] getPrivateImageHubsOK  %+v", 200, o.Payload)
}
func (o *GetPrivateImageHubsOK) GetPayload() []*models.ImageHubSecret {
	return o.Payload
}

func (o *GetPrivateImageHubsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateImageHubsUnauthorized creates a GetPrivateImageHubsUnauthorized with default headers values
func NewGetPrivateImageHubsUnauthorized() *GetPrivateImageHubsUnauthorized {
	return &GetPrivateImageHubsUnauthorized{}
}

/* GetPrivateImageHubsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPrivateImageHubsUnauthorized struct {
}

func (o *GetPrivateImageHubsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/privateimagehubs][%d] getPrivateImageHubsUnauthorized ", 401)
}

func (o *GetPrivateImageHubsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPrivateImageHubsNotFound creates a GetPrivateImageHubsNotFound with default headers values
func NewGetPrivateImageHubsNotFound() *GetPrivateImageHubsNotFound {
	return &GetPrivateImageHubsNotFound{}
}

/* GetPrivateImageHubsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetPrivateImageHubsNotFound struct {
}

func (o *GetPrivateImageHubsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/privateimagehubs][%d] getPrivateImageHubsNotFound ", 404)
}

func (o *GetPrivateImageHubsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
