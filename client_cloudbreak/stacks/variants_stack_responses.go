package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// VariantsStackReader is a Reader for the VariantsStack structure.
type VariantsStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *VariantsStackReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVariantsStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVariantsStackOK creates a VariantsStackOK with default headers values
func NewVariantsStackOK() *VariantsStackOK {
	return &VariantsStackOK{}
}

/*VariantsStackOK handles this case with default header values.

successful operation
*/
type VariantsStackOK struct {
	Payload *models_cloudbreak.PlatformVariantsJSON
}

func (o *VariantsStackOK) Error() string {
	return fmt.Sprintf("[GET /stacks/platformVariants][%d] variantsStackOK  %+v", 200, o.Payload)
}

func (o *VariantsStackOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.PlatformVariantsJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}