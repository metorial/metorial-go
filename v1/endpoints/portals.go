package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals"
)

// PortalsEndpoint provides access to a portal is a branded catalog of integrations and skills for your consumers. Each portal has its own consumers, groups, and access rules.
type PortalsEndpoint struct {
	client *endpoint.Client
}

// NewPortalsEndpoint creates a new PortalsEndpoint.
func NewPortalsEndpoint(client *endpoint.Client) *PortalsEndpoint {
	return &PortalsEndpoint{client: client}
}

// PortalsEndpointListParams contains optional query parameters for List.
type PortalsEndpointListParams struct {
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
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
}

// PortalsEndpointCreateBody contains the request body for Create.
type PortalsEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// AllowedRedirectUrlFilters - Redirect URL patterns consumers may return to after signing in.
	AllowedRedirectUrlFilters *[]map[string]any `json:"allowed_redirect_url_filters,omitempty"`
	// SessionExpiryTimeInSeconds - How long consumer sessions stay valid, in seconds.
	SessionExpiryTimeInSeconds *float64 `json:"session_expiry_time_in_seconds,omitempty"`
	// AllowConsumerSkillAuthoring - Whether consumers can create skills on the portal.
	AllowConsumerSkillAuthoring *bool `json:"allow_consumer_skill_authoring,omitempty"`
	// AllowConsumerSkillPublishing - Whether consumers can publish skills on the portal.
	AllowConsumerSkillPublishing *bool `json:"allow_consumer_skill_publishing,omitempty"`
}

// PortalsEndpointUpdateBody contains the request body for Update.
type PortalsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// AllowedRedirectUrlFilters - Redirect URL patterns consumers may return to after signing in.
	AllowedRedirectUrlFilters *[]map[string]any `json:"allowed_redirect_url_filters,omitempty"`
	// SessionExpiryTimeInSeconds - How long consumer sessions stay valid, in seconds.
	SessionExpiryTimeInSeconds *float64 `json:"session_expiry_time_in_seconds,omitempty"`
	// AllowConsumerSkillAuthoring - Whether consumers can create skills on the portal.
	AllowConsumerSkillAuthoring *bool `json:"allow_consumer_skill_authoring,omitempty"`
	// AllowConsumerSkillPublishing - Whether consumers can publish skills on the portal.
	AllowConsumerSkillPublishing *bool `json:"allow_consumer_skill_publishing,omitempty"`
	// SkillConfiguration - Skill settings for the portal.
	SkillConfiguration *map[string]any `json:"skill_configuration,omitempty"`
}

// List returns a paginated list of portals.
func (e *PortalsEndpoint) List(params *PortalsEndpointListParams) (*portals.PortalsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"portals"},
		Query: query,
	}
	var result portals.PortalsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a portal.
func (e *PortalsEndpoint) Get(portalId string) (*portals.PortalsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId},
	}
	var result portals.PortalsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a portal.
func (e *PortalsEndpoint) Create(body *PortalsEndpointCreateBody) (*portals.PortalsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals"},
		Body: body,
	}
	var result portals.PortalsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a portal. Omitted fields keep their current value.
func (e *PortalsEndpoint) Update(portalId string, body *PortalsEndpointUpdateBody) (*portals.PortalsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId},
		Body: body,
	}
	var result portals.PortalsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a portal.
func (e *PortalsEndpoint) Delete(portalId string) (*portals.PortalsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId},
	}
	var result portals.PortalsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
