// Code generated by go-swagger; DO NOT EDIT.

package smartsensesubscriptions

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

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// NewPostPublicSmartSenseSubscriptionParams creates a new PostPublicSmartSenseSubscriptionParams object
// with the default values initialized.
func NewPostPublicSmartSenseSubscriptionParams() *PostPublicSmartSenseSubscriptionParams {
	var ()
	return &PostPublicSmartSenseSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPublicSmartSenseSubscriptionParamsWithTimeout creates a new PostPublicSmartSenseSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPublicSmartSenseSubscriptionParamsWithTimeout(timeout time.Duration) *PostPublicSmartSenseSubscriptionParams {
	var ()
	return &PostPublicSmartSenseSubscriptionParams{

		timeout: timeout,
	}
}

// NewPostPublicSmartSenseSubscriptionParamsWithContext creates a new PostPublicSmartSenseSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPublicSmartSenseSubscriptionParamsWithContext(ctx context.Context) *PostPublicSmartSenseSubscriptionParams {
	var ()
	return &PostPublicSmartSenseSubscriptionParams{

		Context: ctx,
	}
}

// NewPostPublicSmartSenseSubscriptionParamsWithHTTPClient creates a new PostPublicSmartSenseSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPublicSmartSenseSubscriptionParamsWithHTTPClient(client *http.Client) *PostPublicSmartSenseSubscriptionParams {
	var ()
	return &PostPublicSmartSenseSubscriptionParams{
		HTTPClient: client,
	}
}

/*PostPublicSmartSenseSubscriptionParams contains all the parameters to send to the API endpoint
for the post public smart sense subscription operation typically these are written to a http.Request
*/
type PostPublicSmartSenseSubscriptionParams struct {

	/*Body*/
	Body *models_cloudbreak.SmartSenseSubscriptionJSON

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post public smart sense subscription params
func (o *PostPublicSmartSenseSubscriptionParams) WithTimeout(timeout time.Duration) *PostPublicSmartSenseSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post public smart sense subscription params
func (o *PostPublicSmartSenseSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post public smart sense subscription params
func (o *PostPublicSmartSenseSubscriptionParams) WithContext(ctx context.Context) *PostPublicSmartSenseSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post public smart sense subscription params
func (o *PostPublicSmartSenseSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post public smart sense subscription params
func (o *PostPublicSmartSenseSubscriptionParams) WithHTTPClient(client *http.Client) *PostPublicSmartSenseSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post public smart sense subscription params
func (o *PostPublicSmartSenseSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post public smart sense subscription params
func (o *PostPublicSmartSenseSubscriptionParams) WithBody(body *models_cloudbreak.SmartSenseSubscriptionJSON) *PostPublicSmartSenseSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post public smart sense subscription params
func (o *PostPublicSmartSenseSubscriptionParams) SetBody(body *models_cloudbreak.SmartSenseSubscriptionJSON) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicSmartSenseSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.SmartSenseSubscriptionJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
