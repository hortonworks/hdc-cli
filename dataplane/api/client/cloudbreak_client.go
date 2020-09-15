// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/api/client/authorization"
	"github.com/hortonworks/cb-cli/dataplane/api/client/autoscale"
	"github.com/hortonworks/cb-cli/dataplane/api/client/flow"
	"github.com/hortonworks/cb-cli/dataplane/api/client/flow_public"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v1distrox"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v1internaldistrox"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_audits"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_blueprints"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_blueprints_util"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_cluster_templates"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_databases"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_file_systems"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_imagecatalogs"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_recipes"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_stacks"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4datalake"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4dbconfig"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4diagnostics"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4events"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4info"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4user_profiles"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4utils"
)

// Default cloudbreak HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new cloudbreak HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Cloudbreak {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cloudbreak HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Cloudbreak {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cloudbreak client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Cloudbreak {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Cloudbreak)
	cli.Transport = transport

	cli.Authorization = authorization.New(transport, formats)

	cli.Autoscale = autoscale.New(transport, formats)

	cli.Flow = flow.New(transport, formats)

	cli.FlowPublic = flow_public.New(transport, formats)

	cli.V1distrox = v1distrox.New(transport, formats)

	cli.V1internaldistrox = v1internaldistrox.New(transport, formats)

	cli.V4WorkspaceID = v4_workspace_id.New(transport, formats)

	cli.V4WorkspaceIDAudits = v4_workspace_id_audits.New(transport, formats)

	cli.V4WorkspaceIDBlueprints = v4_workspace_id_blueprints.New(transport, formats)

	cli.V4WorkspaceIDBlueprintsUtil = v4_workspace_id_blueprints_util.New(transport, formats)

	cli.V4WorkspaceIDClusterTemplates = v4_workspace_id_cluster_templates.New(transport, formats)

	cli.V4WorkspaceIDDatabases = v4_workspace_id_databases.New(transport, formats)

	cli.V4WorkspaceIDFileSystems = v4_workspace_id_file_systems.New(transport, formats)

	cli.V4WorkspaceIDImagecatalogs = v4_workspace_id_imagecatalogs.New(transport, formats)

	cli.V4WorkspaceIDRecipes = v4_workspace_id_recipes.New(transport, formats)

	cli.V4WorkspaceIDStacks = v4_workspace_id_stacks.New(transport, formats)

	cli.V4datalake = v4datalake.New(transport, formats)

	cli.V4dbconfig = v4dbconfig.New(transport, formats)

	cli.V4diagnostics = v4diagnostics.New(transport, formats)

	cli.V4events = v4events.New(transport, formats)

	cli.V4info = v4info.New(transport, formats)

	cli.V4userProfiles = v4user_profiles.New(transport, formats)

	cli.V4utils = v4utils.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Cloudbreak is a client for cloudbreak
type Cloudbreak struct {
	Authorization *authorization.Client

	Autoscale *autoscale.Client

	Flow *flow.Client

	FlowPublic *flow_public.Client

	V1distrox *v1distrox.Client

	V1internaldistrox *v1internaldistrox.Client

	V4WorkspaceID *v4_workspace_id.Client

	V4WorkspaceIDAudits *v4_workspace_id_audits.Client

	V4WorkspaceIDBlueprints *v4_workspace_id_blueprints.Client

	V4WorkspaceIDBlueprintsUtil *v4_workspace_id_blueprints_util.Client

	V4WorkspaceIDClusterTemplates *v4_workspace_id_cluster_templates.Client

	V4WorkspaceIDDatabases *v4_workspace_id_databases.Client

	V4WorkspaceIDFileSystems *v4_workspace_id_file_systems.Client

	V4WorkspaceIDImagecatalogs *v4_workspace_id_imagecatalogs.Client

	V4WorkspaceIDRecipes *v4_workspace_id_recipes.Client

	V4WorkspaceIDStacks *v4_workspace_id_stacks.Client

	V4datalake *v4datalake.Client

	V4dbconfig *v4dbconfig.Client

	V4diagnostics *v4diagnostics.Client

	V4events *v4events.Client

	V4info *v4info.Client

	V4userProfiles *v4user_profiles.Client

	V4utils *v4utils.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Cloudbreak) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Authorization.SetTransport(transport)

	c.Autoscale.SetTransport(transport)

	c.Flow.SetTransport(transport)

	c.FlowPublic.SetTransport(transport)

	c.V1distrox.SetTransport(transport)

	c.V1internaldistrox.SetTransport(transport)

	c.V4WorkspaceID.SetTransport(transport)

	c.V4WorkspaceIDAudits.SetTransport(transport)

	c.V4WorkspaceIDBlueprints.SetTransport(transport)

	c.V4WorkspaceIDBlueprintsUtil.SetTransport(transport)

	c.V4WorkspaceIDClusterTemplates.SetTransport(transport)

	c.V4WorkspaceIDDatabases.SetTransport(transport)

	c.V4WorkspaceIDFileSystems.SetTransport(transport)

	c.V4WorkspaceIDImagecatalogs.SetTransport(transport)

	c.V4WorkspaceIDRecipes.SetTransport(transport)

	c.V4WorkspaceIDStacks.SetTransport(transport)

	c.V4datalake.SetTransport(transport)

	c.V4dbconfig.SetTransport(transport)

	c.V4diagnostics.SetTransport(transport)

	c.V4events.SetTransport(transport)

	c.V4info.SetTransport(transport)

	c.V4userProfiles.SetTransport(transport)

	c.V4utils.SetTransport(transport)

}
