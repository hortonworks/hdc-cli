// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CertificatesRotationV4Request certificates rotation v4 request
// swagger:model CertificatesRotationV4Request
type CertificatesRotationV4Request struct {

	// rotate certificates type
	// Enum: [HOST_CERTS]
	RotateCertificatesType string `json:"rotateCertificatesType,omitempty"`
}

// Validate validates this certificates rotation v4 request
func (m *CertificatesRotationV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRotateCertificatesType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var certificatesRotationV4RequestTypeRotateCertificatesTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HOST_CERTS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificatesRotationV4RequestTypeRotateCertificatesTypePropEnum = append(certificatesRotationV4RequestTypeRotateCertificatesTypePropEnum, v)
	}
}

const (

	// CertificatesRotationV4RequestRotateCertificatesTypeHOSTCERTS captures enum value "HOST_CERTS"
	CertificatesRotationV4RequestRotateCertificatesTypeHOSTCERTS string = "HOST_CERTS"
)

// prop value enum
func (m *CertificatesRotationV4Request) validateRotateCertificatesTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, certificatesRotationV4RequestTypeRotateCertificatesTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CertificatesRotationV4Request) validateRotateCertificatesType(formats strfmt.Registry) error {

	if swag.IsZero(m.RotateCertificatesType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRotateCertificatesTypeEnum("rotateCertificatesType", "body", m.RotateCertificatesType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificatesRotationV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificatesRotationV4Request) UnmarshalBinary(b []byte) error {
	var res CertificatesRotationV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}