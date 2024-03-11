// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Step Step holds general information about job step
//
// swagger:model Step
type Step struct {

	// Components associated components
	Components []string `json:"components"`

	// Ended timestamp
	// Format: date-time
	Ended strfmt.DateTime `json:"ended,omitempty"`

	// Name of the step
	// Example: build
	Name string `json:"name,omitempty"`

	// Started timestamp
	// Format: date-time
	Started strfmt.DateTime `json:"started,omitempty"`

	// Status of the step
	// Example: Waiting
	// Enum: [Queued Waiting Running Succeeded Failed Stopped StoppedNoChanges]
	Status string `json:"status,omitempty"`
}

// Validate validates this step
func (m *Step) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Step) validateEnded(formats strfmt.Registry) error {
	if swag.IsZero(m.Ended) { // not required
		return nil
	}

	if err := validate.FormatOf("ended", "body", "date-time", m.Ended.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Step) validateStarted(formats strfmt.Registry) error {
	if swag.IsZero(m.Started) { // not required
		return nil
	}

	if err := validate.FormatOf("started", "body", "date-time", m.Started.String(), formats); err != nil {
		return err
	}

	return nil
}

var stepTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Queued","Waiting","Running","Succeeded","Failed","Stopped","StoppedNoChanges"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stepTypeStatusPropEnum = append(stepTypeStatusPropEnum, v)
	}
}

const (

	// StepStatusQueued captures enum value "Queued"
	StepStatusQueued string = "Queued"

	// StepStatusWaiting captures enum value "Waiting"
	StepStatusWaiting string = "Waiting"

	// StepStatusRunning captures enum value "Running"
	StepStatusRunning string = "Running"

	// StepStatusSucceeded captures enum value "Succeeded"
	StepStatusSucceeded string = "Succeeded"

	// StepStatusFailed captures enum value "Failed"
	StepStatusFailed string = "Failed"

	// StepStatusStopped captures enum value "Stopped"
	StepStatusStopped string = "Stopped"

	// StepStatusStoppedNoChanges captures enum value "StoppedNoChanges"
	StepStatusStoppedNoChanges string = "StoppedNoChanges"
)

// prop value enum
func (m *Step) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stepTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Step) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this step based on context it is used
func (m *Step) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Step) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Step) UnmarshalBinary(b []byte) error {
	var res Step
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
