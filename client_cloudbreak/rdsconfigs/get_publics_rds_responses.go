// Code generated by go-swagger; DO NOT EDIT.

package rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetPublicsRdsReader is a Reader for the GetPublicsRds structure.
type GetPublicsRdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicsRdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsRdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsRdsOK creates a GetPublicsRdsOK with default headers values
func NewGetPublicsRdsOK() *GetPublicsRdsOK {
	return &GetPublicsRdsOK{}
}

/*GetPublicsRdsOK handles this case with default header values.

successful operation
*/
type GetPublicsRdsOK struct {
	Payload []*models_cloudbreak.RDSConfigResponse
}

func (o *GetPublicsRdsOK) Error() string {
	return fmt.Sprintf("[GET /rdsconfigs/account][%d] getPublicsRdsOK  %+v", 200, o.Payload)
}

func (o *GetPublicsRdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
