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

// NewCancelDistroxDiagnosticsCollectionsV1Params creates a new CancelDistroxDiagnosticsCollectionsV1Params object
// with the default values initialized.
func NewCancelDistroxDiagnosticsCollectionsV1Params() *CancelDistroxDiagnosticsCollectionsV1Params {
	var ()
	return &CancelDistroxDiagnosticsCollectionsV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelDistroxDiagnosticsCollectionsV1ParamsWithTimeout creates a new CancelDistroxDiagnosticsCollectionsV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelDistroxDiagnosticsCollectionsV1ParamsWithTimeout(timeout time.Duration) *CancelDistroxDiagnosticsCollectionsV1Params {
	var ()
	return &CancelDistroxDiagnosticsCollectionsV1Params{

		timeout: timeout,
	}
}

// NewCancelDistroxDiagnosticsCollectionsV1ParamsWithContext creates a new CancelDistroxDiagnosticsCollectionsV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCancelDistroxDiagnosticsCollectionsV1ParamsWithContext(ctx context.Context) *CancelDistroxDiagnosticsCollectionsV1Params {
	var ()
	return &CancelDistroxDiagnosticsCollectionsV1Params{

		Context: ctx,
	}
}

// NewCancelDistroxDiagnosticsCollectionsV1ParamsWithHTTPClient creates a new CancelDistroxDiagnosticsCollectionsV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelDistroxDiagnosticsCollectionsV1ParamsWithHTTPClient(client *http.Client) *CancelDistroxDiagnosticsCollectionsV1Params {
	var ()
	return &CancelDistroxDiagnosticsCollectionsV1Params{
		HTTPClient: client,
	}
}

/*CancelDistroxDiagnosticsCollectionsV1Params contains all the parameters to send to the API endpoint
for the cancel distrox diagnostics collections v1 operation typically these are written to a http.Request
*/
type CancelDistroxDiagnosticsCollectionsV1Params struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel distrox diagnostics collections v1 params
func (o *CancelDistroxDiagnosticsCollectionsV1Params) WithTimeout(timeout time.Duration) *CancelDistroxDiagnosticsCollectionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel distrox diagnostics collections v1 params
func (o *CancelDistroxDiagnosticsCollectionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel distrox diagnostics collections v1 params
func (o *CancelDistroxDiagnosticsCollectionsV1Params) WithContext(ctx context.Context) *CancelDistroxDiagnosticsCollectionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel distrox diagnostics collections v1 params
func (o *CancelDistroxDiagnosticsCollectionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel distrox diagnostics collections v1 params
func (o *CancelDistroxDiagnosticsCollectionsV1Params) WithHTTPClient(client *http.Client) *CancelDistroxDiagnosticsCollectionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel distrox diagnostics collections v1 params
func (o *CancelDistroxDiagnosticsCollectionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the cancel distrox diagnostics collections v1 params
func (o *CancelDistroxDiagnosticsCollectionsV1Params) WithCrn(crn string) *CancelDistroxDiagnosticsCollectionsV1Params {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the cancel distrox diagnostics collections v1 params
func (o *CancelDistroxDiagnosticsCollectionsV1Params) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *CancelDistroxDiagnosticsCollectionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}