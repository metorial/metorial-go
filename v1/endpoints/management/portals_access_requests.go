package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/accessrequests"
)

// PortalsAccessRequestsEndpoint provides access to review and resolve access requests for a portal.
type PortalsAccessRequestsEndpoint struct {
	client *endpoint.Client
}

// NewPortalsAccessRequestsEndpoint creates a new PortalsAccessRequestsEndpoint.
func NewPortalsAccessRequestsEndpoint(client *endpoint.Client) *PortalsAccessRequestsEndpoint {
	return &PortalsAccessRequestsEndpoint{client: client}
}

// PortalsAccessRequestsEndpointListParams contains optional query parameters for List.
type PortalsAccessRequestsEndpointListParams struct {
	Limit             *float64 `json:"limit,omitempty"`
	After             *string  `json:"after,omitempty"`
	Before            *string  `json:"before,omitempty"`
	Cursor            *string  `json:"cursor,omitempty"`
	Order             *string  `json:"order,omitempty"`
	Status            *any     `json:"status,omitempty"`
	ConsumerProfileId *any     `json:"consumer_profile_id,omitempty"`
	Search            *string  `json:"search,omitempty"`
}

// PortalsAccessRequestsEndpointUpdateBody contains the request body for Update.
type PortalsAccessRequestsEndpointUpdateBody struct {
	Status            string  `json:"status"`
	ResolutionMessage *string `json:"resolution_message,omitempty"`
	ConsumerGroupId   *string `json:"consumer_group_id,omitempty"`
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

// Get retrieves a access request by ID.
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

// Update approves or rejects a access request.
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
