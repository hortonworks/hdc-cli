// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPutClusterV4Params creates a new PutClusterV4Params object
// with the default values initialized.
func NewPutClusterV4Params() *PutClusterV4Params {
	var ()
	return &PutClusterV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClusterV4ParamsWithTimeout creates a new PutClusterV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClusterV4ParamsWithTimeout(timeout time.Duration) *PutClusterV4Params {
	var ()
	return &PutClusterV4Params{

		timeout: timeout,
	}
}

// NewPutClusterV4ParamsWithContext creates a new PutClusterV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewPutClusterV4ParamsWithContext(ctx context.Context) *PutClusterV4Params {
	var ()
	return &PutClusterV4Params{

		Context: ctx,
	}
}

// NewPutClusterV4ParamsWithHTTPClient creates a new PutClusterV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutClusterV4ParamsWithHTTPClient(client *http.Client) *PutClusterV4Params {
	var ()
	return &PutClusterV4Params{
		HTTPClient: client,
	}
}

/*PutClusterV4Params contains all the parameters to send to the API endpoint
for the put cluster v4 operation typically these are written to a http.Request
*/
type PutClusterV4Params struct {

	/*Body*/
	Body *model.UpdateClusterV4Request
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cluster v4 params
func (o *PutClusterV4Params) WithTimeout(timeout time.Duration) *PutClusterV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cluster v4 params
func (o *PutClusterV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cluster v4 params
func (o *PutClusterV4Params) WithContext(ctx context.Context) *PutClusterV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cluster v4 params
func (o *PutClusterV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cluster v4 params
func (o *PutClusterV4Params) WithHTTPClient(client *http.Client) *PutClusterV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cluster v4 params
func (o *PutClusterV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put cluster v4 params
func (o *PutClusterV4Params) WithBody(body *model.UpdateClusterV4Request) *PutClusterV4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put cluster v4 params
func (o *PutClusterV4Params) SetBody(body *model.UpdateClusterV4Request) {
	o.Body = body
}

// WithName adds the name to the put cluster v4 params
func (o *PutClusterV4Params) WithName(name string) *PutClusterV4Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the put cluster v4 params
func (o *PutClusterV4Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the put cluster v4 params
func (o *PutClusterV4Params) WithWorkspaceID(workspaceID int64) *PutClusterV4Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the put cluster v4 params
func (o *PutClusterV4Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PutClusterV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}