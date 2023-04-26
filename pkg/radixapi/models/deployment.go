// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Deployment Deployment describe an deployment
//
// swagger:model Deployment
type Deployment struct {

	// ActiveFrom Timestamp when the deployment starts (or created)
	// Example: 2006-01-02T15:04:05Z
	ActiveFrom string `json:"activeFrom,omitempty"`

	// ActiveTo Timestamp when the deployment ends
	// Example: 2006-01-02T15:04:05Z
	ActiveTo string `json:"activeTo,omitempty"`

	// Array of components
	Components []*Component `json:"components"`

	// Name of job creating deployment
	CreatedByJob string `json:"createdByJob,omitempty"`

	// Environment the environment this Radix application deployment runs in
	// Example: prod
	Environment string `json:"environment,omitempty"`

	// GitCommitHash the hash of the git commit from which radixconfig.yaml was parsed
	// Example: 4faca8595c5283a9d0f17a623b9255a0d9866a2e
	GitCommitHash string `json:"gitCommitHash,omitempty"`

	// GitTags the git tags that the git commit hash points to
	// Example: \"v1.22.1 v1.22.3\
	GitTags string `json:"gitTags,omitempty"`

	// Name the unique name of the Radix application deployment
	// Example: radix-canary-golang-tzbqi
	Name string `json:"name,omitempty"`

	// Namespace where the deployment is stored
	// Example: radix-canary-golang-dev
	// Required: true
	Namespace *string `json:"namespace"`

	// Repository the GitHub repository that the deployment was built from
	// Example: https://github.com/equinor/radix-canary-golang
	// Required: true
	Repository *string `json:"repository"`
}

// Validate validates this deployment
func (m *Deployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Deployment) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Deployment) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *Deployment) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", m.Repository); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this deployment based on the context it is used
func (m *Deployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Deployment) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Components); i++ {

		if m.Components[i] != nil {
			if err := m.Components[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Deployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Deployment) UnmarshalBinary(b []byte) error {
	var res Deployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
