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

	"github.com/equinor/radix-log-api/pkg/radixapi/models"
)

// NewChangeComponentSecretParams creates a new ChangeComponentSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeComponentSecretParams() *ChangeComponentSecretParams {
	return &ChangeComponentSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeComponentSecretParamsWithTimeout creates a new ChangeComponentSecretParams object
// with the ability to set a timeout on a request.
func NewChangeComponentSecretParamsWithTimeout(timeout time.Duration) *ChangeComponentSecretParams {
	return &ChangeComponentSecretParams{
		timeout: timeout,
	}
}

// NewChangeComponentSecretParamsWithContext creates a new ChangeComponentSecretParams object
// with the ability to set a context for a request.
func NewChangeComponentSecretParamsWithContext(ctx context.Context) *ChangeComponentSecretParams {
	return &ChangeComponentSecretParams{
		Context: ctx,
	}
}

// NewChangeComponentSecretParamsWithHTTPClient creates a new ChangeComponentSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeComponentSecretParamsWithHTTPClient(client *http.Client) *ChangeComponentSecretParams {
	return &ChangeComponentSecretParams{
		HTTPClient: client,
	}
}

/* ChangeComponentSecretParams contains all the parameters to send to the API endpoint
   for the change component secret operation.

   Typically these are written to a http.Request.
*/
type ChangeComponentSecretParams struct {

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

	/* ComponentName.

	   secret component of Radix application
	*/
	ComponentName string

	/* ComponentSecret.

	   New secret value
	*/
	ComponentSecret *models.SecretParameters

	/* EnvName.

	   secret of Radix application
	*/
	EnvName string

	/* SecretName.

	   environment component secret name to be updated
	*/
	SecretName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change component secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeComponentSecretParams) WithDefaults() *ChangeComponentSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change component secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeComponentSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change component secret params
func (o *ChangeComponentSecretParams) WithTimeout(timeout time.Duration) *ChangeComponentSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change component secret params
func (o *ChangeComponentSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change component secret params
func (o *ChangeComponentSecretParams) WithContext(ctx context.Context) *ChangeComponentSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change component secret params
func (o *ChangeComponentSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change component secret params
func (o *ChangeComponentSecretParams) WithHTTPClient(client *http.Client) *ChangeComponentSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change component secret params
func (o *ChangeComponentSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the change component secret params
func (o *ChangeComponentSecretParams) WithImpersonateGroup(impersonateGroup *string) *ChangeComponentSecretParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the change component secret params
func (o *ChangeComponentSecretParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the change component secret params
func (o *ChangeComponentSecretParams) WithImpersonateUser(impersonateUser *string) *ChangeComponentSecretParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the change component secret params
func (o *ChangeComponentSecretParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the change component secret params
func (o *ChangeComponentSecretParams) WithAppName(appName string) *ChangeComponentSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the change component secret params
func (o *ChangeComponentSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithComponentName adds the componentName to the change component secret params
func (o *ChangeComponentSecretParams) WithComponentName(componentName string) *ChangeComponentSecretParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the change component secret params
func (o *ChangeComponentSecretParams) SetComponentName(componentName string) {
	o.ComponentName = componentName
}

// WithComponentSecret adds the componentSecret to the change component secret params
func (o *ChangeComponentSecretParams) WithComponentSecret(componentSecret *models.SecretParameters) *ChangeComponentSecretParams {
	o.SetComponentSecret(componentSecret)
	return o
}

// SetComponentSecret adds the componentSecret to the change component secret params
func (o *ChangeComponentSecretParams) SetComponentSecret(componentSecret *models.SecretParameters) {
	o.ComponentSecret = componentSecret
}

// WithEnvName adds the envName to the change component secret params
func (o *ChangeComponentSecretParams) WithEnvName(envName string) *ChangeComponentSecretParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the change component secret params
func (o *ChangeComponentSecretParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WithSecretName adds the secretName to the change component secret params
func (o *ChangeComponentSecretParams) WithSecretName(secretName string) *ChangeComponentSecretParams {
	o.SetSecretName(secretName)
	return o
}

// SetSecretName adds the secretName to the change component secret params
func (o *ChangeComponentSecretParams) SetSecretName(secretName string) {
	o.SecretName = secretName
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeComponentSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param componentName
	if err := r.SetPathParam("componentName", o.ComponentName); err != nil {
		return err
	}
	if o.ComponentSecret != nil {
		if err := r.SetBodyParam(o.ComponentSecret); err != nil {
			return err
		}
	}

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	// path param secretName
	if err := r.SetPathParam("secretName", o.SecretName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
