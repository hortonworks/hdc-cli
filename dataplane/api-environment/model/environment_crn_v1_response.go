// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EnvironmentCrnV1Response environment crn v1 response
// swagger:model EnvironmentCrnV1Response
type EnvironmentCrnV1Response struct {

	// environment crn
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// environment name
	EnvironmentName string `json:"environmentName,omitempty"`
}

// Validate validates this environment crn v1 response
func (m *EnvironmentCrnV1Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentCrnV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentCrnV1Response) UnmarshalBinary(b []byte) error {
	var res EnvironmentCrnV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}