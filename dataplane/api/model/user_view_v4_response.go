// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UserViewV4Response user view v4 response
// swagger:model UserViewV4Response
type UserViewV4Response struct {

	// user Id
	UserID string `json:"userId,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this user view v4 response
func (m *UserViewV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserViewV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserViewV4Response) UnmarshalBinary(b []byte) error {
	var res UserViewV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}