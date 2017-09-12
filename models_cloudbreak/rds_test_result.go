// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RdsTestResult rds test result
// swagger:model RdsTestResult
type RdsTestResult struct {

	// result of Ldap connection test
	// Required: true
	ConnectionResult *string `json:"connectionResult"`
}

// Validate validates this rds test result
func (m *RdsTestResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionResult(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RdsTestResult) validateConnectionResult(formats strfmt.Registry) error {

	if err := validate.Required("connectionResult", "body", m.ConnectionResult); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RdsTestResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RdsTestResult) UnmarshalBinary(b []byte) error {
	var res RdsTestResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
