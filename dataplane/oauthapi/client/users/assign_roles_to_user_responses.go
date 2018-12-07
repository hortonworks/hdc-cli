// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/oauthapi/model"
)

// AssignRolesToUserReader is a Reader for the AssignRolesToUser structure.
type AssignRolesToUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignRolesToUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAssignRolesToUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAssignRolesToUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAssignRolesToUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAssignRolesToUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAssignRolesToUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAssignRolesToUserOK creates a AssignRolesToUserOK with default headers values
func NewAssignRolesToUserOK() *AssignRolesToUserOK {
	return &AssignRolesToUserOK{}
}

/*AssignRolesToUserOK handles this case with default header values.

successful operation
*/
type AssignRolesToUserOK struct {
	Payload []*model.UserRole
}

func (o *AssignRolesToUserOK) Error() string {
	return fmt.Sprintf("[POST /caas/api/users/{userId}/roles][%d] assignRolesToUserOK  %+v", 200, o.Payload)
}

func (o *AssignRolesToUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignRolesToUserBadRequest creates a AssignRolesToUserBadRequest with default headers values
func NewAssignRolesToUserBadRequest() *AssignRolesToUserBadRequest {
	return &AssignRolesToUserBadRequest{}
}

/*AssignRolesToUserBadRequest handles this case with default header values.

Bad Request
*/
type AssignRolesToUserBadRequest struct {
}

func (o *AssignRolesToUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /caas/api/users/{userId}/roles][%d] assignRolesToUserBadRequest ", 400)
}

func (o *AssignRolesToUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignRolesToUserUnauthorized creates a AssignRolesToUserUnauthorized with default headers values
func NewAssignRolesToUserUnauthorized() *AssignRolesToUserUnauthorized {
	return &AssignRolesToUserUnauthorized{}
}

/*AssignRolesToUserUnauthorized handles this case with default header values.

Unauthorized
*/
type AssignRolesToUserUnauthorized struct {
}

func (o *AssignRolesToUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /caas/api/users/{userId}/roles][%d] assignRolesToUserUnauthorized ", 401)
}

func (o *AssignRolesToUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignRolesToUserForbidden creates a AssignRolesToUserForbidden with default headers values
func NewAssignRolesToUserForbidden() *AssignRolesToUserForbidden {
	return &AssignRolesToUserForbidden{}
}

/*AssignRolesToUserForbidden handles this case with default header values.

Forbidden
*/
type AssignRolesToUserForbidden struct {
}

func (o *AssignRolesToUserForbidden) Error() string {
	return fmt.Sprintf("[POST /caas/api/users/{userId}/roles][%d] assignRolesToUserForbidden ", 403)
}

func (o *AssignRolesToUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignRolesToUserInternalServerError creates a AssignRolesToUserInternalServerError with default headers values
func NewAssignRolesToUserInternalServerError() *AssignRolesToUserInternalServerError {
	return &AssignRolesToUserInternalServerError{}
}

/*AssignRolesToUserInternalServerError handles this case with default header values.

Internal server error
*/
type AssignRolesToUserInternalServerError struct {
}

func (o *AssignRolesToUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /caas/api/users/{userId}/roles][%d] assignRolesToUserInternalServerError ", 500)
}

func (o *AssignRolesToUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
