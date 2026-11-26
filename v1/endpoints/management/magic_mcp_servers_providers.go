package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/magicmcpservers/providers"
)

// MagicMcpServersProvidersEndpoint provides access to a magic MCP server is a stable MCP entrypoint with its own providers and tool filters. Servers can be grouped and bound to tokens, or combined behind an endpoint.
type MagicMcpServersProvidersEndpoint struct {
	client *endpoint.Client
}

// NewMagicMcpServersProvidersEndpoint creates a new MagicMcpServersProvidersEndpoint.
func NewMagicMcpServersProvidersEndpoint(client *endpoint.Client) *MagicMcpServersProvidersEndpoint {
	return &MagicMcpServersProvidersEndpoint{client: client}
}

// MagicMcpServersProvidersEndpointListParams contains optional query parameters for List.
type MagicMcpServersProvidersEndpointListParams struct {
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
	// Status - Only return providers with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return providers with these IDs.
	Id *any `json:"id,omitempty"`
	// ProviderId - Only return providers using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return providers using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// IntegrationInstanceProviderId - Only return providers using one of these integration instance providers.
	IntegrationInstanceProviderId *any `json:"integration_instance_provider_id,omitempty"`
	// ProviderDeploymentId - Only return providers using one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return providers using one of these configurations.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return providers using one of these auth configurations.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Only return results whose magic MCP server provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose magic MCP server provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// MagicMcpServersProvidersEndpointCreateBody contains the request body for Create.
type MagicMcpServersProvidersEndpointCreateBody struct {
	// ProviderId - Provider to add to the server.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Deployment of the provider to use.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Configuration of the provider to use. `null` clears it.
	ProviderConfigId *nullable.Nullable[string] `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Auth configuration of the provider to use. `null` clears it.
	ProviderAuthConfigId *nullable.Nullable[string] `json:"provider_auth_config_id,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// MagicMcpServersProvidersEndpointUpdateBody contains the request body for Update.
type MagicMcpServersProvidersEndpointUpdateBody struct {
	// ProviderDeploymentId - Deployment of the provider to use.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Configuration of the provider to use. `null` clears it.
	ProviderConfigId *nullable.Nullable[string] `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Auth configuration of the provider to use. `null` clears it.
	ProviderAuthConfigId *nullable.Nullable[string] `json:"provider_auth_config_id,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// List returns the backing integration instance providers configured for a magic MCP server.
func (e *MagicMcpServersProvidersEndpoint) List(instanceId string, magicMcpServerId string, params *MagicMcpServersProvidersEndpointListParams) (*providers.MagicMcpServersProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "magic-mcp-servers", magicMcpServerId, "providers"},
		Query: query,
	}
	var result providers.MagicMcpServersProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create adds a provider to a magic MCP server.
func (e *MagicMcpServersProvidersEndpoint) Create(instanceId string, magicMcpServerId string, body *MagicMcpServersProvidersEndpointCreateBody) (*providers.MagicMcpServersProvidersCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-servers", magicMcpServerId, "providers"},
		Body: body,
	}
	var result providers.MagicMcpServersProvidersCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a provider of a magic MCP server.
func (e *MagicMcpServersProvidersEndpoint) Get(instanceId string, magicMcpServerId string, magicMcpServerProviderId string) (*providers.MagicMcpServersProvidersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-servers", magicMcpServerId, "providers", magicMcpServerProviderId},
	}
	var result providers.MagicMcpServersProvidersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a provider of a magic MCP server. Omitted fields keep their current value.
func (e *MagicMcpServersProvidersEndpoint) Update(instanceId string, magicMcpServerId string, magicMcpServerProviderId string, body *MagicMcpServersProvidersEndpointUpdateBody) (*providers.MagicMcpServersProvidersUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-servers", magicMcpServerId, "providers", magicMcpServerProviderId},
		Body: body,
	}
	var result providers.MagicMcpServersProvidersUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a backing integration instance provider from a magic MCP server and removes the shared integration provider when unused.
func (e *MagicMcpServersProvidersEndpoint) Delete(instanceId string, magicMcpServerId string, magicMcpServerProviderId string) (*providers.MagicMcpServersProvidersDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-servers", magicMcpServerId, "providers", magicMcpServerProviderId},
	}
	var result providers.MagicMcpServersProvidersDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
