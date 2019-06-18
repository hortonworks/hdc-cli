// Code generated by go-swagger; DO NOT EDIT.

package v4databaseservers

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

// NewGetDatabaseServerParams creates a new GetDatabaseServerParams object
// with the default values initialized.
func NewGetDatabaseServerParams() *GetDatabaseServerParams {
	var ()
	return &GetDatabaseServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatabaseServerParamsWithTimeout creates a new GetDatabaseServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDatabaseServerParamsWithTimeout(timeout time.Duration) *GetDatabaseServerParams {
	var ()
	return &GetDatabaseServerParams{

		timeout: timeout,
	}
}

// NewGetDatabaseServerParamsWithContext creates a new GetDatabaseServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDatabaseServerParamsWithContext(ctx context.Context) *GetDatabaseServerParams {
	var ()
	return &GetDatabaseServerParams{

		Context: ctx,
	}
}

// NewGetDatabaseServerParamsWithHTTPClient creates a new GetDatabaseServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDatabaseServerParamsWithHTTPClient(client *http.Client) *GetDatabaseServerParams {
	var ()
	return &GetDatabaseServerParams{
		HTTPClient: client,
	}
}

/*GetDatabaseServerParams contains all the parameters to send to the API endpoint
for the get database server operation typically these are written to a http.Request
*/
type GetDatabaseServerParams struct {

	/*EnvironmentID*/
	EnvironmentID string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get database server params
func (o *GetDatabaseServerParams) WithTimeout(timeout time.Duration) *GetDatabaseServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get database server params
func (o *GetDatabaseServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get database server params
func (o *GetDatabaseServerParams) WithContext(ctx context.Context) *GetDatabaseServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get database server params
func (o *GetDatabaseServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get database server params
func (o *GetDatabaseServerParams) WithHTTPClient(client *http.Client) *GetDatabaseServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get database server params
func (o *GetDatabaseServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the get database server params
func (o *GetDatabaseServerParams) WithEnvironmentID(environmentID string) *GetDatabaseServerParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the get database server params
func (o *GetDatabaseServerParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithName adds the name to the get database server params
func (o *GetDatabaseServerParams) WithName(name string) *GetDatabaseServerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get database server params
func (o *GetDatabaseServerParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatabaseServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param environmentId
	qrEnvironmentID := o.EnvironmentID
	qEnvironmentID := qrEnvironmentID
	if qEnvironmentID != "" {
		if err := r.SetQueryParam("environmentId", qEnvironmentID); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}