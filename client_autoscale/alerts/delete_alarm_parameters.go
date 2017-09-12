// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewDeleteAlarmParams creates a new DeleteAlarmParams object
// with the default values initialized.
func NewDeleteAlarmParams() *DeleteAlarmParams {
	var ()
	return &DeleteAlarmParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAlarmParamsWithTimeout creates a new DeleteAlarmParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAlarmParamsWithTimeout(timeout time.Duration) *DeleteAlarmParams {
	var ()
	return &DeleteAlarmParams{

		timeout: timeout,
	}
}

// NewDeleteAlarmParamsWithContext creates a new DeleteAlarmParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAlarmParamsWithContext(ctx context.Context) *DeleteAlarmParams {
	var ()
	return &DeleteAlarmParams{

		Context: ctx,
	}
}

// NewDeleteAlarmParamsWithHTTPClient creates a new DeleteAlarmParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAlarmParamsWithHTTPClient(client *http.Client) *DeleteAlarmParams {
	var ()
	return &DeleteAlarmParams{
		HTTPClient: client,
	}
}

/*DeleteAlarmParams contains all the parameters to send to the API endpoint
for the delete alarm operation typically these are written to a http.Request
*/
type DeleteAlarmParams struct {

	/*AlertID*/
	AlertID int64
	/*ClusterID*/
	ClusterID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete alarm params
func (o *DeleteAlarmParams) WithTimeout(timeout time.Duration) *DeleteAlarmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete alarm params
func (o *DeleteAlarmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete alarm params
func (o *DeleteAlarmParams) WithContext(ctx context.Context) *DeleteAlarmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete alarm params
func (o *DeleteAlarmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete alarm params
func (o *DeleteAlarmParams) WithHTTPClient(client *http.Client) *DeleteAlarmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete alarm params
func (o *DeleteAlarmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertID adds the alertID to the delete alarm params
func (o *DeleteAlarmParams) WithAlertID(alertID int64) *DeleteAlarmParams {
	o.SetAlertID(alertID)
	return o
}

// SetAlertID adds the alertId to the delete alarm params
func (o *DeleteAlarmParams) SetAlertID(alertID int64) {
	o.AlertID = alertID
}

// WithClusterID adds the clusterID to the delete alarm params
func (o *DeleteAlarmParams) WithClusterID(clusterID int64) *DeleteAlarmParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete alarm params
func (o *DeleteAlarmParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAlarmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertId
	if err := r.SetPathParam("alertId", swag.FormatInt64(o.AlertID)); err != nil {
		return err
	}

	// path param clusterId
	if err := r.SetPathParam("clusterId", swag.FormatInt64(o.ClusterID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
