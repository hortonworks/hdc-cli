// Code generated by go-swagger; DO NOT EDIT.

package flexsubscriptions

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

// NewGetFlexSubscriptionByIDParams creates a new GetFlexSubscriptionByIDParams object
// with the default values initialized.
func NewGetFlexSubscriptionByIDParams() *GetFlexSubscriptionByIDParams {
	var ()
	return &GetFlexSubscriptionByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlexSubscriptionByIDParamsWithTimeout creates a new GetFlexSubscriptionByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFlexSubscriptionByIDParamsWithTimeout(timeout time.Duration) *GetFlexSubscriptionByIDParams {
	var ()
	return &GetFlexSubscriptionByIDParams{

		timeout: timeout,
	}
}

// NewGetFlexSubscriptionByIDParamsWithContext creates a new GetFlexSubscriptionByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFlexSubscriptionByIDParamsWithContext(ctx context.Context) *GetFlexSubscriptionByIDParams {
	var ()
	return &GetFlexSubscriptionByIDParams{

		Context: ctx,
	}
}

// NewGetFlexSubscriptionByIDParamsWithHTTPClient creates a new GetFlexSubscriptionByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFlexSubscriptionByIDParamsWithHTTPClient(client *http.Client) *GetFlexSubscriptionByIDParams {
	var ()
	return &GetFlexSubscriptionByIDParams{
		HTTPClient: client,
	}
}

/*GetFlexSubscriptionByIDParams contains all the parameters to send to the API endpoint
for the get flex subscription by Id operation typically these are written to a http.Request
*/
type GetFlexSubscriptionByIDParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get flex subscription by Id params
func (o *GetFlexSubscriptionByIDParams) WithTimeout(timeout time.Duration) *GetFlexSubscriptionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flex subscription by Id params
func (o *GetFlexSubscriptionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flex subscription by Id params
func (o *GetFlexSubscriptionByIDParams) WithContext(ctx context.Context) *GetFlexSubscriptionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flex subscription by Id params
func (o *GetFlexSubscriptionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flex subscription by Id params
func (o *GetFlexSubscriptionByIDParams) WithHTTPClient(client *http.Client) *GetFlexSubscriptionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flex subscription by Id params
func (o *GetFlexSubscriptionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get flex subscription by Id params
func (o *GetFlexSubscriptionByIDParams) WithID(id int64) *GetFlexSubscriptionByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get flex subscription by Id params
func (o *GetFlexSubscriptionByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlexSubscriptionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
