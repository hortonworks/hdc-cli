// Code generated by go-swagger; DO NOT EDIT.

package accountpreferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetAccountPreferencesEndpointReader is a Reader for the GetAccountPreferencesEndpoint structure.
type GetAccountPreferencesEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountPreferencesEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountPreferencesEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountPreferencesEndpointOK creates a GetAccountPreferencesEndpointOK with default headers values
func NewGetAccountPreferencesEndpointOK() *GetAccountPreferencesEndpointOK {
	return &GetAccountPreferencesEndpointOK{}
}

/*GetAccountPreferencesEndpointOK handles this case with default header values.

successful operation
*/
type GetAccountPreferencesEndpointOK struct {
	Payload *models_cloudbreak.AccountPreference
}

func (o *GetAccountPreferencesEndpointOK) Error() string {
	return fmt.Sprintf("[GET /accountpreferences][%d] getAccountPreferencesEndpointOK  %+v", 200, o.Payload)
}

func (o *GetAccountPreferencesEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.AccountPreference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
