// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

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

// NewStartDistroXV1ByCrnsParams creates a new StartDistroXV1ByCrnsParams object
// with the default values initialized.
func NewStartDistroXV1ByCrnsParams() *StartDistroXV1ByCrnsParams {
	var ()
	return &StartDistroXV1ByCrnsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartDistroXV1ByCrnsParamsWithTimeout creates a new StartDistroXV1ByCrnsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartDistroXV1ByCrnsParamsWithTimeout(timeout time.Duration) *StartDistroXV1ByCrnsParams {
	var ()
	return &StartDistroXV1ByCrnsParams{

		timeout: timeout,
	}
}

// NewStartDistroXV1ByCrnsParamsWithContext creates a new StartDistroXV1ByCrnsParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartDistroXV1ByCrnsParamsWithContext(ctx context.Context) *StartDistroXV1ByCrnsParams {
	var ()
	return &StartDistroXV1ByCrnsParams{

		Context: ctx,
	}
}

// NewStartDistroXV1ByCrnsParamsWithHTTPClient creates a new StartDistroXV1ByCrnsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartDistroXV1ByCrnsParamsWithHTTPClient(client *http.Client) *StartDistroXV1ByCrnsParams {
	var ()
	return &StartDistroXV1ByCrnsParams{
		HTTPClient: client,
	}
}

/*StartDistroXV1ByCrnsParams contains all the parameters to send to the API endpoint
for the start distro x v1 by crns operation typically these are written to a http.Request
*/
type StartDistroXV1ByCrnsParams struct {

	/*Crns*/
	Crns []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start distro x v1 by crns params
func (o *StartDistroXV1ByCrnsParams) WithTimeout(timeout time.Duration) *StartDistroXV1ByCrnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start distro x v1 by crns params
func (o *StartDistroXV1ByCrnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start distro x v1 by crns params
func (o *StartDistroXV1ByCrnsParams) WithContext(ctx context.Context) *StartDistroXV1ByCrnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start distro x v1 by crns params
func (o *StartDistroXV1ByCrnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start distro x v1 by crns params
func (o *StartDistroXV1ByCrnsParams) WithHTTPClient(client *http.Client) *StartDistroXV1ByCrnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start distro x v1 by crns params
func (o *StartDistroXV1ByCrnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrns adds the crns to the start distro x v1 by crns params
func (o *StartDistroXV1ByCrnsParams) WithCrns(crns []string) *StartDistroXV1ByCrnsParams {
	o.SetCrns(crns)
	return o
}

// SetCrns adds the crns to the start distro x v1 by crns params
func (o *StartDistroXV1ByCrnsParams) SetCrns(crns []string) {
	o.Crns = crns
}

// WriteToRequest writes these params to a swagger request
func (o *StartDistroXV1ByCrnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesCrns := o.Crns

	joinedCrns := swag.JoinByFormat(valuesCrns, "multi")
	// query array param crns
	if err := r.SetQueryParam("crns", joinedCrns...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}