// Code generated by go-swagger; DO NOT EDIT.

package credentials

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

// NewPrivateInteractiveLoginCredentialParams creates a new PrivateInteractiveLoginCredentialParams object
// with the default values initialized.
func NewPrivateInteractiveLoginCredentialParams() *PrivateInteractiveLoginCredentialParams {
	var ()
	return &PrivateInteractiveLoginCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPrivateInteractiveLoginCredentialParamsWithTimeout creates a new PrivateInteractiveLoginCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPrivateInteractiveLoginCredentialParamsWithTimeout(timeout time.Duration) *PrivateInteractiveLoginCredentialParams {
	var ()
	return &PrivateInteractiveLoginCredentialParams{

		timeout: timeout,
	}
}

// NewPrivateInteractiveLoginCredentialParamsWithContext creates a new PrivateInteractiveLoginCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewPrivateInteractiveLoginCredentialParamsWithContext(ctx context.Context) *PrivateInteractiveLoginCredentialParams {
	var ()
	return &PrivateInteractiveLoginCredentialParams{

		Context: ctx,
	}
}

// NewPrivateInteractiveLoginCredentialParamsWithHTTPClient creates a new PrivateInteractiveLoginCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPrivateInteractiveLoginCredentialParamsWithHTTPClient(client *http.Client) *PrivateInteractiveLoginCredentialParams {
	var ()
	return &PrivateInteractiveLoginCredentialParams{
		HTTPClient: client,
	}
}

/*PrivateInteractiveLoginCredentialParams contains all the parameters to send to the API endpoint
for the private interactive login credential operation typically these are written to a http.Request
*/
type PrivateInteractiveLoginCredentialParams struct {

	/*Body*/
	Body *models_cloudbreak.CredentialRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the private interactive login credential params
func (o *PrivateInteractiveLoginCredentialParams) WithTimeout(timeout time.Duration) *PrivateInteractiveLoginCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the private interactive login credential params
func (o *PrivateInteractiveLoginCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the private interactive login credential params
func (o *PrivateInteractiveLoginCredentialParams) WithContext(ctx context.Context) *PrivateInteractiveLoginCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the private interactive login credential params
func (o *PrivateInteractiveLoginCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the private interactive login credential params
func (o *PrivateInteractiveLoginCredentialParams) WithHTTPClient(client *http.Client) *PrivateInteractiveLoginCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the private interactive login credential params
func (o *PrivateInteractiveLoginCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the private interactive login credential params
func (o *PrivateInteractiveLoginCredentialParams) WithBody(body *models_cloudbreak.CredentialRequest) *PrivateInteractiveLoginCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the private interactive login credential params
func (o *PrivateInteractiveLoginCredentialParams) SetBody(body *models_cloudbreak.CredentialRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PrivateInteractiveLoginCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.CredentialRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
