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

// StackAuthenticationV1Request stack authentication v1 request
// swagger:model StackAuthenticationV1Request
type StackAuthenticationV1Request struct {

	// authentication name for machines
	// Max Length: 32
	// Min Length: 0
	LoginUserName *string `json:"loginUserName,omitempty"`

	// public key for accessing instances
	// Max Length: 2048
	// Min Length: 0
	PublicKey *string `json:"publicKey,omitempty"`

	// public key id for accessing instances
	// Max Length: 255
	// Min Length: 0
	PublicKeyID *string `json:"publicKeyId,omitempty"`
}

// Validate validates this stack authentication v1 request
func (m *StackAuthenticationV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoginUserName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKeyID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackAuthenticationV1Request) validateLoginUserName(formats strfmt.Registry) error {

	if swag.IsZero(m.LoginUserName) { // not required
		return nil
	}

	if err := validate.MinLength("loginUserName", "body", string(*m.LoginUserName), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("loginUserName", "body", string(*m.LoginUserName), 32); err != nil {
		return err
	}

	return nil
}

func (m *StackAuthenticationV1Request) validatePublicKey(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicKey) { // not required
		return nil
	}

	if err := validate.MinLength("publicKey", "body", string(*m.PublicKey), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("publicKey", "body", string(*m.PublicKey), 2048); err != nil {
		return err
	}

	return nil
}

func (m *StackAuthenticationV1Request) validatePublicKeyID(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicKeyID) { // not required
		return nil
	}

	if err := validate.MinLength("publicKeyId", "body", string(*m.PublicKeyID), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("publicKeyId", "body", string(*m.PublicKeyID), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackAuthenticationV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackAuthenticationV1Request) UnmarshalBinary(b []byte) error {
	var res StackAuthenticationV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
