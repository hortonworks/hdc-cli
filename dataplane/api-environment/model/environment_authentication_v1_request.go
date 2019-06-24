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

// EnvironmentAuthenticationV1Request environment authentication v1 request
// swagger:model EnvironmentAuthenticationV1Request
type EnvironmentAuthenticationV1Request struct {

	// User name created on the nodes for SSH access
	// Max Length: 32
	// Min Length: 1
	// Pattern: (^[a-z_]([a-z0-9_-]{0,31}|[a-z0-9_-]{0,30}\$)$)
	LoginUserName string `json:"loginUserName,omitempty"`

	// SSH Public key string.
	PublicKey string `json:"publicKey,omitempty"`

	// Public key ID registered at the cloud provider.
	PublicKeyID string `json:"publicKeyId,omitempty"`
}

// Validate validates this environment authentication v1 request
func (m *EnvironmentAuthenticationV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoginUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentAuthenticationV1Request) validateLoginUserName(formats strfmt.Registry) error {

	if swag.IsZero(m.LoginUserName) { // not required
		return nil
	}

	if err := validate.MinLength("loginUserName", "body", string(m.LoginUserName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("loginUserName", "body", string(m.LoginUserName), 32); err != nil {
		return err
	}

	if err := validate.Pattern("loginUserName", "body", string(m.LoginUserName), `(^[a-z_]([a-z0-9_-]{0,31}|[a-z0-9_-]{0,30}\$)$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentAuthenticationV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentAuthenticationV1Request) UnmarshalBinary(b []byte) error {
	var res EnvironmentAuthenticationV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
