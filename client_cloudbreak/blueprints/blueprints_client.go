package blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new blueprints API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for blueprints API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
GetBlueprintsID retrieves blueprint by id

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetBlueprintsID(params *GetBlueprintsIDParams) (*GetBlueprintsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlueprintsIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetBlueprintsID",
		Method:             "GET",
		PathPattern:        "/blueprints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBlueprintsIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBlueprintsIDOK), nil
}

/*
Delete deletes blueprint by id

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) Delete(params *DeleteParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "delete",
		Method:             "DELETE",
		PathPattern:        "/blueprints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeletePrivate deletes private blueprint by name

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) DeletePrivate(params *DeletePrivateParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrivateParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "deletePrivate",
		Method:             "DELETE",
		PathPattern:        "/blueprints/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePrivateReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeletePublic deletes public owned or private blueprint by name

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) DeletePublic(params *DeletePublicParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublicParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "deletePublic",
		Method:             "DELETE",
		PathPattern:        "/blueprints/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePublicReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
GetPrivate retrieves a private blueprint by name

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetPrivate(params *GetPrivateParams) (*GetPrivateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPrivate",
		Method:             "GET",
		PathPattern:        "/blueprints/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateOK), nil
}

/*
GetPrivates retrieves private blueprints

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetPrivates(params *GetPrivatesParams) (*GetPrivatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivatesParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPrivates",
		Method:             "GET",
		PathPattern:        "/blueprints/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivatesReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivatesOK), nil
}

/*
GetPublic retrieves a public or private owned blueprint by name

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetPublic(params *GetPublicParams) (*GetPublicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPublic",
		Method:             "GET",
		PathPattern:        "/blueprints/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicOK), nil
}

/*
GetPublics retrieves public and private owned blueprints

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetPublics(params *GetPublicsParams) (*GetPublicsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPublics",
		Method:             "GET",
		PathPattern:        "/blueprints/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicsOK), nil
}

/*
PostPrivate creates blueprint as private resource

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) PostPrivate(params *PostPrivateParams) (*PostPrivateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPrivateParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "postPrivate",
		Method:             "POST",
		PathPattern:        "/blueprints/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPrivateReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPrivateOK), nil
}

/*
PostPublic creates blueprint as public resource

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) PostPublic(params *PostPublicParams) (*PostPublicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPublicParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "postPublic",
		Method:             "POST",
		PathPattern:        "/blueprints/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPublicReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPublicOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}