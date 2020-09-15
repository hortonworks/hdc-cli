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
)

// NewStopStackInWorkspaceV4InternalParams creates a new StopStackInWorkspaceV4InternalParams object
// with the default values initialized.
func NewStopStackInWorkspaceV4InternalParams() *StopStackInWorkspaceV4InternalParams {
	var ()
	return &StopStackInWorkspaceV4InternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopStackInWorkspaceV4InternalParamsWithTimeout creates a new StopStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopStackInWorkspaceV4InternalParamsWithTimeout(timeout time.Duration) *StopStackInWorkspaceV4InternalParams {
	var ()
	return &StopStackInWorkspaceV4InternalParams{

		timeout: timeout,
	}
}

// NewStopStackInWorkspaceV4InternalParamsWithContext creates a new StopStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopStackInWorkspaceV4InternalParamsWithContext(ctx context.Context) *StopStackInWorkspaceV4InternalParams {
	var ()
	return &StopStackInWorkspaceV4InternalParams{

		Context: ctx,
	}
}

// NewStopStackInWorkspaceV4InternalParamsWithHTTPClient creates a new StopStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopStackInWorkspaceV4InternalParamsWithHTTPClient(client *http.Client) *StopStackInWorkspaceV4InternalParams {
	var ()
	return &StopStackInWorkspaceV4InternalParams{
		HTTPClient: client,
	}
}

/*StopStackInWorkspaceV4InternalParams contains all the parameters to send to the API endpoint
for the stop stack in workspace v4 internal operation typically these are written to a http.Request
*/
type StopStackInWorkspaceV4InternalParams struct {

	/*InitiatorUserCrn*/
	InitiatorUserCrn *string
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) WithTimeout(timeout time.Duration) *StopStackInWorkspaceV4InternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) WithContext(ctx context.Context) *StopStackInWorkspaceV4InternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) WithHTTPClient(client *http.Client) *StopStackInWorkspaceV4InternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *StopStackInWorkspaceV4InternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithName adds the name to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) WithName(name string) *StopStackInWorkspaceV4InternalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) WithWorkspaceID(workspaceID int64) *StopStackInWorkspaceV4InternalParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the stop stack in workspace v4 internal params
func (o *StopStackInWorkspaceV4InternalParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *StopStackInWorkspaceV4InternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InitiatorUserCrn != nil {

		// query param initiatorUserCrn
		var qrInitiatorUserCrn string
		if o.InitiatorUserCrn != nil {
			qrInitiatorUserCrn = *o.InitiatorUserCrn
		}
		qInitiatorUserCrn := qrInitiatorUserCrn
		if qInitiatorUserCrn != "" {
			if err := r.SetQueryParam("initiatorUserCrn", qInitiatorUserCrn); err != nil {
				return err
			}
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
