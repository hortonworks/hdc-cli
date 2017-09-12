// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePrivateCredentialReader is a Reader for the DeletePrivateCredential structure.
type DeletePrivateCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePrivateCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePrivateCredentialDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePrivateCredentialDefault creates a DeletePrivateCredentialDefault with default headers values
func NewDeletePrivateCredentialDefault(code int) *DeletePrivateCredentialDefault {
	return &DeletePrivateCredentialDefault{
		_statusCode: code,
	}
}

/*DeletePrivateCredentialDefault handles this case with default header values.

successful operation
*/
type DeletePrivateCredentialDefault struct {
	_statusCode int
}

// Code gets the status code for the delete private credential default response
func (o *DeletePrivateCredentialDefault) Code() int {
	return o._statusCode
}

func (o *DeletePrivateCredentialDefault) Error() string {
	return fmt.Sprintf("[DELETE /credentials/user/{name}][%d] deletePrivateCredential default ", o._statusCode)
}

func (o *DeletePrivateCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
