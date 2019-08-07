// Code generated by go-swagger; DO NOT EDIT.

package databases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
)

// DeleteMultipleDatabasesReader is a Reader for the DeleteMultipleDatabases structure.
type DeleteMultipleDatabasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMultipleDatabasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMultipleDatabasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteMultipleDatabasesOK creates a DeleteMultipleDatabasesOK with default headers values
func NewDeleteMultipleDatabasesOK() *DeleteMultipleDatabasesOK {
	return &DeleteMultipleDatabasesOK{}
}

/*DeleteMultipleDatabasesOK handles this case with default header values.

successful operation
*/
type DeleteMultipleDatabasesOK struct {
	Payload *model.DatabaseV4Responses
}

func (o *DeleteMultipleDatabasesOK) Error() string {
	return fmt.Sprintf("[DELETE /v4/databases][%d] deleteMultipleDatabasesOK  %+v", 200, o.Payload)
}

func (o *DeleteMultipleDatabasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DatabaseV4Responses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}