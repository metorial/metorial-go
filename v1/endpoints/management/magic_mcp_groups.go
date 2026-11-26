package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/magicmcpgroups"
)

// MagicMcpGroupsEndpoint provides access to a magic MCP group is a set of magic MCP servers that a token can be scoped to.
type MagicMcpGroupsEndpoint struct {
	client *endpoint.Client
}

// NewMagicMcpGroupsEndpoint creates a new MagicMcpGroupsEndpoint.
func NewMagicMcpGroupsEndpoint(client *endpoint.Client) *MagicMcpGroupsEndpoint {
	return &MagicMcpGroupsEndpoint{client: client}
}

// MagicMcpGroupsEndpointListParams contains optional query parameters for List.
type MagicMcpGroupsEndpointListParams struct {
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
	// Status - Only return groups with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Search - Search by name or description.
	Search *string `json:"search,omitempty"`
	// Id - Only return groups with these IDs.
	Id *any `json:"id,omitempty"`
}

// MagicMcpGroupsEndpointCreateBody contains the request body for Create.
type MagicMcpGroupsEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the group. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MagicMcpGroupsEndpointUpdateBody contains the request body for Update.
type MagicMcpGroupsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the group. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MagicMcpGroupsEndpointAddServersBody contains the request body for AddServers.
type MagicMcpGroupsEndpointAddServersBody struct {
	// MagicMcpServerIds - Servers to add to the group.
	MagicMcpServerIds []string `json:"magic_mcp_server_ids"`
}

// MagicMcpGroupsEndpointRemoveServersBody contains the request body for RemoveServers.
type MagicMcpGroupsEndpointRemoveServersBody struct {
	// MagicMcpServerIds - Servers to remove from the group.
	MagicMcpServerIds []string `json:"magic_mcp_server_ids"`
}

// List returns a paginated list of magic MCP groups.
func (e *MagicMcpGroupsEndpoint) List(instanceId string, params *MagicMcpGroupsEndpointListParams) (*magicmcpgroups.MagicMcpGroupsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "magic-mcp-groups"},
		Query: query,
	}
	var result magicmcpgroups.MagicMcpGroupsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific magic MCP group.
func (e *MagicMcpGroupsEndpoint) Get(instanceId string, magicMcpGroupId string) (*magicmcpgroups.MagicMcpGroupsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-groups", magicMcpGroupId},
	}
	var result magicmcpgroups.MagicMcpGroupsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a magic MCP group.
func (e *MagicMcpGroupsEndpoint) Create(instanceId string, body *MagicMcpGroupsEndpointCreateBody) (*magicmcpgroups.MagicMcpGroupsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-groups"},
		Body: body,
	}
	var result magicmcpgroups.MagicMcpGroupsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a magic MCP group. Fails while the group is the only group of an active magic MCP token, so that deleting a group never widens what a token can access.
func (e *MagicMcpGroupsEndpoint) Delete(instanceId string, magicMcpGroupId string) (*magicmcpgroups.MagicMcpGroupsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-groups", magicMcpGroupId},
	}
	var result magicmcpgroups.MagicMcpGroupsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a magic MCP group. Omitted fields keep their current value.
func (e *MagicMcpGroupsEndpoint) Update(instanceId string, magicMcpGroupId string, body *MagicMcpGroupsEndpointUpdateBody) (*magicmcpgroups.MagicMcpGroupsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-groups", magicMcpGroupId},
		Body: body,
	}
	var result magicmcpgroups.MagicMcpGroupsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AddServers adds magic MCP servers to a group.
func (e *MagicMcpGroupsEndpoint) AddServers(instanceId string, magicMcpGroupId string, body *MagicMcpGroupsEndpointAddServersBody) (*magicmcpgroups.MagicMcpGroupsAddServersOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-groups", magicMcpGroupId, "add-servers"},
		Body: body,
	}
	var result magicmcpgroups.MagicMcpGroupsAddServersOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RemoveServers removes magic MCP servers from a group.
func (e *MagicMcpGroupsEndpoint) RemoveServers(instanceId string, magicMcpGroupId string, body *MagicMcpGroupsEndpointRemoveServersBody) (*magicmcpgroups.MagicMcpGroupsRemoveServersOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "magic-mcp-groups", magicMcpGroupId, "remove-servers"},
		Body: body,
	}
	var result magicmcpgroups.MagicMcpGroupsRemoveServersOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
