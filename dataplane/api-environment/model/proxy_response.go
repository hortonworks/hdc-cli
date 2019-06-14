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

// ProxyResponse Cloudbreak allows you to save your existing proxy configuration information as an external source so that you can provide the proxy information to multiple clusters that you create with Cloudbreak
// swagger:model ProxyResponse
type ProxyResponse struct {

	// crn of the creator
	Creator string `json:"creator,omitempty"`

	// proxy configuration id for the cluster
	Crn string `json:"crn,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// host or IP address of proxy server
	// Required: true
	// Max Length: 255
	// Min Length: 1
	// Pattern: (^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$)
	Host *string `json:"host"`

	// Name of the proxy configuration resource
	// Required: true
	// Max Length: 100
	// Min Length: 4
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// Password to use for basic authentication
	Password *SecretResponse `json:"password,omitempty"`

	// port of proxy server (typically: 3128 or 8080)
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Port *int32 `json:"port"`

	// determines the protocol (http or https)
	// Required: true
	// Pattern: ^http(s)?$
	Protocol *string `json:"protocol"`

	// Username to use for basic authentication
	UserName *SecretResponse `json:"userName,omitempty"`
}

// Validate validates this proxy response
func (m *ProxyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxyResponse) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *ProxyResponse) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	if err := validate.MinLength("host", "body", string(*m.Host), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("host", "body", string(*m.Host), 255); err != nil {
		return err
	}

	if err := validate.Pattern("host", "body", string(*m.Host), `(^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$)`); err != nil {
		return err
	}

	return nil
}

func (m *ProxyResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 4); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *ProxyResponse) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if m.Password != nil {
		if err := m.Password.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("password")
			}
			return err
		}
	}

	return nil
}

func (m *ProxyResponse) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *ProxyResponse) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	if err := validate.Pattern("protocol", "body", string(*m.Protocol), `^http(s)?$`); err != nil {
		return err
	}

	return nil
}

func (m *ProxyResponse) validateUserName(formats strfmt.Registry) error {

	if swag.IsZero(m.UserName) { // not required
		return nil
	}

	if m.UserName != nil {
		if err := m.UserName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProxyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxyResponse) UnmarshalBinary(b []byte) error {
	var res ProxyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
