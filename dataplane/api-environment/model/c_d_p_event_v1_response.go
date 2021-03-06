// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CDPEventV1Response c d p event v1 response
// swagger:model CDPEventV1Response
type CDPEventV1Response struct {

	// availability zone
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// blueprint Id
	BlueprintID int64 `json:"blueprintId,omitempty"`

	// blueprint name
	BlueprintName string `json:"blueprintName,omitempty"`

	// cloud
	Cloud string `json:"cloud,omitempty"`

	// cluster Id
	ClusterID int64 `json:"clusterId,omitempty"`

	// cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// cluster status
	ClusterStatus string `json:"clusterStatus,omitempty"`

	// event message
	EventMessage string `json:"eventMessage,omitempty"`

	// event timestamp
	EventTimestamp int64 `json:"eventTimestamp,omitempty"`

	// event type
	EventType string `json:"eventType,omitempty"`

	// instance group
	InstanceGroup string `json:"instanceGroup,omitempty"`

	// ldap details
	LdapDetails *LdapDetails `json:"ldapDetails,omitempty"`

	// node count
	NodeCount int32 `json:"nodeCount,omitempty"`

	// notification type
	NotificationType string `json:"notificationType,omitempty"`

	// rds details
	RdsDetails *RdsDetails `json:"rdsDetails,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// stack crn
	StackCrn string `json:"stackCrn,omitempty"`

	// stack name
	StackName string `json:"stackName,omitempty"`

	// stack status
	StackStatus string `json:"stackStatus,omitempty"`

	// tenant name
	TenantName string `json:"tenantName,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`

	// workspace Id
	WorkspaceID int64 `json:"workspaceId,omitempty"`
}

// Validate validates this c d p event v1 response
func (m *CDPEventV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLdapDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRdsDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CDPEventV1Response) validateLdapDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.LdapDetails) { // not required
		return nil
	}

	if m.LdapDetails != nil {
		if err := m.LdapDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldapDetails")
			}
			return err
		}
	}

	return nil
}

func (m *CDPEventV1Response) validateRdsDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.RdsDetails) { // not required
		return nil
	}

	if m.RdsDetails != nil {
		if err := m.RdsDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rdsDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CDPEventV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CDPEventV1Response) UnmarshalBinary(b []byte) error {
	var res CDPEventV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
