// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// SetRangerCloudIdentityMappingReader is a Reader for the SetRangerCloudIdentityMapping structure.
type SetRangerCloudIdentityMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRangerCloudIdentityMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetRangerCloudIdentityMappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetRangerCloudIdentityMappingOK creates a SetRangerCloudIdentityMappingOK with default headers values
func NewSetRangerCloudIdentityMappingOK() *SetRangerCloudIdentityMappingOK {
	return &SetRangerCloudIdentityMappingOK{}
}

/*SetRangerCloudIdentityMappingOK handles this case with default header values.

successful operation
*/
type SetRangerCloudIdentityMappingOK struct {
	Payload *model.RangerCloudIdentitySyncStatus
}

func (o *SetRangerCloudIdentityMappingOK) Error() string {
	return fmt.Sprintf("[POST /sdx/envcrn/{envCrn}/ranger_cloud_identity_mapping][%d] setRangerCloudIdentityMappingOK  %+v", 200, o.Payload)
}

func (o *SetRangerCloudIdentityMappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RangerCloudIdentitySyncStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}