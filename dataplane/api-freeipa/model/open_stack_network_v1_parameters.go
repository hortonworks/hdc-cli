// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OpenStackNetworkV1Parameters open stack network v1 parameters
// swagger:model OpenStackNetworkV1Parameters
type OpenStackNetworkV1Parameters struct {

	// network Id
	NetworkID string `json:"networkId,omitempty"`

	// networking option
	NetworkingOption string `json:"networkingOption,omitempty"`

	// public net Id
	PublicNetID string `json:"publicNetId,omitempty"`

	// router Id
	RouterID string `json:"routerId,omitempty"`

	// subnet Id
	SubnetID string `json:"subnetId,omitempty"`
}

// Validate validates this open stack network v1 parameters
func (m *OpenStackNetworkV1Parameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenStackNetworkV1Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenStackNetworkV1Parameters) UnmarshalBinary(b []byte) error {
	var res OpenStackNetworkV1Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
