// Code generated by go-swagger; DO NOT EDIT.

package flow_logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLastFlowByResourceNameParams creates a new GetLastFlowByResourceNameParams object
// with the default values initialized.
func NewGetLastFlowByResourceNameParams() *GetLastFlowByResourceNameParams {
	var ()
	return &GetLastFlowByResourceNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLastFlowByResourceNameParamsWithTimeout creates a new GetLastFlowByResourceNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLastFlowByResourceNameParamsWithTimeout(timeout time.Duration) *GetLastFlowByResourceNameParams {
	var ()
	return &GetLastFlowByResourceNameParams{

		timeout: timeout,
	}
}

// NewGetLastFlowByResourceNameParamsWithContext creates a new GetLastFlowByResourceNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLastFlowByResourceNameParamsWithContext(ctx context.Context) *GetLastFlowByResourceNameParams {
	var ()
	return &GetLastFlowByResourceNameParams{

		Context: ctx,
	}
}

// NewGetLastFlowByResourceNameParamsWithHTTPClient creates a new GetLastFlowByResourceNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLastFlowByResourceNameParamsWithHTTPClient(client *http.Client) *GetLastFlowByResourceNameParams {
	var ()
	return &GetLastFlowByResourceNameParams{
		HTTPClient: client,
	}
}

/*GetLastFlowByResourceNameParams contains all the parameters to send to the API endpoint
for the get last flow by resource name operation typically these are written to a http.Request
*/
type GetLastFlowByResourceNameParams struct {

	/*ResourceName*/
	ResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get last flow by resource name params
func (o *GetLastFlowByResourceNameParams) WithTimeout(timeout time.Duration) *GetLastFlowByResourceNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get last flow by resource name params
func (o *GetLastFlowByResourceNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get last flow by resource name params
func (o *GetLastFlowByResourceNameParams) WithContext(ctx context.Context) *GetLastFlowByResourceNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get last flow by resource name params
func (o *GetLastFlowByResourceNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get last flow by resource name params
func (o *GetLastFlowByResourceNameParams) WithHTTPClient(client *http.Client) *GetLastFlowByResourceNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get last flow by resource name params
func (o *GetLastFlowByResourceNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceName adds the resourceName to the get last flow by resource name params
func (o *GetLastFlowByResourceNameParams) WithResourceName(resourceName string) *GetLastFlowByResourceNameParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the get last flow by resource name params
func (o *GetLastFlowByResourceNameParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetLastFlowByResourceNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceName
	if err := r.SetPathParam("resourceName", o.ResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
