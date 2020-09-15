// Code generated by go-swagger; DO NOT EDIT.

package diagnostics

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

// NewGetSdxCmRolesParams creates a new GetSdxCmRolesParams object
// with the default values initialized.
func NewGetSdxCmRolesParams() *GetSdxCmRolesParams {
	var ()
	return &GetSdxCmRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSdxCmRolesParamsWithTimeout creates a new GetSdxCmRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSdxCmRolesParamsWithTimeout(timeout time.Duration) *GetSdxCmRolesParams {
	var ()
	return &GetSdxCmRolesParams{

		timeout: timeout,
	}
}

// NewGetSdxCmRolesParamsWithContext creates a new GetSdxCmRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSdxCmRolesParamsWithContext(ctx context.Context) *GetSdxCmRolesParams {
	var ()
	return &GetSdxCmRolesParams{

		Context: ctx,
	}
}

// NewGetSdxCmRolesParamsWithHTTPClient creates a new GetSdxCmRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSdxCmRolesParamsWithHTTPClient(client *http.Client) *GetSdxCmRolesParams {
	var ()
	return &GetSdxCmRolesParams{
		HTTPClient: client,
	}
}

/*GetSdxCmRolesParams contains all the parameters to send to the API endpoint
for the get sdx cm roles operation typically these are written to a http.Request
*/
type GetSdxCmRolesParams struct {

	/*StackCrn*/
	StackCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sdx cm roles params
func (o *GetSdxCmRolesParams) WithTimeout(timeout time.Duration) *GetSdxCmRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sdx cm roles params
func (o *GetSdxCmRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sdx cm roles params
func (o *GetSdxCmRolesParams) WithContext(ctx context.Context) *GetSdxCmRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sdx cm roles params
func (o *GetSdxCmRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sdx cm roles params
func (o *GetSdxCmRolesParams) WithHTTPClient(client *http.Client) *GetSdxCmRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sdx cm roles params
func (o *GetSdxCmRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStackCrn adds the stackCrn to the get sdx cm roles params
func (o *GetSdxCmRolesParams) WithStackCrn(stackCrn string) *GetSdxCmRolesParams {
	o.SetStackCrn(stackCrn)
	return o
}

// SetStackCrn adds the stackCrn to the get sdx cm roles params
func (o *GetSdxCmRolesParams) SetStackCrn(stackCrn string) {
	o.StackCrn = stackCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetSdxCmRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stackCrn
	if err := r.SetPathParam("stackCrn", o.StackCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
