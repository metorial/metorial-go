package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/consumergroups"
)

// PortalsConsumerGroupsEndpoint provides access to a portal consumer group is a set of consumers on a portal. Groups decide which listings each consumer sees.
type PortalsConsumerGroupsEndpoint struct {
	client *endpoint.Client
}

// NewPortalsConsumerGroupsEndpoint creates a new PortalsConsumerGroupsEndpoint.
func NewPortalsConsumerGroupsEndpoint(client *endpoint.Client) *PortalsConsumerGroupsEndpoint {
	return &PortalsConsumerGroupsEndpoint{client: client}
}

// PortalsConsumerGroupsEndpointListParams contains optional query parameters for List.
type PortalsConsumerGroupsEndpointListParams struct {
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
}

// PortalsConsumerGroupsEndpointCreateBody contains the request body for Create.
type PortalsConsumerGroupsEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDefault - Whether this is a default group on the portal.
	IsDefault *bool `json:"is_default,omitempty"`
}

// PortalsConsumerGroupsEndpointUpdateBody contains the request body for Update.
type PortalsConsumerGroupsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDefault - Whether this is a default group on the portal.
	IsDefault *bool `json:"is_default,omitempty"`
}

// List returns a paginated list of consumer groups for a portal.
func (e *PortalsConsumerGroupsEndpoint) List(portalId string, params *PortalsConsumerGroupsEndpointListParams) (*consumergroups.PortalsConsumerGroupsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"portals", portalId, "consumer-groups"},
		Query: query,
	}
	var result consumergroups.PortalsConsumerGroupsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a portal consumer group.
func (e *PortalsConsumerGroupsEndpoint) Get(portalId string, consumerGroupId string) (*consumergroups.PortalsConsumerGroupsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId, "consumer-groups", consumerGroupId},
	}
	var result consumergroups.PortalsConsumerGroupsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a consumer group for a portal.
func (e *PortalsConsumerGroupsEndpoint) Create(portalId string, body *PortalsConsumerGroupsEndpointCreateBody) (*consumergroups.PortalsConsumerGroupsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId, "consumer-groups"},
		Body: body,
	}
	var result consumergroups.PortalsConsumerGroupsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a portal consumer group. Omitted fields keep their current value.
func (e *PortalsConsumerGroupsEndpoint) Update(portalId string, consumerGroupId string, body *PortalsConsumerGroupsEndpointUpdateBody) (*consumergroups.PortalsConsumerGroupsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId, "consumer-groups", consumerGroupId},
		Body: body,
	}
	var result consumergroups.PortalsConsumerGroupsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a portal consumer group.
func (e *PortalsConsumerGroupsEndpoint) Delete(portalId string, consumerGroupId string) (*consumergroups.PortalsConsumerGroupsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId, "consumer-groups", consumerGroupId},
	}
	var result consumergroups.PortalsConsumerGroupsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
