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

// NewSetScalingParams creates a new SetScalingParams object
// with the default values initialized.
func NewSetScalingParams() *SetScalingParams {
	var ()
	return &SetScalingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetScalingParamsWithTimeout creates a new SetScalingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetScalingParamsWithTimeout(timeout time.Duration) *SetScalingParams {
	var ()
	return &SetScalingParams{

		timeout: timeout,
	}
}

// NewSetScalingParamsWithContext creates a new SetScalingParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetScalingParamsWithContext(ctx context.Context) *SetScalingParams {
	var ()
	return &SetScalingParams{

		Context: ctx,
	}
}

// NewSetScalingParamsWithHTTPClient creates a new SetScalingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetScalingParamsWithHTTPClient(client *http.Client) *SetScalingParams {
	var ()
	return &SetScalingParams{
		HTTPClient: client,
	}
}

/*SetScalingParams contains all the parameters to send to the API endpoint
for the set scaling operation typically these are written to a http.Request
*/
type SetScalingParams struct {

	/*Body*/
	Body *models_autoscale.ScalingPolicy
	/*ClusterID*/
	ClusterID int64
	/*PolicyID*/
	PolicyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set scaling params
func (o *SetScalingParams) WithTimeout(timeout time.Duration) *SetScalingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set scaling params
func (o *SetScalingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set scaling params
func (o *SetScalingParams) WithContext(ctx context.Context) *SetScalingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set scaling params
func (o *SetScalingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set scaling params
func (o *SetScalingParams) WithHTTPClient(client *http.Client) *SetScalingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set scaling params
func (o *SetScalingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set scaling params
func (o *SetScalingParams) WithBody(body *models_autoscale.ScalingPolicy) *SetScalingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set scaling params
func (o *SetScalingParams) SetBody(body *models_autoscale.ScalingPolicy) {
	o.Body = body
}

// WithClusterID adds the clusterID to the set scaling params
func (o *SetScalingParams) WithClusterID(clusterID int64) *SetScalingParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the set scaling params
func (o *SetScalingParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WithPolicyID adds the policyID to the set scaling params
func (o *SetScalingParams) WithPolicyID(policyID int64) *SetScalingParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the set scaling params
func (o *SetScalingParams) SetPolicyID(policyID int64) {
	o.PolicyID = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *SetScalingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param policyId
	if err := r.SetPathParam("policyId", swag.FormatInt64(o.PolicyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
