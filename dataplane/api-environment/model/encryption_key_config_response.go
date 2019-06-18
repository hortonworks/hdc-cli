// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EncryptionKeyConfigResponse encryption key config response
// swagger:model EncryptionKeyConfigResponse
type EncryptionKeyConfigResponse struct {

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// properties
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// Validate validates this encryption key config response
func (m *EncryptionKeyConfigResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionKeyConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionKeyConfigResponse) UnmarshalBinary(b []byte) error {
	var res EncryptionKeyConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}