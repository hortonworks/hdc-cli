// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DescribeFreeIpaV1Response describe free ipa v1 response
// swagger:model DescribeFreeIpaV1Response
type DescribeFreeIpaV1Response struct {

	// freeipa stack related authentication
	// Required: true
	Authentication *StackAuthenticationV1Response `json:"authentication"`

	// crn
	// Required: true
	Crn *string `json:"crn"`

	// The crn of the environment
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// settings for freeipa server
	// Required: true
	FreeIpa *FreeIpaServerV1Response `json:"freeIpa"`

	// settings for custom images
	Image *ImageSettingsV1Response `json:"image,omitempty"`

	// collection of instance groupst
	// Required: true
	InstanceGroups []*InstanceGroupV1Response `json:"instanceGroups"`

	// name of the freeipa stack
	// Required: true
	Name *string `json:"name"`

	// freeipa stack related network
	Network *NetworkV1Response `json:"network,omitempty"`

	// placement configuration parameters for a cluster (e.g. 'region', 'availabilityZone')
	// Required: true
	Placement *PlacementV1Response `json:"placement"`

	// status
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED]
	Status string `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this describe free ipa v1 response
func (m *DescribeFreeIpaV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeIpa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlacement(formats); err != nil {
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

func (m *DescribeFreeIpaV1Response) validateAuthentication(formats strfmt.Registry) error {

	if err := validate.Required("authentication", "body", m.Authentication); err != nil {
		return err
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *DescribeFreeIpaV1Response) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *DescribeFreeIpaV1Response) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *DescribeFreeIpaV1Response) validateFreeIpa(formats strfmt.Registry) error {

	if err := validate.Required("freeIpa", "body", m.FreeIpa); err != nil {
		return err
	}

	if m.FreeIpa != nil {
		if err := m.FreeIpa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeIpa")
			}
			return err
		}
	}

	return nil
}

func (m *DescribeFreeIpaV1Response) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *DescribeFreeIpaV1Response) validateInstanceGroups(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroups", "body", m.InstanceGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.InstanceGroups); i++ {
		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {
			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DescribeFreeIpaV1Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DescribeFreeIpaV1Response) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *DescribeFreeIpaV1Response) validatePlacement(formats strfmt.Registry) error {

	if err := validate.Required("placement", "body", m.Placement); err != nil {
		return err
	}

	if m.Placement != nil {
		if err := m.Placement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("placement")
			}
			return err
		}
	}

	return nil
}

var describeFreeIpaV1ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		describeFreeIpaV1ResponseTypeStatusPropEnum = append(describeFreeIpaV1ResponseTypeStatusPropEnum, v)
	}
}

const (

	// DescribeFreeIpaV1ResponseStatusREQUESTED captures enum value "REQUESTED"
	DescribeFreeIpaV1ResponseStatusREQUESTED string = "REQUESTED"

	// DescribeFreeIpaV1ResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	DescribeFreeIpaV1ResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// DescribeFreeIpaV1ResponseStatusAVAILABLE captures enum value "AVAILABLE"
	DescribeFreeIpaV1ResponseStatusAVAILABLE string = "AVAILABLE"

	// DescribeFreeIpaV1ResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	DescribeFreeIpaV1ResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// DescribeFreeIpaV1ResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	DescribeFreeIpaV1ResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// DescribeFreeIpaV1ResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	DescribeFreeIpaV1ResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// DescribeFreeIpaV1ResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	DescribeFreeIpaV1ResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// DescribeFreeIpaV1ResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	DescribeFreeIpaV1ResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// DescribeFreeIpaV1ResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	DescribeFreeIpaV1ResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// DescribeFreeIpaV1ResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	DescribeFreeIpaV1ResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// DescribeFreeIpaV1ResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	DescribeFreeIpaV1ResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// DescribeFreeIpaV1ResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	DescribeFreeIpaV1ResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// DescribeFreeIpaV1ResponseStatusSTOPPED captures enum value "STOPPED"
	DescribeFreeIpaV1ResponseStatusSTOPPED string = "STOPPED"

	// DescribeFreeIpaV1ResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	DescribeFreeIpaV1ResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// DescribeFreeIpaV1ResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	DescribeFreeIpaV1ResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// DescribeFreeIpaV1ResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	DescribeFreeIpaV1ResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// DescribeFreeIpaV1ResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	DescribeFreeIpaV1ResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// DescribeFreeIpaV1ResponseStatusSTARTFAILED captures enum value "START_FAILED"
	DescribeFreeIpaV1ResponseStatusSTARTFAILED string = "START_FAILED"

	// DescribeFreeIpaV1ResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	DescribeFreeIpaV1ResponseStatusSTOPFAILED string = "STOP_FAILED"

	// DescribeFreeIpaV1ResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	DescribeFreeIpaV1ResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// DescribeFreeIpaV1ResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	DescribeFreeIpaV1ResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"
)

// prop value enum
func (m *DescribeFreeIpaV1Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, describeFreeIpaV1ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DescribeFreeIpaV1Response) validateStatus(formats strfmt.Registry) error {

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
func (m *DescribeFreeIpaV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeFreeIpaV1Response) UnmarshalBinary(b []byte) error {
	var res DescribeFreeIpaV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}