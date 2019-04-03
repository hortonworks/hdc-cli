// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_cluster_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// ListClusterDefinitionsByWorkspaceReader is a Reader for the ListClusterDefinitionsByWorkspace structure.
type ListClusterDefinitionsByWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterDefinitionsByWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListClusterDefinitionsByWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListClusterDefinitionsByWorkspaceOK creates a ListClusterDefinitionsByWorkspaceOK with default headers values
func NewListClusterDefinitionsByWorkspaceOK() *ListClusterDefinitionsByWorkspaceOK {
	return &ListClusterDefinitionsByWorkspaceOK{}
}

/*ListClusterDefinitionsByWorkspaceOK handles this case with default header values.

successful operation
*/
type ListClusterDefinitionsByWorkspaceOK struct {
	Payload *model.ClusterDefinitionV4ViewResponses
}

func (o *ListClusterDefinitionsByWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/cluster_definitions][%d] listClusterDefinitionsByWorkspaceOK  %+v", 200, o.Payload)
}

func (o *ListClusterDefinitionsByWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ClusterDefinitionV4ViewResponses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}