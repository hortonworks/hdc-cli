// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RetryDistroXV1ByCrnReader is a Reader for the RetryDistroXV1ByCrn structure.
type RetryDistroXV1ByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetryDistroXV1ByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewRetryDistroXV1ByCrnDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewRetryDistroXV1ByCrnDefault creates a RetryDistroXV1ByCrnDefault with default headers values
func NewRetryDistroXV1ByCrnDefault(code int) *RetryDistroXV1ByCrnDefault {
	return &RetryDistroXV1ByCrnDefault{
		_statusCode: code,
	}
}

/*RetryDistroXV1ByCrnDefault handles this case with default header values.

successful operation
*/
type RetryDistroXV1ByCrnDefault struct {
	_statusCode int
}

// Code gets the status code for the retry distro x v1 by crn default response
func (o *RetryDistroXV1ByCrnDefault) Code() int {
	return o._statusCode
}

func (o *RetryDistroXV1ByCrnDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/crn/{crn}/retry][%d] retryDistroXV1ByCrn default ", o._statusCode)
}

func (o *RetryDistroXV1ByCrnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}