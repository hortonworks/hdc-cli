// Code generated by go-swagger; DO NOT EDIT.

package templates

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

// NewDeletePublicTemplateParams creates a new DeletePublicTemplateParams object
// with the default values initialized.
func NewDeletePublicTemplateParams() *DeletePublicTemplateParams {
	var ()
	return &DeletePublicTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePublicTemplateParamsWithTimeout creates a new DeletePublicTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePublicTemplateParamsWithTimeout(timeout time.Duration) *DeletePublicTemplateParams {
	var ()
	return &DeletePublicTemplateParams{

		timeout: timeout,
	}
}

// NewDeletePublicTemplateParamsWithContext creates a new DeletePublicTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePublicTemplateParamsWithContext(ctx context.Context) *DeletePublicTemplateParams {
	var ()
	return &DeletePublicTemplateParams{

		Context: ctx,
	}
}

// NewDeletePublicTemplateParamsWithHTTPClient creates a new DeletePublicTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePublicTemplateParamsWithHTTPClient(client *http.Client) *DeletePublicTemplateParams {
	var ()
	return &DeletePublicTemplateParams{
		HTTPClient: client,
	}
}

/*DeletePublicTemplateParams contains all the parameters to send to the API endpoint
for the delete public template operation typically these are written to a http.Request
*/
type DeletePublicTemplateParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete public template params
func (o *DeletePublicTemplateParams) WithTimeout(timeout time.Duration) *DeletePublicTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete public template params
func (o *DeletePublicTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete public template params
func (o *DeletePublicTemplateParams) WithContext(ctx context.Context) *DeletePublicTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete public template params
func (o *DeletePublicTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete public template params
func (o *DeletePublicTemplateParams) WithHTTPClient(client *http.Client) *DeletePublicTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete public template params
func (o *DeletePublicTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete public template params
func (o *DeletePublicTemplateParams) WithName(name string) *DeletePublicTemplateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete public template params
func (o *DeletePublicTemplateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePublicTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
