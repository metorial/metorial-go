package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/access"
)

// PortalsAccessEndpoint provides access to a portal access rule grants one consumer group access to one listing target, such as an integration, skill, or skill set.
type PortalsAccessEndpoint struct {
	client *endpoint.Client
}

// NewPortalsAccessEndpoint creates a new PortalsAccessEndpoint.
func NewPortalsAccessEndpoint(client *endpoint.Client) *PortalsAccessEndpoint {
	return &PortalsAccessEndpoint{client: client}
}

// PortalsAccessEndpointListParams contains optional query parameters for List.
type PortalsAccessEndpointListParams struct {
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
	// Search - Search by name or description.
	Search *string `json:"search,omitempty"`
	// ConsumerGroupId - Only return rules for these consumer groups.
	ConsumerGroupId *any `json:"consumer_group_id,omitempty"`
	// ProviderTemplateId - Only return rules for these provider templates.
	ProviderTemplateId *any `json:"provider_template_id,omitempty"`
	// MagicMcpServerId - Only return rules for these magic MCP servers.
	MagicMcpServerId *any `json:"magic_mcp_server_id,omitempty"`
	// SkillId - Only return rules for these skills.
	SkillId *any `json:"skill_id,omitempty"`
	// SkillTemplateId - Only return rules for these skill templates.
	SkillTemplateId *any `json:"skill_template_id,omitempty"`
	// SkillGroupId - Only return rules for these skill groups.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
	// SkillMarketplaceId - Only return rules for these skill marketplaces.
	SkillMarketplaceId *any `json:"skill_marketplace_id,omitempty"`
	// SkillPluginId - Only return rules for these skill plugins.
	SkillPluginId *any `json:"skill_plugin_id,omitempty"`
	// ConsumerAccessListingId - Only return rules attached to these listings.
	ConsumerAccessListingId *any `json:"consumer_access_listing_id,omitempty"`
	// Type - Only return rules of these types.
	Type *any `json:"type,omitempty"`
}

// PortalsAccessEndpointCreateBody contains the request body for Create.
type PortalsAccessEndpointCreateBody struct {
	// ConsumerGroupId - Consumer group the rule grants access to.
	ConsumerGroupId string `json:"consumer_group_id"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Readme - Longer guide shown to consumers.
	Readme *string `json:"readme,omitempty"`
	// Access - The integration, skill, or skill set the rule grants access to.
	Access any `json:"access"`
}

// PortalsAccessEndpointUpdateBody contains the request body for Update.
type PortalsAccessEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Readme - Longer guide shown to consumers.
	Readme *string `json:"readme,omitempty"`
	// Permission - What the group can do with the skill or skill group.
	Permission *string `json:"permission,omitempty"`
}

// List returns a paginated list of consumer access rules for a portal.
func (e *PortalsAccessEndpoint) List(instanceId string, portalId string, params *PortalsAccessEndpointListParams) (*access.PortalsAccessListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "portals", portalId, "access"},
		Query: query,
	}
	var result access.PortalsAccessListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a portal access rule.
func (e *PortalsAccessEndpoint) Get(instanceId string, portalId string, accessId string) (*access.PortalsAccessGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "access", accessId},
	}
	var result access.PortalsAccessGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a consumer access rule for a portal.
func (e *PortalsAccessEndpoint) Create(instanceId string, portalId string, body *PortalsAccessEndpointCreateBody) (*access.PortalsAccessCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "access"},
		Body: body,
	}
	var result access.PortalsAccessCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a portal access rule. Only skill and skill group rules use `permission`. Omitted fields keep their current value.
func (e *PortalsAccessEndpoint) Update(instanceId string, portalId string, accessId string, body *PortalsAccessEndpointUpdateBody) (*access.PortalsAccessUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "access", accessId},
		Body: body,
	}
	var result access.PortalsAccessUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a portal access rule.
func (e *PortalsAccessEndpoint) Delete(instanceId string, portalId string, accessId string) (*access.PortalsAccessDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "access", accessId},
	}
	var result access.PortalsAccessDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
