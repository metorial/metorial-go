package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/consumerprofiles"
)

// PortalsConsumerProfilesEndpoint provides access to manage the consumers and effective group assignments for a portal.
type PortalsConsumerProfilesEndpoint struct {
	client *endpoint.Client
}

// NewPortalsConsumerProfilesEndpoint creates a new PortalsConsumerProfilesEndpoint.
func NewPortalsConsumerProfilesEndpoint(client *endpoint.Client) *PortalsConsumerProfilesEndpoint {
	return &PortalsConsumerProfilesEndpoint{client: client}
}

// PortalsConsumerProfilesEndpointListParams contains optional query parameters for List.
type PortalsConsumerProfilesEndpointListParams struct {
	Limit           *float64 `json:"limit,omitempty"`
	After           *string  `json:"after,omitempty"`
	Before          *string  `json:"before,omitempty"`
	Cursor          *string  `json:"cursor,omitempty"`
	Order           *string  `json:"order,omitempty"`
	Search          *string  `json:"search,omitempty"`
	ConsumerGroupId *string  `json:"consumer_group_id,omitempty"`
	Status          *any     `json:"status,omitempty"`
}

// PortalsConsumerProfilesEndpointAssignGroupsBody contains the request body for AssignGroups.
type PortalsConsumerProfilesEndpointAssignGroupsBody struct {
	GroupIds []string `json:"group_ids"`
}

// PortalsConsumerProfilesEndpointUnassignGroupsBody contains the request body for UnassignGroups.
type PortalsConsumerProfilesEndpointUnassignGroupsBody struct {
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

// Get retrieves a portal consumer profile by ID.
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
