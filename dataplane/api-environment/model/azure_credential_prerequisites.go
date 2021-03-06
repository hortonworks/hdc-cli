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

// AzureCredentialPrerequisites azure credential prerequisites
// swagger:model AzureCredentialPrerequisites
type AzureCredentialPrerequisites struct {

	// Azure CLI command to create Azure AD Application as prerequisite for credential creation. The field is base64 encoded.
	// Required: true
	AppCreationCommand *string `json:"appCreationCommand"`

	// Azure specific JSON file that is base64 encoded and describes the necessary Azure role for cloud resource provisioning.
	// Required: true
	RoleDefitionJSON *string `json:"roleDefitionJson"`
}

// Validate validates this azure credential prerequisites
func (m *AzureCredentialPrerequisites) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppCreationCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleDefitionJSON(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureCredentialPrerequisites) validateAppCreationCommand(formats strfmt.Registry) error {

	if err := validate.Required("appCreationCommand", "body", m.AppCreationCommand); err != nil {
		return err
	}

	return nil
}

func (m *AzureCredentialPrerequisites) validateRoleDefitionJSON(formats strfmt.Registry) error {

	if err := validate.Required("roleDefitionJson", "body", m.RoleDefitionJSON); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureCredentialPrerequisites) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCredentialPrerequisites) UnmarshalBinary(b []byte) error {
	var res AzureCredentialPrerequisites
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
