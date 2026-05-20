package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/magicmcpservers/providers"
)

// MagicMcpServersProvidersEndpoint provides access to magic MCP servers are stable MCP entrypoints backed by one Subspace session template.
type MagicMcpServersProvidersEndpoint struct {
	client *endpoint.Client
}

// NewMagicMcpServersProvidersEndpoint creates a new MagicMcpServersProvidersEndpoint.
func NewMagicMcpServersProvidersEndpoint(client *endpoint.Client) *MagicMcpServersProvidersEndpoint {
	return &MagicMcpServersProvidersEndpoint{client: client}
}

// MagicMcpServersProvidersEndpointListParams contains optional query parameters for List.
type MagicMcpServersProvidersEndpointListParams struct {
	Limit                         *float64 `json:"limit,omitempty"`
	After                         *string  `json:"after,omitempty"`
	Before                        *string  `json:"before,omitempty"`
	Cursor                        *string  `json:"cursor,omitempty"`
	Order                         *string  `json:"order,omitempty"`
	Status                        *any     `json:"status,omitempty"`
	Id                            *any     `json:"id,omitempty"`
	ProviderId                    *any     `json:"provider_id,omitempty"`
	IntegrationProviderId         *any     `json:"integration_provider_id,omitempty"`
	IntegrationInstanceProviderId *any     `json:"integration_instance_provider_id,omitempty"`
	ProviderDeploymentId          *any     `json:"provider_deployment_id,omitempty"`
	ProviderConfigId              *any     `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId          *any     `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Filter magic MCP server provider creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter magic MCP server provider last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// MagicMcpServersProvidersEndpointCreateBody contains the request body for Create.
type MagicMcpServersProvidersEndpointCreateBody struct {
	ProviderId           string  `json:"provider_id"`
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	ProviderConfigId     *string `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	ToolFilters          *any    `json:"tool_filters,omitempty"`
}

// MagicMcpServersProvidersEndpointUpdateBody contains the request body for Update.
type MagicMcpServersProvidersEndpointUpdateBody struct {
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	ProviderConfigId     *string `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	ToolFilters          *any    `json:"tool_filters,omitempty"`
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

// Create creates a configurable provider row for a magic MCP server.
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

// Get retrieves a specific backing integration instance provider for a magic MCP server.
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

// Update updates a backing integration provider and then sets the corresponding integration instance provider for a magic MCP server.
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
