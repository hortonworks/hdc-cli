// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HostGroupAdjustment host group adjustment
// swagger:model HostGroupAdjustment
type HostGroupAdjustment struct {

	// name of the host group
	// Required: true
	HostGroup *string `json:"hostGroup"`

	// scaling adjustment of the host groups
	// Required: true
	ScalingAdjustment *int32 `json:"scalingAdjustment"`

	// validate node count during downscale
	ValidateNodeCount *bool `json:"validateNodeCount,omitempty"`

	// on cluster update, update stack too
	WithStackUpdate *bool `json:"withStackUpdate,omitempty"`
}

// Validate validates this host group adjustment
func (m *HostGroupAdjustment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScalingAdjustment(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostGroupAdjustment) validateHostGroup(formats strfmt.Registry) error {

	if err := validate.Required("hostGroup", "body", m.HostGroup); err != nil {
		return err
	}

	return nil
}

func (m *HostGroupAdjustment) validateScalingAdjustment(formats strfmt.Registry) error {

	if err := validate.Required("scalingAdjustment", "body", m.ScalingAdjustment); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostGroupAdjustment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostGroupAdjustment) UnmarshalBinary(b []byte) error {
	var res HostGroupAdjustment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
