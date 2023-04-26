// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewDisableApplicationAlertingParams creates a new DisableApplicationAlertingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisableApplicationAlertingParams() *DisableApplicationAlertingParams {
	return &DisableApplicationAlertingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisableApplicationAlertingParamsWithTimeout creates a new DisableApplicationAlertingParams object
// with the ability to set a timeout on a request.
func NewDisableApplicationAlertingParamsWithTimeout(timeout time.Duration) *DisableApplicationAlertingParams {
	return &DisableApplicationAlertingParams{
		timeout: timeout,
	}
}

// NewDisableApplicationAlertingParamsWithContext creates a new DisableApplicationAlertingParams object
// with the ability to set a context for a request.
func NewDisableApplicationAlertingParamsWithContext(ctx context.Context) *DisableApplicationAlertingParams {
	return &DisableApplicationAlertingParams{
		Context: ctx,
	}
}

// NewDisableApplicationAlertingParamsWithHTTPClient creates a new DisableApplicationAlertingParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisableApplicationAlertingParamsWithHTTPClient(client *http.Client) *DisableApplicationAlertingParams {
	return &DisableApplicationAlertingParams{
		HTTPClient: client,
	}
}

/* DisableApplicationAlertingParams contains all the parameters to send to the API endpoint
   for the disable application alerting operation.

   Typically these are written to a http.Request.
*/
type DisableApplicationAlertingParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disable application alerting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableApplicationAlertingParams) WithDefaults() *DisableApplicationAlertingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disable application alerting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableApplicationAlertingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disable application alerting params
func (o *DisableApplicationAlertingParams) WithTimeout(timeout time.Duration) *DisableApplicationAlertingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable application alerting params
func (o *DisableApplicationAlertingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable application alerting params
func (o *DisableApplicationAlertingParams) WithContext(ctx context.Context) *DisableApplicationAlertingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable application alerting params
func (o *DisableApplicationAlertingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable application alerting params
func (o *DisableApplicationAlertingParams) WithHTTPClient(client *http.Client) *DisableApplicationAlertingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable application alerting params
func (o *DisableApplicationAlertingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the disable application alerting params
func (o *DisableApplicationAlertingParams) WithImpersonateGroup(impersonateGroup *string) *DisableApplicationAlertingParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the disable application alerting params
func (o *DisableApplicationAlertingParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the disable application alerting params
func (o *DisableApplicationAlertingParams) WithImpersonateUser(impersonateUser *string) *DisableApplicationAlertingParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the disable application alerting params
func (o *DisableApplicationAlertingParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the disable application alerting params
func (o *DisableApplicationAlertingParams) WithAppName(appName string) *DisableApplicationAlertingParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the disable application alerting params
func (o *DisableApplicationAlertingParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *DisableApplicationAlertingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
