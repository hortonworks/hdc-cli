// Code generated by go-swagger; DO NOT EDIT.

package ldap

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

// NewPostPublicLdapParams creates a new PostPublicLdapParams object
// with the default values initialized.
func NewPostPublicLdapParams() *PostPublicLdapParams {
	var ()
	return &PostPublicLdapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPublicLdapParamsWithTimeout creates a new PostPublicLdapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPublicLdapParamsWithTimeout(timeout time.Duration) *PostPublicLdapParams {
	var ()
	return &PostPublicLdapParams{

		timeout: timeout,
	}
}

// NewPostPublicLdapParamsWithContext creates a new PostPublicLdapParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPublicLdapParamsWithContext(ctx context.Context) *PostPublicLdapParams {
	var ()
	return &PostPublicLdapParams{

		Context: ctx,
	}
}

// NewPostPublicLdapParamsWithHTTPClient creates a new PostPublicLdapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPublicLdapParamsWithHTTPClient(client *http.Client) *PostPublicLdapParams {
	var ()
	return &PostPublicLdapParams{
		HTTPClient: client,
	}
}

/*PostPublicLdapParams contains all the parameters to send to the API endpoint
for the post public ldap operation typically these are written to a http.Request
*/
type PostPublicLdapParams struct {

	/*Body*/
	Body *models_cloudbreak.LdapConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post public ldap params
func (o *PostPublicLdapParams) WithTimeout(timeout time.Duration) *PostPublicLdapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post public ldap params
func (o *PostPublicLdapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post public ldap params
func (o *PostPublicLdapParams) WithContext(ctx context.Context) *PostPublicLdapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post public ldap params
func (o *PostPublicLdapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post public ldap params
func (o *PostPublicLdapParams) WithHTTPClient(client *http.Client) *PostPublicLdapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post public ldap params
func (o *PostPublicLdapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post public ldap params
func (o *PostPublicLdapParams) WithBody(body *models_cloudbreak.LdapConfigRequest) *PostPublicLdapParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post public ldap params
func (o *PostPublicLdapParams) SetBody(body *models_cloudbreak.LdapConfigRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicLdapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.LdapConfigRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
