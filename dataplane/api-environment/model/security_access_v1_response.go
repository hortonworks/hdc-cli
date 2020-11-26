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

// SecurityAccessV1Response security access v1 response
// swagger:model SecurityAccessV1Response
type SecurityAccessV1Response struct {

	// CIDR range which is allowed for inbound traffic. Either IPv4 or IPv6 is allowed.
	// Max Length: 4000
	// Min Length: 5
	Cidr string `json:"cidr,omitempty"`

	// Security group where all other hosts are placed.
	// Max Length: 255
	// Min Length: 1
	DefaultSecurityGroupID string `json:"defaultSecurityGroupId,omitempty"`

	// Security group where Knox-enabled hosts are placed.
	// Max Length: 255
	// Min Length: 1
	SecurityGroupIDForKnox string `json:"securityGroupIdForKnox,omitempty"`
}

// Validate validates this security access v1 response
func (m *SecurityAccessV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultSecurityGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupIDForKnox(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAccessV1Response) validateCidr(formats strfmt.Registry) error {

	if swag.IsZero(m.Cidr) { // not required
		return nil
	}

	if err := validate.MinLength("cidr", "body", string(m.Cidr), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("cidr", "body", string(m.Cidr), 4000); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAccessV1Response) validateDefaultSecurityGroupID(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultSecurityGroupID) { // not required
		return nil
	}

	if err := validate.MinLength("defaultSecurityGroupId", "body", string(m.DefaultSecurityGroupID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("defaultSecurityGroupId", "body", string(m.DefaultSecurityGroupID), 255); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAccessV1Response) validateSecurityGroupIDForKnox(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroupIDForKnox) { // not required
		return nil
	}

	if err := validate.MinLength("securityGroupIdForKnox", "body", string(m.SecurityGroupIDForKnox), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("securityGroupIdForKnox", "body", string(m.SecurityGroupIDForKnox), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAccessV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAccessV1Response) UnmarshalBinary(b []byte) error {
	var res SecurityAccessV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
