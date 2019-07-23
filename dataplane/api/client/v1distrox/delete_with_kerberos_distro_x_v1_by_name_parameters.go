// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteWithKerberosDistroXV1ByNameParams creates a new DeleteWithKerberosDistroXV1ByNameParams object
// with the default values initialized.
func NewDeleteWithKerberosDistroXV1ByNameParams() *DeleteWithKerberosDistroXV1ByNameParams {
	var ()
	return &DeleteWithKerberosDistroXV1ByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWithKerberosDistroXV1ByNameParamsWithTimeout creates a new DeleteWithKerberosDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWithKerberosDistroXV1ByNameParamsWithTimeout(timeout time.Duration) *DeleteWithKerberosDistroXV1ByNameParams {
	var ()
	return &DeleteWithKerberosDistroXV1ByNameParams{

		timeout: timeout,
	}
}

// NewDeleteWithKerberosDistroXV1ByNameParamsWithContext creates a new DeleteWithKerberosDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWithKerberosDistroXV1ByNameParamsWithContext(ctx context.Context) *DeleteWithKerberosDistroXV1ByNameParams {
	var ()
	return &DeleteWithKerberosDistroXV1ByNameParams{

		Context: ctx,
	}
}

// NewDeleteWithKerberosDistroXV1ByNameParamsWithHTTPClient creates a new DeleteWithKerberosDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWithKerberosDistroXV1ByNameParamsWithHTTPClient(client *http.Client) *DeleteWithKerberosDistroXV1ByNameParams {
	var ()
	return &DeleteWithKerberosDistroXV1ByNameParams{
		HTTPClient: client,
	}
}

/*DeleteWithKerberosDistroXV1ByNameParams contains all the parameters to send to the API endpoint
for the delete with kerberos distro x v1 by name operation typically these are written to a http.Request
*/
type DeleteWithKerberosDistroXV1ByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete with kerberos distro x v1 by name params
func (o *DeleteWithKerberosDistroXV1ByNameParams) WithTimeout(timeout time.Duration) *DeleteWithKerberosDistroXV1ByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete with kerberos distro x v1 by name params
func (o *DeleteWithKerberosDistroXV1ByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete with kerberos distro x v1 by name params
func (o *DeleteWithKerberosDistroXV1ByNameParams) WithContext(ctx context.Context) *DeleteWithKerberosDistroXV1ByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete with kerberos distro x v1 by name params
func (o *DeleteWithKerberosDistroXV1ByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete with kerberos distro x v1 by name params
func (o *DeleteWithKerberosDistroXV1ByNameParams) WithHTTPClient(client *http.Client) *DeleteWithKerberosDistroXV1ByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete with kerberos distro x v1 by name params
func (o *DeleteWithKerberosDistroXV1ByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete with kerberos distro x v1 by name params
func (o *DeleteWithKerberosDistroXV1ByNameParams) WithName(name string) *DeleteWithKerberosDistroXV1ByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete with kerberos distro x v1 by name params
func (o *DeleteWithKerberosDistroXV1ByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWithKerberosDistroXV1ByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
