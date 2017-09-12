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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPlatformsParams creates a new GetPlatformsParams object
// with the default values initialized.
func NewGetPlatformsParams() *GetPlatformsParams {
	var ()
	return &GetPlatformsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlatformsParamsWithTimeout creates a new GetPlatformsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlatformsParamsWithTimeout(timeout time.Duration) *GetPlatformsParams {
	var ()
	return &GetPlatformsParams{

		timeout: timeout,
	}
}

// NewGetPlatformsParamsWithContext creates a new GetPlatformsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlatformsParamsWithContext(ctx context.Context) *GetPlatformsParams {
	var ()
	return &GetPlatformsParams{

		Context: ctx,
	}
}

// NewGetPlatformsParamsWithHTTPClient creates a new GetPlatformsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlatformsParamsWithHTTPClient(client *http.Client) *GetPlatformsParams {
	var ()
	return &GetPlatformsParams{
		HTTPClient: client,
	}
}

/*GetPlatformsParams contains all the parameters to send to the API endpoint
for the get platforms operation typically these are written to a http.Request
*/
type GetPlatformsParams struct {

	/*Extended*/
	Extended *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get platforms params
func (o *GetPlatformsParams) WithTimeout(timeout time.Duration) *GetPlatformsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get platforms params
func (o *GetPlatformsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get platforms params
func (o *GetPlatformsParams) WithContext(ctx context.Context) *GetPlatformsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get platforms params
func (o *GetPlatformsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get platforms params
func (o *GetPlatformsParams) WithHTTPClient(client *http.Client) *GetPlatformsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get platforms params
func (o *GetPlatformsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtended adds the extended to the get platforms params
func (o *GetPlatformsParams) WithExtended(extended *bool) *GetPlatformsParams {
	o.SetExtended(extended)
	return o
}

// SetExtended adds the extended to the get platforms params
func (o *GetPlatformsParams) SetExtended(extended *bool) {
	o.Extended = extended
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlatformsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Extended != nil {

		// query param extended
		var qrExtended bool
		if o.Extended != nil {
			qrExtended = *o.Extended
		}
		qExtended := swag.FormatBool(qrExtended)
		if qExtended != "" {
			if err := r.SetQueryParam("extended", qExtended); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
