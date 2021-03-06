// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DiagnosticsCollection diagnostics collection
// swagger:model DiagnosticsCollection
type DiagnosticsCollection struct {

	// created
	Created int64 `json:"created,omitempty"`

	// current flow status
	CurrentFlowStatus string `json:"currentFlowStatus,omitempty"`

	// flow Id
	FlowID string `json:"flowId,omitempty"`

	// progress percentage
	ProgressPercentage int32 `json:"progressPercentage,omitempty"`

	// properties
	Properties map[string]interface{} `json:"properties,omitempty"`

	// status
	// Enum: [IN_PROGRESS FAILED CANCELLED FINISHED]
	Status string `json:"status,omitempty"`
}

// Validate validates this diagnostics collection
func (m *DiagnosticsCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var diagnosticsCollectionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","FAILED","CANCELLED","FINISHED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diagnosticsCollectionTypeStatusPropEnum = append(diagnosticsCollectionTypeStatusPropEnum, v)
	}
}

const (

	// DiagnosticsCollectionStatusINPROGRESS captures enum value "IN_PROGRESS"
	DiagnosticsCollectionStatusINPROGRESS string = "IN_PROGRESS"

	// DiagnosticsCollectionStatusFAILED captures enum value "FAILED"
	DiagnosticsCollectionStatusFAILED string = "FAILED"

	// DiagnosticsCollectionStatusCANCELLED captures enum value "CANCELLED"
	DiagnosticsCollectionStatusCANCELLED string = "CANCELLED"

	// DiagnosticsCollectionStatusFINISHED captures enum value "FINISHED"
	DiagnosticsCollectionStatusFINISHED string = "FINISHED"
)

// prop value enum
func (m *DiagnosticsCollection) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, diagnosticsCollectionTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DiagnosticsCollection) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiagnosticsCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiagnosticsCollection) UnmarshalBinary(b []byte) error {
	var res DiagnosticsCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
