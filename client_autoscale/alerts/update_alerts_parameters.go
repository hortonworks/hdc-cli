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

	"github.com/hortonworks/hdc-cli/models_autoscale"
)

// NewUpdateAlertsParams creates a new UpdateAlertsParams object
// with the default values initialized.
func NewUpdateAlertsParams() *UpdateAlertsParams {
	var ()
	return &UpdateAlertsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAlertsParamsWithTimeout creates a new UpdateAlertsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAlertsParamsWithTimeout(timeout time.Duration) *UpdateAlertsParams {
	var ()
	return &UpdateAlertsParams{

		timeout: timeout,
	}
}

// NewUpdateAlertsParamsWithContext creates a new UpdateAlertsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAlertsParamsWithContext(ctx context.Context) *UpdateAlertsParams {
	var ()
	return &UpdateAlertsParams{

		Context: ctx,
	}
}

// NewUpdateAlertsParamsWithHTTPClient creates a new UpdateAlertsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAlertsParamsWithHTTPClient(client *http.Client) *UpdateAlertsParams {
	var ()
	return &UpdateAlertsParams{
		HTTPClient: client,
	}
}

/*UpdateAlertsParams contains all the parameters to send to the API endpoint
for the update alerts operation typically these are written to a http.Request
*/
type UpdateAlertsParams struct {

	/*AlertID*/
	AlertID int64
	/*Body*/
	Body *models_autoscale.MetricAlert
	/*ClusterID*/
	ClusterID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update alerts params
func (o *UpdateAlertsParams) WithTimeout(timeout time.Duration) *UpdateAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update alerts params
func (o *UpdateAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update alerts params
func (o *UpdateAlertsParams) WithContext(ctx context.Context) *UpdateAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update alerts params
func (o *UpdateAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update alerts params
func (o *UpdateAlertsParams) WithHTTPClient(client *http.Client) *UpdateAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update alerts params
func (o *UpdateAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertID adds the alertID to the update alerts params
func (o *UpdateAlertsParams) WithAlertID(alertID int64) *UpdateAlertsParams {
	o.SetAlertID(alertID)
	return o
}

// SetAlertID adds the alertId to the update alerts params
func (o *UpdateAlertsParams) SetAlertID(alertID int64) {
	o.AlertID = alertID
}

// WithBody adds the body to the update alerts params
func (o *UpdateAlertsParams) WithBody(body *models_autoscale.MetricAlert) *UpdateAlertsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update alerts params
func (o *UpdateAlertsParams) SetBody(body *models_autoscale.MetricAlert) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update alerts params
func (o *UpdateAlertsParams) WithClusterID(clusterID int64) *UpdateAlertsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update alerts params
func (o *UpdateAlertsParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertId
	if err := r.SetPathParam("alertId", swag.FormatInt64(o.AlertID)); err != nil {
		return err
	}

	if o.Body == nil {
		o.Body = new(models_autoscale.MetricAlert)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
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
