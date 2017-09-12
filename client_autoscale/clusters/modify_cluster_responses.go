// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_autoscale"
)

// ModifyClusterReader is a Reader for the ModifyCluster structure.
type ModifyClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyClusterOK creates a ModifyClusterOK with default headers values
func NewModifyClusterOK() *ModifyClusterOK {
	return &ModifyClusterOK{}
}

/*ModifyClusterOK handles this case with default header values.

successful operation
*/
type ModifyClusterOK struct {
	Payload *models_autoscale.ClusterSummary
}

func (o *ModifyClusterOK) Error() string {
	return fmt.Sprintf("[PUT /clusters/{clusterId}][%d] modifyClusterOK  %+v", 200, o.Payload)
}

func (o *ModifyClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_autoscale.ClusterSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
