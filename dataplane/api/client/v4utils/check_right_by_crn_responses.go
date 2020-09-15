// Code generated by go-swagger; DO NOT EDIT.

package v4utils

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// CheckRightByCrnReader is a Reader for the CheckRightByCrn structure.
type CheckRightByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckRightByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckRightByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCheckRightByCrnOK creates a CheckRightByCrnOK with default headers values
func NewCheckRightByCrnOK() *CheckRightByCrnOK {
	return &CheckRightByCrnOK{}
}

/*CheckRightByCrnOK handles this case with default header values.

successful operation
*/
type CheckRightByCrnOK struct {
	Payload *model.CheckResourceRightsV4Response
}

func (o *CheckRightByCrnOK) Error() string {
	return fmt.Sprintf("[POST /v4/utils/check_right_by_crn][%d] checkRightByCrnOK  %+v", 200, o.Payload)
}

func (o *CheckRightByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CheckResourceRightsV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
