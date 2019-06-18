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

// EnvironmentNetworkAwsV1Params environment network aws v1 params
// swagger:model EnvironmentNetworkAwsV1Params
type EnvironmentNetworkAwsV1Params struct {

	// Subnet ids of the specified networks
	// Required: true
	VpcID *string `json:"vpcId"`
}

// Validate validates this environment network aws v1 params
func (m *EnvironmentNetworkAwsV1Params) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVpcID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentNetworkAwsV1Params) validateVpcID(formats strfmt.Registry) error {

	if err := validate.Required("vpcId", "body", m.VpcID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentNetworkAwsV1Params) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentNetworkAwsV1Params) UnmarshalBinary(b []byte) error {
	var res EnvironmentNetworkAwsV1Params
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}