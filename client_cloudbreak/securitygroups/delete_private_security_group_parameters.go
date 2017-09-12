// Code generated by go-swagger; DO NOT EDIT.

package securitygroups

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

// NewDeletePrivateSecurityGroupParams creates a new DeletePrivateSecurityGroupParams object
// with the default values initialized.
func NewDeletePrivateSecurityGroupParams() *DeletePrivateSecurityGroupParams {
	var ()
	return &DeletePrivateSecurityGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePrivateSecurityGroupParamsWithTimeout creates a new DeletePrivateSecurityGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePrivateSecurityGroupParamsWithTimeout(timeout time.Duration) *DeletePrivateSecurityGroupParams {
	var ()
	return &DeletePrivateSecurityGroupParams{

		timeout: timeout,
	}
}

// NewDeletePrivateSecurityGroupParamsWithContext creates a new DeletePrivateSecurityGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePrivateSecurityGroupParamsWithContext(ctx context.Context) *DeletePrivateSecurityGroupParams {
	var ()
	return &DeletePrivateSecurityGroupParams{

		Context: ctx,
	}
}

// NewDeletePrivateSecurityGroupParamsWithHTTPClient creates a new DeletePrivateSecurityGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePrivateSecurityGroupParamsWithHTTPClient(client *http.Client) *DeletePrivateSecurityGroupParams {
	var ()
	return &DeletePrivateSecurityGroupParams{
		HTTPClient: client,
	}
}

/*DeletePrivateSecurityGroupParams contains all the parameters to send to the API endpoint
for the delete private security group operation typically these are written to a http.Request
*/
type DeletePrivateSecurityGroupParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete private security group params
func (o *DeletePrivateSecurityGroupParams) WithTimeout(timeout time.Duration) *DeletePrivateSecurityGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete private security group params
func (o *DeletePrivateSecurityGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete private security group params
func (o *DeletePrivateSecurityGroupParams) WithContext(ctx context.Context) *DeletePrivateSecurityGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete private security group params
func (o *DeletePrivateSecurityGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete private security group params
func (o *DeletePrivateSecurityGroupParams) WithHTTPClient(client *http.Client) *DeletePrivateSecurityGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete private security group params
func (o *DeletePrivateSecurityGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete private security group params
func (o *DeletePrivateSecurityGroupParams) WithName(name string) *DeletePrivateSecurityGroupParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete private security group params
func (o *DeletePrivateSecurityGroupParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePrivateSecurityGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
