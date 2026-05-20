package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/listings"
)

// PortalsListingsEndpoint provides access to read the shared listings available on a portal surface.
type PortalsListingsEndpoint struct {
	client *endpoint.Client
}

// NewPortalsListingsEndpoint creates a new PortalsListingsEndpoint.
func NewPortalsListingsEndpoint(client *endpoint.Client) *PortalsListingsEndpoint {
	return &PortalsListingsEndpoint{client: client}
}

// PortalsListingsEndpointListParams contains optional query parameters for List.
type PortalsListingsEndpointListParams struct {
	Limit                          *float64 `json:"limit,omitempty"`
	After                          *string  `json:"after,omitempty"`
	Before                         *string  `json:"before,omitempty"`
	Cursor                         *string  `json:"cursor,omitempty"`
	Order                          *string  `json:"order,omitempty"`
	Search                         *string  `json:"search,omitempty"`
	ConsumerSurfaceProviderGroupId *any     `json:"consumer_surface_provider_group_id,omitempty"`
	ProviderTemplateId             *any     `json:"provider_template_id,omitempty"`
	MagicMcpServerId               *any     `json:"magic_mcp_server_id,omitempty"`
	SkillId                        *any     `json:"skill_id,omitempty"`
	SkillTemplateId                *any     `json:"skill_template_id,omitempty"`
	SkillGroupId                   *any     `json:"skill_group_id,omitempty"`
	SkillMarketplaceId             *any     `json:"skill_marketplace_id,omitempty"`
	Type                           *any     `json:"type,omitempty"`
}

// PortalsListingsEndpointCreateBody contains the request body for Create.
type PortalsListingsEndpointCreateBody struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Readme      *string `json:"readme,omitempty"`
	Access      any     `json:"access"`
}

// PortalsListingsEndpointUpdateBody contains the request body for Update.
type PortalsListingsEndpointUpdateBody struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Readme      *string `json:"readme,omitempty"`
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

// Get retrieves one shared listing for a portal.
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

// Create creates a shared listing for a portal.
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

// Update updates listing metadata for a portal listing.
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
