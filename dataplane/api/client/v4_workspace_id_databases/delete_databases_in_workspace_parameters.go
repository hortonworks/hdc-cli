// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_databases

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
)

// NewDeleteDatabasesInWorkspaceParams creates a new DeleteDatabasesInWorkspaceParams object
// with the default values initialized.
func NewDeleteDatabasesInWorkspaceParams() *DeleteDatabasesInWorkspaceParams {
	var ()
	return &DeleteDatabasesInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDatabasesInWorkspaceParamsWithTimeout creates a new DeleteDatabasesInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDatabasesInWorkspaceParamsWithTimeout(timeout time.Duration) *DeleteDatabasesInWorkspaceParams {
	var ()
	return &DeleteDatabasesInWorkspaceParams{

		timeout: timeout,
	}
}

// NewDeleteDatabasesInWorkspaceParamsWithContext creates a new DeleteDatabasesInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDatabasesInWorkspaceParamsWithContext(ctx context.Context) *DeleteDatabasesInWorkspaceParams {
	var ()
	return &DeleteDatabasesInWorkspaceParams{

		Context: ctx,
	}
}

// NewDeleteDatabasesInWorkspaceParamsWithHTTPClient creates a new DeleteDatabasesInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDatabasesInWorkspaceParamsWithHTTPClient(client *http.Client) *DeleteDatabasesInWorkspaceParams {
	var ()
	return &DeleteDatabasesInWorkspaceParams{
		HTTPClient: client,
	}
}

/*DeleteDatabasesInWorkspaceParams contains all the parameters to send to the API endpoint
for the delete databases in workspace operation typically these are written to a http.Request
*/
type DeleteDatabasesInWorkspaceParams struct {

	/*Body*/
	Body []string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete databases in workspace params
func (o *DeleteDatabasesInWorkspaceParams) WithTimeout(timeout time.Duration) *DeleteDatabasesInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete databases in workspace params
func (o *DeleteDatabasesInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete databases in workspace params
func (o *DeleteDatabasesInWorkspaceParams) WithContext(ctx context.Context) *DeleteDatabasesInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete databases in workspace params
func (o *DeleteDatabasesInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete databases in workspace params
func (o *DeleteDatabasesInWorkspaceParams) WithHTTPClient(client *http.Client) *DeleteDatabasesInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete databases in workspace params
func (o *DeleteDatabasesInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete databases in workspace params
func (o *DeleteDatabasesInWorkspaceParams) WithBody(body []string) *DeleteDatabasesInWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete databases in workspace params
func (o *DeleteDatabasesInWorkspaceParams) SetBody(body []string) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the delete databases in workspace params
func (o *DeleteDatabasesInWorkspaceParams) WithWorkspaceID(workspaceID int64) *DeleteDatabasesInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete databases in workspace params
func (o *DeleteDatabasesInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDatabasesInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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