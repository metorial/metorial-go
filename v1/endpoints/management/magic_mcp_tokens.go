package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/magicmcptokens"
)

// MagicMcpTokensEndpoint provides access to a magic MCP token is a credential that authorizes access to magic MCP servers and endpoints.
type MagicMcpTokensEndpoint struct {
	client *endpoint.Client
}

// NewMagicMcpTokensEndpoint creates a new MagicMcpTokensEndpoint.
func NewMagicMcpTokensEndpoint(client *endpoint.Client) *MagicMcpTokensEndpoint {
	return &MagicMcpTokensEndpoint{client: client}
}

// MagicMcpTokensEndpointListParams contains optional query parameters for List.
type MagicMcpTokensEndpointListParams struct {
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
	// Status - Only return tokens with one of these statuses.
	Status *any `json:"status,omitempty"`
	// MagicMcpGroupId - Only return tokens scoped to one of these groups.
	MagicMcpGroupId *any `json:"magic_mcp_group_id,omitempty"`
	// MagicMcpServerId - Only return tokens bound to one of these servers.
	MagicMcpServerId *any `json:"magic_mcp_server_id,omitempty"`
	// MagicMcpEndpointId - Only return tokens bound to one of these endpoints.
	MagicMcpEndpointId *any `json:"magic_mcp_endpoint_id,omitempty"`
}

// MagicMcpTokensEndpointCreateBody contains the request body for Create.
type MagicMcpTokensEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the token. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// MagicMcpGroupIds - Groups to scope the token to.
	MagicMcpGroupIds *[]string `json:"magic_mcp_group_ids,omitempty"`
	// MagicMcpServerId - Server to bind the token to.
	MagicMcpServerId *string `json:"magic_mcp_server_id,omitempty"`
	// MagicMcpEndpointId - Endpoint to bind the token to.
	MagicMcpEndpointId *string `json:"magic_mcp_endpoint_id,omitempty"`
}

// MagicMcpTokensEndpointUpdateBody contains the request body for Update.
type MagicMcpTokensEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the token. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MagicMcpTokensEndpointAddGroupsBody contains the request body for AddGroups.
type MagicMcpTokensEndpointAddGroupsBody struct {
	// MagicMcpGroupIds - Groups to scope the token to.
	MagicMcpGroupIds []string `json:"magic_mcp_group_ids"`
}

// MagicMcpTokensEndpointRemoveGroupsBody contains the request body for RemoveGroups.
type MagicMcpTokensEndpointRemoveGroupsBody struct {
	// MagicMcpGroupIds - Groups to remove from the token.
	MagicMcpGroupIds []string `json:"magic_mcp_group_ids"`
}

// List returns a paginated list of magic MCP tokens.
func (e *MagicMcpTokensEndpoint) List(instanceId string, params *MagicMcpTokensEndpointListParams) (*magicmcptokens.MagicMcpTokensListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "magic-mcp-tokens"},
		Query: query,
	}
	var result magicmcptokens.MagicMcpTokensListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific magic MCP token.
func (e *MagicMcpTokensEndpoint) Get(instanceId string, magicMcpTokenId string) (*magicmcptokens.MagicMcpTokensGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-tokens", magicMcpTokenId},
	}
	var result magicmcptokens.MagicMcpTokensGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a magic MCP token.
func (e *MagicMcpTokensEndpoint) Create(instanceId string, body *MagicMcpTokensEndpointCreateBody) (*magicmcptokens.MagicMcpTokensCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-tokens"},
		Body: body,
	}
	var result magicmcptokens.MagicMcpTokensCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a magic MCP token.
func (e *MagicMcpTokensEndpoint) Delete(instanceId string, magicMcpTokenId string) (*magicmcptokens.MagicMcpTokensDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-tokens", magicMcpTokenId},
	}
	var result magicmcptokens.MagicMcpTokensDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a magic MCP token. Omitted fields keep their current value.
func (e *MagicMcpTokensEndpoint) Update(instanceId string, magicMcpTokenId string, body *MagicMcpTokensEndpointUpdateBody) (*magicmcptokens.MagicMcpTokensUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-tokens", magicMcpTokenId},
		Body: body,
	}
	var result magicmcptokens.MagicMcpTokensUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AddGroups adds groups to a magic MCP token.
func (e *MagicMcpTokensEndpoint) AddGroups(instanceId string, magicMcpTokenId string, body *MagicMcpTokensEndpointAddGroupsBody) (*magicmcptokens.MagicMcpTokensAddGroupsOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-tokens", magicMcpTokenId, "add-groups"},
		Body: body,
	}
	var result magicmcptokens.MagicMcpTokensAddGroupsOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RemoveGroups removes groups from a magic MCP token.
func (e *MagicMcpTokensEndpoint) RemoveGroups(instanceId string, magicMcpTokenId string, body *MagicMcpTokensEndpointRemoveGroupsBody) (*magicmcptokens.MagicMcpTokensRemoveGroupsOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-tokens", magicMcpTokenId, "remove-groups"},
		Body: body,
	}
	var result magicmcptokens.MagicMcpTokensRemoveGroupsOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
