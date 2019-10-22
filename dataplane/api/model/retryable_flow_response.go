// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RetryableFlowResponse retryable flow response
// swagger:model RetryableFlowResponse
type RetryableFlowResponse struct {

	// Date when the operation went failed.
	FailDate int64 `json:"failDate,omitempty"`

	// Name of the failed operation, that is also retryable.
	Name string `json:"name,omitempty"`
}

// Validate validates this retryable flow response
func (m *RetryableFlowResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetryableFlowResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetryableFlowResponse) UnmarshalBinary(b []byte) error {
	var res RetryableFlowResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}