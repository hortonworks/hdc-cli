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

// NewPutDefaultFlexSubscriptionByIDParams creates a new PutDefaultFlexSubscriptionByIDParams object
// with the default values initialized.
func NewPutDefaultFlexSubscriptionByIDParams() *PutDefaultFlexSubscriptionByIDParams {
	var ()
	return &PutDefaultFlexSubscriptionByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDefaultFlexSubscriptionByIDParamsWithTimeout creates a new PutDefaultFlexSubscriptionByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDefaultFlexSubscriptionByIDParamsWithTimeout(timeout time.Duration) *PutDefaultFlexSubscriptionByIDParams {
	var ()
	return &PutDefaultFlexSubscriptionByIDParams{

		timeout: timeout,
	}
}

// NewPutDefaultFlexSubscriptionByIDParamsWithContext creates a new PutDefaultFlexSubscriptionByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDefaultFlexSubscriptionByIDParamsWithContext(ctx context.Context) *PutDefaultFlexSubscriptionByIDParams {
	var ()
	return &PutDefaultFlexSubscriptionByIDParams{

		Context: ctx,
	}
}

// NewPutDefaultFlexSubscriptionByIDParamsWithHTTPClient creates a new PutDefaultFlexSubscriptionByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDefaultFlexSubscriptionByIDParamsWithHTTPClient(client *http.Client) *PutDefaultFlexSubscriptionByIDParams {
	var ()
	return &PutDefaultFlexSubscriptionByIDParams{
		HTTPClient: client,
	}
}

/*PutDefaultFlexSubscriptionByIDParams contains all the parameters to send to the API endpoint
for the put default flex subscription by Id operation typically these are written to a http.Request
*/
type PutDefaultFlexSubscriptionByIDParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put default flex subscription by Id params
func (o *PutDefaultFlexSubscriptionByIDParams) WithTimeout(timeout time.Duration) *PutDefaultFlexSubscriptionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put default flex subscription by Id params
func (o *PutDefaultFlexSubscriptionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put default flex subscription by Id params
func (o *PutDefaultFlexSubscriptionByIDParams) WithContext(ctx context.Context) *PutDefaultFlexSubscriptionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put default flex subscription by Id params
func (o *PutDefaultFlexSubscriptionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put default flex subscription by Id params
func (o *PutDefaultFlexSubscriptionByIDParams) WithHTTPClient(client *http.Client) *PutDefaultFlexSubscriptionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put default flex subscription by Id params
func (o *PutDefaultFlexSubscriptionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put default flex subscription by Id params
func (o *PutDefaultFlexSubscriptionByIDParams) WithID(id int64) *PutDefaultFlexSubscriptionByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put default flex subscription by Id params
func (o *PutDefaultFlexSubscriptionByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDefaultFlexSubscriptionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
