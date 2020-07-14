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

// AttachedFreeIpaRequestAwsSpotParameters attached free ipa request aws spot parameters
// swagger:model AttachedFreeIpaRequestAwsSpotParameters
type AttachedFreeIpaRequestAwsSpotParameters struct {

	// Max price per hour of spot instances launched in FreeIpa instance group
	// Maximum: 255
	// Minimum: 0.001
	MaxPrice float64 `json:"maxPrice,omitempty"`

	// Percentage of spot instances launched in FreeIpa instance group
	// Maximum: 100
	// Minimum: 0
	Percentage *int32 `json:"percentage,omitempty"`
}

// Validate validates this attached free ipa request aws spot parameters
func (m *AttachedFreeIpaRequestAwsSpotParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AttachedFreeIpaRequestAwsSpotParameters) validateMaxPrice(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxPrice) { // not required
		return nil
	}

	if err := validate.Minimum("maxPrice", "body", float64(m.MaxPrice), 0.001, false); err != nil {
		return err
	}

	if err := validate.Maximum("maxPrice", "body", float64(m.MaxPrice), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *AttachedFreeIpaRequestAwsSpotParameters) validatePercentage(formats strfmt.Registry) error {

	if swag.IsZero(m.Percentage) { // not required
		return nil
	}

	if err := validate.MinimumInt("percentage", "body", int64(*m.Percentage), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("percentage", "body", int64(*m.Percentage), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttachedFreeIpaRequestAwsSpotParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttachedFreeIpaRequestAwsSpotParameters) UnmarshalBinary(b []byte) error {
	var res AttachedFreeIpaRequestAwsSpotParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
