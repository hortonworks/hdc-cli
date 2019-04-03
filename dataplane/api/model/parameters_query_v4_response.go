// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ParametersQueryV4Response parameters query v4 response
// swagger:model ParametersQueryV4Response
type ParametersQueryV4Response struct {

	// AmbariKerberosDescriptor parameters as a json
	// Required: true
	Custom map[string]string `json:"custom"`
}

// Validate validates this parameters query v4 response
func (m *ParametersQueryV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParametersQueryV4Response) validateCustom(formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *ParametersQueryV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParametersQueryV4Response) UnmarshalBinary(b []byte) error {
	var res ParametersQueryV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}