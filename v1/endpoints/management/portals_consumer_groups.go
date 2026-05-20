package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/consumergroups"
)

// PortalsConsumerGroupsEndpoint provides access to manage the consumer groups that drive portal visibility and access rules.
type PortalsConsumerGroupsEndpoint struct {
	client *endpoint.Client
}

// NewPortalsConsumerGroupsEndpoint creates a new PortalsConsumerGroupsEndpoint.
func NewPortalsConsumerGroupsEndpoint(client *endpoint.Client) *PortalsConsumerGroupsEndpoint {
	return &PortalsConsumerGroupsEndpoint{client: client}
}

// PortalsConsumerGroupsEndpointListParams contains optional query parameters for List.
type PortalsConsumerGroupsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Status *any     `json:"status,omitempty"`
	Search *string  `json:"search,omitempty"`
}

// PortalsConsumerGroupsEndpointCreateBody contains the request body for Create.
type PortalsConsumerGroupsEndpointCreateBody struct {
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	SsoGroupIds *[]string `json:"sso_group_ids,omitempty"`
	IsDefault   *bool     `json:"is_default,omitempty"`
}

// PortalsConsumerGroupsEndpointUpdateBody contains the request body for Update.
type PortalsConsumerGroupsEndpointUpdateBody struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	SsoGroupIds *[]string `json:"sso_group_ids,omitempty"`
	IsDefault   *bool     `json:"is_default,omitempty"`
}

// List returns a paginated list of consumer groups for a portal.
func (e *PortalsConsumerGroupsEndpoint) List(instanceId string, portalId string, params *PortalsConsumerGroupsEndpointListParams) (*consumergroups.PortalsConsumerGroupsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "portals", portalId, "consumer-groups"},
		Query: query,
	}
	var result consumergroups.PortalsConsumerGroupsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a portal consumer group by ID.
func (e *PortalsConsumerGroupsEndpoint) Get(instanceId string, portalId string, consumerGroupId string) (*consumergroups.PortalsConsumerGroupsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "consumer-groups", consumerGroupId},
	}
	var result consumergroups.PortalsConsumerGroupsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new consumer group for the portal.
func (e *PortalsConsumerGroupsEndpoint) Create(instanceId string, portalId string, body *PortalsConsumerGroupsEndpointCreateBody) (*consumergroups.PortalsConsumerGroupsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "consumer-groups"},
		Body: body,
	}
	var result consumergroups.PortalsConsumerGroupsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a consumer group for the portal.
func (e *PortalsConsumerGroupsEndpoint) Update(instanceId string, portalId string, consumerGroupId string, body *PortalsConsumerGroupsEndpointUpdateBody) (*consumergroups.PortalsConsumerGroupsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "consumer-groups", consumerGroupId},
		Body: body,
	}
	var result consumergroups.PortalsConsumerGroupsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a consumer group for the portal.
func (e *PortalsConsumerGroupsEndpoint) Delete(instanceId string, portalId string, consumerGroupId string) (*consumergroups.PortalsConsumerGroupsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "consumer-groups", consumerGroupId},
	}
	var result consumergroups.PortalsConsumerGroupsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
