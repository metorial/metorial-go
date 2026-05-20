package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/consumerinvites"
)

// PortalsConsumerInvitesEndpoint provides access to list and inspect consumer invites for a portal.
type PortalsConsumerInvitesEndpoint struct {
	client *endpoint.Client
}

// NewPortalsConsumerInvitesEndpoint creates a new PortalsConsumerInvitesEndpoint.
func NewPortalsConsumerInvitesEndpoint(client *endpoint.Client) *PortalsConsumerInvitesEndpoint {
	return &PortalsConsumerInvitesEndpoint{client: client}
}

// PortalsConsumerInvitesEndpointListParams contains optional query parameters for List.
type PortalsConsumerInvitesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Search *string  `json:"search,omitempty"`
	Status *any     `json:"status,omitempty"`
}

// PortalsConsumerInvitesEndpointCreateBody contains the request body for Create.
type PortalsConsumerInvitesEndpointCreateBody struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Message *string `json:"message,omitempty"`
}

// List returns a paginated list of invites for a portal.
func (e *PortalsConsumerInvitesEndpoint) List(portalId string, params *PortalsConsumerInvitesEndpointListParams) (*consumerinvites.PortalsConsumerInvitesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"portals", portalId, "consumer-invites"},
		Query: query,
	}
	var result consumerinvites.PortalsConsumerInvitesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create invites a consumer to a portal.
func (e *PortalsConsumerInvitesEndpoint) Create(portalId string, body *PortalsConsumerInvitesEndpointCreateBody) (*consumerinvites.PortalsConsumerInvitesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId, "consumer-invites"},
		Body: body,
	}
	var result consumerinvites.PortalsConsumerInvitesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a portal consumer invite by ID.
func (e *PortalsConsumerInvitesEndpoint) Get(portalId string, consumerInviteId string) (*consumerinvites.PortalsConsumerInvitesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"portals", portalId, "consumer-invites", consumerInviteId},
	}
	var result consumerinvites.PortalsConsumerInvitesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
