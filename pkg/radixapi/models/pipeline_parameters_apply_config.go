// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PipelineParametersApplyConfig PipelineParametersApplyConfig describes base info
//
// swagger:model PipelineParametersApplyConfig
type PipelineParametersApplyConfig struct {

	// TriggeredBy of the job - if empty will use user token upn (user principle name)
	// Example: a_user@equinor.com
	TriggeredBy string `json:"triggeredBy,omitempty"`
}

// Validate validates this pipeline parameters apply config
func (m *PipelineParametersApplyConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pipeline parameters apply config based on context it is used
func (m *PipelineParametersApplyConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineParametersApplyConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineParametersApplyConfig) UnmarshalBinary(b []byte) error {
	var res PipelineParametersApplyConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}