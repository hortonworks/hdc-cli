package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*SecurityGroupResponse security group response

swagger:model SecurityGroupResponse
*/
type SecurityGroupResponse struct {

	/* account id of the resource owner that is provided by OAuth provider
	 */
	Account *string `json:"account,omitempty"`

	/* description of the resource

	Max Length: 1000
	Min Length: 0
	*/
	Description *string `json:"description,omitempty"`

	/* id of the resource
	 */
	ID *int64 `json:"id,omitempty"`

	/* name of the resource

	Required: true
	Max Length: 100
	Min Length: 1
	Pattern: ([a-z][-a-z0-9]*[a-z0-9])
	*/
	Name string `json:"name"`

	/* id of the resource owner that is provided by OAuth provider
	 */
	Owner *string `json:"owner,omitempty"`

	/* resource is visible in account

	Required: true
	*/
	PublicInAccount bool `json:"publicInAccount"`

	/* list of security rules that relates to the security group
	 */
	SecurityRules []*SecurityRuleResponse `json:"securityRules,omitempty"`
}

// Validate validates this security group response
func (m *SecurityGroupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublicInAccount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecurityRules(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupResponse) validateDescription(formats strfmt.Registry) error {

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

func (m *SecurityGroupResponse) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `([a-z][-a-z0-9]*[a-z0-9])`); err != nil {
		return err
	}

	return nil
}

func (m *SecurityGroupResponse) validatePublicInAccount(formats strfmt.Registry) error {

	if err := validate.Required("publicInAccount", "body", bool(m.PublicInAccount)); err != nil {
		return err
	}

	return nil
}

func (m *SecurityGroupResponse) validateSecurityRules(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityRules) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityRules); i++ {

		if m.SecurityRules[i] != nil {

			if err := m.SecurityRules[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}