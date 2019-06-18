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

// SynchronizeUsersV1Request synchronize users v1 request
// swagger:model SynchronizeUsersV1Request
type SynchronizeUsersV1Request struct {

	// Optional environments to sync
	// Unique: true
	Environments []string `json:"environments"`

	// Optional users to sync
	// Unique: true
	Users []string `json:"users"`
}

// Validate validates this synchronize users v1 request
func (m *SynchronizeUsersV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SynchronizeUsersV1Request) validateEnvironments(formats strfmt.Registry) error {

	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	if err := validate.UniqueItems("environments", "body", m.Environments); err != nil {
		return err
	}

	return nil
}

func (m *SynchronizeUsersV1Request) validateUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.Users) { // not required
		return nil
	}

	if err := validate.UniqueItems("users", "body", m.Users); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SynchronizeUsersV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SynchronizeUsersV1Request) UnmarshalBinary(b []byte) error {
	var res SynchronizeUsersV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}