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

// RegenerateDeployKeyReader is a Reader for the RegenerateDeployKey structure.
type RegenerateDeployKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegenerateDeployKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegenerateDeployKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRegenerateDeployKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRegenerateDeployKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegenerateDeployKeyOK creates a RegenerateDeployKeyOK with default headers values
func NewRegenerateDeployKeyOK() *RegenerateDeployKeyOK {
	return &RegenerateDeployKeyOK{}
}

/* RegenerateDeployKeyOK describes a response with status code 200, with default header values.

Successful regenerate machine-user token
*/
type RegenerateDeployKeyOK struct {
	Payload *models.DeployKeyAndSecret
}

func (o *RegenerateDeployKeyOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-deploy-key][%d] regenerateDeployKeyOK  %+v", 200, o.Payload)
}
func (o *RegenerateDeployKeyOK) GetPayload() *models.DeployKeyAndSecret {
	return o.Payload
}

func (o *RegenerateDeployKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeployKeyAndSecret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateDeployKeyUnauthorized creates a RegenerateDeployKeyUnauthorized with default headers values
func NewRegenerateDeployKeyUnauthorized() *RegenerateDeployKeyUnauthorized {
	return &RegenerateDeployKeyUnauthorized{}
}

/* RegenerateDeployKeyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RegenerateDeployKeyUnauthorized struct {
}

func (o *RegenerateDeployKeyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-deploy-key][%d] regenerateDeployKeyUnauthorized ", 401)
}

func (o *RegenerateDeployKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateDeployKeyNotFound creates a RegenerateDeployKeyNotFound with default headers values
func NewRegenerateDeployKeyNotFound() *RegenerateDeployKeyNotFound {
	return &RegenerateDeployKeyNotFound{}
}

/* RegenerateDeployKeyNotFound describes a response with status code 404, with default header values.

Not found
*/
type RegenerateDeployKeyNotFound struct {
}

func (o *RegenerateDeployKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-deploy-key][%d] regenerateDeployKeyNotFound ", 404)
}

func (o *RegenerateDeployKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}