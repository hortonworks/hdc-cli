// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AttachedClusterInfoV4Response attached cluster info v4 response
// swagger:model AttachedClusterInfoV4Response
type AttachedClusterInfoV4Response struct {

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// name of the resource
	Name string `json:"name,omitempty"`
}

// Validate validates this attached cluster info v4 response
func (m *AttachedClusterInfoV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AttachedClusterInfoV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttachedClusterInfoV4Response) UnmarshalBinary(b []byte) error {
	var res AttachedClusterInfoV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}