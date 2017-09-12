// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_autoscale"
)

// GetHistoryByIDReader is a Reader for the GetHistoryByID structure.
type GetHistoryByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHistoryByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHistoryByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetHistoryByIDOK creates a GetHistoryByIDOK with default headers values
func NewGetHistoryByIDOK() *GetHistoryByIDOK {
	return &GetHistoryByIDOK{}
}

/*GetHistoryByIDOK handles this case with default header values.

successful operation
*/
type GetHistoryByIDOK struct {
	Payload *models_autoscale.History
}

func (o *GetHistoryByIDOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{clusterId}/history/{historyId}][%d] getHistoryByIdOK  %+v", 200, o.Payload)
}

func (o *GetHistoryByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_autoscale.History)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
