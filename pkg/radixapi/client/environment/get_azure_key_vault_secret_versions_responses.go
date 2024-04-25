// Code generated by go-swagger; DO NOT EDIT.

package environment

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

// GetAzureKeyVaultSecretVersionsReader is a Reader for the GetAzureKeyVaultSecretVersions structure.
type GetAzureKeyVaultSecretVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAzureKeyVaultSecretVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAzureKeyVaultSecretVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAzureKeyVaultSecretVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAzureKeyVaultSecretVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAzureKeyVaultSecretVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAzureKeyVaultSecretVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAzureKeyVaultSecretVersionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAzureKeyVaultSecretVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}] getAzureKeyVaultSecretVersions", response, response.Code())
	}
}

// NewGetAzureKeyVaultSecretVersionsOK creates a GetAzureKeyVaultSecretVersionsOK with default headers values
func NewGetAzureKeyVaultSecretVersionsOK() *GetAzureKeyVaultSecretVersionsOK {
	return &GetAzureKeyVaultSecretVersionsOK{}
}

/*
GetAzureKeyVaultSecretVersionsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetAzureKeyVaultSecretVersionsOK struct {
	Payload []*models.AzureKeyVaultSecretVersion
}

// IsSuccess returns true when this get azure key vault secret versions o k response has a 2xx status code
func (o *GetAzureKeyVaultSecretVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get azure key vault secret versions o k response has a 3xx status code
func (o *GetAzureKeyVaultSecretVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure key vault secret versions o k response has a 4xx status code
func (o *GetAzureKeyVaultSecretVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get azure key vault secret versions o k response has a 5xx status code
func (o *GetAzureKeyVaultSecretVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure key vault secret versions o k response a status code equal to that given
func (o *GetAzureKeyVaultSecretVersionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get azure key vault secret versions o k response
func (o *GetAzureKeyVaultSecretVersionsOK) Code() int {
	return 200
}

func (o *GetAzureKeyVaultSecretVersionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsOK %s", 200, payload)
}

func (o *GetAzureKeyVaultSecretVersionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsOK %s", 200, payload)
}

func (o *GetAzureKeyVaultSecretVersionsOK) GetPayload() []*models.AzureKeyVaultSecretVersion {
	return o.Payload
}

func (o *GetAzureKeyVaultSecretVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureKeyVaultSecretVersionsBadRequest creates a GetAzureKeyVaultSecretVersionsBadRequest with default headers values
func NewGetAzureKeyVaultSecretVersionsBadRequest() *GetAzureKeyVaultSecretVersionsBadRequest {
	return &GetAzureKeyVaultSecretVersionsBadRequest{}
}

/*
GetAzureKeyVaultSecretVersionsBadRequest describes a response with status code 400, with default header values.

Invalid application
*/
type GetAzureKeyVaultSecretVersionsBadRequest struct {
}

