// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/api/model"
)

// PostPrivateCredentialReader is a Reader for the PostPrivateCredential structure.
type PostPrivateCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPrivateCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPrivateCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPrivateCredentialOK creates a PostPrivateCredentialOK with default headers values
func NewPostPrivateCredentialOK() *PostPrivateCredentialOK {
	return &PostPrivateCredentialOK{}
}

/*PostPrivateCredentialOK handles this case with default header values.

successful operation
*/
type PostPrivateCredentialOK struct {
	Payload *model.CredentialResponse
}

func (o *PostPrivateCredentialOK) Error() string {
	return fmt.Sprintf("[POST /v1/credentials/user][%d] postPrivateCredentialOK  %+v", 200, o.Payload)
}

func (o *PostPrivateCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}