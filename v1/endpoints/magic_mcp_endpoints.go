package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/magicmcpendpoints"
)

// MagicMcpEndpointsEndpoint provides access to a magic MCP endpoint is one connection target that routes to multiple magic MCP servers.
type MagicMcpEndpointsEndpoint struct {
	client *endpoint.Client
}

// NewMagicMcpEndpointsEndpoint creates a new MagicMcpEndpointsEndpoint.
func NewMagicMcpEndpointsEndpoint(client *endpoint.Client) *MagicMcpEndpointsEndpoint {
	return &MagicMcpEndpointsEndpoint{client: client}
}

// MagicMcpEndpointsEndpointListParams contains optional query parameters for List.
type MagicMcpEndpointsEndpointListParams struct {
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
	// Status - Only return endpoints with one of these statuses.
	Status *any `json:"status,omitempty"`
	// MagicMcpServerId - Only return endpoints routing to one of these servers.
	MagicMcpServerId *any `json:"magic_mcp_server_id,omitempty"`
	// Search - Search by name or description.
	Search *string `json:"search,omitempty"`
}

// MagicMcpEndpointsEndpointCreateBody contains the request body for Create.
type MagicMcpEndpointsEndpointCreateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the endpoint. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ConsumerProfileId - Consumer profile that will own the endpoint.
	ConsumerProfileId *string `json:"consumer_profile_id,omitempty"`
	// SkillPluginId - Skill plugin the endpoint is created for.
	SkillPluginId *string `json:"skill_plugin_id,omitempty"`
	// EnableSkillEditing - Expose skill editing tools on this endpoint so the connecting consumer can read and edit the skills they have access to. Only connections authenticated with a magic MCP token bound to a single consumer get the tools; connections using a publishable key do not. Requires skills to be enabled for the instance and consumer skill editing to be allowed for the project. The tools disappear if either is turned off later.
	EnableSkillEditing *bool `json:"enable_skill_editing,omitempty"`
	// EnableSkillResources - Serve skills on this endpoint over MCP (the skills extension: skill:// resources, skills/list and skills/get) so the connecting consumer can read the skills they have access to. Only connections authenticated with a magic MCP token bound to a single consumer get the skills. Requires skills to be enabled for the instance.
	EnableSkillResources *bool `json:"enable_skill_resources,omitempty"`
	// SkillSelection - Which skills the endpoint exposes, both as skill resources and through the skill editing tools. Defaults to all skills the consumer has access to. A `selected` selection on a consumer-owned endpoint requires the project to allow skill selection.
	SkillSelection *any `json:"skill_selection,omitempty"`
	// MagicMcpServers - Servers the endpoint routes to.
	MagicMcpServers *[]map[string]any `json:"magic_mcp_servers,omitempty"`
}

// MagicMcpEndpointsEndpointUpdateBody contains the request body for Update.
type MagicMcpEndpointsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the endpoint. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// EnableSkillEditing - Expose skill editing tools on this endpoint so the connecting consumer can read and edit the skills they have access to. Only connections authenticated with a magic MCP token bound to a single consumer get the tools; connections using a publishable key do not. Requires skills to be enabled for the instance and consumer skill editing to be allowed for the project. The tools disappear if either is turned off later.
	EnableSkillEditing *bool `json:"enable_skill_editing,omitempty"`
	// EnableSkillResources - Serve skills on this endpoint over MCP (the skills extension: skill:// resources, skills/list and skills/get) so the connecting consumer can read the skills they have access to. Only connections authenticated with a magic MCP token bound to a single consumer get the skills. Requires skills to be enabled for the instance.
	EnableSkillResources *bool `json:"enable_skill_resources,omitempty"`
	// SkillSelection - Which skills the endpoint exposes, both as skill resources and through the skill editing tools. Defaults to all skills the consumer has access to. A `selected` selection on a consumer-owned endpoint requires the project to allow skill selection.
	SkillSelection *any `json:"skill_selection,omitempty"`
}

// MagicMcpEndpointsEndpointAddServersBody contains the request body for AddServers.
type MagicMcpEndpointsEndpointAddServersBody struct {
	// MagicMcpServers - Servers to add to the endpoint.
	MagicMcpServers *[]map[string]any `json:"magic_mcp_servers,omitempty"`
}

// MagicMcpEndpointsEndpointRemoveServersBody contains the request body for RemoveServers.
type MagicMcpEndpointsEndpointRemoveServersBody struct {
	// MagicMcpServerIds - Servers to remove from the endpoint.
	MagicMcpServerIds []string `json:"magic_mcp_server_ids"`
}

// List returns a paginated list of magic MCP endpoints.
func (e *MagicMcpEndpointsEndpoint) List(params *MagicMcpEndpointsEndpointListParams) (*magicmcpendpoints.MagicMcpEndpointsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"magic-mcp-endpoints"},
		Query: query,
	}
	var result magicmcpendpoints.MagicMcpEndpointsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific magic MCP endpoint.
func (e *MagicMcpEndpointsEndpoint) Get(magicMcpEndpointId string) (*magicmcpendpoints.MagicMcpEndpointsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints", magicMcpEndpointId},
	}
	var result magicmcpendpoints.MagicMcpEndpointsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a magic MCP endpoint.
func (e *MagicMcpEndpointsEndpoint) Create(body *MagicMcpEndpointsEndpointCreateBody) (*magicmcpendpoints.MagicMcpEndpointsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints"},
		Body: body,
	}
	var result magicmcpendpoints.MagicMcpEndpointsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a magic MCP endpoint.
func (e *MagicMcpEndpointsEndpoint) Delete(magicMcpEndpointId string) (*magicmcpendpoints.MagicMcpEndpointsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints", magicMcpEndpointId},
	}
	var result magicmcpendpoints.MagicMcpEndpointsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a magic MCP endpoint. Omitted fields keep their current value.
func (e *MagicMcpEndpointsEndpoint) Update(magicMcpEndpointId string, body *MagicMcpEndpointsEndpointUpdateBody) (*magicmcpendpoints.MagicMcpEndpointsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints", magicMcpEndpointId},
		Body: body,
	}
	var result magicmcpendpoints.MagicMcpEndpointsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AddServers adds magic MCP servers to a magic MCP endpoint.
func (e *MagicMcpEndpointsEndpoint) AddServers(magicMcpEndpointId string, body *MagicMcpEndpointsEndpointAddServersBody) (*magicmcpendpoints.MagicMcpEndpointsAddServersOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints", magicMcpEndpointId, "add-servers"},
		Body: body,
	}
	var result magicmcpendpoints.MagicMcpEndpointsAddServersOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RemoveServers removes magic MCP servers from a magic MCP endpoint.
func (e *MagicMcpEndpointsEndpoint) RemoveServers(magicMcpEndpointId string, body *MagicMcpEndpointsEndpointRemoveServersBody) (*magicmcpendpoints.MagicMcpEndpointsRemoveServersOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints", magicMcpEndpointId, "remove-servers"},
		Body: body,
	}
	var result magicmcpendpoints.MagicMcpEndpointsRemoveServersOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
