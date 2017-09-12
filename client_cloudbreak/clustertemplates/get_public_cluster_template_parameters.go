// Code generated by go-swagger; DO NOT EDIT.

package clustertemplates

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

// NewGetPublicClusterTemplateParams creates a new GetPublicClusterTemplateParams object
// with the default values initialized.
func NewGetPublicClusterTemplateParams() *GetPublicClusterTemplateParams {
	var ()
	return &GetPublicClusterTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicClusterTemplateParamsWithTimeout creates a new GetPublicClusterTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicClusterTemplateParamsWithTimeout(timeout time.Duration) *GetPublicClusterTemplateParams {
	var ()
	return &GetPublicClusterTemplateParams{

		timeout: timeout,
	}
}

// NewGetPublicClusterTemplateParamsWithContext creates a new GetPublicClusterTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicClusterTemplateParamsWithContext(ctx context.Context) *GetPublicClusterTemplateParams {
	var ()
	return &GetPublicClusterTemplateParams{

		Context: ctx,
	}
}

// NewGetPublicClusterTemplateParamsWithHTTPClient creates a new GetPublicClusterTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicClusterTemplateParamsWithHTTPClient(client *http.Client) *GetPublicClusterTemplateParams {
	var ()
	return &GetPublicClusterTemplateParams{
		HTTPClient: client,
	}
}

/*GetPublicClusterTemplateParams contains all the parameters to send to the API endpoint
for the get public cluster template operation typically these are written to a http.Request
*/
type GetPublicClusterTemplateParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public cluster template params
func (o *GetPublicClusterTemplateParams) WithTimeout(timeout time.Duration) *GetPublicClusterTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public cluster template params
func (o *GetPublicClusterTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public cluster template params
func (o *GetPublicClusterTemplateParams) WithContext(ctx context.Context) *GetPublicClusterTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public cluster template params
func (o *GetPublicClusterTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public cluster template params
func (o *GetPublicClusterTemplateParams) WithHTTPClient(client *http.Client) *GetPublicClusterTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public cluster template params
func (o *GetPublicClusterTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get public cluster template params
func (o *GetPublicClusterTemplateParams) WithName(name string) *GetPublicClusterTemplateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get public cluster template params
func (o *GetPublicClusterTemplateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicClusterTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
