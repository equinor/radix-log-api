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

// AzureKeyVaultSecretVersion AzureKeyVaultSecretVersion holds a version of a Azure Key vault secret
//
// swagger:model AzureKeyVaultSecretVersion
type AzureKeyVaultSecretVersion struct {

	// BatchCreated which uses the secret
	// Example: 2006-01-02T15:04:05Z
	BatchCreated string `json:"batchCreated,omitempty"`

	// BatchName which uses the secret
	// Example: batch-abc
	BatchName string `json:"batchName,omitempty"`

	// JobCreated which uses the secret
	// Example: 2006-01-02T15:04:05Z
	JobCreated string `json:"jobCreated,omitempty"`

	// JobName which uses the secret
	// Example: job-abc
	JobName string `json:"jobName,omitempty"`

	// ReplicaCreated which uses the secret
	// Example: 2006-01-02T15:04:05Z
	// Required: true
	ReplicaCreated *string `json:"replicaCreated"`

	// ReplicaName which uses the secret
	// Example: abcdf
	// Required: true
	ReplicaName *string `json:"replicaName"`

	// Version of the secret
	// Example: 0123456789
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this azure key vault secret version
func (m *AzureKeyVaultSecretVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReplicaCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureKeyVaultSecretVersion) validateReplicaCreated(formats strfmt.Registry) error {

	if err := validate.Required("replicaCreated", "body", m.ReplicaCreated); err != nil {
		return err
	}

	return nil
}

func (m *AzureKeyVaultSecretVersion) validateReplicaName(formats strfmt.Registry) error {

	if err := validate.Required("replicaName", "body", m.ReplicaName); err != nil {
		return err
	}

	return nil
}

func (m *AzureKeyVaultSecretVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this azure key vault secret version based on context it is used
func (m *AzureKeyVaultSecretVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureKeyVaultSecretVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureKeyVaultSecretVersion) UnmarshalBinary(b []byte) error {
	var res AzureKeyVaultSecretVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
