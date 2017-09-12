// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HardwareInfoResponse hardware info response
// swagger:model HardwareInfoResponse
type HardwareInfoResponse struct {

	// metadata of hosts
	HostMetadata *HostMetadata `json:"hostMetadata,omitempty"`

	// metadata of instances
	InstanceMetaData *InstanceMetaData `json:"instanceMetaData,omitempty"`
}

// Validate validates this hardware info response
func (m *HardwareInfoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstanceMetaData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HardwareInfoResponse) validateHostMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.HostMetadata) { // not required
		return nil
	}

	if m.HostMetadata != nil {

		if err := m.HostMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *HardwareInfoResponse) validateInstanceMetaData(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceMetaData) { // not required
		return nil
	}

	if m.InstanceMetaData != nil {

		if err := m.InstanceMetaData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instanceMetaData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HardwareInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HardwareInfoResponse) UnmarshalBinary(b []byte) error {
	var res HardwareInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
