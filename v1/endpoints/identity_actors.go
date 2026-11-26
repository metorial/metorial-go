package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/identityactors"
)

// IdentityActorsEndpoint provides access to an identity actor is a person or agent that can own identities and take part in delegations.
type IdentityActorsEndpoint struct {
	client *endpoint.Client
}

// NewIdentityActorsEndpoint creates a new IdentityActorsEndpoint.
func NewIdentityActorsEndpoint(client *endpoint.Client) *IdentityActorsEndpoint {
	return &IdentityActorsEndpoint{client: client}
}

// IdentityActorsEndpointListParams contains optional query parameters for List.
type IdentityActorsEndpointListParams struct {
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
	// Search - Filter actors by name or description.
	Search *string `json:"search,omitempty"`
	// Status - Only return identity actors with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return identity actors with one of these IDs.
	Id *any `json:"id,omitempty"`
	// AgentId - Only return identity actors linked to one of these agents.
	AgentId *any `json:"agent_id,omitempty"`
	// ConsumerProfileId - Only return identity actors linked to one of these consumer profiles.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// ConsumerId - Only return identity actors linked to one of these consumers.
	ConsumerId *any `json:"consumer_id,omitempty"`
	// CreatedAt - Only return results whose identity actor creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity actor last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IdentityActorsEndpointCreateBody contains the request body for Create.
type IdentityActorsEndpointCreateBody struct {
	// Type - Whether this actor is a person or an agent.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the identity actor.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// IdentityActorsEndpointUpdateBody contains the request body for Update.
type IdentityActorsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the identity actor. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of identity actors for the instance.
func (e *IdentityActorsEndpoint) List(params *IdentityActorsEndpointListParams) (*identityactors.IdentityActorsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"identity-actors"},
		Query: query,
	}
	var result identityactors.IdentityActorsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an identity actor.
func (e *IdentityActorsEndpoint) Get(identityActorId string) (*identityactors.IdentityActorsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identity-actors", identityActorId},
	}
	var result identityactors.IdentityActorsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an identity actor.
func (e *IdentityActorsEndpoint) Create(body *IdentityActorsEndpointCreateBody) (*identityactors.IdentityActorsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identity-actors"},
		Body: body,
	}
	var result identityactors.IdentityActorsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an identity actor. Omitted fields keep their current value.
func (e *IdentityActorsEndpoint) Update(identityActorId string, body *IdentityActorsEndpointUpdateBody) (*identityactors.IdentityActorsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identity-actors", identityActorId},
		Body: body,
	}
	var result identityactors.IdentityActorsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives an identity actor. It stays listable with `status=archived`.
func (e *IdentityActorsEndpoint) Delete(identityActorId string) (*identityactors.IdentityActorsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identity-actors", identityActorId},
	}
	var result identityactors.IdentityActorsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
