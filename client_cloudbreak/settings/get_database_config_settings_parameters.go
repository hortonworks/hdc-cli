// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// NewGetDatabaseConfigSettingsParams creates a new GetDatabaseConfigSettingsParams object
// with the default values initialized.
func NewGetDatabaseConfigSettingsParams() *GetDatabaseConfigSettingsParams {

	return &GetDatabaseConfigSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatabaseConfigSettingsParamsWithTimeout creates a new GetDatabaseConfigSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDatabaseConfigSettingsParamsWithTimeout(timeout time.Duration) *GetDatabaseConfigSettingsParams {

	return &GetDatabaseConfigSettingsParams{

		timeout: timeout,
	}
}

// NewGetDatabaseConfigSettingsParamsWithContext creates a new GetDatabaseConfigSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDatabaseConfigSettingsParamsWithContext(ctx context.Context) *GetDatabaseConfigSettingsParams {

	return &GetDatabaseConfigSettingsParams{

		Context: ctx,
	}
}

// NewGetDatabaseConfigSettingsParamsWithHTTPClient creates a new GetDatabaseConfigSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDatabaseConfigSettingsParamsWithHTTPClient(client *http.Client) *GetDatabaseConfigSettingsParams {

	return &GetDatabaseConfigSettingsParams{
		HTTPClient: client,
	}
}

/*GetDatabaseConfigSettingsParams contains all the parameters to send to the API endpoint
for the get database config settings operation typically these are written to a http.Request
*/
type GetDatabaseConfigSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get database config settings params
func (o *GetDatabaseConfigSettingsParams) WithTimeout(timeout time.Duration) *GetDatabaseConfigSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get database config settings params
func (o *GetDatabaseConfigSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get database config settings params
func (o *GetDatabaseConfigSettingsParams) WithContext(ctx context.Context) *GetDatabaseConfigSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get database config settings params
func (o *GetDatabaseConfigSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get database config settings params
func (o *GetDatabaseConfigSettingsParams) WithHTTPClient(client *http.Client) *GetDatabaseConfigSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get database config settings params
func (o *GetDatabaseConfigSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatabaseConfigSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
