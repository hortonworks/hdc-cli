// Code generated by go-swagger; DO NOT EDIT.

package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetPublicsTemplateReader is a Reader for the GetPublicsTemplate structure.
type GetPublicsTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicsTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsTemplateOK creates a GetPublicsTemplateOK with default headers values
func NewGetPublicsTemplateOK() *GetPublicsTemplateOK {
	return &GetPublicsTemplateOK{}
}

/*GetPublicsTemplateOK handles this case with default header values.

successful operation
*/
type GetPublicsTemplateOK struct {
	Payload []*models_cloudbreak.TemplateResponse
}

func (o *GetPublicsTemplateOK) Error() string {
	return fmt.Sprintf("[GET /templates/account][%d] getPublicsTemplateOK  %+v", 200, o.Payload)
}

func (o *GetPublicsTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
