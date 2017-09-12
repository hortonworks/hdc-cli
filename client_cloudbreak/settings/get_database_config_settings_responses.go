// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetDatabaseConfigSettingsReader is a Reader for the GetDatabaseConfigSettings structure.
type GetDatabaseConfigSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatabaseConfigSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDatabaseConfigSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDatabaseConfigSettingsOK creates a GetDatabaseConfigSettingsOK with default headers values
func NewGetDatabaseConfigSettingsOK() *GetDatabaseConfigSettingsOK {
	return &GetDatabaseConfigSettingsOK{}
}

/*GetDatabaseConfigSettingsOK handles this case with default header values.

successful operation
*/
type GetDatabaseConfigSettingsOK struct {
	Payload GetDatabaseConfigSettingsOKBody
}

func (o *GetDatabaseConfigSettingsOK) Error() string {
	return fmt.Sprintf("[GET /settings/database][%d] getDatabaseConfigSettingsOK  %+v", 200, o.Payload)
}

func (o *GetDatabaseConfigSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDatabaseConfigSettingsOKBody get database config settings o k body
swagger:model GetDatabaseConfigSettingsOKBody
*/
type GetDatabaseConfigSettingsOKBody map[string]interface{}

// Validate validates this get database config settings o k body
func (o GetDatabaseConfigSettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("getDatabaseConfigSettingsOK", "body", o); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
