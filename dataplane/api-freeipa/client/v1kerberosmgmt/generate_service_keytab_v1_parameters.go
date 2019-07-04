// Code generated by go-swagger; DO NOT EDIT.

package v1kerberosmgmt

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

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// NewGenerateServiceKeytabV1Params creates a new GenerateServiceKeytabV1Params object
// with the default values initialized.
func NewGenerateServiceKeytabV1Params() *GenerateServiceKeytabV1Params {
	var ()
	return &GenerateServiceKeytabV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateServiceKeytabV1ParamsWithTimeout creates a new GenerateServiceKeytabV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGenerateServiceKeytabV1ParamsWithTimeout(timeout time.Duration) *GenerateServiceKeytabV1Params {
	var ()
	return &GenerateServiceKeytabV1Params{

		timeout: timeout,
	}
}

// NewGenerateServiceKeytabV1ParamsWithContext creates a new GenerateServiceKeytabV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGenerateServiceKeytabV1ParamsWithContext(ctx context.Context) *GenerateServiceKeytabV1Params {
	var ()
	return &GenerateServiceKeytabV1Params{

		Context: ctx,
	}
}

// NewGenerateServiceKeytabV1ParamsWithHTTPClient creates a new GenerateServiceKeytabV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGenerateServiceKeytabV1ParamsWithHTTPClient(client *http.Client) *GenerateServiceKeytabV1Params {
	var ()
	return &GenerateServiceKeytabV1Params{
		HTTPClient: client,
	}
}

/*GenerateServiceKeytabV1Params contains all the parameters to send to the API endpoint
for the generate service keytab v1 operation typically these are written to a http.Request
*/
type GenerateServiceKeytabV1Params struct {

	/*Body*/
	Body *model.ServiceKeytabV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the generate service keytab v1 params
func (o *GenerateServiceKeytabV1Params) WithTimeout(timeout time.Duration) *GenerateServiceKeytabV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate service keytab v1 params
func (o *GenerateServiceKeytabV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate service keytab v1 params
func (o *GenerateServiceKeytabV1Params) WithContext(ctx context.Context) *GenerateServiceKeytabV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate service keytab v1 params
func (o *GenerateServiceKeytabV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate service keytab v1 params
func (o *GenerateServiceKeytabV1Params) WithHTTPClient(client *http.Client) *GenerateServiceKeytabV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate service keytab v1 params
func (o *GenerateServiceKeytabV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the generate service keytab v1 params
func (o *GenerateServiceKeytabV1Params) WithBody(body *model.ServiceKeytabV1Request) *GenerateServiceKeytabV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the generate service keytab v1 params
func (o *GenerateServiceKeytabV1Params) SetBody(body *model.ServiceKeytabV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateServiceKeytabV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}