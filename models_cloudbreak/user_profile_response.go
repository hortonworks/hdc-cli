// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserProfileResponse user profile response
// swagger:model UserProfileResponse
type UserProfileResponse struct {

	// account
	Account string `json:"account,omitempty"`

	// credential
	Credential *CredentialResponse `json:"credential,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// ui properties
	UIProperties map[string]interface{} `json:"uiProperties,omitempty"`
}

// Validate validates this user profile response
func (m *UserProfileResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredential(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserProfileResponse) validateCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {

		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserProfileResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserProfileResponse) UnmarshalBinary(b []byte) error {
	var res UserProfileResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
