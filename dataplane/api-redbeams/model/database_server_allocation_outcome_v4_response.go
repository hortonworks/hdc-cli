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

// DatabaseServerAllocationOutcomeV4Response database server allocation outcome v4 response
// swagger:model DatabaseServerAllocationOutcomeV4Response
type DatabaseServerAllocationOutcomeV4Response struct {

	// CRN of the environment of the database server
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// Name of the database server
	// Required: true
	Name *string `json:"name"`

	// CRN of the database server
	// Required: true
	ResourceCrn *string `json:"resourceCrn"`

	// status
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED DELETE_REQUESTED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED]
	Status string `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this database server allocation outcome v4 response
func (m *DatabaseServerAllocationOutcomeV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseServerAllocationOutcomeV4Response) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseServerAllocationOutcomeV4Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseServerAllocationOutcomeV4Response) validateResourceCrn(formats strfmt.Registry) error {

	if err := validate.Required("resourceCrn", "body", m.ResourceCrn); err != nil {
		return err
	}

	return nil
}

var databaseServerAllocationOutcomeV4ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","DELETE_REQUESTED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		databaseServerAllocationOutcomeV4ResponseTypeStatusPropEnum = append(databaseServerAllocationOutcomeV4ResponseTypeStatusPropEnum, v)
	}
}

const (

	// DatabaseServerAllocationOutcomeV4ResponseStatusREQUESTED captures enum value "REQUESTED"
	DatabaseServerAllocationOutcomeV4ResponseStatusREQUESTED string = "REQUESTED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	DatabaseServerAllocationOutcomeV4ResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// DatabaseServerAllocationOutcomeV4ResponseStatusAVAILABLE captures enum value "AVAILABLE"
	DatabaseServerAllocationOutcomeV4ResponseStatusAVAILABLE string = "AVAILABLE"

	// DatabaseServerAllocationOutcomeV4ResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	DatabaseServerAllocationOutcomeV4ResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// DatabaseServerAllocationOutcomeV4ResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	DatabaseServerAllocationOutcomeV4ResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	DatabaseServerAllocationOutcomeV4ResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	DatabaseServerAllocationOutcomeV4ResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	DatabaseServerAllocationOutcomeV4ResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusDELETEREQUESTED captures enum value "DELETE_REQUESTED"
	DatabaseServerAllocationOutcomeV4ResponseStatusDELETEREQUESTED string = "DELETE_REQUESTED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	DatabaseServerAllocationOutcomeV4ResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// DatabaseServerAllocationOutcomeV4ResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	DatabaseServerAllocationOutcomeV4ResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// DatabaseServerAllocationOutcomeV4ResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	DatabaseServerAllocationOutcomeV4ResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	DatabaseServerAllocationOutcomeV4ResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusSTOPPED captures enum value "STOPPED"
	DatabaseServerAllocationOutcomeV4ResponseStatusSTOPPED string = "STOPPED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	DatabaseServerAllocationOutcomeV4ResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	DatabaseServerAllocationOutcomeV4ResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	DatabaseServerAllocationOutcomeV4ResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// DatabaseServerAllocationOutcomeV4ResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	DatabaseServerAllocationOutcomeV4ResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// DatabaseServerAllocationOutcomeV4ResponseStatusSTARTFAILED captures enum value "START_FAILED"
	DatabaseServerAllocationOutcomeV4ResponseStatusSTARTFAILED string = "START_FAILED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	DatabaseServerAllocationOutcomeV4ResponseStatusSTOPFAILED string = "STOP_FAILED"

	// DatabaseServerAllocationOutcomeV4ResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	DatabaseServerAllocationOutcomeV4ResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// DatabaseServerAllocationOutcomeV4ResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	DatabaseServerAllocationOutcomeV4ResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"
)

// prop value enum
func (m *DatabaseServerAllocationOutcomeV4Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, databaseServerAllocationOutcomeV4ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DatabaseServerAllocationOutcomeV4Response) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseServerAllocationOutcomeV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseServerAllocationOutcomeV4Response) UnmarshalBinary(b []byte) error {
	var res DatabaseServerAllocationOutcomeV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}