// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CloudbreakDetailsV4Response cloudbreak details v4 response
// swagger:model CloudbreakDetailsV4Response
type CloudbreakDetailsV4Response struct {

	// version of the Cloudbreak that provisioned the stack
	Version string `json:"version,omitempty"`
}

// Validate validates this cloudbreak details v4 response
func (m *CloudbreakDetailsV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudbreakDetailsV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudbreakDetailsV4Response) UnmarshalBinary(b []byte) error {
	var res CloudbreakDetailsV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
