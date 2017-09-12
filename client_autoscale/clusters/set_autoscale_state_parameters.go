// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// NewSetAutoscaleStateParams creates a new SetAutoscaleStateParams object
// with the default values initialized.
func NewSetAutoscaleStateParams() *SetAutoscaleStateParams {
	var ()
	return &SetAutoscaleStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetAutoscaleStateParamsWithTimeout creates a new SetAutoscaleStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetAutoscaleStateParamsWithTimeout(timeout time.Duration) *SetAutoscaleStateParams {
	var ()
	return &SetAutoscaleStateParams{

		timeout: timeout,
	}
}

// NewSetAutoscaleStateParamsWithContext creates a new SetAutoscaleStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetAutoscaleStateParamsWithContext(ctx context.Context) *SetAutoscaleStateParams {
	var ()
	return &SetAutoscaleStateParams{

		Context: ctx,
	}
}

// NewSetAutoscaleStateParamsWithHTTPClient creates a new SetAutoscaleStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetAutoscaleStateParamsWithHTTPClient(client *http.Client) *SetAutoscaleStateParams {
	var ()
	return &SetAutoscaleStateParams{
		HTTPClient: client,
	}
}

/*SetAutoscaleStateParams contains all the parameters to send to the API endpoint
for the set autoscale state operation typically these are written to a http.Request
*/
type SetAutoscaleStateParams struct {

	/*Body*/
	Body *models_autoscale.ClusterAutoscaleState
	/*ClusterID*/
	ClusterID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set autoscale state params
func (o *SetAutoscaleStateParams) WithTimeout(timeout time.Duration) *SetAutoscaleStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set autoscale state params
func (o *SetAutoscaleStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set autoscale state params
func (o *SetAutoscaleStateParams) WithContext(ctx context.Context) *SetAutoscaleStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set autoscale state params
func (o *SetAutoscaleStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set autoscale state params
func (o *SetAutoscaleStateParams) WithHTTPClient(client *http.Client) *SetAutoscaleStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set autoscale state params
func (o *SetAutoscaleStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set autoscale state params
func (o *SetAutoscaleStateParams) WithBody(body *models_autoscale.ClusterAutoscaleState) *SetAutoscaleStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set autoscale state params
func (o *SetAutoscaleStateParams) SetBody(body *models_autoscale.ClusterAutoscaleState) {
	o.Body = body
}

// WithClusterID adds the clusterID to the set autoscale state params
func (o *SetAutoscaleStateParams) WithClusterID(clusterID int64) *SetAutoscaleStateParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the set autoscale state params
func (o *SetAutoscaleStateParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *SetAutoscaleStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_autoscale.ClusterAutoscaleState)
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
