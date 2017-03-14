package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*BlueprintRequest blueprint request

swagger:model BlueprintRequest
*/
type BlueprintRequest struct {

	// TODO WARNING: do not replace it with string, otherwise it cannot be serialized
	/* ambari blueprint JSON, set this or the url field

	Required: true
	*/
	AmbariBlueprint AmbariBlueprint `json:"ambariBlueprint"`

	/* description of the resource

	Max Length: 1000
	Min Length: 0
	*/
	Description *string `json:"description,omitempty"`

	/* name of the resource

	Required: true
	Max Length: 100
	Min Length: 1
	*/
	Name string `json:"name"`

	/* properties to extend the blueprint with
	 */
	Properties []Configurations `json:"properties,omitempty"`

	/* url source of an ambari blueprint, set this or the ambariBlueprint field
	 */
	URL *string `json:"url,omitempty"`
}

// Validate validates this blueprint request
func (m *BlueprintRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlueprintRequest) validateDescription(formats strfmt.Registry) error {

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

func (m *BlueprintRequest) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	for i := 0; i < len(m.Properties); i++ {

		if err := validate.Required("properties"+"."+strconv.Itoa(i), "body", m.Properties[i]); err != nil {
			return err
		}

	}

	return nil
}