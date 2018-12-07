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

// RoleWithPermission role with permission
// swagger:model RoleWithPermission

type RoleWithPermission struct {

	// display name
	// Required: true
	DisplayName *string `json:"display_name"`

	// id
	ID strfmt.UUID `json:"id,omitempty"`

	// is default
	IsDefault interface{} `json:"is_default,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// permissions
	// Required: true
	Permissions []string `json:"permissions"`

	// service
	Service string `json:"service,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

/* polymorph RoleWithPermission display_name false */

/* polymorph RoleWithPermission id false */

/* polymorph RoleWithPermission is_default false */

/* polymorph RoleWithPermission name false */

/* polymorph RoleWithPermission permissions false */

/* polymorph RoleWithPermission service false */

/* polymorph RoleWithPermission type false */

// Validate validates this role with permission
func (m *RoleWithPermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleWithPermission) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("display_name", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *RoleWithPermission) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RoleWithPermission) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoleWithPermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleWithPermission) UnmarshalBinary(b []byte) error {
	var res RoleWithPermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
