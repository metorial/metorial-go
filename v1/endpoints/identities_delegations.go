package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/identities/delegations"
)

// IdentitiesDelegationsEndpoint provides access to an identity delegation is a grant of provider permissions from an identity owner to another actor. It can narrow permissions per credential.
type IdentitiesDelegationsEndpoint struct {
	client *endpoint.Client
}

// NewIdentitiesDelegationsEndpoint creates a new IdentitiesDelegationsEndpoint.
func NewIdentitiesDelegationsEndpoint(client *endpoint.Client) *IdentitiesDelegationsEndpoint {
	return &IdentitiesDelegationsEndpoint{client: client}
}

// IdentitiesDelegationsEndpointListParams contains optional query parameters for List.
type IdentitiesDelegationsEndpointListParams struct {
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
	// Status - Only return identity delegations with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Permissions - Only return identity delegations granting one of these permissions.
	Permissions *any `json:"permissions,omitempty"`
	// Id - Only return identity delegations with one of these IDs.
	Id *any `json:"id,omitempty"`
	// OwnerActorId - Only return identity delegations owned by one of these actors.
	OwnerActorId *any `json:"owner_actor_id,omitempty"`
	// DelegatorActorId - Only return identity delegations started by one of these actors.
	DelegatorActorId *any `json:"delegator_actor_id,omitempty"`
	// DelegateeActorId - Only return identity delegations received by one of these actors.
	DelegateeActorId *any `json:"delegatee_actor_id,omitempty"`
	// IdentityId - Only return identity delegations for one of these identities.
	IdentityId *any `json:"identity_id,omitempty"`
	// CreatedAt - Only return results whose identity delegation creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity delegation last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IdentitiesDelegationsEndpointCreateBody contains the request body for Create.
type IdentitiesDelegationsEndpointCreateBody struct {
	// IdentityId - Identity to delegate.
	IdentityId string `json:"identity_id"`
	// DelegatorActorId - Actor initiating the delegation, if different from the owner.
	DelegatorActorId *string `json:"delegator_actor_id,omitempty"`
	// DelegateeActorId - Actor receiving the delegation.
	DelegateeActorId string `json:"delegatee_actor_id"`
	// Permissions - Permissions to grant as part of the delegation.
	Permissions *[]string `json:"permissions,omitempty"`
	// ExpiresAt - Optional expiration timestamp for the delegation.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// DelegationConfigId - Delegation config to use for this delegation.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// CredentialOverrides - Optional per-credential permission overrides.
	CredentialOverrides *[]map[string]any `json:"credential_overrides,omitempty"`
	// Note - Optional human-readable note for the delegation.
	Note *string `json:"note,omitempty"`
	// Metadata - Custom data to identify or organize the identity delegation.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of identity delegations for the instance.
func (e *IdentitiesDelegationsEndpoint) List(params *IdentitiesDelegationsEndpointListParams) (*delegations.IdentitiesDelegationsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"identity-delegations"},
		Query: query,
	}
	var result delegations.IdentitiesDelegationsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an identity delegation.
func (e *IdentitiesDelegationsEndpoint) Get(identityDelegationId string) (*delegations.IdentitiesDelegationsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identity-delegations", identityDelegationId},
	}
	var result delegations.IdentitiesDelegationsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an identity delegation.
func (e *IdentitiesDelegationsEndpoint) Create(body *IdentitiesDelegationsEndpointCreateBody) (*delegations.IdentitiesDelegationsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identity-delegations"},
		Body: body,
	}
	var result delegations.IdentitiesDelegationsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Revoke revokes an identity delegation.
func (e *IdentitiesDelegationsEndpoint) Revoke(identityDelegationId string) (*delegations.IdentitiesDelegationsRevokeOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identity-delegations", identityDelegationId, "revoke"},
	}
	var result delegations.IdentitiesDelegationsRevokeOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
