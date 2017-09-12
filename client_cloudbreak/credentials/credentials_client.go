// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new credentials API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for credentials API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCredential deletes credential by id

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) DeleteCredential(params *DeleteCredentialParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCredentialParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCredential",
		Method:             "DELETE",
		PathPattern:        "/credentials/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePrivateCredential deletes private credential by name

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) DeletePrivateCredential(params *DeletePrivateCredentialParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrivateCredentialParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePrivateCredential",
		Method:             "DELETE",
		PathPattern:        "/credentials/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePrivateCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePublicCredential deletes public owned or private credential by name

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) DeletePublicCredential(params *DeletePublicCredentialParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublicCredentialParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePublicCredential",
		Method:             "DELETE",
		PathPattern:        "/credentials/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePublicCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetCredential retrieves credential by id

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetCredential(params *GetCredentialParams) (*GetCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCredential",
		Method:             "GET",
		PathPattern:        "/credentials/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCredentialOK), nil

}

/*
GetPrivateCredential retrieves a private credential by name

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetPrivateCredential(params *GetPrivateCredentialParams) (*GetPrivateCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivateCredential",
		Method:             "GET",
		PathPattern:        "/credentials/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateCredentialOK), nil

}

/*
GetPrivatesCredential retrieves private credentials

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetPrivatesCredential(params *GetPrivatesCredentialParams) (*GetPrivatesCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivatesCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivatesCredential",
		Method:             "GET",
		PathPattern:        "/credentials/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivatesCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivatesCredentialOK), nil

}

/*
GetPublicCredential retrieves a public or private owned credential by name

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetPublicCredential(params *GetPublicCredentialParams) (*GetPublicCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicCredential",
		Method:             "GET",
		PathPattern:        "/credentials/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicCredentialOK), nil

}

/*
GetPublicsCredential retrieves public and private owned credentials

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetPublicsCredential(params *GetPublicsCredentialParams) (*GetPublicsCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicsCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicsCredential",
		Method:             "GET",
		PathPattern:        "/credentials/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicsCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicsCredentialOK), nil

}

/*
PostPrivateCredential creates credential as private resource

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) PostPrivateCredential(params *PostPrivateCredentialParams) (*PostPrivateCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPrivateCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPrivateCredential",
		Method:             "POST",
		PathPattern:        "/credentials/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPrivateCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPrivateCredentialOK), nil

}

/*
PostPublicCredential creates credential as public resource

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) PostPublicCredential(params *PostPublicCredentialParams) (*PostPublicCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPublicCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPublicCredential",
		Method:             "POST",
		PathPattern:        "/credentials/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPublicCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPublicCredentialOK), nil

}

/*
PrivateInteractiveLoginCredential interactives login

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) PrivateInteractiveLoginCredential(params *PrivateInteractiveLoginCredentialParams) (*PrivateInteractiveLoginCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrivateInteractiveLoginCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "privateInteractiveLoginCredential",
		Method:             "POST",
		PathPattern:        "/credentials/userinteractivelogin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PrivateInteractiveLoginCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PrivateInteractiveLoginCredentialOK), nil

}

/*
PublicInteractiveLoginCredential interactives login

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) PublicInteractiveLoginCredential(params *PublicInteractiveLoginCredentialParams) (*PublicInteractiveLoginCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicInteractiveLoginCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicInteractiveLoginCredential",
		Method:             "POST",
		PathPattern:        "/credentials/accountinteractivelogin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PublicInteractiveLoginCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PublicInteractiveLoginCredentialOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
