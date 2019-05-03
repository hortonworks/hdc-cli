// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// ListProxyConfigsByWorkspaceReader is a Reader for the ListProxyConfigsByWorkspace structure.
type ListProxyConfigsByWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProxyConfigsByWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListProxyConfigsByWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListProxyConfigsByWorkspaceOK creates a ListProxyConfigsByWorkspaceOK with default headers values
func NewListProxyConfigsByWorkspaceOK() *ListProxyConfigsByWorkspaceOK {
	return &ListProxyConfigsByWorkspaceOK{}
}

/*ListProxyConfigsByWorkspaceOK handles this case with default header values.

successful operation
*/
type ListProxyConfigsByWorkspaceOK struct {
	Payload *model.ProxyV4Responses
}

func (o *ListProxyConfigsByWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/proxies][%d] listProxyConfigsByWorkspaceOK  %+v", 200, o.Payload)
}

func (o *ListProxyConfigsByWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ProxyV4Responses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}