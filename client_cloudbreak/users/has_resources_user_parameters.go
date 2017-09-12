// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewHasResourcesUserParams creates a new HasResourcesUserParams object
// with the default values initialized.
func NewHasResourcesUserParams() *HasResourcesUserParams {
	var ()
	return &HasResourcesUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHasResourcesUserParamsWithTimeout creates a new HasResourcesUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHasResourcesUserParamsWithTimeout(timeout time.Duration) *HasResourcesUserParams {
	var ()
	return &HasResourcesUserParams{

		timeout: timeout,
	}
}

// NewHasResourcesUserParamsWithContext creates a new HasResourcesUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewHasResourcesUserParamsWithContext(ctx context.Context) *HasResourcesUserParams {
	var ()
	return &HasResourcesUserParams{

		Context: ctx,
	}
}

// NewHasResourcesUserParamsWithHTTPClient creates a new HasResourcesUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHasResourcesUserParamsWithHTTPClient(client *http.Client) *HasResourcesUserParams {
	var ()
	return &HasResourcesUserParams{
		HTTPClient: client,
	}
}

/*HasResourcesUserParams contains all the parameters to send to the API endpoint
for the has resources user operation typically these are written to a http.Request
*/
type HasResourcesUserParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the has resources user params
func (o *HasResourcesUserParams) WithTimeout(timeout time.Duration) *HasResourcesUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the has resources user params
func (o *HasResourcesUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the has resources user params
func (o *HasResourcesUserParams) WithContext(ctx context.Context) *HasResourcesUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the has resources user params
func (o *HasResourcesUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the has resources user params
func (o *HasResourcesUserParams) WithHTTPClient(client *http.Client) *HasResourcesUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the has resources user params
func (o *HasResourcesUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the has resources user params
func (o *HasResourcesUserParams) WithID(id string) *HasResourcesUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the has resources user params
func (o *HasResourcesUserParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *HasResourcesUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
