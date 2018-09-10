// Code generated by go-swagger; DO NOT EDIT.

package v3workspaces

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

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewAddWorkspaceUsersParams creates a new AddWorkspaceUsersParams object
// with the default values initialized.
func NewAddWorkspaceUsersParams() *AddWorkspaceUsersParams {
	var ()
	return &AddWorkspaceUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddWorkspaceUsersParamsWithTimeout creates a new AddWorkspaceUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddWorkspaceUsersParamsWithTimeout(timeout time.Duration) *AddWorkspaceUsersParams {
	var ()
	return &AddWorkspaceUsersParams{

		timeout: timeout,
	}
}

// NewAddWorkspaceUsersParamsWithContext creates a new AddWorkspaceUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddWorkspaceUsersParamsWithContext(ctx context.Context) *AddWorkspaceUsersParams {
	var ()
	return &AddWorkspaceUsersParams{

		Context: ctx,
	}
}

// NewAddWorkspaceUsersParamsWithHTTPClient creates a new AddWorkspaceUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddWorkspaceUsersParamsWithHTTPClient(client *http.Client) *AddWorkspaceUsersParams {
	var ()
	return &AddWorkspaceUsersParams{
		HTTPClient: client,
	}
}

/*AddWorkspaceUsersParams contains all the parameters to send to the API endpoint
for the add workspace users operation typically these are written to a http.Request
*/
type AddWorkspaceUsersParams struct {

	/*Body*/
	Body []*models_cloudbreak.ChangeWorkspaceUsersJSON
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add workspace users params
func (o *AddWorkspaceUsersParams) WithTimeout(timeout time.Duration) *AddWorkspaceUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add workspace users params
func (o *AddWorkspaceUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add workspace users params
func (o *AddWorkspaceUsersParams) WithContext(ctx context.Context) *AddWorkspaceUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add workspace users params
func (o *AddWorkspaceUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add workspace users params
func (o *AddWorkspaceUsersParams) WithHTTPClient(client *http.Client) *AddWorkspaceUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add workspace users params
func (o *AddWorkspaceUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add workspace users params
func (o *AddWorkspaceUsersParams) WithBody(body []*models_cloudbreak.ChangeWorkspaceUsersJSON) *AddWorkspaceUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add workspace users params
func (o *AddWorkspaceUsersParams) SetBody(body []*models_cloudbreak.ChangeWorkspaceUsersJSON) {
	o.Body = body
}

// WithName adds the name to the add workspace users params
func (o *AddWorkspaceUsersParams) WithName(name string) *AddWorkspaceUsersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the add workspace users params
func (o *AddWorkspaceUsersParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *AddWorkspaceUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
