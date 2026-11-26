package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/consumerprofiles"
)

// PortalsConsumerProfilesEndpoint provides access to a portal consumer profile is the identity of one person on a portal. Group assignments decide what the consumer can access.
type PortalsConsumerProfilesEndpoint struct {
	client *endpoint.Client
}

// NewPortalsConsumerProfilesEndpoint creates a new PortalsConsumerProfilesEndpoint.
func NewPortalsConsumerProfilesEndpoint(client *endpoint.Client) *PortalsConsumerProfilesEndpoint {
	return &PortalsConsumerProfilesEndpoint{client: client}
}

// PortalsConsumerProfilesEndpointListParams contains optional query parameters for List.
type PortalsConsumerProfilesEndpointListParams struct {
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
	// Search - Search by name or email.
	Search *string `json:"search,omitempty"`
	// Email - Only return profiles with one of these emails.
	Email *any `json:"email,omitempty"`
	// ConsumerGroupId - Only return profiles in this consumer group.
	ConsumerGroupId *string `json:"consumer_group_id,omitempty"`
	// Source - Only return profiles created manually or through SSO.
	Source *string `json:"source,omitempty"`
	// Status - Only return profiles with one of these statuses.
	Status *any `json:"status,omitempty"`
}

// PortalsConsumerProfilesEndpointCreateBody contains the request body for Create.
type PortalsConsumerProfilesEndpointCreateBody struct {
	// Email - Email address of the consumer.
	Email string `json:"email"`
	// Name - Display name.
	Name string `json:"name"`
}

// PortalsConsumerProfilesEndpointAssignGroupsBody contains the request body for AssignGroups.
type PortalsConsumerProfilesEndpointAssignGroupsBody struct {
	// GroupIds - Groups to assign the profile to.
	GroupIds []string `json:"group_ids"`
}

// PortalsConsumerProfilesEndpointUnassignGroupsBody contains the request body for UnassignGroups.
type PortalsConsumerProfilesEndpointUnassignGroupsBody struct {
	// GroupIds - Groups to remove the profile from.
	GroupIds []string `json:"group_ids"`
}

// List returns a paginated list of consumer profiles for a portal.
func (e *PortalsConsumerProfilesEndpoint) List(portalId string, params *PortalsConsumerProfilesEndpointListParams) (*consumerprofiles.PortalsConsumerProfilesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"portals", portalId, "consumer-profile"},
		Query: query,
	}
	var result consumerprofiles.PortalsConsumerProfilesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a portal consumer profile.
func (e *PortalsConsumerProfilesEndpoint) Get(portalId string, consumerProfileId string) (*consumerprofiles.PortalsConsumerProfilesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId, "consumer-profile", consumerProfileId},
	}
	var result consumerprofiles.PortalsConsumerProfilesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a portal consumer profile.
func (e *PortalsConsumerProfilesEndpoint) Create(portalId string, body *PortalsConsumerProfilesEndpointCreateBody) (*consumerprofiles.PortalsConsumerProfilesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId, "consumer-profile"},
		Body: body,
	}
	var result consumerprofiles.PortalsConsumerProfilesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete soft-deletes a portal consumer profile.
func (e *PortalsConsumerProfilesEndpoint) Delete(portalId string, consumerProfileId string) (*consumerprofiles.PortalsConsumerProfilesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId, "consumer-profile", consumerProfileId},
	}
	var result consumerprofiles.PortalsConsumerProfilesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AssignGroups assigns one or more groups to a portal consumer profile.
func (e *PortalsConsumerProfilesEndpoint) AssignGroups(portalId string, consumerProfileId string, body *PortalsConsumerProfilesEndpointAssignGroupsBody) (*consumerprofiles.PortalsConsumerProfilesAssignGroupsOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId, "consumer-profile", consumerProfileId, "assign-groups"},
		Body: body,
	}
	var result consumerprofiles.PortalsConsumerProfilesAssignGroupsOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UnassignGroups removes one or more groups from a portal consumer profile.
func (e *PortalsConsumerProfilesEndpoint) UnassignGroups(portalId string, consumerProfileId string, body *PortalsConsumerProfilesEndpointUnassignGroupsBody) (*consumerprofiles.PortalsConsumerProfilesUnassignGroupsOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId, "consumer-profile", consumerProfileId, "unassign-groups"},
		Body: body,
	}
	var result consumerprofiles.PortalsConsumerProfilesUnassignGroupsOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
