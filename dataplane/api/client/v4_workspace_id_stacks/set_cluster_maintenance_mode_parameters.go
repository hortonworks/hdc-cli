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

// NewSetClusterMaintenanceModeParams creates a new SetClusterMaintenanceModeParams object
// with the default values initialized.
func NewSetClusterMaintenanceModeParams() *SetClusterMaintenanceModeParams {
	var ()
	return &SetClusterMaintenanceModeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetClusterMaintenanceModeParamsWithTimeout creates a new SetClusterMaintenanceModeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetClusterMaintenanceModeParamsWithTimeout(timeout time.Duration) *SetClusterMaintenanceModeParams {
	var ()
	return &SetClusterMaintenanceModeParams{

		timeout: timeout,
	}
}

// NewSetClusterMaintenanceModeParamsWithContext creates a new SetClusterMaintenanceModeParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetClusterMaintenanceModeParamsWithContext(ctx context.Context) *SetClusterMaintenanceModeParams {
	var ()
	return &SetClusterMaintenanceModeParams{

		Context: ctx,
	}
}

// NewSetClusterMaintenanceModeParamsWithHTTPClient creates a new SetClusterMaintenanceModeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetClusterMaintenanceModeParamsWithHTTPClient(client *http.Client) *SetClusterMaintenanceModeParams {
	var ()
	return &SetClusterMaintenanceModeParams{
		HTTPClient: client,
	}
}

/*SetClusterMaintenanceModeParams contains all the parameters to send to the API endpoint
for the set cluster maintenance mode operation typically these are written to a http.Request
*/
type SetClusterMaintenanceModeParams struct {

	/*AccountID*/
	AccountID *string
	/*Body*/
	Body *model.MaintenanceModeV4Request
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) WithTimeout(timeout time.Duration) *SetClusterMaintenanceModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) WithContext(ctx context.Context) *SetClusterMaintenanceModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) WithHTTPClient(client *http.Client) *SetClusterMaintenanceModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) WithAccountID(accountID *string) *SetClusterMaintenanceModeParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithBody adds the body to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) WithBody(body *model.MaintenanceModeV4Request) *SetClusterMaintenanceModeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) SetBody(body *model.MaintenanceModeV4Request) {
	o.Body = body
}

// WithName adds the name to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) WithName(name string) *SetClusterMaintenanceModeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) WithWorkspaceID(workspaceID int64) *SetClusterMaintenanceModeParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the set cluster maintenance mode params
func (o *SetClusterMaintenanceModeParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *SetClusterMaintenanceModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID string
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
				return err
			}
		}

	}

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
