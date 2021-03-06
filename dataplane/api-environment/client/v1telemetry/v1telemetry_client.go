// Code generated by go-swagger; DO NOT EDIT.

package v1telemetry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1telemetry API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1telemetry API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAccountTelemetryV1 gets account specific telemetry settings

Gather telemetry related settings (features and anonymization rules) these are used as global settings for environments
*/
func (a *Client) GetAccountTelemetryV1(params *GetAccountTelemetryV1Params) (*GetAccountTelemetryV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountTelemetryV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccountTelemetryV1",
		Method:             "GET",
		PathPattern:        "/v1/telemetry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccountTelemetryV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountTelemetryV1OK), nil

}

/*
GetDefaultAccountTelemetryV1 gets default account specific telemetry settings

Gather default telemetry account level settings - default rules contains email/card number regex patterns
*/
func (a *Client) GetDefaultAccountTelemetryV1(params *GetDefaultAccountTelemetryV1Params) (*GetDefaultAccountTelemetryV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultAccountTelemetryV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDefaultAccountTelemetryV1",
		Method:             "GET",
		PathPattern:        "/v1/telemetry/default",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDefaultAccountTelemetryV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDefaultAccountTelemetryV1OK), nil

}

/*
ListFeaturesV1 gets account specific telemetry settings

Gather telemetry related settings (features and anonymization rules) these are used as global settings for environments
*/
func (a *Client) ListFeaturesV1(params *ListFeaturesV1Params) (*ListFeaturesV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFeaturesV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listFeaturesV1",
		Method:             "GET",
		PathPattern:        "/v1/telemetry/features",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListFeaturesV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListFeaturesV1OK), nil

}

/*
ListRulesInAccountV1 gets account specific telemetry settings

Gather telemetry related settings (features and anonymization rules) these are used as global settings for environments
*/
func (a *Client) ListRulesInAccountV1(params *ListRulesInAccountV1Params) (*ListRulesInAccountV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRulesInAccountV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRulesInAccountV1",
		Method:             "GET",
		PathPattern:        "/v1/telemetry/rules/{accountId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRulesInAccountV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRulesInAccountV1OK), nil

}

/*
ListRulesV1 gets account specific telemetry settings

Gather telemetry related settings (features and anonymization rules) these are used as global settings for environments
*/
func (a *Client) ListRulesV1(params *ListRulesV1Params) (*ListRulesV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRulesV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRulesV1",
		Method:             "GET",
		PathPattern:        "/v1/telemetry/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRulesV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRulesV1OK), nil

}

/*
TestRuleV1 tests anonymization pattern

Testing anonymization pattern - check the pattern is valid against an input and found matches
*/
func (a *Client) TestRuleV1(params *TestRuleV1Params) (*TestRuleV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestRuleV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "testRuleV1",
		Method:             "POST",
		PathPattern:        "/v1/telemetry/rules/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TestRuleV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TestRuleV1OK), nil

}

/*
UpdateAccountTelemetryV1 updates account specific telemetry settings

Update account level telemetry settings - will override already existing settings and archive old one
*/
func (a *Client) UpdateAccountTelemetryV1(params *UpdateAccountTelemetryV1Params) (*UpdateAccountTelemetryV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAccountTelemetryV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAccountTelemetryV1",
		Method:             "POST",
		PathPattern:        "/v1/telemetry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateAccountTelemetryV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateAccountTelemetryV1OK), nil

}

/*
UpdateRulesV1 gets account specific telemetry feature settings

For account level telemetry settings, different features can be defined globally, these features can be overriden by environemnt level telemetry settings
*/
func (a *Client) UpdateRulesV1(params *UpdateRulesV1Params) (*UpdateRulesV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRulesV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRulesV1",
		Method:             "POST",
		PathPattern:        "/v1/telemetry/features",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateRulesV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateRulesV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
