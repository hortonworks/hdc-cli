// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SetDistroXMaintenanceModeReader is a Reader for the SetDistroXMaintenanceMode structure.
type SetDistroXMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDistroXMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewSetDistroXMaintenanceModeDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewSetDistroXMaintenanceModeDefault creates a SetDistroXMaintenanceModeDefault with default headers values
func NewSetDistroXMaintenanceModeDefault(code int) *SetDistroXMaintenanceModeDefault {
	return &SetDistroXMaintenanceModeDefault{
		_statusCode: code,
	}
}

/*SetDistroXMaintenanceModeDefault handles this case with default header values.

successful operation
*/
type SetDistroXMaintenanceModeDefault struct {
	_statusCode int
}

// Code gets the status code for the set distro x maintenance mode default response
func (o *SetDistroXMaintenanceModeDefault) Code() int {
	return o._statusCode
}

func (o *SetDistroXMaintenanceModeDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/{name}/maintenance][%d] setDistroXMaintenanceMode default ", o._statusCode)
}

func (o *SetDistroXMaintenanceModeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}