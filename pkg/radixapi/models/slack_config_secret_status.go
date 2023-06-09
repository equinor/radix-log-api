// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SlackConfigSecretStatus SlackConfigSecretStatus
//
// swagger:model SlackConfigSecretStatus
type SlackConfigSecretStatus struct {

	// WebhookURLConfigured flag indicates if a Slack webhook URL is set
	WebhookURLConfigured bool `json:"webhookUrlConfigured,omitempty"`
}

// Validate validates this slack config secret status
func (m *SlackConfigSecretStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this slack config secret status based on context it is used
func (m *SlackConfigSecretStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SlackConfigSecretStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SlackConfigSecretStatus) UnmarshalBinary(b []byte) error {
	var res SlackConfigSecretStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
