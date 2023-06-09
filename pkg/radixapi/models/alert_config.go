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

// AlertConfig AlertConfig defines a mapping between a pre-defined alert name and a receiver
//
// swagger:model AlertConfig
type AlertConfig struct {

	// Alert defines the name of a predefined alert
	// Required: true
	Alert *string `json:"alert"`

	// Receiver is the name of the receiver that will handle this alert
	// Required: true
	Receiver *string `json:"receiver"`
}

// Validate validates this alert config
func (m *AlertConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceiver(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertConfig) validateAlert(formats strfmt.Registry) error {

	if err := validate.Required("alert", "body", m.Alert); err != nil {
		return err
	}

	return nil
}

func (m *AlertConfig) validateReceiver(formats strfmt.Registry) error {

	if err := validate.Required("receiver", "body", m.Receiver); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this alert config based on context it is used
func (m *AlertConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertConfig) UnmarshalBinary(b []byte) error {
	var res AlertConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
