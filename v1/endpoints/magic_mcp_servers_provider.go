package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/magicmcpservers/provider"
)

// MagicMcpServersProviderEndpoint provides access to magic MCP server providers define which providers are included in the setup session template backing a magic MCP server.
type MagicMcpServersProviderEndpoint struct {
	client *endpoint.Client
}

// NewMagicMcpServersProviderEndpoint creates a new MagicMcpServersProviderEndpoint.
func NewMagicMcpServersProviderEndpoint(client *endpoint.Client) *MagicMcpServersProviderEndpoint {
	return &MagicMcpServersProviderEndpoint{client: client}
}

// MagicMcpServersProviderEndpointListParams contains optional query parameters for List.
type MagicMcpServersProviderEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Status *any     `json:"status,omitempty"`
	// Id - Filter by magic MCP server provider ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter by provider config ID(s)
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Filter by provider auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Filter magic MCP server provider creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter magic MCP server provider last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// MagicMcpServersProviderEndpointCreateBody contains the request body for Create.
type MagicMcpServersProviderEndpointCreateBody struct {
	ProviderDeploymentId  *string `json:"provider_deployment_id,omitempty"`
	ProviderConfigId      *string `json:"provider_config_id,omitempty"`
	ProviderConfigVaultId *string `json:"provider_config_vault_id,omitempty"`
	ProviderAuthConfigId  *string `json:"provider_auth_config_id,omitempty"`
	ToolFilters           *any    `json:"tool_filters,omitempty"`
}

// MagicMcpServersProviderEndpointUpdateBody contains the request body for Update.
type MagicMcpServersProviderEndpointUpdateBody struct {
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// List returns a paginated list of providers configured for a magic MCP server.
func (e *MagicMcpServersProviderEndpoint) List(magicMcpServerId string, params *MagicMcpServersProviderEndpointListParams) (*provider.MagicMcpServersProviderListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"magic-mcp-servers", magicMcpServerId, "provider"},
		Query: query,
	}
	var result provider.MagicMcpServersProviderListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider configuration from a magic MCP server.
func (e *MagicMcpServersProviderEndpoint) Get(magicMcpServerId string, magicMcpServerProviderId string) (*provider.MagicMcpServersProviderGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-servers", magicMcpServerId, "provider", magicMcpServerProviderId},
	}
	var result provider.MagicMcpServersProviderGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create adds a new provider configuration to a magic MCP server.
func (e *MagicMcpServersProviderEndpoint) Create(magicMcpServerId string, body *MagicMcpServersProviderEndpointCreateBody) (*provider.MagicMcpServersProviderCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-servers", magicMcpServerId, "provider"},
		Body: body,
	}
	var result provider.MagicMcpServersProviderCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a provider configuration in a magic MCP server.
func (e *MagicMcpServersProviderEndpoint) Update(magicMcpServerId string, magicMcpServerProviderId string, body *MagicMcpServersProviderEndpointUpdateBody) (*provider.MagicMcpServersProviderUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-servers", magicMcpServerId, "provider", magicMcpServerProviderId},
		Body: body,
	}
	var result provider.MagicMcpServersProviderUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a provider configuration from a magic MCP server.
func (e *MagicMcpServersProviderEndpoint) Delete(magicMcpServerId string, magicMcpServerProviderId string) (*provider.MagicMcpServersProviderDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-servers", magicMcpServerId, "provider", magicMcpServerProviderId},
	}
	var result provider.MagicMcpServersProviderDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
