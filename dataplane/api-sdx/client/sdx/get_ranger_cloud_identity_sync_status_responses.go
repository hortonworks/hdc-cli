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

// GetRangerCloudIdentitySyncStatusReader is a Reader for the GetRangerCloudIdentitySyncStatus structure.
type GetRangerCloudIdentitySyncStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRangerCloudIdentitySyncStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRangerCloudIdentitySyncStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRangerCloudIdentitySyncStatusOK creates a GetRangerCloudIdentitySyncStatusOK with default headers values
func NewGetRangerCloudIdentitySyncStatusOK() *GetRangerCloudIdentitySyncStatusOK {
	return &GetRangerCloudIdentitySyncStatusOK{}
}

/*GetRangerCloudIdentitySyncStatusOK handles this case with default header values.

successful operation
*/
type GetRangerCloudIdentitySyncStatusOK struct {
	Payload *model.RangerCloudIdentitySyncStatus
}

func (o *GetRangerCloudIdentitySyncStatusOK) Error() string {
	return fmt.Sprintf("[GET /sdx/envcrn/{envCrn}/ranger_cloud_identity_sync_status][%d] getRangerCloudIdentitySyncStatusOK  %+v", 200, o.Payload)
}

func (o *GetRangerCloudIdentitySyncStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RangerCloudIdentitySyncStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
