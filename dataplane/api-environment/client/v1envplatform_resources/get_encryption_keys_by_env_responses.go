// Code generated by go-swagger; DO NOT EDIT.

package v1envplatform_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// GetEncryptionKeysByEnvReader is a Reader for the GetEncryptionKeysByEnv structure.
type GetEncryptionKeysByEnvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEncryptionKeysByEnvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEncryptionKeysByEnvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEncryptionKeysByEnvOK creates a GetEncryptionKeysByEnvOK with default headers values
func NewGetEncryptionKeysByEnvOK() *GetEncryptionKeysByEnvOK {
	return &GetEncryptionKeysByEnvOK{}
}

/*GetEncryptionKeysByEnvOK handles this case with default header values.

successful operation
*/
type GetEncryptionKeysByEnvOK struct {
	Payload *model.PlatformEncryptionKeysResponse
}

func (o *GetEncryptionKeysByEnvOK) Error() string {
	return fmt.Sprintf("[GET /v1/env/platform_resources/encryption_keys][%d] getEncryptionKeysByEnvOK  %+v", 200, o.Payload)
}

func (o *GetEncryptionKeysByEnvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.PlatformEncryptionKeysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
