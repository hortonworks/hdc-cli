// Code generated by go-swagger; DO NOT EDIT.

package securitygroups

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

// NewPostPublicSecurityGroupParams creates a new PostPublicSecurityGroupParams object
// with the default values initialized.
func NewPostPublicSecurityGroupParams() *PostPublicSecurityGroupParams {
	var ()
	return &PostPublicSecurityGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPublicSecurityGroupParamsWithTimeout creates a new PostPublicSecurityGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPublicSecurityGroupParamsWithTimeout(timeout time.Duration) *PostPublicSecurityGroupParams {
	var ()
	return &PostPublicSecurityGroupParams{

		timeout: timeout,
	}
}

// NewPostPublicSecurityGroupParamsWithContext creates a new PostPublicSecurityGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPublicSecurityGroupParamsWithContext(ctx context.Context) *PostPublicSecurityGroupParams {
	var ()
	return &PostPublicSecurityGroupParams{

		Context: ctx,
	}
}

// NewPostPublicSecurityGroupParamsWithHTTPClient creates a new PostPublicSecurityGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPublicSecurityGroupParamsWithHTTPClient(client *http.Client) *PostPublicSecurityGroupParams {
	var ()
	return &PostPublicSecurityGroupParams{
		HTTPClient: client,
	}
}

/*PostPublicSecurityGroupParams contains all the parameters to send to the API endpoint
for the post public security group operation typically these are written to a http.Request
*/
type PostPublicSecurityGroupParams struct {

	/*Body*/
	Body *models_cloudbreak.SecurityGroupRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post public security group params
func (o *PostPublicSecurityGroupParams) WithTimeout(timeout time.Duration) *PostPublicSecurityGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post public security group params
func (o *PostPublicSecurityGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post public security group params
func (o *PostPublicSecurityGroupParams) WithContext(ctx context.Context) *PostPublicSecurityGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post public security group params
func (o *PostPublicSecurityGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post public security group params
func (o *PostPublicSecurityGroupParams) WithHTTPClient(client *http.Client) *PostPublicSecurityGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post public security group params
func (o *PostPublicSecurityGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post public security group params
func (o *PostPublicSecurityGroupParams) WithBody(body *models_cloudbreak.SecurityGroupRequest) *PostPublicSecurityGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post public security group params
func (o *PostPublicSecurityGroupParams) SetBody(body *models_cloudbreak.SecurityGroupRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicSecurityGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.SecurityGroupRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
