// Code generated by go-swagger; DO NOT EDIT.

package v1proxies

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

// NewGetProxyConfigByEnvironmentCrnV1Params creates a new GetProxyConfigByEnvironmentCrnV1Params object
// with the default values initialized.
func NewGetProxyConfigByEnvironmentCrnV1Params() *GetProxyConfigByEnvironmentCrnV1Params {
	var ()
	return &GetProxyConfigByEnvironmentCrnV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProxyConfigByEnvironmentCrnV1ParamsWithTimeout creates a new GetProxyConfigByEnvironmentCrnV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProxyConfigByEnvironmentCrnV1ParamsWithTimeout(timeout time.Duration) *GetProxyConfigByEnvironmentCrnV1Params {
	var ()
	return &GetProxyConfigByEnvironmentCrnV1Params{

		timeout: timeout,
	}
}

// NewGetProxyConfigByEnvironmentCrnV1ParamsWithContext creates a new GetProxyConfigByEnvironmentCrnV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetProxyConfigByEnvironmentCrnV1ParamsWithContext(ctx context.Context) *GetProxyConfigByEnvironmentCrnV1Params {
	var ()
	return &GetProxyConfigByEnvironmentCrnV1Params{

		Context: ctx,
	}
}

// NewGetProxyConfigByEnvironmentCrnV1ParamsWithHTTPClient creates a new GetProxyConfigByEnvironmentCrnV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProxyConfigByEnvironmentCrnV1ParamsWithHTTPClient(client *http.Client) *GetProxyConfigByEnvironmentCrnV1Params {
	var ()
	return &GetProxyConfigByEnvironmentCrnV1Params{
		HTTPClient: client,
	}
}

/*GetProxyConfigByEnvironmentCrnV1Params contains all the parameters to send to the API endpoint
for the get proxy config by environment crn v1 operation typically these are written to a http.Request
*/
type GetProxyConfigByEnvironmentCrnV1Params struct {

	/*EnvironmentCrn*/
	EnvironmentCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get proxy config by environment crn v1 params
func (o *GetProxyConfigByEnvironmentCrnV1Params) WithTimeout(timeout time.Duration) *GetProxyConfigByEnvironmentCrnV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get proxy config by environment crn v1 params
func (o *GetProxyConfigByEnvironmentCrnV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get proxy config by environment crn v1 params
func (o *GetProxyConfigByEnvironmentCrnV1Params) WithContext(ctx context.Context) *GetProxyConfigByEnvironmentCrnV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get proxy config by environment crn v1 params
func (o *GetProxyConfigByEnvironmentCrnV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get proxy config by environment crn v1 params
func (o *GetProxyConfigByEnvironmentCrnV1Params) WithHTTPClient(client *http.Client) *GetProxyConfigByEnvironmentCrnV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get proxy config by environment crn v1 params
func (o *GetProxyConfigByEnvironmentCrnV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCrn adds the environmentCrn to the get proxy config by environment crn v1 params
func (o *GetProxyConfigByEnvironmentCrnV1Params) WithEnvironmentCrn(environmentCrn string) *GetProxyConfigByEnvironmentCrnV1Params {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get proxy config by environment crn v1 params
func (o *GetProxyConfigByEnvironmentCrnV1Params) SetEnvironmentCrn(environmentCrn string) {
	o.EnvironmentCrn = environmentCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetProxyConfigByEnvironmentCrnV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentCrn
	if err := r.SetPathParam("environmentCrn", o.EnvironmentCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
