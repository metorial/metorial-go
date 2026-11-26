package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/identities"
)

// IdentitiesEndpoint provides access to an identity is a set of provider credentials owned by one actor. It lets you manage and delegate provider access consistently.
type IdentitiesEndpoint struct {
	client *endpoint.Client
}

// NewIdentitiesEndpoint creates a new IdentitiesEndpoint.
func NewIdentitiesEndpoint(client *endpoint.Client) *IdentitiesEndpoint {
	return &IdentitiesEndpoint{client: client}
}

// IdentitiesEndpointListParams contains optional query parameters for List.
type IdentitiesEndpointListParams struct {
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
	// Search - Filter identities by name or description.
	Search *string `json:"search,omitempty"`
	// Status - Only return identities with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return identities with one of these IDs.
	Id *any `json:"id,omitempty"`
	// AgentId - Only return identities owned by one of these agents.
	AgentId *any `json:"agent_id,omitempty"`
	// ActorId - Only return identities owned by one of these identity actors.
	ActorId *any `json:"actor_id,omitempty"`
	// CreatedAt - Only return results whose identity creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IdentitiesEndpointCreateBody contains the request body for Create.
type IdentitiesEndpointCreateBody struct {
	// ActorId - Identity actor that will own the new identity.
	ActorId string `json:"actor_id"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the identity.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Credentials - Credentials to create and attach as part of identity creation.
	Credentials *[]map[string]any `json:"credentials,omitempty"`
}

// IdentitiesEndpointUpdateBody contains the request body for Update.
type IdentitiesEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the identity. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of identities for the instance.
func (e *IdentitiesEndpoint) List(params *IdentitiesEndpointListParams) (*identities.IdentitiesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"identities"},
		Query: query,
	}
	var result identities.IdentitiesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an identity.
func (e *IdentitiesEndpoint) Get(identityId string) (*identities.IdentitiesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identities", identityId},
	}
	var result identities.IdentitiesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an identity.
func (e *IdentitiesEndpoint) Create(body *IdentitiesEndpointCreateBody) (*identities.IdentitiesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identities"},
		Body: body,
	}
	var result identities.IdentitiesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an identity. Omitted fields keep their current value.
func (e *IdentitiesEndpoint) Update(identityId string, body *IdentitiesEndpointUpdateBody) (*identities.IdentitiesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identities", identityId},
		Body: body,
	}
	var result identities.IdentitiesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives an identity. It stays listable with `status=archived`.
func (e *IdentitiesEndpoint) Delete(identityId string) (*identities.IdentitiesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identities", identityId},
	}
	var result identities.IdentitiesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
