// Code generated by go-swagger; DO NOT EDIT.

package usages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetDeployerUsageReader is a Reader for the GetDeployerUsage structure.
type GetDeployerUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeployerUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeployerUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeployerUsageOK creates a GetDeployerUsageOK with default headers values
func NewGetDeployerUsageOK() *GetDeployerUsageOK {
	return &GetDeployerUsageOK{}
}

/*GetDeployerUsageOK handles this case with default header values.

successful operation
*/
type GetDeployerUsageOK struct {
	Payload []*models_cloudbreak.CloudbreakUsage
}

func (o *GetDeployerUsageOK) Error() string {
	return fmt.Sprintf("[GET /usages][%d] getDeployerUsageOK  %+v", 200, o.Payload)
}

func (o *GetDeployerUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
