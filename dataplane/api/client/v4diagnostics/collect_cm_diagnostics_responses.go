// Code generated by go-swagger; DO NOT EDIT.

package v4diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// CollectCmDiagnosticsReader is a Reader for the CollectCmDiagnostics structure.
type CollectCmDiagnosticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectCmDiagnosticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCollectCmDiagnosticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCollectCmDiagnosticsOK creates a CollectCmDiagnosticsOK with default headers values
func NewCollectCmDiagnosticsOK() *CollectCmDiagnosticsOK {
	return &CollectCmDiagnosticsOK{}
}

/*CollectCmDiagnosticsOK handles this case with default header values.

successful operation
*/
type CollectCmDiagnosticsOK struct {
	Payload *model.FlowIdentifier
}

func (o *CollectCmDiagnosticsOK) Error() string {
	return fmt.Sprintf("[POST /v4/diagnostics/cm][%d] collectCmDiagnosticsOK  %+v", 200, o.Payload)
}

func (o *CollectCmDiagnosticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}