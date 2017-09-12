// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigsResponse configs response
// swagger:model ConfigsResponse
type ConfigsResponse struct {

	// response object
	// Required: true
	// Unique: true
	Inputs []*BlueprintInput `json:"inputs"`
}

// Validate validates this configs response
func (m *ConfigsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigsResponse) validateInputs(formats strfmt.Registry) error {

	if err := validate.Required("inputs", "body", m.Inputs); err != nil {
		return err
	}

	if err := validate.UniqueItems("inputs", "body", m.Inputs); err != nil {
		return err
	}

	for i := 0; i < len(m.Inputs); i++ {

		if swag.IsZero(m.Inputs[i]) { // not required
			continue
		}

		if m.Inputs[i] != nil {

			if err := m.Inputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigsResponse) UnmarshalBinary(b []byte) error {
	var res ConfigsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
