// Code generated by go-swagger; DO NOT EDIT.

package strategies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/oauthapi/model"
)

// ListStrategiesReader is a Reader for the ListStrategies structure.
type ListStrategiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListStrategiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListStrategiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListStrategiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListStrategiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListStrategiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListStrategiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListStrategiesOK creates a ListStrategiesOK with default headers values
func NewListStrategiesOK() *ListStrategiesOK {
	return &ListStrategiesOK{}
}

/*ListStrategiesOK handles this case with default header values.

successful operation
*/
type ListStrategiesOK struct {
	Payload []*model.Strategy
}

func (o *ListStrategiesOK) Error() string {
	return fmt.Sprintf("[GET /caas/api/strategies][%d] listStrategiesOK  %+v", 200, o.Payload)
}

func (o *ListStrategiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListStrategiesBadRequest creates a ListStrategiesBadRequest with default headers values
func NewListStrategiesBadRequest() *ListStrategiesBadRequest {
	return &ListStrategiesBadRequest{}
}

/*ListStrategiesBadRequest handles this case with default header values.

Bad Request
*/
type ListStrategiesBadRequest struct {
}

func (o *ListStrategiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /caas/api/strategies][%d] listStrategiesBadRequest ", 400)
}

func (o *ListStrategiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListStrategiesUnauthorized creates a ListStrategiesUnauthorized with default headers values
func NewListStrategiesUnauthorized() *ListStrategiesUnauthorized {
	return &ListStrategiesUnauthorized{}
}

/*ListStrategiesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListStrategiesUnauthorized struct {
}

func (o *ListStrategiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /caas/api/strategies][%d] listStrategiesUnauthorized ", 401)
}

func (o *ListStrategiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListStrategiesForbidden creates a ListStrategiesForbidden with default headers values
func NewListStrategiesForbidden() *ListStrategiesForbidden {
	return &ListStrategiesForbidden{}
}

/*ListStrategiesForbidden handles this case with default header values.

Forbidden
*/
type ListStrategiesForbidden struct {
}

func (o *ListStrategiesForbidden) Error() string {
	return fmt.Sprintf("[GET /caas/api/strategies][%d] listStrategiesForbidden ", 403)
}

func (o *ListStrategiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListStrategiesInternalServerError creates a ListStrategiesInternalServerError with default headers values
func NewListStrategiesInternalServerError() *ListStrategiesInternalServerError {
	return &ListStrategiesInternalServerError{}
}

/*ListStrategiesInternalServerError handles this case with default header values.

Internal server error
*/
type ListStrategiesInternalServerError struct {
}

func (o *ListStrategiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /caas/api/strategies][%d] listStrategiesInternalServerError ", 500)
}

func (o *ListStrategiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}