// IsSuccess returns true when this get azure key vault secret versions bad request response has a 2xx status code
func (o *GetAzureKeyVaultSecretVersionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get azure key vault secret versions bad request response has a 3xx status code
func (o *GetAzureKeyVaultSecretVersionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure key vault secret versions bad request response has a 4xx status code
func (o *GetAzureKeyVaultSecretVersionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get azure key vault secret versions bad request response has a 5xx status code
func (o *GetAzureKeyVaultSecretVersionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure key vault secret versions bad request response a status code equal to that given
func (o *GetAzureKeyVaultSecretVersionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get azure key vault secret versions bad request response
func (o *GetAzureKeyVaultSecretVersionsBadRequest) Code() int {
	return 400
}

func (o *GetAzureKeyVaultSecretVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsBadRequest", 400)
}

func (o *GetAzureKeyVaultSecretVersionsBadRequest) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsBadRequest", 400)
}

func (o *GetAzureKeyVaultSecretVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAzureKeyVaultSecretVersionsUnauthorized creates a GetAzureKeyVaultSecretVersionsUnauthorized with default headers values
func NewGetAzureKeyVaultSecretVersionsUnauthorized() *GetAzureKeyVaultSecretVersionsUnauthorized {
	return &GetAzureKeyVaultSecretVersionsUnauthorized{}
}

/*
GetAzureKeyVaultSecretVersionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAzureKeyVaultSecretVersionsUnauthorized struct {
}

// IsSuccess returns true when this get azure key vault secret versions unauthorized response has a 2xx status code
func (o *GetAzureKeyVaultSecretVersionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get azure key vault secret versions unauthorized response has a 3xx status code
func (o *GetAzureKeyVaultSecretVersionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure key vault secret versions unauthorized response has a 4xx status code
func (o *GetAzureKeyVaultSecretVersionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get azure key vault secret versions unauthorized response has a 5xx status code
func (o *GetAzureKeyVaultSecretVersionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure key vault secret versions unauthorized response a status code equal to that given
func (o *GetAzureKeyVaultSecretVersionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get azure key vault secret versions unauthorized response
func (o *GetAzureKeyVaultSecretVersionsUnauthorized) Code() int {
	return 401
}

func (o *GetAzureKeyVaultSecretVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsUnauthorized", 401)
}

func (o *GetAzureKeyVaultSecretVersionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsUnauthorized", 401)
}

func (o *GetAzureKeyVaultSecretVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAzureKeyVaultSecretVersionsForbidden creates a GetAzureKeyVaultSecretVersionsForbidden with default headers values
func NewGetAzureKeyVaultSecretVersionsForbidden() *GetAzureKeyVaultSecretVersionsForbidden {
	return &GetAzureKeyVaultSecretVersionsForbidden{}
}

/*
GetAzureKeyVaultSecretVersionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAzureKeyVaultSecretVersionsForbidden struct {
}

// IsSuccess returns true when this get azure key vault secret versions forbidden response has a 2xx status code
func (o *GetAzureKeyVaultSecretVersionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get azure key vault secret versions forbidden response has a 3xx status code
func (o *GetAzureKeyVaultSecretVersionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure key vault secret versions forbidden response has a 4xx status code
func (o *GetAzureKeyVaultSecretVersionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get azure key vault secret versions forbidden response has a 5xx status code
func (o *GetAzureKeyVaultSecretVersionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure key vault secret versions forbidden response a status code equal to that given
func (o *GetAzureKeyVaultSecretVersionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get azure key vault secret versions forbidden response
func (o *GetAzureKeyVaultSecretVersionsForbidden) Code() int {
	return 403
}

func (o *GetAzureKeyVaultSecretVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsForbidden", 403)
}

func (o *GetAzureKeyVaultSecretVersionsForbidden) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsForbidden", 403)
}

func (o *GetAzureKeyVaultSecretVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAzureKeyVaultSecretVersionsNotFound creates a GetAzureKeyVaultSecretVersionsNotFound with default headers values
func NewGetAzureKeyVaultSecretVersionsNotFound() *GetAzureKeyVaultSecretVersionsNotFound {
	return &GetAzureKeyVaultSecretVersionsNotFound{}
}

/*
GetAzureKeyVaultSecretVersionsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetAzureKeyVaultSecretVersionsNotFound struct {
}

// IsSuccess returns true when this get azure key vault secret versions not found response has a 2xx status code
func (o *GetAzureKeyVaultSecretVersionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get azure key vault secret versions not found response has a 3xx status code
func (o *GetAzureKeyVaultSecretVersionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure key vault secret versions not found response has a 4xx status code
func (o *GetAzureKeyVaultSecretVersionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get azure key vault secret versions not found response has a 5xx status code
func (o *GetAzureKeyVaultSecretVersionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure key vault secret versions not found response a status code equal to that given
func (o *GetAzureKeyVaultSecretVersionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get azure key vault secret versions not found response
func (o *GetAzureKeyVaultSecretVersionsNotFound) Code() int {
	return 404
}

func (o *GetAzureKeyVaultSecretVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsNotFound", 404)
}

func (o *GetAzureKeyVaultSecretVersionsNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsNotFound", 404)
}

func (o *GetAzureKeyVaultSecretVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAzureKeyVaultSecretVersionsConflict creates a GetAzureKeyVaultSecretVersionsConflict with default headers values
func NewGetAzureKeyVaultSecretVersionsConflict() *GetAzureKeyVaultSecretVersionsConflict {
	return &GetAzureKeyVaultSecretVersionsConflict{}
}

/*
GetAzureKeyVaultSecretVersionsConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAzureKeyVaultSecretVersionsConflict struct {
}

// IsSuccess returns true when this get azure key vault secret versions conflict response has a 2xx status code
func (o *GetAzureKeyVaultSecretVersionsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get azure key vault secret versions conflict response has a 3xx status code
func (o *GetAzureKeyVaultSecretVersionsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure key vault secret versions conflict response has a 4xx status code
func (o *GetAzureKeyVaultSecretVersionsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get azure key vault secret versions conflict response has a 5xx status code
func (o *GetAzureKeyVaultSecretVersionsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure key vault secret versions conflict response a status code equal to that given
func (o *GetAzureKeyVaultSecretVersionsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get azure key vault secret versions conflict response
func (o *GetAzureKeyVaultSecretVersionsConflict) Code() int {
	return 409
}

func (o *GetAzureKeyVaultSecretVersionsConflict) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsConflict", 409)
}

func (o *GetAzureKeyVaultSecretVersionsConflict) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsConflict", 409)
}

func (o *GetAzureKeyVaultSecretVersionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAzureKeyVaultSecretVersionsInternalServerError creates a GetAzureKeyVaultSecretVersionsInternalServerError with default headers values
func NewGetAzureKeyVaultSecretVersionsInternalServerError() *GetAzureKeyVaultSecretVersionsInternalServerError {
	return &GetAzureKeyVaultSecretVersionsInternalServerError{}
}

/*
GetAzureKeyVaultSecretVersionsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAzureKeyVaultSecretVersionsInternalServerError struct {
}

// IsSuccess returns true when this get azure key vault secret versions internal server error response has a 2xx status code
func (o *GetAzureKeyVaultSecretVersionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get azure key vault secret versions internal server error response has a 3xx status code
func (o *GetAzureKeyVaultSecretVersionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure key vault secret versions internal server error response has a 4xx status code
func (o *GetAzureKeyVaultSecretVersionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get azure key vault secret versions internal server error response has a 5xx status code
func (o *GetAzureKeyVaultSecretVersionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get azure key vault secret versions internal server error response a status code equal to that given
func (o *GetAzureKeyVaultSecretVersionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get azure key vault secret versions internal server error response
func (o *GetAzureKeyVaultSecretVersionsInternalServerError) Code() int {
	return 500
}

func (o *GetAzureKeyVaultSecretVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsInternalServerError", 500)
}

func (o *GetAzureKeyVaultSecretVersionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}][%d] getAzureKeyVaultSecretVersionsInternalServerError", 500)
}

func (o *GetAzureKeyVaultSecretVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
