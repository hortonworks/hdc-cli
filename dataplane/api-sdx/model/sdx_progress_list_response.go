// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SdxProgressListResponse sdx progress list response
// swagger:model SdxProgressListResponse
type SdxProgressListResponse struct {

	// recent flow operations
	RecentFlowOperations []*FlowProgressResponse `json:"recentFlowOperations"`

	// recent internal flow operations
	RecentInternalFlowOperations []*FlowProgressResponse `json:"recentInternalFlowOperations"`
}

// Validate validates this sdx progress list response
func (m *SdxProgressListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecentFlowOperations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecentInternalFlowOperations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SdxProgressListResponse) validateRecentFlowOperations(formats strfmt.Registry) error {

	if swag.IsZero(m.RecentFlowOperations) { // not required
		return nil
	}

	for i := 0; i < len(m.RecentFlowOperations); i++ {
		if swag.IsZero(m.RecentFlowOperations[i]) { // not required
			continue
		}

		if m.RecentFlowOperations[i] != nil {
			if err := m.RecentFlowOperations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recentFlowOperations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SdxProgressListResponse) validateRecentInternalFlowOperations(formats strfmt.Registry) error {

	if swag.IsZero(m.RecentInternalFlowOperations) { // not required
		return nil
	}

	for i := 0; i < len(m.RecentInternalFlowOperations); i++ {
		if swag.IsZero(m.RecentInternalFlowOperations[i]) { // not required
			continue
		}

		if m.RecentInternalFlowOperations[i] != nil {
			if err := m.RecentInternalFlowOperations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recentInternalFlowOperations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SdxProgressListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxProgressListResponse) UnmarshalBinary(b []byte) error {
	var res SdxProgressListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}