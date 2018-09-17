// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_smartsensesubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// CreateSmartSenseSubscriptionInWorkspaceReader is a Reader for the CreateSmartSenseSubscriptionInWorkspace structure.
type CreateSmartSenseSubscriptionInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSmartSenseSubscriptionInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSmartSenseSubscriptionInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSmartSenseSubscriptionInWorkspaceOK creates a CreateSmartSenseSubscriptionInWorkspaceOK with default headers values
func NewCreateSmartSenseSubscriptionInWorkspaceOK() *CreateSmartSenseSubscriptionInWorkspaceOK {
	return &CreateSmartSenseSubscriptionInWorkspaceOK{}
}

/*CreateSmartSenseSubscriptionInWorkspaceOK handles this case with default header values.

successful operation
*/
type CreateSmartSenseSubscriptionInWorkspaceOK struct {
	Payload *models_cloudbreak.SmartSenseSubscriptionJSON
}

func (o *CreateSmartSenseSubscriptionInWorkspaceOK) Error() string {
	return fmt.Sprintf("[POST /v3/{workspaceId}/smartsensesubscriptions][%d] createSmartSenseSubscriptionInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *CreateSmartSenseSubscriptionInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SmartSenseSubscriptionJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}