// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeploymentSummaryPipelineJobInfo deployment summary pipeline job info
//
// swagger:model DeploymentSummaryPipelineJobInfo
type DeploymentSummaryPipelineJobInfo struct {

	// CommitID the commit ID of the branch to build
	// Example: 4faca8595c5283a9d0f17a623b9255a0d9866a2e
	CommitID string `json:"commitID,omitempty"`

	// Name of job creating deployment
	CreatedByJob string `json:"createdByJob,omitempty"`

	// Type of pipeline job
	// Example: build-deploy
	PipelineJobType string `json:"pipelineJobType,omitempty"`

	// Name of the environment the deployment was promoted from
	// Applies only for pipeline jobs of type 'promote'
	// Example: qa
	PromotedFromEnvironment string `json:"promotedFromEnvironment,omitempty"`
}

// Validate validates this deployment summary pipeline job info
func (m *DeploymentSummaryPipelineJobInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this deployment summary pipeline job info based on context it is used
func (m *DeploymentSummaryPipelineJobInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentSummaryPipelineJobInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentSummaryPipelineJobInfo) UnmarshalBinary(b []byte) error {
	var res DeploymentSummaryPipelineJobInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
