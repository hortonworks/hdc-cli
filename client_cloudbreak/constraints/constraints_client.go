// Code generated by go-swagger; DO NOT EDIT.

package constraints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new constraints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for constraints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteConstraint deletes constraint template by id

A constraint template tells Cloudbreak the resource constraints (cpu, memory, disk) of the Ambari containers that will be deployed to the cluster. A constraint template must be created onenvironments where there is no one-to-one mapping between containers and nodes, like Mesos.
*/
func (a *Client) DeleteConstraint(params *DeleteConstraintParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConstraintParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteConstraint",
		Method:             "DELETE",
		PathPattern:        "/constraints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteConstraintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePrivateConstraint deletes private constraint template by name

A constraint template tells Cloudbreak the resource constraints (cpu, memory, disk) of the Ambari containers that will be deployed to the cluster. A constraint template must be created onenvironments where there is no one-to-one mapping between containers and nodes, like Mesos.
*/
func (a *Client) DeletePrivateConstraint(params *DeletePrivateConstraintParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrivateConstraintParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePrivateConstraint",
		Method:             "DELETE",
		PathPattern:        "/constraints/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePrivateConstraintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePublicConstraint deletes public owned or private constraint template by name

A constraint template tells Cloudbreak the resource constraints (cpu, memory, disk) of the Ambari containers that will be deployed to the cluster. A constraint template must be created onenvironments where there is no one-to-one mapping between containers and nodes, like Mesos.
*/
func (a *Client) DeletePublicConstraint(params *DeletePublicConstraintParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublicConstraintParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePublicConstraint",
		Method:             "DELETE",
		PathPattern:        "/constraints/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePublicConstraintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetConstraint retrieves constraint template by id

A constraint template tells Cloudbreak the resource constraints (cpu, memory, disk) of the Ambari containers that will be deployed to the cluster. A constraint template must be created onenvironments where there is no one-to-one mapping between containers and nodes, like Mesos.
*/
func (a *Client) GetConstraint(params *GetConstraintParams) (*GetConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConstraintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConstraint",
		Method:             "GET",
		PathPattern:        "/constraints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConstraintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConstraintOK), nil

}

/*
GetPrivateConstraint retrieves a private constraint template by name

A constraint template tells Cloudbreak the resource constraints (cpu, memory, disk) of the Ambari containers that will be deployed to the cluster. A constraint template must be created onenvironments where there is no one-to-one mapping between containers and nodes, like Mesos.
*/
func (a *Client) GetPrivateConstraint(params *GetPrivateConstraintParams) (*GetPrivateConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateConstraintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivateConstraint",
		Method:             "GET",
		PathPattern:        "/constraints/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateConstraintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateConstraintOK), nil

}

/*
GetPrivatesConstraint retrieves private constraint templates

A constraint template tells Cloudbreak the resource constraints (cpu, memory, disk) of the Ambari containers that will be deployed to the cluster. A constraint template must be created onenvironments where there is no one-to-one mapping between containers and nodes, like Mesos.
*/
func (a *Client) GetPrivatesConstraint(params *GetPrivatesConstraintParams) (*GetPrivatesConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivatesConstraintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivatesConstraint",
		Method:             "GET",
		PathPattern:        "/constraints/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivatesConstraintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivatesConstraintOK), nil

}

/*
GetPublicConstraint retrieves a public or private owned constraint template by name

A constraint template tells Cloudbreak the resource constraints (cpu, memory, disk) of the Ambari containers that will be deployed to the cluster. A constraint template must be created onenvironments where there is no one-to-one mapping between containers and nodes, like Mesos.
*/
func (a *Client) GetPublicConstraint(params *GetPublicConstraintParams) (*GetPublicConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicConstraintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicConstraint",
		Method:             "GET",
		PathPattern:        "/constraints/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicConstraintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicConstraintOK), nil

}

/*
GetPublicsConstraint retrieves public and private owned constraint templates

A constraint template tells Cloudbreak the resource constraints (cpu, memory, disk) of the Ambari containers that will be deployed to the cluster. A constraint template must be created onenvironments where there is no one-to-one mapping between containers and nodes, like Mesos.
*/
func (a *Client) GetPublicsConstraint(params *GetPublicsConstraintParams) (*GetPublicsConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicsConstraintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicsConstraint",
		Method:             "GET",
		PathPattern:        "/constraints/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicsConstraintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicsConstraintOK), nil

}

/*
PostPrivateConstraint creates constraint template as private resource

A constraint template tells Cloudbreak the resource constraints (cpu, memory, disk) of the Ambari containers that will be deployed to the cluster. A constraint template must be created onenvironments where there is no one-to-one mapping between containers and nodes, like Mesos.
*/
func (a *Client) PostPrivateConstraint(params *PostPrivateConstraintParams) (*PostPrivateConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPrivateConstraintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPrivateConstraint",
		Method:             "POST",
		PathPattern:        "/constraints/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPrivateConstraintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPrivateConstraintOK), nil

}

/*
PostPublicConstraint creates constraint template as public resource

A constraint template tells Cloudbreak the resource constraints (cpu, memory, disk) of the Ambari containers that will be deployed to the cluster. A constraint template must be created onenvironments where there is no one-to-one mapping between containers and nodes, like Mesos.
*/
func (a *Client) PostPublicConstraint(params *PostPublicConstraintParams) (*PostPublicConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPublicConstraintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPublicConstraint",
		Method:             "POST",
		PathPattern:        "/constraints/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPublicConstraintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPublicConstraintOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
