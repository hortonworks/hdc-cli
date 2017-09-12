// Code generated by go-swagger; DO NOT EDIT.

package connectors

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

// NewGetRegionRByTypeParams creates a new GetRegionRByTypeParams object
// with the default values initialized.
func NewGetRegionRByTypeParams() *GetRegionRByTypeParams {
	var ()
	return &GetRegionRByTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegionRByTypeParamsWithTimeout creates a new GetRegionRByTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRegionRByTypeParamsWithTimeout(timeout time.Duration) *GetRegionRByTypeParams {
	var ()
	return &GetRegionRByTypeParams{

		timeout: timeout,
	}
}

// NewGetRegionRByTypeParamsWithContext creates a new GetRegionRByTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRegionRByTypeParamsWithContext(ctx context.Context) *GetRegionRByTypeParams {
	var ()
	return &GetRegionRByTypeParams{

		Context: ctx,
	}
}

// NewGetRegionRByTypeParamsWithHTTPClient creates a new GetRegionRByTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRegionRByTypeParamsWithHTTPClient(client *http.Client) *GetRegionRByTypeParams {
	var ()
	return &GetRegionRByTypeParams{
		HTTPClient: client,
	}
}

/*GetRegionRByTypeParams contains all the parameters to send to the API endpoint
for the get region r by type operation typically these are written to a http.Request
*/
type GetRegionRByTypeParams struct {

	/*Type*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get region r by type params
func (o *GetRegionRByTypeParams) WithTimeout(timeout time.Duration) *GetRegionRByTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get region r by type params
func (o *GetRegionRByTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get region r by type params
func (o *GetRegionRByTypeParams) WithContext(ctx context.Context) *GetRegionRByTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get region r by type params
func (o *GetRegionRByTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get region r by type params
func (o *GetRegionRByTypeParams) WithHTTPClient(client *http.Client) *GetRegionRByTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get region r by type params
func (o *GetRegionRByTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithType adds the typeVar to the get region r by type params
func (o *GetRegionRByTypeParams) WithType(typeVar string) *GetRegionRByTypeParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get region r by type params
func (o *GetRegionRByTypeParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegionRByTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
