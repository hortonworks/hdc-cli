// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AccountMappingBase account mapping base
// swagger:model AccountMappingBase
type AccountMappingBase struct {

	// group mappings
	GroupMappings map[string]string `json:"groupMappings,omitempty"`

	// user mappings
	UserMappings map[string]string `json:"userMappings,omitempty"`
}

// Validate validates this account mapping base
func (m *AccountMappingBase) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountMappingBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountMappingBase) UnmarshalBinary(b []byte) error {
	var res AccountMappingBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}