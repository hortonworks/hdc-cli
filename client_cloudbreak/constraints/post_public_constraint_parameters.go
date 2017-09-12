// Code generated by go-swagger; DO NOT EDIT.

package constraints

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

// NewPostPublicConstraintParams creates a new PostPublicConstraintParams object
// with the default values initialized.
func NewPostPublicConstraintParams() *PostPublicConstraintParams {
	var ()
	return &PostPublicConstraintParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPublicConstraintParamsWithTimeout creates a new PostPublicConstraintParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPublicConstraintParamsWithTimeout(timeout time.Duration) *PostPublicConstraintParams {
	var ()
	return &PostPublicConstraintParams{

		timeout: timeout,
	}
}

// NewPostPublicConstraintParamsWithContext creates a new PostPublicConstraintParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPublicConstraintParamsWithContext(ctx context.Context) *PostPublicConstraintParams {
	var ()
	return &PostPublicConstraintParams{

		Context: ctx,
	}
}

// NewPostPublicConstraintParamsWithHTTPClient creates a new PostPublicConstraintParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPublicConstraintParamsWithHTTPClient(client *http.Client) *PostPublicConstraintParams {
	var ()
	return &PostPublicConstraintParams{
		HTTPClient: client,
	}
}

/*PostPublicConstraintParams contains all the parameters to send to the API endpoint
for the post public constraint operation typically these are written to a http.Request
*/
type PostPublicConstraintParams struct {

	/*Body*/
	Body *models_cloudbreak.ConstraintTemplateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post public constraint params
func (o *PostPublicConstraintParams) WithTimeout(timeout time.Duration) *PostPublicConstraintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post public constraint params
func (o *PostPublicConstraintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post public constraint params
func (o *PostPublicConstraintParams) WithContext(ctx context.Context) *PostPublicConstraintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post public constraint params
func (o *PostPublicConstraintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post public constraint params
func (o *PostPublicConstraintParams) WithHTTPClient(client *http.Client) *PostPublicConstraintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post public constraint params
func (o *PostPublicConstraintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post public constraint params
func (o *PostPublicConstraintParams) WithBody(body *models_cloudbreak.ConstraintTemplateRequest) *PostPublicConstraintParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post public constraint params
func (o *PostPublicConstraintParams) SetBody(body *models_cloudbreak.ConstraintTemplateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicConstraintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.ConstraintTemplateRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
