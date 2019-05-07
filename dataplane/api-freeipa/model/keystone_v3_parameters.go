// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// KeystoneV3Parameters keystone v3 parameters
// swagger:model KeystoneV3Parameters
type KeystoneV3Parameters struct {

	// domain
	Domain *DomainKeystoneV3Parameters `json:"domain,omitempty"`

	// project
	Project *ProjectKeystoneV3Parameters `json:"project,omitempty"`
}

// Validate validates this keystone v3 parameters
func (m *KeystoneV3Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeystoneV3Parameters) validateDomain(formats strfmt.Registry) error {

	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if m.Domain != nil {
		if err := m.Domain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

func (m *KeystoneV3Parameters) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeystoneV3Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeystoneV3Parameters) UnmarshalBinary(b []byte) error {
	var res KeystoneV3Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
