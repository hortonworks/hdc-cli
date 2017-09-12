// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewAddScalingParams creates a new AddScalingParams object
// with the default values initialized.
func NewAddScalingParams() *AddScalingParams {
	var ()
	return &AddScalingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddScalingParamsWithTimeout creates a new AddScalingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddScalingParamsWithTimeout(timeout time.Duration) *AddScalingParams {
	var ()
	return &AddScalingParams{

		timeout: timeout,
	}
}

// NewAddScalingParamsWithContext creates a new AddScalingParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddScalingParamsWithContext(ctx context.Context) *AddScalingParams {
	var ()
	return &AddScalingParams{

		Context: ctx,
	}
}

// NewAddScalingParamsWithHTTPClient creates a new AddScalingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddScalingParamsWithHTTPClient(client *http.Client) *AddScalingParams {
	var ()
	return &AddScalingParams{
		HTTPClient: client,
	}
}

/*AddScalingParams contains all the parameters to send to the API endpoint
for the add scaling operation typically these are written to a http.Request
*/
type AddScalingParams struct {

	/*Body*/
	Body *models_autoscale.ScalingPolicy
	/*ClusterID*/
	ClusterID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add scaling params
func (o *AddScalingParams) WithTimeout(timeout time.Duration) *AddScalingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add scaling params
func (o *AddScalingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add scaling params
func (o *AddScalingParams) WithContext(ctx context.Context) *AddScalingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add scaling params
func (o *AddScalingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add scaling params
func (o *AddScalingParams) WithHTTPClient(client *http.Client) *AddScalingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add scaling params
func (o *AddScalingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add scaling params
func (o *AddScalingParams) WithBody(body *models_autoscale.ScalingPolicy) *AddScalingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add scaling params
func (o *AddScalingParams) SetBody(body *models_autoscale.ScalingPolicy) {
	o.Body = body
}

// WithClusterID adds the clusterID to the add scaling params
func (o *AddScalingParams) WithClusterID(clusterID int64) *AddScalingParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the add scaling params
func (o *AddScalingParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *AddScalingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_autoscale.ScalingPolicy)
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
