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

// AwsPlacementGroupV4Parameters aws placement group v4 parameters
// swagger:model AwsPlacementGroupV4Parameters
type AwsPlacementGroupV4Parameters struct {

	// aws placement group strategy for vm
	// Enum: [NONE PARTITION SPREAD CLUSTER]
	Strategy string `json:"strategy,omitempty"`
}

// Validate validates this aws placement group v4 parameters
func (m *AwsPlacementGroupV4Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var awsPlacementGroupV4ParametersTypeStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","PARTITION","SPREAD","CLUSTER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsPlacementGroupV4ParametersTypeStrategyPropEnum = append(awsPlacementGroupV4ParametersTypeStrategyPropEnum, v)
	}
}

const (

	// AwsPlacementGroupV4ParametersStrategyNONE captures enum value "NONE"
	AwsPlacementGroupV4ParametersStrategyNONE string = "NONE"

	// AwsPlacementGroupV4ParametersStrategyPARTITION captures enum value "PARTITION"
	AwsPlacementGroupV4ParametersStrategyPARTITION string = "PARTITION"

	// AwsPlacementGroupV4ParametersStrategySPREAD captures enum value "SPREAD"
	AwsPlacementGroupV4ParametersStrategySPREAD string = "SPREAD"

	// AwsPlacementGroupV4ParametersStrategyCLUSTER captures enum value "CLUSTER"
	AwsPlacementGroupV4ParametersStrategyCLUSTER string = "CLUSTER"
)

// prop value enum
func (m *AwsPlacementGroupV4Parameters) validateStrategyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, awsPlacementGroupV4ParametersTypeStrategyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AwsPlacementGroupV4Parameters) validateStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.Strategy) { // not required
		return nil
	}

	// value enum
	if err := m.validateStrategyEnum("strategy", "body", m.Strategy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsPlacementGroupV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsPlacementGroupV4Parameters) UnmarshalBinary(b []byte) error {
	var res AwsPlacementGroupV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
