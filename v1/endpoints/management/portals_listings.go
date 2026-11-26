package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/listings"
)

// PortalsListingsEndpoint provides access to a portal listing is the shared catalog entry for one integration, skill, or skill set on a portal. Access rules decide which consumer groups see it.
type PortalsListingsEndpoint struct {
	client *endpoint.Client
}

// NewPortalsListingsEndpoint creates a new PortalsListingsEndpoint.
func NewPortalsListingsEndpoint(client *endpoint.Client) *PortalsListingsEndpoint {
	return &PortalsListingsEndpoint{client: client}
}

// PortalsListingsEndpointListParams contains optional query parameters for List.
type PortalsListingsEndpointListParams struct {
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
	// ConsumerSurfaceProviderGroupId - Only return listings in these provider groups.
	ConsumerSurfaceProviderGroupId *any `json:"consumer_surface_provider_group_id,omitempty"`
	// ProviderTemplateId - Only return listings for these provider templates.
	ProviderTemplateId *any `json:"provider_template_id,omitempty"`
	// MagicMcpServerId - Only return listings for these magic MCP servers.
	MagicMcpServerId *any `json:"magic_mcp_server_id,omitempty"`
	// SkillId - Only return listings for these skills.
	SkillId *any `json:"skill_id,omitempty"`
	// SkillTemplateId - Only return listings for these skill templates.
	SkillTemplateId *any `json:"skill_template_id,omitempty"`
	// SkillGroupId - Only return listings for these skill groups.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
	// SkillMarketplaceId - Only return listings for these skill marketplaces.
	SkillMarketplaceId *any `json:"skill_marketplace_id,omitempty"`
	// Type - Only return listings of these types.
	Type *any `json:"type,omitempty"`
	// HasAccess - Only return listings that have at least one consumer access grant (true) or none (false).
	HasAccess *bool `json:"has_access,omitempty"`
}

// PortalsListingsEndpointCreateBody contains the request body for Create.
type PortalsListingsEndpointCreateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Readme - Longer guide shown to consumers.
	Readme *string `json:"readme,omitempty"`
	// IsDiscoverable - Whether every consumer on the portal can find this listing in the catalog and request access. When false, only consumers who already have access through one of their groups see it. Defaults to true.
	IsDiscoverable *bool `json:"is_discoverable,omitempty"`
	// Access - The integration, skill, or skill set the listing is for.
	Access any `json:"access"`
}

// PortalsListingsEndpointUpdateBody contains the request body for Update.
type PortalsListingsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Readme - Longer guide shown to consumers.
	Readme *string `json:"readme,omitempty"`
	// IsDiscoverable - Whether every consumer on the portal can find this listing in the catalog and request access. When false, only consumers who already have access through one of their groups see it.
	IsDiscoverable *bool `json:"is_discoverable,omitempty"`
}

// List returns a paginated list of shared listings for a portal.
func (e *PortalsListingsEndpoint) List(instanceId string, portalId string, params *PortalsListingsEndpointListParams) (*listings.PortalsListingsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "portals", portalId, "listings"},
		Query: query,
	}
	var result listings.PortalsListingsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a shared listing for a portal.
func (e *PortalsListingsEndpoint) Get(instanceId string, portalId string, listingId string) (*listings.PortalsListingsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "listings", listingId},
	}
	var result listings.PortalsListingsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates or updates the shared listing for one target on a portal. A portal has at most one listing per target.
func (e *PortalsListingsEndpoint) Create(instanceId string, portalId string, body *PortalsListingsEndpointCreateBody) (*listings.PortalsListingsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "listings"},
		Body: body,
	}
	var result listings.PortalsListingsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a portal listing. Omitted fields keep their current value.
func (e *PortalsListingsEndpoint) Update(instanceId string, portalId string, listingId string, body *PortalsListingsEndpointUpdateBody) (*listings.PortalsListingsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "listings", listingId},
		Body: body,
	}
	var result listings.PortalsListingsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a portal listing and all consumer access attached to it.
func (e *PortalsListingsEndpoint) Delete(instanceId string, portalId string, listingId string) (*listings.PortalsListingsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "listings", listingId},
	}
	var result listings.PortalsListingsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
