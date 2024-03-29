// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DNSExternalAlias DNSExternalAlias holds public external DNS alias information
//
// swagger:model DNSExternalAlias
type DNSExternalAlias struct {

	// ComponentName the component exposing the endpoint
	// Example: frontend
	// Required: true
	ComponentName *string `json:"componentName"`

	// EnvironmentName the environment hosting the endpoint
	// Example: prod
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// URL the public endpoint
	// Example: https://my-app.equinor.com
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this DNS external alias
func (m *DNSExternalAlias) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSExternalAlias) validateComponentName(formats strfmt.Registry) error {

	if err := validate.Required("componentName", "body", m.ComponentName); err != nil {
		return err
	}

	return nil
}

func (m *DNSExternalAlias) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *DNSExternalAlias) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this DNS external alias based on context it is used
func (m *DNSExternalAlias) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSExternalAlias) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSExternalAlias) UnmarshalBinary(b []byte) error {
	var res DNSExternalAlias
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
