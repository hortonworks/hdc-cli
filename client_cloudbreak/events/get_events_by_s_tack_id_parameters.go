// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewGetEventsBySTackIDParams creates a new GetEventsBySTackIDParams object
// with the default values initialized.
func NewGetEventsBySTackIDParams() *GetEventsBySTackIDParams {
	var ()
	return &GetEventsBySTackIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventsBySTackIDParamsWithTimeout creates a new GetEventsBySTackIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventsBySTackIDParamsWithTimeout(timeout time.Duration) *GetEventsBySTackIDParams {
	var ()
	return &GetEventsBySTackIDParams{

		timeout: timeout,
	}
}

// NewGetEventsBySTackIDParamsWithContext creates a new GetEventsBySTackIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventsBySTackIDParamsWithContext(ctx context.Context) *GetEventsBySTackIDParams {
	var ()
	return &GetEventsBySTackIDParams{

		Context: ctx,
	}
}

// NewGetEventsBySTackIDParamsWithHTTPClient creates a new GetEventsBySTackIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventsBySTackIDParamsWithHTTPClient(client *http.Client) *GetEventsBySTackIDParams {
	var ()
	return &GetEventsBySTackIDParams{
		HTTPClient: client,
	}
}

/*GetEventsBySTackIDParams contains all the parameters to send to the API endpoint
for the get events by s tack Id operation typically these are written to a http.Request
*/
type GetEventsBySTackIDParams struct {

	/*StackID*/
	StackID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get events by s tack Id params
func (o *GetEventsBySTackIDParams) WithTimeout(timeout time.Duration) *GetEventsBySTackIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get events by s tack Id params
func (o *GetEventsBySTackIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get events by s tack Id params
func (o *GetEventsBySTackIDParams) WithContext(ctx context.Context) *GetEventsBySTackIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get events by s tack Id params
func (o *GetEventsBySTackIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get events by s tack Id params
func (o *GetEventsBySTackIDParams) WithHTTPClient(client *http.Client) *GetEventsBySTackIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get events by s tack Id params
func (o *GetEventsBySTackIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStackID adds the stackID to the get events by s tack Id params
func (o *GetEventsBySTackIDParams) WithStackID(stackID int64) *GetEventsBySTackIDParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get events by s tack Id params
func (o *GetEventsBySTackIDParams) SetStackID(stackID int64) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventsBySTackIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stackId
	if err := r.SetPathParam("stackId", swag.FormatInt64(o.StackID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
