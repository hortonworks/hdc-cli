// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateKerberosConfigV1Request create kerberos config v1 request
// swagger:model CreateKerberosConfigV1Request
type CreateKerberosConfigV1Request struct {

	// active directory
	ActiveDirectory *ActiveDirectoryKerberosV1Descriptor `json:"activeDirectory,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// The crn of the environment
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// free ipa
	FreeIpa *FreeIPAKerberosV1Descriptor `json:"freeIpa,omitempty"`

	// mit
	Mit *MITKerberosV1Descriptor `json:"mit,omitempty"`

	// the name of the kerberos configuration
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this create kerberos config v1 request
func (m *CreateKerberosConfigV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveDirectory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeIpa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateKerberosConfigV1Request) validateActiveDirectory(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveDirectory) { // not required
		return nil
	}

	if m.ActiveDirectory != nil {
		if err := m.ActiveDirectory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activeDirectory")
			}
			return err
		}
	}

	return nil
}

func (m *CreateKerberosConfigV1Request) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *CreateKerberosConfigV1Request) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *CreateKerberosConfigV1Request) validateFreeIpa(formats strfmt.Registry) error {

	if swag.IsZero(m.FreeIpa) { // not required
		return nil
	}

	if m.FreeIpa != nil {
		if err := m.FreeIpa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeIpa")
			}
			return err
		}
	}

	return nil
}

func (m *CreateKerberosConfigV1Request) validateMit(formats strfmt.Registry) error {

	if swag.IsZero(m.Mit) { // not required
		return nil
	}

	if m.Mit != nil {
		if err := m.Mit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mit")
			}
			return err
		}
	}

	return nil
}

func (m *CreateKerberosConfigV1Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateKerberosConfigV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateKerberosConfigV1Request) UnmarshalBinary(b []byte) error {
	var res CreateKerberosConfigV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
