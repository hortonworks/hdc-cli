// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_audits

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

// NewGetAuditEventsZipInOrganizationParams creates a new GetAuditEventsZipInOrganizationParams object
// with the default values initialized.
func NewGetAuditEventsZipInOrganizationParams() *GetAuditEventsZipInOrganizationParams {
	var ()
	return &GetAuditEventsZipInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuditEventsZipInOrganizationParamsWithTimeout creates a new GetAuditEventsZipInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuditEventsZipInOrganizationParamsWithTimeout(timeout time.Duration) *GetAuditEventsZipInOrganizationParams {
	var ()
	return &GetAuditEventsZipInOrganizationParams{

		timeout: timeout,
	}
}

// NewGetAuditEventsZipInOrganizationParamsWithContext creates a new GetAuditEventsZipInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuditEventsZipInOrganizationParamsWithContext(ctx context.Context) *GetAuditEventsZipInOrganizationParams {
	var ()
	return &GetAuditEventsZipInOrganizationParams{

		Context: ctx,
	}
}

// NewGetAuditEventsZipInOrganizationParamsWithHTTPClient creates a new GetAuditEventsZipInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuditEventsZipInOrganizationParamsWithHTTPClient(client *http.Client) *GetAuditEventsZipInOrganizationParams {
	var ()
	return &GetAuditEventsZipInOrganizationParams{
		HTTPClient: client,
	}
}

/*GetAuditEventsZipInOrganizationParams contains all the parameters to send to the API endpoint
for the get audit events zip in organization operation typically these are written to a http.Request
*/
type GetAuditEventsZipInOrganizationParams struct {

	/*OrganizationID*/
	OrganizationID int64
	/*ResourceID*/
	ResourceID int64
	/*ResourceType*/
	ResourceType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) WithTimeout(timeout time.Duration) *GetAuditEventsZipInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) WithContext(ctx context.Context) *GetAuditEventsZipInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) WithHTTPClient(client *http.Client) *GetAuditEventsZipInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) WithOrganizationID(organizationID int64) *GetAuditEventsZipInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WithResourceID adds the resourceID to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) WithResourceID(resourceID int64) *GetAuditEventsZipInOrganizationParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) SetResourceID(resourceID int64) {
	o.ResourceID = resourceID
}

// WithResourceType adds the resourceType to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) WithResourceType(resourceType string) *GetAuditEventsZipInOrganizationParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get audit events zip in organization params
func (o *GetAuditEventsZipInOrganizationParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuditEventsZipInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	// path param resourceId
	if err := r.SetPathParam("resourceId", swag.FormatInt64(o.ResourceID)); err != nil {
		return err
	}

	// path param resourceType
	if err := r.SetPathParam("resourceType", o.ResourceType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}