// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AmbariStackRepoDetailsV4Response ambari stack repo details v4 response
// swagger:model AmbariStackRepoDetailsV4Response
type AmbariStackRepoDetailsV4Response struct {

	// stack
	Stack map[string]string `json:"stack,omitempty"`

	// util
	Util map[string]string `json:"util,omitempty"`
}

// Validate validates this ambari stack repo details v4 response
func (m *AmbariStackRepoDetailsV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AmbariStackRepoDetailsV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmbariStackRepoDetailsV4Response) UnmarshalBinary(b []byte) error {
	var res AmbariStackRepoDetailsV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}