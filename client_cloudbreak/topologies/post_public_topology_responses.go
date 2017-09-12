// Code generated by go-swagger; DO NOT EDIT.

package topologies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// PostPublicTopologyReader is a Reader for the PostPublicTopology structure.
type PostPublicTopologyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPublicTopologyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPublicTopologyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPublicTopologyOK creates a PostPublicTopologyOK with default headers values
func NewPostPublicTopologyOK() *PostPublicTopologyOK {
	return &PostPublicTopologyOK{}
}

/*PostPublicTopologyOK handles this case with default header values.

successful operation
*/
type PostPublicTopologyOK struct {
	Payload *models_cloudbreak.TopologyResponse
}

func (o *PostPublicTopologyOK) Error() string {
	return fmt.Sprintf("[POST /topologies/account][%d] postPublicTopologyOK  %+v", 200, o.Payload)
}

func (o *PostPublicTopologyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.TopologyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
