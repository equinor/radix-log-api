// Code generated by go-swagger; DO NOT EDIT.

package pipeline_job

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
)

// NewGetPipelineJobStepLogsParams creates a new GetPipelineJobStepLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPipelineJobStepLogsParams() *GetPipelineJobStepLogsParams {
	return &GetPipelineJobStepLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPipelineJobStepLogsParamsWithTimeout creates a new GetPipelineJobStepLogsParams object
// with the ability to set a timeout on a request.
func NewGetPipelineJobStepLogsParamsWithTimeout(timeout time.Duration) *GetPipelineJobStepLogsParams {
	return &GetPipelineJobStepLogsParams{
		timeout: timeout,
	}
}

// NewGetPipelineJobStepLogsParamsWithContext creates a new GetPipelineJobStepLogsParams object
// with the ability to set a context for a request.
func NewGetPipelineJobStepLogsParamsWithContext(ctx context.Context) *GetPipelineJobStepLogsParams {
	return &GetPipelineJobStepLogsParams{
		Context: ctx,
	}
}

// NewGetPipelineJobStepLogsParamsWithHTTPClient creates a new GetPipelineJobStepLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPipelineJobStepLogsParamsWithHTTPClient(client *http.Client) *GetPipelineJobStepLogsParams {
	return &GetPipelineJobStepLogsParams{
		HTTPClient: client,
	}
}

/*
GetPipelineJobStepLogsParams contains all the parameters to send to the API endpoint

	for the get pipeline job step logs operation.

	Typically these are written to a http.Request.
*/
type GetPipelineJobStepLogsParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup []string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   name of Radix application
	*/
	AppName string

	/* File.

	   Get log as a file if true

	   Format: boolean
	*/
	File *string

	/* JobName.

	   Name of the pipeline job
	*/
	JobName string

	/* Lines.

	   Get log lines (example 1000)

	   Format: number
	*/
	Lines *string

	/* SinceTime.

	   Get log only from sinceTime (example 2020-03-18T07:20:41+00:00)

	   Format: date-time
	*/
	SinceTime *strfmt.DateTime

	/* StepName.

	   Name of the pipeline job step
	*/
	StepName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get pipeline job step logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPipelineJobStepLogsParams) WithDefaults() *GetPipelineJobStepLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get pipeline job step logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPipelineJobStepLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) WithTimeout(timeout time.Duration) *GetPipelineJobStepLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) WithContext(ctx context.Context) *GetPipelineJobStepLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) WithHTTPClient(client *http.Client) *GetPipelineJobStepLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) WithImpersonateGroup(impersonateGroup []string) *GetPipelineJobStepLogsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) WithImpersonateUser(impersonateUser *string) *GetPipelineJobStepLogsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) WithAppName(appName string) *GetPipelineJobStepLogsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithFile adds the file to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) WithFile(file *string) *GetPipelineJobStepLogsParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) SetFile(file *string) {
	o.File = file
}

// WithJobName adds the jobName to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) WithJobName(jobName string) *GetPipelineJobStepLogsParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WithLines adds the lines to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) WithLines(lines *string) *GetPipelineJobStepLogsParams {
	o.SetLines(lines)
	return o
}

// SetLines adds the lines to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) SetLines(lines *string) {
	o.Lines = lines
}

// WithSinceTime adds the sinceTime to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) WithSinceTime(sinceTime *strfmt.DateTime) *GetPipelineJobStepLogsParams {
	o.SetSinceTime(sinceTime)
	return o
}

// SetSinceTime adds the sinceTime to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) SetSinceTime(sinceTime *strfmt.DateTime) {
	o.SinceTime = sinceTime
}

// WithStepName adds the stepName to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) WithStepName(stepName string) *GetPipelineJobStepLogsParams {
	o.SetStepName(stepName)
	return o
}

// SetStepName adds the stepName to the get pipeline job step logs params
func (o *GetPipelineJobStepLogsParams) SetStepName(stepName string) {
	o.StepName = stepName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPipelineJobStepLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.File != nil {

		// query param file
		var qrFile string

		if o.File != nil {
			qrFile = *o.File
		}
		qFile := qrFile
		if qFile != "" {

			if err := r.SetQueryParam("file", qFile); err != nil {
				return err
			}
		}
	}

	// path param jobName
	if err := r.SetPathParam("jobName", o.JobName); err != nil {
		return err
	}

	if o.Lines != nil {

		// query param lines
		var qrLines string

		if o.Lines != nil {
			qrLines = *o.Lines
		}
		qLines := qrLines
		if qLines != "" {

			if err := r.SetQueryParam("lines", qLines); err != nil {
				return err
			}
		}
	}

	if o.SinceTime != nil {

		// query param sinceTime
		var qrSinceTime strfmt.DateTime

		if o.SinceTime != nil {
			qrSinceTime = *o.SinceTime
		}
		qSinceTime := qrSinceTime.String()
		if qSinceTime != "" {

			if err := r.SetQueryParam("sinceTime", qSinceTime); err != nil {
				return err
			}
		}
	}

	// path param stepName
	if err := r.SetPathParam("stepName", o.StepName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetPipelineJobStepLogs binds the parameter Impersonate-Group
func (o *GetPipelineJobStepLogsParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
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
