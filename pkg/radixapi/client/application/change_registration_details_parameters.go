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
	"github.com/go-openapi/swag"

	"github.com/equinor/radix-log-api/pkg/radixapi/models"
)

// NewChangeRegistrationDetailsParams creates a new ChangeRegistrationDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeRegistrationDetailsParams() *ChangeRegistrationDetailsParams {
	return &ChangeRegistrationDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeRegistrationDetailsParamsWithTimeout creates a new ChangeRegistrationDetailsParams object
// with the ability to set a timeout on a request.
func NewChangeRegistrationDetailsParamsWithTimeout(timeout time.Duration) *ChangeRegistrationDetailsParams {
	return &ChangeRegistrationDetailsParams{
		timeout: timeout,
	}
}

// NewChangeRegistrationDetailsParamsWithContext creates a new ChangeRegistrationDetailsParams object
// with the ability to set a context for a request.
func NewChangeRegistrationDetailsParamsWithContext(ctx context.Context) *ChangeRegistrationDetailsParams {
	return &ChangeRegistrationDetailsParams{
		Context: ctx,
	}
}

// NewChangeRegistrationDetailsParamsWithHTTPClient creates a new ChangeRegistrationDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeRegistrationDetailsParamsWithHTTPClient(client *http.Client) *ChangeRegistrationDetailsParams {
	return &ChangeRegistrationDetailsParams{
		HTTPClient: client,
	}
}

/*
ChangeRegistrationDetailsParams contains all the parameters to send to the API endpoint

	for the change registration details operation.

	Typically these are written to a http.Request.
*/
type ChangeRegistrationDetailsParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup []string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   Name of application
	*/
	AppName string

	/* ApplicationRegistration.

	   request for Application to change
	*/
	ApplicationRegistration *models.ApplicationRegistrationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change registration details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeRegistrationDetailsParams) WithDefaults() *ChangeRegistrationDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change registration details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeRegistrationDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change registration details params
func (o *ChangeRegistrationDetailsParams) WithTimeout(timeout time.Duration) *ChangeRegistrationDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change registration details params
func (o *ChangeRegistrationDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change registration details params
func (o *ChangeRegistrationDetailsParams) WithContext(ctx context.Context) *ChangeRegistrationDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change registration details params
func (o *ChangeRegistrationDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change registration details params
func (o *ChangeRegistrationDetailsParams) WithHTTPClient(client *http.Client) *ChangeRegistrationDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change registration details params
func (o *ChangeRegistrationDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the change registration details params
func (o *ChangeRegistrationDetailsParams) WithImpersonateGroup(impersonateGroup []string) *ChangeRegistrationDetailsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the change registration details params
func (o *ChangeRegistrationDetailsParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the change registration details params
func (o *ChangeRegistrationDetailsParams) WithImpersonateUser(impersonateUser *string) *ChangeRegistrationDetailsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the change registration details params
func (o *ChangeRegistrationDetailsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the change registration details params
func (o *ChangeRegistrationDetailsParams) WithAppName(appName string) *ChangeRegistrationDetailsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the change registration details params
func (o *ChangeRegistrationDetailsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithApplicationRegistration adds the applicationRegistration to the change registration details params
func (o *ChangeRegistrationDetailsParams) WithApplicationRegistration(applicationRegistration *models.ApplicationRegistrationRequest) *ChangeRegistrationDetailsParams {
	o.SetApplicationRegistration(applicationRegistration)
	return o
}

// SetApplicationRegistration adds the applicationRegistration to the change registration details params
func (o *ChangeRegistrationDetailsParams) SetApplicationRegistration(applicationRegistration *models.ApplicationRegistrationRequest) {
	o.ApplicationRegistration = applicationRegistration
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeRegistrationDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImpersonateGroup != nil {

		// binding items for Impersonate-Group
		joinedImpersonateGroup := o.bindParamImpersonateGroup(reg)

		// header array param Impersonate-Group
		if len(joinedImpersonateGroup) > 0 {
			if err := r.SetHeaderParam("Impersonate-Group", joinedImpersonateGroup[0]); err != nil {
				return err
			}
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
	if o.ApplicationRegistration != nil {
		if err := r.SetBodyParam(o.ApplicationRegistration); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamChangeRegistrationDetails binds the parameter Impersonate-Group
func (o *ChangeRegistrationDetailsParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
	impersonateGroupIR := o.ImpersonateGroup

	var impersonateGroupIC []string
	for _, impersonateGroupIIR := range impersonateGroupIR { // explode []string

		impersonateGroupIIV := impersonateGroupIIR // string as string
		impersonateGroupIC = append(impersonateGroupIC, impersonateGroupIIV)
	}

	// items.CollectionFormat: ""
	impersonateGroupIS := swag.JoinByFormat(impersonateGroupIC, "")

	return impersonateGroupIS
}
