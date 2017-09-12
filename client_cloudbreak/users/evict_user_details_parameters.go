// Code generated by go-swagger; DO NOT EDIT.

package users

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

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// NewEvictUserDetailsParams creates a new EvictUserDetailsParams object
// with the default values initialized.
func NewEvictUserDetailsParams() *EvictUserDetailsParams {
	var ()
	return &EvictUserDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEvictUserDetailsParamsWithTimeout creates a new EvictUserDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEvictUserDetailsParamsWithTimeout(timeout time.Duration) *EvictUserDetailsParams {
	var ()
	return &EvictUserDetailsParams{

		timeout: timeout,
	}
}

// NewEvictUserDetailsParamsWithContext creates a new EvictUserDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewEvictUserDetailsParamsWithContext(ctx context.Context) *EvictUserDetailsParams {
	var ()
	return &EvictUserDetailsParams{

		Context: ctx,
	}
}

// NewEvictUserDetailsParamsWithHTTPClient creates a new EvictUserDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEvictUserDetailsParamsWithHTTPClient(client *http.Client) *EvictUserDetailsParams {
	var ()
	return &EvictUserDetailsParams{
		HTTPClient: client,
	}
}

/*EvictUserDetailsParams contains all the parameters to send to the API endpoint
for the evict user details operation typically these are written to a http.Request
*/
type EvictUserDetailsParams struct {

	/*Body*/
	Body *models_cloudbreak.User
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the evict user details params
func (o *EvictUserDetailsParams) WithTimeout(timeout time.Duration) *EvictUserDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the evict user details params
func (o *EvictUserDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the evict user details params
func (o *EvictUserDetailsParams) WithContext(ctx context.Context) *EvictUserDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the evict user details params
func (o *EvictUserDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the evict user details params
func (o *EvictUserDetailsParams) WithHTTPClient(client *http.Client) *EvictUserDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the evict user details params
func (o *EvictUserDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the evict user details params
func (o *EvictUserDetailsParams) WithBody(body *models_cloudbreak.User) *EvictUserDetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the evict user details params
func (o *EvictUserDetailsParams) SetBody(body *models_cloudbreak.User) {
	o.Body = body
}

// WithID adds the id to the evict user details params
func (o *EvictUserDetailsParams) WithID(id string) *EvictUserDetailsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the evict user details params
func (o *EvictUserDetailsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EvictUserDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.User)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
