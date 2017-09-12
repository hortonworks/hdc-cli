// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PlatformResourceRequestJSON platform resource request Json
// swagger:model PlatformResourceRequestJson
type PlatformResourceRequestJSON struct {

	// credential resource id for the request
	CredentialID int64 `json:"credentialId,omitempty"`

	// credential resource name for the request
	CredentialName string `json:"credentialName,omitempty"`

	// filter for resources
	Filters map[string]string `json:"filters,omitempty"`

	// cloud provider api variant
	PlatformVariant string `json:"platformVariant,omitempty"`

	// Related region
	Region string `json:"region,omitempty"`
}

// Validate validates this platform resource request Json
func (m *PlatformResourceRequestJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PlatformResourceRequestJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformResourceRequestJSON) UnmarshalBinary(b []byte) error {
	var res PlatformResourceRequestJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
