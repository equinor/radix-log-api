// Code generated by go-swagger; DO NOT EDIT.

package environment

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

// NewGetEnvironmentSummaryParams creates a new GetEnvironmentSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEnvironmentSummaryParams() *GetEnvironmentSummaryParams {
	return &GetEnvironmentSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnvironmentSummaryParamsWithTimeout creates a new GetEnvironmentSummaryParams object
// with the ability to set a timeout on a request.
func NewGetEnvironmentSummaryParamsWithTimeout(timeout time.Duration) *GetEnvironmentSummaryParams {
	return &GetEnvironmentSummaryParams{
		timeout: timeout,
	}
}

// NewGetEnvironmentSummaryParamsWithContext creates a new GetEnvironmentSummaryParams object
// with the ability to set a context for a request.
func NewGetEnvironmentSummaryParamsWithContext(ctx context.Context) *GetEnvironmentSummaryParams {
	return &GetEnvironmentSummaryParams{
		Context: ctx,
	}
}

// NewGetEnvironmentSummaryParamsWithHTTPClient creates a new GetEnvironmentSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEnvironmentSummaryParamsWithHTTPClient(client *http.Client) *GetEnvironmentSummaryParams {
	return &GetEnvironmentSummaryParams{
		HTTPClient: client,
	}
}

/*
GetEnvironmentSummaryParams contains all the parameters to send to the API endpoint

	for the get environment summary operation.

	Typically these are written to a http.Request.
*/
type GetEnvironmentSummaryParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   name of Radix application
	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get environment summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEnvironmentSummaryParams) WithDefaults() *GetEnvironmentSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get environment summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEnvironmentSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get environment summary params
func (o *GetEnvironmentSummaryParams) WithTimeout(timeout time.Duration) *GetEnvironmentSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get environment summary params
func (o *GetEnvironmentSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get environment summary params
func (o *GetEnvironmentSummaryParams) WithContext(ctx context.Context) *GetEnvironmentSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get environment summary params
func (o *GetEnvironmentSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get environment summary params
func (o *GetEnvironmentSummaryParams) WithHTTPClient(client *http.Client) *GetEnvironmentSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get environment summary params
func (o *GetEnvironmentSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get environment summary params
func (o *GetEnvironmentSummaryParams) WithImpersonateGroup(impersonateGroup *string) *GetEnvironmentSummaryParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get environment summary params
func (o *GetEnvironmentSummaryParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get environment summary params
func (o *GetEnvironmentSummaryParams) WithImpersonateUser(impersonateUser *string) *GetEnvironmentSummaryParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get environment summary params
func (o *GetEnvironmentSummaryParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get environment summary params
func (o *GetEnvironmentSummaryParams) WithAppName(appName string) *GetEnvironmentSummaryParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get environment summary params
func (o *GetEnvironmentSummaryParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnvironmentSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
