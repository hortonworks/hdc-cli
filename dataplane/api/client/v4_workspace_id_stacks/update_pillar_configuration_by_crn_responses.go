// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// UpdatePillarConfigurationByCrnReader is a Reader for the UpdatePillarConfigurationByCrn structure.
type UpdatePillarConfigurationByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePillarConfigurationByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdatePillarConfigurationByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePillarConfigurationByCrnOK creates a UpdatePillarConfigurationByCrnOK with default headers values
func NewUpdatePillarConfigurationByCrnOK() *UpdatePillarConfigurationByCrnOK {
	return &UpdatePillarConfigurationByCrnOK{}
}

/*UpdatePillarConfigurationByCrnOK handles this case with default header values.

successful operation
*/
type UpdatePillarConfigurationByCrnOK struct {
	Payload *model.FlowIdentifier
}

func (o *UpdatePillarConfigurationByCrnOK) Error() string {
	return fmt.Sprintf("[PUT /v4/{workspaceId}/stacks/crn/{crn}/pillar_config_update][%d] updatePillarConfigurationByCrnOK  %+v", 200, o.Payload)
}

func (o *UpdatePillarConfigurationByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
