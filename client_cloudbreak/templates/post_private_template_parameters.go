// Code generated by go-swagger; DO NOT EDIT.

package templates

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

// NewPostPrivateTemplateParams creates a new PostPrivateTemplateParams object
// with the default values initialized.
func NewPostPrivateTemplateParams() *PostPrivateTemplateParams {
	var ()
	return &PostPrivateTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPrivateTemplateParamsWithTimeout creates a new PostPrivateTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPrivateTemplateParamsWithTimeout(timeout time.Duration) *PostPrivateTemplateParams {
	var ()
	return &PostPrivateTemplateParams{

		timeout: timeout,
	}
}

// NewPostPrivateTemplateParamsWithContext creates a new PostPrivateTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPrivateTemplateParamsWithContext(ctx context.Context) *PostPrivateTemplateParams {
	var ()
	return &PostPrivateTemplateParams{

		Context: ctx,
	}
}

// NewPostPrivateTemplateParamsWithHTTPClient creates a new PostPrivateTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPrivateTemplateParamsWithHTTPClient(client *http.Client) *PostPrivateTemplateParams {
	var ()
	return &PostPrivateTemplateParams{
		HTTPClient: client,
	}
}

/*PostPrivateTemplateParams contains all the parameters to send to the API endpoint
for the post private template operation typically these are written to a http.Request
*/
type PostPrivateTemplateParams struct {

	/*Body*/
	Body *models_cloudbreak.TemplateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post private template params
func (o *PostPrivateTemplateParams) WithTimeout(timeout time.Duration) *PostPrivateTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post private template params
func (o *PostPrivateTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post private template params
func (o *PostPrivateTemplateParams) WithContext(ctx context.Context) *PostPrivateTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post private template params
func (o *PostPrivateTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post private template params
func (o *PostPrivateTemplateParams) WithHTTPClient(client *http.Client) *PostPrivateTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post private template params
func (o *PostPrivateTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post private template params
func (o *PostPrivateTemplateParams) WithBody(body *models_cloudbreak.TemplateRequest) *PostPrivateTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post private template params
func (o *PostPrivateTemplateParams) SetBody(body *models_cloudbreak.TemplateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPrivateTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.TemplateRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
