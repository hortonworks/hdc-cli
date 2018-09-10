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

// GetSmartSenseSubscriptionInWorkspaceReader is a Reader for the GetSmartSenseSubscriptionInWorkspace structure.
type GetSmartSenseSubscriptionInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSmartSenseSubscriptionInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSmartSenseSubscriptionInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSmartSenseSubscriptionInWorkspaceOK creates a GetSmartSenseSubscriptionInWorkspaceOK with default headers values
func NewGetSmartSenseSubscriptionInWorkspaceOK() *GetSmartSenseSubscriptionInWorkspaceOK {
	return &GetSmartSenseSubscriptionInWorkspaceOK{}
}

/*GetSmartSenseSubscriptionInWorkspaceOK handles this case with default header values.

successful operation
*/
type GetSmartSenseSubscriptionInWorkspaceOK struct {
	Payload *models_cloudbreak.SmartSenseSubscriptionJSON
}

func (o *GetSmartSenseSubscriptionInWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v3/{workspaceId}/smartsensesubscriptions/{name}][%d] getSmartSenseSubscriptionInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetSmartSenseSubscriptionInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SmartSenseSubscriptionJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
