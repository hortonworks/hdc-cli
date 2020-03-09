// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// HasFlowRunningByFlowIDReader is a Reader for the HasFlowRunningByFlowID structure.
type HasFlowRunningByFlowIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HasFlowRunningByFlowIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHasFlowRunningByFlowIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHasFlowRunningByFlowIDOK creates a HasFlowRunningByFlowIDOK with default headers values
func NewHasFlowRunningByFlowIDOK() *HasFlowRunningByFlowIDOK {
	return &HasFlowRunningByFlowIDOK{}
}

/*HasFlowRunningByFlowIDOK handles this case with default header values.

successful operation
*/
type HasFlowRunningByFlowIDOK struct {
	Payload *model.FlowCheckResponse
}

func (o *HasFlowRunningByFlowIDOK) Error() string {
	return fmt.Sprintf("[GET /flow/check/flowId/{flowId}][%d] hasFlowRunningByFlowIdOK  %+v", 200, o.Payload)
}

func (o *HasFlowRunningByFlowIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}