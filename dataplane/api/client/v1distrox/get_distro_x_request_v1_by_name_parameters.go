// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDistroXRequestV1ByNameParams creates a new GetDistroXRequestV1ByNameParams object
// with the default values initialized.
func NewGetDistroXRequestV1ByNameParams() *GetDistroXRequestV1ByNameParams {
	var ()
	return &GetDistroXRequestV1ByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDistroXRequestV1ByNameParamsWithTimeout creates a new GetDistroXRequestV1ByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDistroXRequestV1ByNameParamsWithTimeout(timeout time.Duration) *GetDistroXRequestV1ByNameParams {
	var ()
	return &GetDistroXRequestV1ByNameParams{

		timeout: timeout,
	}
}

// NewGetDistroXRequestV1ByNameParamsWithContext creates a new GetDistroXRequestV1ByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDistroXRequestV1ByNameParamsWithContext(ctx context.Context) *GetDistroXRequestV1ByNameParams {
	var ()
	return &GetDistroXRequestV1ByNameParams{

		Context: ctx,
	}
}

// NewGetDistroXRequestV1ByNameParamsWithHTTPClient creates a new GetDistroXRequestV1ByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDistroXRequestV1ByNameParamsWithHTTPClient(client *http.Client) *GetDistroXRequestV1ByNameParams {
	var ()
	return &GetDistroXRequestV1ByNameParams{
		HTTPClient: client,
	}
}

/*GetDistroXRequestV1ByNameParams contains all the parameters to send to the API endpoint
for the get distro x request v1 by name operation typically these are written to a http.Request
*/
type GetDistroXRequestV1ByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get distro x request v1 by name params
func (o *GetDistroXRequestV1ByNameParams) WithTimeout(timeout time.Duration) *GetDistroXRequestV1ByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get distro x request v1 by name params
func (o *GetDistroXRequestV1ByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get distro x request v1 by name params
func (o *GetDistroXRequestV1ByNameParams) WithContext(ctx context.Context) *GetDistroXRequestV1ByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get distro x request v1 by name params
func (o *GetDistroXRequestV1ByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get distro x request v1 by name params
func (o *GetDistroXRequestV1ByNameParams) WithHTTPClient(client *http.Client) *GetDistroXRequestV1ByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get distro x request v1 by name params
func (o *GetDistroXRequestV1ByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get distro x request v1 by name params
func (o *GetDistroXRequestV1ByNameParams) WithName(name string) *GetDistroXRequestV1ByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get distro x request v1 by name params
func (o *GetDistroXRequestV1ByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetDistroXRequestV1ByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}