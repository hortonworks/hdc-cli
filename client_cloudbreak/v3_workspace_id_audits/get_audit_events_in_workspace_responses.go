// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_audits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetAuditEventsInWorkspaceReader is a Reader for the GetAuditEventsInWorkspace structure.
type GetAuditEventsInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditEventsInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuditEventsInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAuditEventsInWorkspaceOK creates a GetAuditEventsInWorkspaceOK with default headers values
func NewGetAuditEventsInWorkspaceOK() *GetAuditEventsInWorkspaceOK {
	return &GetAuditEventsInWorkspaceOK{}
}

/*GetAuditEventsInWorkspaceOK handles this case with default header values.

successful operation
*/
type GetAuditEventsInWorkspaceOK struct {
	Payload []*models_cloudbreak.AuditEvent
}

func (o *GetAuditEventsInWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v3/{workspaceId}/audits/events/{resourceType}/{resourceId}][%d] getAuditEventsInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetAuditEventsInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
