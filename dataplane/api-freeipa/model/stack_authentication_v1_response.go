// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// StackAuthenticationV1Response stack authentication v1 response
// swagger:model StackAuthenticationV1Response
type StackAuthenticationV1Response struct {

	// authentication name for machines
	LoginUserName string `json:"loginUserName,omitempty"`

	// public key for accessing instances
	PublicKey string `json:"publicKey,omitempty"`

	// public key id for accessing instances
	PublicKeyID string `json:"publicKeyId,omitempty"`
}

// Validate validates this stack authentication v1 response
func (m *StackAuthenticationV1Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StackAuthenticationV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackAuthenticationV1Response) UnmarshalBinary(b []byte) error {
	var res StackAuthenticationV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}