// Code generated by go-swagger; DO NOT EDIT.

package dbconfig

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

// NewGetDatalakeDbConfigParams creates a new GetDatalakeDbConfigParams object
// with the default values initialized.
func NewGetDatalakeDbConfigParams() *GetDatalakeDbConfigParams {
	var ()
	return &GetDatalakeDbConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatalakeDbConfigParamsWithTimeout creates a new GetDatalakeDbConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDatalakeDbConfigParamsWithTimeout(timeout time.Duration) *GetDatalakeDbConfigParams {
	var ()
	return &GetDatalakeDbConfigParams{

		timeout: timeout,
	}
}

// NewGetDatalakeDbConfigParamsWithContext creates a new GetDatalakeDbConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDatalakeDbConfigParamsWithContext(ctx context.Context) *GetDatalakeDbConfigParams {
	var ()
	return &GetDatalakeDbConfigParams{

		Context: ctx,
	}
}

// NewGetDatalakeDbConfigParamsWithHTTPClient creates a new GetDatalakeDbConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDatalakeDbConfigParamsWithHTTPClient(client *http.Client) *GetDatalakeDbConfigParams {
	var ()
	return &GetDatalakeDbConfigParams{
		HTTPClient: client,
	}
}

/*GetDatalakeDbConfigParams contains all the parameters to send to the API endpoint
for the get datalake db config operation typically these are written to a http.Request
*/
type GetDatalakeDbConfigParams struct {

	/*DatabaseType*/
	DatabaseType string
	/*DatalakeCrn*/
	DatalakeCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get datalake db config params
func (o *GetDatalakeDbConfigParams) WithTimeout(timeout time.Duration) *GetDatalakeDbConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datalake db config params
func (o *GetDatalakeDbConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datalake db config params
func (o *GetDatalakeDbConfigParams) WithContext(ctx context.Context) *GetDatalakeDbConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datalake db config params
func (o *GetDatalakeDbConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datalake db config params
func (o *GetDatalakeDbConfigParams) WithHTTPClient(client *http.Client) *GetDatalakeDbConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datalake db config params
func (o *GetDatalakeDbConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatabaseType adds the databaseType to the get datalake db config params
func (o *GetDatalakeDbConfigParams) WithDatabaseType(databaseType string) *GetDatalakeDbConfigParams {
	o.SetDatabaseType(databaseType)
	return o
}

// SetDatabaseType adds the databaseType to the get datalake db config params
func (o *GetDatalakeDbConfigParams) SetDatabaseType(databaseType string) {
	o.DatabaseType = databaseType
}

// WithDatalakeCrn adds the datalakeCrn to the get datalake db config params
func (o *GetDatalakeDbConfigParams) WithDatalakeCrn(datalakeCrn string) *GetDatalakeDbConfigParams {
	o.SetDatalakeCrn(datalakeCrn)
	return o
}

// SetDatalakeCrn adds the datalakeCrn to the get datalake db config params
func (o *GetDatalakeDbConfigParams) SetDatalakeCrn(datalakeCrn string) {
	o.DatalakeCrn = datalakeCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatalakeDbConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param databaseType
	qrDatabaseType := o.DatabaseType
	qDatabaseType := qrDatabaseType
	if qDatabaseType != "" {
		if err := r.SetQueryParam("databaseType", qDatabaseType); err != nil {
			return err
		}
	}

	// query param datalakeCrn
	qrDatalakeCrn := o.DatalakeCrn
	qDatalakeCrn := qrDatalakeCrn
	if qDatalakeCrn != "" {
		if err := r.SetQueryParam("datalakeCrn", qDatalakeCrn); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
