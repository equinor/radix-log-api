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

	"github.com/equinor/radix-log-api/pkg/radixapi/models"
)

// NewTriggerPipelineBuildParams creates a new TriggerPipelineBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTriggerPipelineBuildParams() *TriggerPipelineBuildParams {
	return &TriggerPipelineBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTriggerPipelineBuildParamsWithTimeout creates a new TriggerPipelineBuildParams object
// with the ability to set a timeout on a request.
func NewTriggerPipelineBuildParamsWithTimeout(timeout time.Duration) *TriggerPipelineBuildParams {
	return &TriggerPipelineBuildParams{
		timeout: timeout,
	}
}

// NewTriggerPipelineBuildParamsWithContext creates a new TriggerPipelineBuildParams object
// with the ability to set a context for a request.
func NewTriggerPipelineBuildParamsWithContext(ctx context.Context) *TriggerPipelineBuildParams {
	return &TriggerPipelineBuildParams{
		Context: ctx,
	}
}

// NewTriggerPipelineBuildParamsWithHTTPClient creates a new TriggerPipelineBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewTriggerPipelineBuildParamsWithHTTPClient(client *http.Client) *TriggerPipelineBuildParams {
	return &TriggerPipelineBuildParams{
		HTTPClient: client,
	}
}

/*
TriggerPipelineBuildParams contains all the parameters to send to the API endpoint

	for the trigger pipeline build operation.

	Typically these are written to a http.Request.
*/
type TriggerPipelineBuildParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* PipelineParametersBuild.

	   Pipeline parameters
	*/
	PipelineParametersBuild *models.PipelineParametersBuild

	/* AppName.

	   Name of application
	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the trigger pipeline build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerPipelineBuildParams) WithDefaults() *TriggerPipelineBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the trigger pipeline build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerPipelineBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) WithTimeout(timeout time.Duration) *TriggerPipelineBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) WithContext(ctx context.Context) *TriggerPipelineBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) WithHTTPClient(client *http.Client) *TriggerPipelineBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) WithImpersonateGroup(impersonateGroup *string) *TriggerPipelineBuildParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) WithImpersonateUser(impersonateUser *string) *TriggerPipelineBuildParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithPipelineParametersBuild adds the pipelineParametersBuild to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) WithPipelineParametersBuild(pipelineParametersBuild *models.PipelineParametersBuild) *TriggerPipelineBuildParams {
	o.SetPipelineParametersBuild(pipelineParametersBuild)
	return o
}

// SetPipelineParametersBuild adds the pipelineParametersBuild to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) SetPipelineParametersBuild(pipelineParametersBuild *models.PipelineParametersBuild) {
	o.PipelineParametersBuild = pipelineParametersBuild
}

// WithAppName adds the appName to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) WithAppName(appName string) *TriggerPipelineBuildParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the trigger pipeline build params
func (o *TriggerPipelineBuildParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *TriggerPipelineBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.PipelineParametersBuild != nil {
		if err := r.SetBodyParam(o.PipelineParametersBuild); err != nil {
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
