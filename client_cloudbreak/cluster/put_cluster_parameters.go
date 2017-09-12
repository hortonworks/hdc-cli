// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// NewPutClusterParams creates a new PutClusterParams object
// with the default values initialized.
func NewPutClusterParams() *PutClusterParams {
	var ()
	return &PutClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClusterParamsWithTimeout creates a new PutClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClusterParamsWithTimeout(timeout time.Duration) *PutClusterParams {
	var ()
	return &PutClusterParams{

		timeout: timeout,
	}
}

// NewPutClusterParamsWithContext creates a new PutClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutClusterParamsWithContext(ctx context.Context) *PutClusterParams {
	var ()
	return &PutClusterParams{

		Context: ctx,
	}
}

// NewPutClusterParamsWithHTTPClient creates a new PutClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutClusterParamsWithHTTPClient(client *http.Client) *PutClusterParams {
	var ()
	return &PutClusterParams{
		HTTPClient: client,
	}
}

/*PutClusterParams contains all the parameters to send to the API endpoint
for the put cluster operation typically these are written to a http.Request
*/
type PutClusterParams struct {

	/*Body*/
	Body *models_cloudbreak.UpdateCluster
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cluster params
func (o *PutClusterParams) WithTimeout(timeout time.Duration) *PutClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cluster params
func (o *PutClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cluster params
func (o *PutClusterParams) WithContext(ctx context.Context) *PutClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cluster params
func (o *PutClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cluster params
func (o *PutClusterParams) WithHTTPClient(client *http.Client) *PutClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cluster params
func (o *PutClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put cluster params
func (o *PutClusterParams) WithBody(body *models_cloudbreak.UpdateCluster) *PutClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put cluster params
func (o *PutClusterParams) SetBody(body *models_cloudbreak.UpdateCluster) {
	o.Body = body
}

// WithID adds the id to the put cluster params
func (o *PutClusterParams) WithID(id int64) *PutClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put cluster params
func (o *PutClusterParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.UpdateCluster)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
