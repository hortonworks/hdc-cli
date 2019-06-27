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

// GetSdxDetailReader is a Reader for the GetSdxDetail structure.
type GetSdxDetailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSdxDetailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSdxDetailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSdxDetailOK creates a GetSdxDetailOK with default headers values
func NewGetSdxDetailOK() *GetSdxDetailOK {
	return &GetSdxDetailOK{}
}

/*GetSdxDetailOK handles this case with default header values.

successful operation
*/
type GetSdxDetailOK struct {
	Payload *model.SdxClusterDetailResponse
}

func (o *GetSdxDetailOK) Error() string {
	return fmt.Sprintf("[GET /sdx/{name}/detail][%d] getSdxDetailOK  %+v", 200, o.Payload)
}

func (o *GetSdxDetailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SdxClusterDetailResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
