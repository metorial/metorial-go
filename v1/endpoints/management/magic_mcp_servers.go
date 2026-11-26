package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/magicmcpservers"
)

// MagicMcpServersEndpoint provides access to a magic MCP server is a stable MCP entrypoint with its own providers and tool filters. Servers can be grouped and bound to tokens, or combined behind an endpoint.
type MagicMcpServersEndpoint struct {
	client *endpoint.Client
}

// NewMagicMcpServersEndpoint creates a new MagicMcpServersEndpoint.
func NewMagicMcpServersEndpoint(client *endpoint.Client) *MagicMcpServersEndpoint {
	return &MagicMcpServersEndpoint{client: client}
}

// MagicMcpServersEndpointListParams contains optional query parameters for List.
type MagicMcpServersEndpointListParams struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Ownership - Which resources to return under the resource ownership rules. 'shared' returns only resources that aren't owned, 'usable' returns shared resources plus owned resources the caller can use, 'all' returns every resource the caller can see. Defaults to 'usable' for consumers, and otherwise to 'shared' when ownership is enforced for the project and 'all' when it isn't.
	Ownership *string `json:"ownership,omitempty"`
	// Status - Only return servers with one of these statuses.
	Status *any `json:"status,omitempty"`
	// MagicMcpGroupId - Only return servers in these groups.
	MagicMcpGroupId *any `json:"magic_mcp_group_id,omitempty"`
	// ProviderTemplateId - Only return servers built from these provider templates.
	ProviderTemplateId *any `json:"provider_template_id,omitempty"`
	// IntegrationInstanceId - Only return servers built from these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// Owner - Only return organization-owned or consumer-owned servers.
	Owner *any `json:"owner,omitempty"`
	// ProviderId - Only return servers using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ConsumerId - Only return servers owned by these consumers.
	ConsumerId *any `json:"consumer_id,omitempty"`
	// ConsumerProfileId - Only return servers owned by these consumer profiles.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// Search - Search by name or description.
	Search *string `json:"search,omitempty"`
	// Id - Only return servers with these IDs.
	Id *any `json:"id,omitempty"`
	// PreconfiguredOnly - Only return preconfigured servers.
	PreconfiguredOnly *bool `json:"preconfigured_only,omitempty"`
}

// MagicMcpServersEndpointCreateBody contains the request body for Create.
type MagicMcpServersEndpointCreateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the server. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderTemplateId - Provider template to build the server from.
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	// IntegrationInstanceId - Integration instance to build the server from.
	IntegrationInstanceId *string `json:"integration_instance_id,omitempty"`
	// ConsumerProfileId - Consumer profile that will own the server.
	ConsumerProfileId *string `json:"consumer_profile_id,omitempty"`
}

// MagicMcpServersEndpointUpdateBody contains the request body for Update.
type MagicMcpServersEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the server. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Aliases - Connection slugs for the server.
	Aliases *[]string `json:"aliases,omitempty"`
}

// List returns a paginated list of magic MCP servers.
func (e *MagicMcpServersEndpoint) List(instanceId string, params *MagicMcpServersEndpointListParams) (*magicmcpservers.MagicMcpServersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "magic-mcp-servers"},
		Query: query,
	}
	var result magicmcpservers.MagicMcpServersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific magic MCP server.
func (e *MagicMcpServersEndpoint) Get(instanceId string, magicMcpServerId string) (*magicmcpservers.MagicMcpServersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-servers", magicMcpServerId},
	}
	var result magicmcpservers.MagicMcpServersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Tools returns the effective set of tools available through the providers backing a magic MCP server.
func (e *MagicMcpServersEndpoint) Tools(instanceId string, magicMcpServerId string) (*magicmcpservers.MagicMcpServersToolsOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-servers", magicMcpServerId, "tools"},
	}
	var result magicmcpservers.MagicMcpServersToolsOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a magic MCP server. A session is created automatically on first connection and then reused.
func (e *MagicMcpServersEndpoint) Create(instanceId string, body *MagicMcpServersEndpointCreateBody) (*magicmcpservers.MagicMcpServersCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-servers"},
		Body: body,
	}
	var result magicmcpservers.MagicMcpServersCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a magic MCP server.
func (e *MagicMcpServersEndpoint) Delete(instanceId string, magicMcpServerId string) (*magicmcpservers.MagicMcpServersDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-servers", magicMcpServerId},
	}
	var result magicmcpservers.MagicMcpServersDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a magic MCP server. Omitted fields keep their current value.
func (e *MagicMcpServersEndpoint) Update(instanceId string, magicMcpServerId string, body *MagicMcpServersEndpointUpdateBody) (*magicmcpservers.MagicMcpServersUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-servers", magicMcpServerId},
		Body: body,
	}
	var result magicmcpservers.MagicMcpServersUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
