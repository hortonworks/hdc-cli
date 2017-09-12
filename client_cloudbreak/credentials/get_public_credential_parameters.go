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

// NewGetPublicCredentialParams creates a new GetPublicCredentialParams object
// with the default values initialized.
func NewGetPublicCredentialParams() *GetPublicCredentialParams {
	var ()
	return &GetPublicCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicCredentialParamsWithTimeout creates a new GetPublicCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicCredentialParamsWithTimeout(timeout time.Duration) *GetPublicCredentialParams {
	var ()
	return &GetPublicCredentialParams{

		timeout: timeout,
	}
}

// NewGetPublicCredentialParamsWithContext creates a new GetPublicCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicCredentialParamsWithContext(ctx context.Context) *GetPublicCredentialParams {
	var ()
	return &GetPublicCredentialParams{

		Context: ctx,
	}
}

// NewGetPublicCredentialParamsWithHTTPClient creates a new GetPublicCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicCredentialParamsWithHTTPClient(client *http.Client) *GetPublicCredentialParams {
	var ()
	return &GetPublicCredentialParams{
		HTTPClient: client,
	}
}

/*GetPublicCredentialParams contains all the parameters to send to the API endpoint
for the get public credential operation typically these are written to a http.Request
*/
type GetPublicCredentialParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public credential params
func (o *GetPublicCredentialParams) WithTimeout(timeout time.Duration) *GetPublicCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public credential params
func (o *GetPublicCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public credential params
func (o *GetPublicCredentialParams) WithContext(ctx context.Context) *GetPublicCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public credential params
func (o *GetPublicCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public credential params
func (o *GetPublicCredentialParams) WithHTTPClient(client *http.Client) *GetPublicCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public credential params
func (o *GetPublicCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get public credential params
func (o *GetPublicCredentialParams) WithName(name string) *GetPublicCredentialParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get public credential params
func (o *GetPublicCredentialParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
