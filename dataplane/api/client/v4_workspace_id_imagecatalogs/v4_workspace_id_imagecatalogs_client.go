// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_imagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4 workspace id imagecatalogs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4 workspace id imagecatalogs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateImageCatalogInWorkspace creates image catalog in workspace

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) CreateImageCatalogInWorkspace(params *CreateImageCatalogInWorkspaceParams) (*CreateImageCatalogInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateImageCatalogInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createImageCatalogInWorkspace",
		Method:             "POST",
		PathPattern:        "/v4/{workspaceId}/image_catalogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateImageCatalogInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateImageCatalogInWorkspaceOK), nil

}

/*
DeleteImageCatalogByCrnInWorkspace deletes image catalog by crn in workspace

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) DeleteImageCatalogByCrnInWorkspace(params *DeleteImageCatalogByCrnInWorkspaceParams) (*DeleteImageCatalogByCrnInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImageCatalogByCrnInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteImageCatalogByCrnInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/image_catalogs/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteImageCatalogByCrnInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteImageCatalogByCrnInWorkspaceOK), nil

}

/*
DeleteImageCatalogInWorkspace deletes image catalog by name in workspace

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) DeleteImageCatalogInWorkspace(params *DeleteImageCatalogInWorkspaceParams) (*DeleteImageCatalogInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImageCatalogInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteImageCatalogInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/image_catalogs/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteImageCatalogInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteImageCatalogInWorkspaceOK), nil

}

/*
DeleteImageCatalogsInWorkspace deletes multiple image catalogs by name in workspace

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) DeleteImageCatalogsInWorkspace(params *DeleteImageCatalogsInWorkspaceParams) (*DeleteImageCatalogsInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImageCatalogsInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteImageCatalogsInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/image_catalogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteImageCatalogsInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteImageCatalogsInWorkspaceOK), nil

}

/*
GetImageCatalogByCrnInWorkspace gets image catalog by crn in workspace

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) GetImageCatalogByCrnInWorkspace(params *GetImageCatalogByCrnInWorkspaceParams) (*GetImageCatalogByCrnInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageCatalogByCrnInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getImageCatalogByCrnInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/image_catalogs/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageCatalogByCrnInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetImageCatalogByCrnInWorkspaceOK), nil

}

/*
GetImageCatalogInWorkspace gets image catalog by name in workspace

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) GetImageCatalogInWorkspace(params *GetImageCatalogInWorkspaceParams) (*GetImageCatalogInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageCatalogInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getImageCatalogInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/image_catalogs/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageCatalogInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetImageCatalogInWorkspaceOK), nil

}

/*
GetImageCatalogRequestFromNameInWorkspace retrieves imagecatalog request by imagecatalog name

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) GetImageCatalogRequestFromNameInWorkspace(params *GetImageCatalogRequestFromNameInWorkspaceParams) (*GetImageCatalogRequestFromNameInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageCatalogRequestFromNameInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getImageCatalogRequestFromNameInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/image_catalogs/{name}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageCatalogRequestFromNameInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetImageCatalogRequestFromNameInWorkspaceOK), nil

}

/*
GetImagesByNameInWorkspace determines available images for the given stack or platformfrom the given imagecatalog name

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) GetImagesByNameInWorkspace(params *GetImagesByNameInWorkspaceParams) (*GetImagesByNameInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImagesByNameInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getImagesByNameInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/image_catalogs/{name}/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImagesByNameInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetImagesByNameInWorkspaceOK), nil

}

/*
GetImagesInWorkspace determines available images for the given stack or platformfrom the default image catalog

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) GetImagesInWorkspace(params *GetImagesInWorkspaceParams) (*GetImagesInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImagesInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getImagesInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/image_catalogs/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImagesInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetImagesInWorkspaceOK), nil

}

/*
ListImageCatalogsByWorkspace lists image catalogs for the given workspace

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) ListImageCatalogsByWorkspace(params *ListImageCatalogsByWorkspaceParams) (*ListImageCatalogsByWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListImageCatalogsByWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listImageCatalogsByWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/image_catalogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListImageCatalogsByWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListImageCatalogsByWorkspaceOK), nil

}

/*
UpdateImageCatalogInWorkspace updates image catalog by id

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) UpdateImageCatalogInWorkspace(params *UpdateImageCatalogInWorkspaceParams) (*UpdateImageCatalogInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateImageCatalogInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateImageCatalogInWorkspace",
		Method:             "PUT",
		PathPattern:        "/v4/{workspaceId}/image_catalogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateImageCatalogInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateImageCatalogInWorkspaceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
