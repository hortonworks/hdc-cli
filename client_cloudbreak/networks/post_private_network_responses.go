// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// PostPrivateNetworkReader is a Reader for the PostPrivateNetwork structure.
type PostPrivateNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPrivateNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPrivateNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPrivateNetworkOK creates a PostPrivateNetworkOK with default headers values
func NewPostPrivateNetworkOK() *PostPrivateNetworkOK {
	return &PostPrivateNetworkOK{}
}

/*PostPrivateNetworkOK handles this case with default header values.

successful operation
*/
type PostPrivateNetworkOK struct {
	Payload *models_cloudbreak.NetworkResponse
}

func (o *PostPrivateNetworkOK) Error() string {
	return fmt.Sprintf("[POST /networks/user][%d] postPrivateNetworkOK  %+v", 200, o.Payload)
}

func (o *PostPrivateNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.NetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
