// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FailurePolicyRequest failure policy request
// swagger:model FailurePolicyRequest
type FailurePolicyRequest struct {

	// type of  adjustment
	// Required: true
	AdjustmentType *string `json:"adjustmentType"`

	// threshold of failure policy
	Threshold int64 `json:"threshold,omitempty"`
}

// Validate validates this failure policy request
func (m *FailurePolicyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdjustmentType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var failurePolicyRequestTypeAdjustmentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXACT","PERCENTAGE","BEST_EFFORT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		failurePolicyRequestTypeAdjustmentTypePropEnum = append(failurePolicyRequestTypeAdjustmentTypePropEnum, v)
	}
}

const (
	// FailurePolicyRequestAdjustmentTypeEXACT captures enum value "EXACT"
	FailurePolicyRequestAdjustmentTypeEXACT string = "EXACT"
	// FailurePolicyRequestAdjustmentTypePERCENTAGE captures enum value "PERCENTAGE"
	FailurePolicyRequestAdjustmentTypePERCENTAGE string = "PERCENTAGE"
	// FailurePolicyRequestAdjustmentTypeBESTEFFORT captures enum value "BEST_EFFORT"
	FailurePolicyRequestAdjustmentTypeBESTEFFORT string = "BEST_EFFORT"
)

// prop value enum
func (m *FailurePolicyRequest) validateAdjustmentTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, failurePolicyRequestTypeAdjustmentTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FailurePolicyRequest) validateAdjustmentType(formats strfmt.Registry) error {

	if err := validate.Required("adjustmentType", "body", m.AdjustmentType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAdjustmentTypeEnum("adjustmentType", "body", *m.AdjustmentType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FailurePolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FailurePolicyRequest) UnmarshalBinary(b []byte) error {
	var res FailurePolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
