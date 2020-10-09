// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// RotateAutoTLSCertificatesByNameReader is a Reader for the RotateAutoTLSCertificatesByName structure.
type RotateAutoTLSCertificatesByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RotateAutoTLSCertificatesByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRotateAutoTLSCertificatesByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRotateAutoTLSCertificatesByNameOK creates a RotateAutoTLSCertificatesByNameOK with default headers values
func NewRotateAutoTLSCertificatesByNameOK() *RotateAutoTLSCertificatesByNameOK {
	return &RotateAutoTLSCertificatesByNameOK{}
}

/*RotateAutoTLSCertificatesByNameOK handles this case with default header values.

successful operation
*/
type RotateAutoTLSCertificatesByNameOK struct {
	Payload *model.CertificatesRotationV4Response
}

func (o *RotateAutoTLSCertificatesByNameOK) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/name/{name}/rotate_autotls_certificates][%d] rotateAutoTlsCertificatesByNameOK  %+v", 200, o.Payload)
}

func (o *RotateAutoTLSCertificatesByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CertificatesRotationV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}