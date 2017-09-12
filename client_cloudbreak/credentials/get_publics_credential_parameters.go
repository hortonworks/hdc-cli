// Code generated by go-swagger; DO NOT EDIT.

package credentials

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

// NewGetPublicsCredentialParams creates a new GetPublicsCredentialParams object
// with the default values initialized.
func NewGetPublicsCredentialParams() *GetPublicsCredentialParams {

	return &GetPublicsCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicsCredentialParamsWithTimeout creates a new GetPublicsCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicsCredentialParamsWithTimeout(timeout time.Duration) *GetPublicsCredentialParams {

	return &GetPublicsCredentialParams{

		timeout: timeout,
	}
}

// NewGetPublicsCredentialParamsWithContext creates a new GetPublicsCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicsCredentialParamsWithContext(ctx context.Context) *GetPublicsCredentialParams {

	return &GetPublicsCredentialParams{

		Context: ctx,
	}
}

// NewGetPublicsCredentialParamsWithHTTPClient creates a new GetPublicsCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicsCredentialParamsWithHTTPClient(client *http.Client) *GetPublicsCredentialParams {

	return &GetPublicsCredentialParams{
		HTTPClient: client,
	}
}

/*GetPublicsCredentialParams contains all the parameters to send to the API endpoint
for the get publics credential operation typically these are written to a http.Request
*/
type GetPublicsCredentialParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get publics credential params
func (o *GetPublicsCredentialParams) WithTimeout(timeout time.Duration) *GetPublicsCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get publics credential params
func (o *GetPublicsCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get publics credential params
func (o *GetPublicsCredentialParams) WithContext(ctx context.Context) *GetPublicsCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get publics credential params
func (o *GetPublicsCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get publics credential params
func (o *GetPublicsCredentialParams) WithHTTPClient(client *http.Client) *GetPublicsCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get publics credential params
func (o *GetPublicsCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicsCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
