// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RepairSdxNodeByCrnReader is a Reader for the RepairSdxNodeByCrn structure.
type RepairSdxNodeByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepairSdxNodeByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewRepairSdxNodeByCrnDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewRepairSdxNodeByCrnDefault creates a RepairSdxNodeByCrnDefault with default headers values
func NewRepairSdxNodeByCrnDefault(code int) *RepairSdxNodeByCrnDefault {
	return &RepairSdxNodeByCrnDefault{
		_statusCode: code,
	}
}

/*RepairSdxNodeByCrnDefault handles this case with default header values.

successful operation
*/
type RepairSdxNodeByCrnDefault struct {
	_statusCode int
}

// Code gets the status code for the repair sdx node by crn default response
func (o *RepairSdxNodeByCrnDefault) Code() int {
	return o._statusCode
}

func (o *RepairSdxNodeByCrnDefault) Error() string {
	return fmt.Sprintf("[POST /sdx/crn/{crn}/manual_repair][%d] repairSdxNodeByCrn default ", o._statusCode)
}

func (o *RepairSdxNodeByCrnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
