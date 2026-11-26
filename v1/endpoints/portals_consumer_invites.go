package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/consumerinvites"
)

// PortalsConsumerInvitesEndpoint provides access to a portal consumer invite is an email invitation for one person to join a portal.
type PortalsConsumerInvitesEndpoint struct {
	client *endpoint.Client
}

// NewPortalsConsumerInvitesEndpoint creates a new PortalsConsumerInvitesEndpoint.
func NewPortalsConsumerInvitesEndpoint(client *endpoint.Client) *PortalsConsumerInvitesEndpoint {
	return &PortalsConsumerInvitesEndpoint{client: client}
}

// PortalsConsumerInvitesEndpointListParams contains optional query parameters for List.
type PortalsConsumerInvitesEndpointListParams struct {
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
	// Email - Only return invites sent to one of these emails.
	Email *any `json:"email,omitempty"`
	// Status - Only return invites with one of these statuses.
	Status *any `json:"status,omitempty"`
}

// PortalsConsumerInvitesEndpointCreateBody contains the request body for Create.
type PortalsConsumerInvitesEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address to invite.
	Email string `json:"email"`
	// Message - Personal message included in the invite.
	Message *string `json:"message,omitempty"`
}

// List returns a paginated list of consumer invites for a portal.
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

// Get retrieves a portal consumer invite.
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
