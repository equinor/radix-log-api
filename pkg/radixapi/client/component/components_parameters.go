// Code generated by go-swagger; DO NOT EDIT.

package component

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

// NewComponentsParams creates a new ComponentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewComponentsParams() *ComponentsParams {
	return &ComponentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewComponentsParamsWithTimeout creates a new ComponentsParams object
// with the ability to set a timeout on a request.
func NewComponentsParamsWithTimeout(timeout time.Duration) *ComponentsParams {
	return &ComponentsParams{
		timeout: timeout,
	}
}

// NewComponentsParamsWithContext creates a new ComponentsParams object
// with the ability to set a context for a request.
func NewComponentsParamsWithContext(ctx context.Context) *ComponentsParams {
	return &ComponentsParams{
		Context: ctx,
	}
}

// NewComponentsParamsWithHTTPClient creates a new ComponentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewComponentsParamsWithHTTPClient(client *http.Client) *ComponentsParams {
	return &ComponentsParams{
		HTTPClient: client,
	}
}

/* ComponentsParams contains all the parameters to send to the API endpoint
   for the components operation.

   Typically these are written to a http.Request.
*/
type ComponentsParams struct {

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

	/* DeploymentName.

	   Name of deployment
	*/
	DeploymentName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ComponentsParams) WithDefaults() *ComponentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ComponentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the components params
func (o *ComponentsParams) WithTimeout(timeout time.Duration) *ComponentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the components params
func (o *ComponentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the components params
func (o *ComponentsParams) WithContext(ctx context.Context) *ComponentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the components params
func (o *ComponentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the components params
func (o *ComponentsParams) WithHTTPClient(client *http.Client) *ComponentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the components params
func (o *ComponentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the components params
func (o *ComponentsParams) WithImpersonateGroup(impersonateGroup *string) *ComponentsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the components params
func (o *ComponentsParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the components params
func (o *ComponentsParams) WithImpersonateUser(impersonateUser *string) *ComponentsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the components params
func (o *ComponentsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the components params
func (o *ComponentsParams) WithAppName(appName string) *ComponentsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the components params
func (o *ComponentsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithDeploymentName adds the deploymentName to the components params
func (o *ComponentsParams) WithDeploymentName(deploymentName string) *ComponentsParams {
	o.SetDeploymentName(deploymentName)
	return o
}

// SetDeploymentName adds the deploymentName to the components params
func (o *ComponentsParams) SetDeploymentName(deploymentName string) {
	o.DeploymentName = deploymentName
}

// WriteToRequest writes these params to a swagger request
func (o *ComponentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deploymentName
	if err := r.SetPathParam("deploymentName", o.DeploymentName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
