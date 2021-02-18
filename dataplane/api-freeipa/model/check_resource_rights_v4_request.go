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

// CheckResourceRightsV4Request check resource rights v4 request
// swagger:model CheckResourceRightsV4Request
type CheckResourceRightsV4Request struct {

	// resource rights
	ResourceRights []*ResourceRightsV4 `json:"resourceRights"`

	// rights
	Rights []string `json:"rights"`
}

// Validate validates this check resource rights v4 request
func (m *CheckResourceRightsV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceRights(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckResourceRightsV4Request) validateResourceRights(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceRights) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceRights); i++ {
		if swag.IsZero(m.ResourceRights[i]) { // not required
			continue
		}

		if m.ResourceRights[i] != nil {
			if err := m.ResourceRights[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceRights" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var checkResourceRightsV4RequestRightsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENV_CREATE","ENV_START","ENV_STOP","ENV_DELETE","ENV_DESCRIBE","CHANGE_CRED","DH_CREATE","DH_START","DH_STOP","DH_DELETE","DH_REPAIR","DH_RESIZE","DH_RETRY","SDX_UPGRADE","SDX_REPAIR","SDX_RETRY","LIST_ASSIGNED_ROLES","DISTROX_READ","DISTROX_WRITE","SDX_READ","SDX_WRITE","ENVIRONMENT_READ","ENVIRONMENT_WRITE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		checkResourceRightsV4RequestRightsItemsEnum = append(checkResourceRightsV4RequestRightsItemsEnum, v)
	}
}

func (m *CheckResourceRightsV4Request) validateRightsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, checkResourceRightsV4RequestRightsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *CheckResourceRightsV4Request) validateRights(formats strfmt.Registry) error {

	if swag.IsZero(m.Rights) { // not required
		return nil
	}

	for i := 0; i < len(m.Rights); i++ {

		// value enum
		if err := m.validateRightsItemsEnum("rights"+"."+strconv.Itoa(i), "body", m.Rights[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckResourceRightsV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckResourceRightsV4Request) UnmarshalBinary(b []byte) error {
	var res CheckResourceRightsV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}