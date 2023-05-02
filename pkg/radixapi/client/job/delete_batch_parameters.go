// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteBatchParams creates a new DeleteBatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBatchParams() *DeleteBatchParams {
	return &DeleteBatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBatchParamsWithTimeout creates a new DeleteBatchParams object
// with the ability to set a timeout on a request.
func NewDeleteBatchParamsWithTimeout(timeout time.Duration) *DeleteBatchParams {
	return &DeleteBatchParams{
		timeout: timeout,
	}
}

// NewDeleteBatchParamsWithContext creates a new DeleteBatchParams object
// with the ability to set a context for a request.
func NewDeleteBatchParamsWithContext(ctx context.Context) *DeleteBatchParams {
	return &DeleteBatchParams{
		Context: ctx,
	}
}

// NewDeleteBatchParamsWithHTTPClient creates a new DeleteBatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBatchParamsWithHTTPClient(client *http.Client) *DeleteBatchParams {
	return &DeleteBatchParams{
		HTTPClient: client,
	}
}

/* DeleteBatchParams contains all the parameters to send to the API endpoint
   for the delete batch operation.

   Typically these are written to a http.Request.
*/
type DeleteBatchParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   Name of application
	*/
	AppName string

	/* BatchName.

	   Name of batch
	*/
	BatchName string

	/* EnvName.

	   Name of environment
	*/
	EnvName string

	/* JobComponentName.

	   Name of job-component
	*/
	JobComponentName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBatchParams) WithDefaults() *DeleteBatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete batch params
func (o *DeleteBatchParams) WithTimeout(timeout time.Duration) *DeleteBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete batch params
func (o *DeleteBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete batch params
func (o *DeleteBatchParams) WithContext(ctx context.Context) *DeleteBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete batch params
func (o *DeleteBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete batch params
func (o *DeleteBatchParams) WithHTTPClient(client *http.Client) *DeleteBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete batch params
func (o *DeleteBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the delete batch params
func (o *DeleteBatchParams) WithImpersonateGroup(impersonateGroup *string) *DeleteBatchParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the delete batch params
func (o *DeleteBatchParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the delete batch params
func (o *DeleteBatchParams) WithImpersonateUser(impersonateUser *string) *DeleteBatchParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the delete batch params
func (o *DeleteBatchParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the delete batch params
func (o *DeleteBatchParams) WithAppName(appName string) *DeleteBatchParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the delete batch params
func (o *DeleteBatchParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBatchName adds the batchName to the delete batch params
func (o *DeleteBatchParams) WithBatchName(batchName string) *DeleteBatchParams {
	o.SetBatchName(batchName)
	return o
}

// SetBatchName adds the batchName to the delete batch params
func (o *DeleteBatchParams) SetBatchName(batchName string) {
	o.BatchName = batchName
}

// WithEnvName adds the envName to the delete batch params
func (o *DeleteBatchParams) WithEnvName(envName string) *DeleteBatchParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the delete batch params
func (o *DeleteBatchParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WithJobComponentName adds the jobComponentName to the delete batch params
func (o *DeleteBatchParams) WithJobComponentName(jobComponentName string) *DeleteBatchParams {
	o.SetJobComponentName(jobComponentName)
	return o
}

// SetJobComponentName adds the jobComponentName to the delete batch params
func (o *DeleteBatchParams) SetJobComponentName(jobComponentName string) {
	o.JobComponentName = jobComponentName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImpersonateGroup != nil {

		// header param Impersonate-Group
		if err := r.SetHeaderParam("Impersonate-Group", *o.ImpersonateGroup); err != nil {
			return err
		}
	}

	if o.ImpersonateUser != nil {

		// header param Impersonate-User
		if err := r.SetHeaderParam("Impersonate-User", *o.ImpersonateUser); err != nil {
			return err
		}
	}

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	// path param batchName
	if err := r.SetPathParam("batchName", o.BatchName); err != nil {
		return err
	}

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	// path param jobComponentName
	if err := r.SetPathParam("jobComponentName", o.JobComponentName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}