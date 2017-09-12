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

// NewCreateAlertsParams creates a new CreateAlertsParams object
// with the default values initialized.
func NewCreateAlertsParams() *CreateAlertsParams {
	var ()
	return &CreateAlertsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAlertsParamsWithTimeout creates a new CreateAlertsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAlertsParamsWithTimeout(timeout time.Duration) *CreateAlertsParams {
	var ()
	return &CreateAlertsParams{

		timeout: timeout,
	}
}

// NewCreateAlertsParamsWithContext creates a new CreateAlertsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAlertsParamsWithContext(ctx context.Context) *CreateAlertsParams {
	var ()
	return &CreateAlertsParams{

		Context: ctx,
	}
}

// NewCreateAlertsParamsWithHTTPClient creates a new CreateAlertsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAlertsParamsWithHTTPClient(client *http.Client) *CreateAlertsParams {
	var ()
	return &CreateAlertsParams{
		HTTPClient: client,
	}
}

/*CreateAlertsParams contains all the parameters to send to the API endpoint
for the create alerts operation typically these are written to a http.Request
*/
type CreateAlertsParams struct {

	/*Body*/
	Body *models_autoscale.MetricAlert
	/*ClusterID*/
	ClusterID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create alerts params
func (o *CreateAlertsParams) WithTimeout(timeout time.Duration) *CreateAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create alerts params
func (o *CreateAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create alerts params
func (o *CreateAlertsParams) WithContext(ctx context.Context) *CreateAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create alerts params
func (o *CreateAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create alerts params
func (o *CreateAlertsParams) WithHTTPClient(client *http.Client) *CreateAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create alerts params
func (o *CreateAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create alerts params
func (o *CreateAlertsParams) WithBody(body *models_autoscale.MetricAlert) *CreateAlertsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create alerts params
func (o *CreateAlertsParams) SetBody(body *models_autoscale.MetricAlert) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create alerts params
func (o *CreateAlertsParams) WithClusterID(clusterID int64) *CreateAlertsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create alerts params
func (o *CreateAlertsParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
