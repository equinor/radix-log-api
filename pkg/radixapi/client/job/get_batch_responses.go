// Code generated by go-swagger; DO NOT EDIT.

package job

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

// GetBatchReader is a Reader for the GetBatch structure.
type GetBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}] getBatch", response, response.Code())
	}
}

// NewGetBatchOK creates a GetBatchOK with default headers values
func NewGetBatchOK() *GetBatchOK {
	return &GetBatchOK{}
}

/*
GetBatchOK describes a response with status code 200, with default header values.

scheduled batch
*/
type GetBatchOK struct {
	Payload *models.ScheduledBatchSummary
}

// IsSuccess returns true when this get batch o k response has a 2xx status code
func (o *GetBatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get batch o k response has a 3xx status code
func (o *GetBatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get batch o k response has a 4xx status code
func (o *GetBatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get batch o k response has a 5xx status code
func (o *GetBatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get batch o k response a status code equal to that given
func (o *GetBatchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get batch o k response
func (o *GetBatchOK) Code() int {
	return 200
}

func (o *GetBatchOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}][%d] getBatchOK %s", 200, payload)
}

func (o *GetBatchOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}][%d] getBatchOK %s", 200, payload)
}

func (o *GetBatchOK) GetPayload() *models.ScheduledBatchSummary {
	return o.Payload
}

func (o *GetBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduledBatchSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBatchNotFound creates a GetBatchNotFound with default headers values
func NewGetBatchNotFound() *GetBatchNotFound {
	return &GetBatchNotFound{}
}

/*
GetBatchNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetBatchNotFound struct {
}

// IsSuccess returns true when this get batch not found response has a 2xx status code
func (o *GetBatchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get batch not found response has a 3xx status code
func (o *GetBatchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get batch not found response has a 4xx status code
func (o *GetBatchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get batch not found response has a 5xx status code
func (o *GetBatchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get batch not found response a status code equal to that given
func (o *GetBatchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get batch not found response
func (o *GetBatchNotFound) Code() int {
	return 404
}

func (o *GetBatchNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}][%d] getBatchNotFound", 404)
}

func (o *GetBatchNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}][%d] getBatchNotFound", 404)
}

func (o *GetBatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
