// Code generated by go-swagger; DO NOT EDIT.

package history

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

// NewGetHistoryParams creates a new GetHistoryParams object
// with the default values initialized.
func NewGetHistoryParams() *GetHistoryParams {
	var ()
	return &GetHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHistoryParamsWithTimeout creates a new GetHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHistoryParamsWithTimeout(timeout time.Duration) *GetHistoryParams {
	var ()
	return &GetHistoryParams{

		timeout: timeout,
	}
}

// NewGetHistoryParamsWithContext creates a new GetHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHistoryParamsWithContext(ctx context.Context) *GetHistoryParams {
	var ()
	return &GetHistoryParams{

		Context: ctx,
	}
}

// NewGetHistoryParamsWithHTTPClient creates a new GetHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHistoryParamsWithHTTPClient(client *http.Client) *GetHistoryParams {
	var ()
	return &GetHistoryParams{
		HTTPClient: client,
	}
}

/*GetHistoryParams contains all the parameters to send to the API endpoint
for the get history operation typically these are written to a http.Request
*/
type GetHistoryParams struct {

	/*ClusterID*/
	ClusterID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get history params
func (o *GetHistoryParams) WithTimeout(timeout time.Duration) *GetHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get history params
func (o *GetHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get history params
func (o *GetHistoryParams) WithContext(ctx context.Context) *GetHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get history params
func (o *GetHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get history params
func (o *GetHistoryParams) WithHTTPClient(client *http.Client) *GetHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get history params
func (o *GetHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get history params
func (o *GetHistoryParams) WithClusterID(clusterID int64) *GetHistoryParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get history params
func (o *GetHistoryParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", swag.FormatInt64(o.ClusterID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
