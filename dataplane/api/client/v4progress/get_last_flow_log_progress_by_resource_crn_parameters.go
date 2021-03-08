// Code generated by go-swagger; DO NOT EDIT.

package v4progress

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLastFlowLogProgressByResourceCrnParams creates a new GetLastFlowLogProgressByResourceCrnParams object
// with the default values initialized.
func NewGetLastFlowLogProgressByResourceCrnParams() *GetLastFlowLogProgressByResourceCrnParams {
	var ()
	return &GetLastFlowLogProgressByResourceCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLastFlowLogProgressByResourceCrnParamsWithTimeout creates a new GetLastFlowLogProgressByResourceCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLastFlowLogProgressByResourceCrnParamsWithTimeout(timeout time.Duration) *GetLastFlowLogProgressByResourceCrnParams {
	var ()
	return &GetLastFlowLogProgressByResourceCrnParams{

		timeout: timeout,
	}
}

// NewGetLastFlowLogProgressByResourceCrnParamsWithContext creates a new GetLastFlowLogProgressByResourceCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLastFlowLogProgressByResourceCrnParamsWithContext(ctx context.Context) *GetLastFlowLogProgressByResourceCrnParams {
	var ()
	return &GetLastFlowLogProgressByResourceCrnParams{

		Context: ctx,
	}
}

// NewGetLastFlowLogProgressByResourceCrnParamsWithHTTPClient creates a new GetLastFlowLogProgressByResourceCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLastFlowLogProgressByResourceCrnParamsWithHTTPClient(client *http.Client) *GetLastFlowLogProgressByResourceCrnParams {
	var ()
	return &GetLastFlowLogProgressByResourceCrnParams{
		HTTPClient: client,
	}
}

/*GetLastFlowLogProgressByResourceCrnParams contains all the parameters to send to the API endpoint
for the get last flow log progress by resource crn operation typically these are written to a http.Request
*/
type GetLastFlowLogProgressByResourceCrnParams struct {

	/*ResourceCrn*/
	ResourceCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get last flow log progress by resource crn params
func (o *GetLastFlowLogProgressByResourceCrnParams) WithTimeout(timeout time.Duration) *GetLastFlowLogProgressByResourceCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get last flow log progress by resource crn params
func (o *GetLastFlowLogProgressByResourceCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get last flow log progress by resource crn params
func (o *GetLastFlowLogProgressByResourceCrnParams) WithContext(ctx context.Context) *GetLastFlowLogProgressByResourceCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get last flow log progress by resource crn params
func (o *GetLastFlowLogProgressByResourceCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get last flow log progress by resource crn params
func (o *GetLastFlowLogProgressByResourceCrnParams) WithHTTPClient(client *http.Client) *GetLastFlowLogProgressByResourceCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get last flow log progress by resource crn params
func (o *GetLastFlowLogProgressByResourceCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceCrn adds the resourceCrn to the get last flow log progress by resource crn params
func (o *GetLastFlowLogProgressByResourceCrnParams) WithResourceCrn(resourceCrn string) *GetLastFlowLogProgressByResourceCrnParams {
	o.SetResourceCrn(resourceCrn)
	return o
}

// SetResourceCrn adds the resourceCrn to the get last flow log progress by resource crn params
func (o *GetLastFlowLogProgressByResourceCrnParams) SetResourceCrn(resourceCrn string) {
	o.ResourceCrn = resourceCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetLastFlowLogProgressByResourceCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceCrn
	if err := r.SetPathParam("resourceCrn", o.ResourceCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}