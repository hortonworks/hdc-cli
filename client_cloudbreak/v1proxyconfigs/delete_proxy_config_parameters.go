// Code generated by go-swagger; DO NOT EDIT.

package v1proxyconfigs

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

// NewDeleteProxyConfigParams creates a new DeleteProxyConfigParams object
// with the default values initialized.
func NewDeleteProxyConfigParams() *DeleteProxyConfigParams {
	var ()
	return &DeleteProxyConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProxyConfigParamsWithTimeout creates a new DeleteProxyConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProxyConfigParamsWithTimeout(timeout time.Duration) *DeleteProxyConfigParams {
	var ()
	return &DeleteProxyConfigParams{

		timeout: timeout,
	}
}

// NewDeleteProxyConfigParamsWithContext creates a new DeleteProxyConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProxyConfigParamsWithContext(ctx context.Context) *DeleteProxyConfigParams {
	var ()
	return &DeleteProxyConfigParams{

		Context: ctx,
	}
}

// NewDeleteProxyConfigParamsWithHTTPClient creates a new DeleteProxyConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProxyConfigParamsWithHTTPClient(client *http.Client) *DeleteProxyConfigParams {
	var ()
	return &DeleteProxyConfigParams{
		HTTPClient: client,
	}
}

/*DeleteProxyConfigParams contains all the parameters to send to the API endpoint
for the delete proxy config operation typically these are written to a http.Request
*/
type DeleteProxyConfigParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete proxy config params
func (o *DeleteProxyConfigParams) WithTimeout(timeout time.Duration) *DeleteProxyConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete proxy config params
func (o *DeleteProxyConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete proxy config params
func (o *DeleteProxyConfigParams) WithContext(ctx context.Context) *DeleteProxyConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete proxy config params
func (o *DeleteProxyConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete proxy config params
func (o *DeleteProxyConfigParams) WithHTTPClient(client *http.Client) *DeleteProxyConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete proxy config params
func (o *DeleteProxyConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete proxy config params
func (o *DeleteProxyConfigParams) WithID(id int64) *DeleteProxyConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete proxy config params
func (o *DeleteProxyConfigParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProxyConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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