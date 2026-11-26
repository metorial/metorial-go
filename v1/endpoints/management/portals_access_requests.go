package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/accessrequests"
)

// PortalsAccessRequestsEndpoint provides access to a portal access request is a consumer's request for access to a listing. Approving it grants access through a consumer group.
type PortalsAccessRequestsEndpoint struct {
	client *endpoint.Client
}

// NewPortalsAccessRequestsEndpoint creates a new PortalsAccessRequestsEndpoint.
func NewPortalsAccessRequestsEndpoint(client *endpoint.Client) *PortalsAccessRequestsEndpoint {
	return &PortalsAccessRequestsEndpoint{client: client}
}

// PortalsAccessRequestsEndpointListParams contains optional query parameters for List.
type PortalsAccessRequestsEndpointListParams struct {
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
	// Status - Only return requests with one of these statuses.
	Status *any `json:"status,omitempty"`
	// ConsumerProfileId - Only return requests from these consumer profiles.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// Search - Search by name or description.
	Search *string `json:"search,omitempty"`
}

// PortalsAccessRequestsEndpointUpdateBody contains the request body for Update.
type PortalsAccessRequestsEndpointUpdateBody struct {
	// Status - Whether to approve or reject the request.
	Status string `json:"status"`
	// ResolutionMessage - Message shown to the consumer about the decision.
	ResolutionMessage *string `json:"resolution_message,omitempty"`
	// ConsumerGroupId - Consumer group the approval grants access through.
	ConsumerGroupId *string `json:"consumer_group_id,omitempty"`
}

// List returns a paginated list of access requests for a portal.
func (e *PortalsAccessRequestsEndpoint) List(instanceId string, portalId string, params *PortalsAccessRequestsEndpointListParams) (*accessrequests.PortalsAccessRequestsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "portals", portalId, "access-requests"},
		Query: query,
	}
	var result accessrequests.PortalsAccessRequestsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a portal access request.
func (e *PortalsAccessRequestsEndpoint) Get(instanceId string, portalId string, accessRequestId string) (*accessrequests.PortalsAccessRequestsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "access-requests", accessRequestId},
	}
	var result accessrequests.PortalsAccessRequestsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update approves or rejects a portal access request. Approving with a consumer group grants access through that group.
func (e *PortalsAccessRequestsEndpoint) Update(instanceId string, portalId string, accessRequestId string, body *PortalsAccessRequestsEndpointUpdateBody) (*accessrequests.PortalsAccessRequestsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "access-requests", accessRequestId},
		Body: body,
	}
	var result accessrequests.PortalsAccessRequestsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
