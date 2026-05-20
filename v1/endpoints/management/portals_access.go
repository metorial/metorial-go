package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/access"
)

// PortalsAccessEndpoint provides access to manage which consumer groups can access portal provider templates and MCP servers.
type PortalsAccessEndpoint struct {
	client *endpoint.Client
}

// NewPortalsAccessEndpoint creates a new PortalsAccessEndpoint.
func NewPortalsAccessEndpoint(client *endpoint.Client) *PortalsAccessEndpoint {
	return &PortalsAccessEndpoint{client: client}
}

// PortalsAccessEndpointListParams contains optional query parameters for List.
type PortalsAccessEndpointListParams struct {
	Limit                   *float64 `json:"limit,omitempty"`
	After                   *string  `json:"after,omitempty"`
	Before                  *string  `json:"before,omitempty"`
	Cursor                  *string  `json:"cursor,omitempty"`
	Order                   *string  `json:"order,omitempty"`
	Search                  *string  `json:"search,omitempty"`
	ConsumerGroupId         *any     `json:"consumer_group_id,omitempty"`
	ProviderTemplateId      *any     `json:"provider_template_id,omitempty"`
	MagicMcpServerId        *any     `json:"magic_mcp_server_id,omitempty"`
	SkillId                 *any     `json:"skill_id,omitempty"`
	SkillTemplateId         *any     `json:"skill_template_id,omitempty"`
	SkillGroupId            *any     `json:"skill_group_id,omitempty"`
	SkillMarketplaceId      *any     `json:"skill_marketplace_id,omitempty"`
	ConsumerAccessListingId *any     `json:"consumer_access_listing_id,omitempty"`
	Type                    *any     `json:"type,omitempty"`
}

// PortalsAccessEndpointCreateBody contains the request body for Create.
type PortalsAccessEndpointCreateBody struct {
	ConsumerGroupId string  `json:"consumer_group_id"`
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	Readme          *string `json:"readme,omitempty"`
	Access          any     `json:"access"`
}

// PortalsAccessEndpointUpdateBody contains the request body for Update.
type PortalsAccessEndpointUpdateBody struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Readme      *string `json:"readme,omitempty"`
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

// Get retrieves a portal access rule by ID.
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

// Create creates a new consumer access rule for the portal.
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

// Update updates the shared listing fields for a portal access rule.
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

// Delete deletes a consumer access rule from the portal.
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
