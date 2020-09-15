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

// DiagnosticsCollectionV1Request diagnostics collection v1 request
// swagger:model DiagnosticsCollectionV1Request
type DiagnosticsCollectionV1Request struct {

	// Additional log path and label pairs that will be sent in the diagnostics collection
	AdditionalLogs []*VMLog `json:"additionalLogs"`

	// description of the diagnostics collection
	Description string `json:"description,omitempty"`

	// Destination for the diagnostic collection request.
	// Required: true
	// Enum: [LOCAL CLOUD_STORAGE SUPPORT ENG]
	Destination *string `json:"destination"`

	// END time for the time interval of the diagnostic collection request.
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// Host groups (instance groups), used it to run diagnostics collection only those hosts that are included the specific host groups
	// Unique: true
	HostGroups []string `json:"hostGroups"`

	// Host (fqdn) filter, use it to run diagnostics collection on only specific hosts
	// Unique: true
	Hosts []string `json:"hosts"`

	// Include salt logs in the diagnostic collections
	IncludeSaltLogs bool `json:"includeSaltLogs,omitempty"`

	// Issue number or JIRA ticket number related to this diagnostic collection request.
	Issue string `json:"issue,omitempty"`

	// With labels you can filter what kind of logs you'd like to collect.
	Labels []string `json:"labels"`

	// Skip cloud storage write operation testing or databus connection check (depends on the destination) during init stage.
	SkipValidation bool `json:"skipValidation,omitempty"`

	// the unique crn of the resource
	// Required: true
	StackCrn *string `json:"stackCrn"`

	// Start time for the time interval of the diagnostic collection request.
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// Upgrade or install required telemetry cli tool on the nodes (works only with network)
	UpdatePackage bool `json:"updatePackage,omitempty"`
}

// Validate validates this diagnostics collection v1 request
func (m *DiagnosticsCollectionV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiagnosticsCollectionV1Request) validateAdditionalLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalLogs); i++ {
		if swag.IsZero(m.AdditionalLogs[i]) { // not required
			continue
		}

		if m.AdditionalLogs[i] != nil {
			if err := m.AdditionalLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var diagnosticsCollectionV1RequestTypeDestinationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LOCAL","CLOUD_STORAGE","SUPPORT","ENG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diagnosticsCollectionV1RequestTypeDestinationPropEnum = append(diagnosticsCollectionV1RequestTypeDestinationPropEnum, v)
	}
}

const (

	// DiagnosticsCollectionV1RequestDestinationLOCAL captures enum value "LOCAL"
	DiagnosticsCollectionV1RequestDestinationLOCAL string = "LOCAL"

	// DiagnosticsCollectionV1RequestDestinationCLOUDSTORAGE captures enum value "CLOUD_STORAGE"
	DiagnosticsCollectionV1RequestDestinationCLOUDSTORAGE string = "CLOUD_STORAGE"

	// DiagnosticsCollectionV1RequestDestinationSUPPORT captures enum value "SUPPORT"
	DiagnosticsCollectionV1RequestDestinationSUPPORT string = "SUPPORT"

	// DiagnosticsCollectionV1RequestDestinationENG captures enum value "ENG"
	DiagnosticsCollectionV1RequestDestinationENG string = "ENG"
)

// prop value enum
func (m *DiagnosticsCollectionV1Request) validateDestinationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, diagnosticsCollectionV1RequestTypeDestinationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DiagnosticsCollectionV1Request) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	// value enum
	if err := m.validateDestinationEnum("destination", "body", *m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *DiagnosticsCollectionV1Request) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DiagnosticsCollectionV1Request) validateHostGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.HostGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("hostGroups", "body", m.HostGroups); err != nil {
		return err
	}

	return nil
}

func (m *DiagnosticsCollectionV1Request) validateHosts(formats strfmt.Registry) error {

	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	if err := validate.UniqueItems("hosts", "body", m.Hosts); err != nil {
		return err
	}

	return nil
}

func (m *DiagnosticsCollectionV1Request) validateStackCrn(formats strfmt.Registry) error {

	if err := validate.Required("stackCrn", "body", m.StackCrn); err != nil {
		return err
	}

	return nil
}

func (m *DiagnosticsCollectionV1Request) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiagnosticsCollectionV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiagnosticsCollectionV1Request) UnmarshalBinary(b []byte) error {
	var res DiagnosticsCollectionV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
