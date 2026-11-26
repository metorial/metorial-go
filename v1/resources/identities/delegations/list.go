package delegations

import (
	"encoding/json"
	"time"
)

// IdentitiesDelegationsListOutputItemsAttestation - How the delegation was approved. `null` when not recorded.
type IdentitiesDelegationsListOutputItemsAttestation struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the delegation was approved.
	Type string `json:"type"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationsListOutputItemsIdentity - Identity the delegation grants access to.
type IdentitiesDelegationsListOutputItemsIdentity struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description string `json:"description"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// IdentitiesDelegationsListOutputItemsPartiesActor represents the identities delegations list output items parties actor type.
type IdentitiesDelegationsListOutputItemsPartiesActor struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - Type of actor that owns or participates in identities.
	Type string `json:"type"`
	// Status - Current lifecycle status of the identity actor. `archived` once you archive the actor.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// AgentId - Linked agent identifier when this actor represents an agent.
	AgentId *string `json:"agent_id,omitempty"`
	// ConsumerId - Identifier of the consumer that owns this actor when its type is consumer.
	ConsumerId *string `json:"consumer_id,omitempty"`
	// DefaultIdentityId - Identifier of the actor's default identity, which is its oldest active identity.
	DefaultIdentityId *string `json:"default_identity_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IdentitiesDelegationsListOutputItemsParties represents the identities delegations list output items parties type.
type IdentitiesDelegationsListOutputItemsParties struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Roles - Roles this actor has in the delegation.
	Roles []string                                         `json:"roles"`
	Actor IdentitiesDelegationsListOutputItemsPartiesActor `json:"actor"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationsListOutputItemsRequestRequester represents the identities delegations list output items request requester type.
type IdentitiesDelegationsListOutputItemsRequestRequester struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - Type of actor that owns or participates in identities.
	Type string `json:"type"`
	// Status - Current lifecycle status of the identity actor. `archived` once you archive the actor.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// AgentId - Linked agent identifier when this actor represents an agent.
	AgentId *string `json:"agent_id,omitempty"`
	// ConsumerId - Identifier of the consumer that owns this actor when its type is consumer.
	ConsumerId *string `json:"consumer_id,omitempty"`
	// DefaultIdentityId - Identifier of the actor's default identity, which is its oldest active identity.
	DefaultIdentityId *string `json:"default_identity_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IdentitiesDelegationsListOutputItemsRequest - Request that created this delegation. `null` for direct delegations.
type IdentitiesDelegationsListOutputItemsRequest struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current status of the related delegation request.
	Status string `json:"status"`
	// DeniedReason - Reason the request was denied, if applicable.
	DeniedReason *string                                              `json:"denied_reason,omitempty"`
	Requester    IdentitiesDelegationsListOutputItemsRequestRequester `json:"requester"`
	// IdentityId - Identity targeted by the request.
	IdentityId string `json:"identity_id"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationsListOutputItemsCredentialOverrides represents the identities delegations list output items credential overrides type.
type IdentitiesDelegationsListOutputItemsCredentialOverrides struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current status of the credential override. `expired` once past `expires_at`.
	Status string `json:"status"`
	// Permissions - Permissions granted for this credential override.
	Permissions []string `json:"permissions"`
	// CredentialId - Credential receiving the override.
	CredentialId string `json:"credential_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - When the override expires. `null` when it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// IdentitiesDelegationsListOutputItems represents the identities delegations list output items type.
type IdentitiesDelegationsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current status of the delegation. `waiting_for_consent` until the delegation is approved.
	Status string `json:"status"`
	// DeniedReason - Reason the delegation was denied, if applicable.
	DeniedReason *string `json:"denied_reason,omitempty"`
	// DelegationLevel - Depth level of this delegation in the delegation chain.
	DelegationLevel float64 `json:"delegation_level"`
	// Permissions - Permissions granted by this delegation.
	Permissions []string `json:"permissions"`
	// Attestation - How the delegation was approved. `null` when not recorded.
	Attestation *IdentitiesDelegationsListOutputItemsAttestation `json:"attestation,omitempty"`
	// Note - Optional note explaining the delegation.
	Note *string `json:"note,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Identity - Identity the delegation grants access to.
	Identity IdentitiesDelegationsListOutputItemsIdentity `json:"identity"`
	// DelegationConfigId - Delegation config used to evaluate this delegation.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// Parties - Actors involved in the delegation and their roles.
	Parties []IdentitiesDelegationsListOutputItemsParties `json:"parties"`
	// Request - Request that created this delegation. `null` for direct delegations.
	Request *IdentitiesDelegationsListOutputItemsRequest `json:"request,omitempty"`
	// CredentialOverrides - Per-credential permission overrides attached to the delegation.
	CredentialOverrides []IdentitiesDelegationsListOutputItemsCredentialOverrides `json:"credential_overrides"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - When the delegation expires. `null` when it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// RevokedAt - When the delegation was revoked. `null` unless revoked.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

// IdentitiesDelegationsListOutputPagination represents the identities delegations list output pagination type.
type IdentitiesDelegationsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IdentitiesDelegationsListOutput represents the identities delegations list output type.
type IdentitiesDelegationsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                    `json:"object"`
	Items      []IdentitiesDelegationsListOutputItems    `json:"items"`
	Pagination IdentitiesDelegationsListOutputPagination `json:"pagination"`
}

// MapIdentitiesDelegationsListOutputFromJSON deserializes JSON data into a IdentitiesDelegationsListOutput.
func MapIdentitiesDelegationsListOutputFromJSON(data []byte) (*IdentitiesDelegationsListOutput, error) {
	var v IdentitiesDelegationsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationsListOutputToJSON serializes a IdentitiesDelegationsListOutput to JSON.
func MapIdentitiesDelegationsListOutputToJSON(v *IdentitiesDelegationsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentitiesDelegationsListQueryCreatedAt - Only return results whose identity delegation creation time is in this range. Set `gt`, `lt`, or both.
type IdentitiesDelegationsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentitiesDelegationsListQueryUpdatedAt - Only return results whose identity delegation last update time is in this range. Set `gt`, `lt`, or both.
type IdentitiesDelegationsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentitiesDelegationsListQuery represents the identities delegations list query type.
type IdentitiesDelegationsListQuery struct {
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
	CreatedAt *IdentitiesDelegationsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity delegation last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IdentitiesDelegationsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIdentitiesDelegationsListQueryFromJSON deserializes JSON data into a IdentitiesDelegationsListQuery.
func MapIdentitiesDelegationsListQueryFromJSON(data []byte) (*IdentitiesDelegationsListQuery, error) {
	var v IdentitiesDelegationsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationsListQueryToJSON serializes a IdentitiesDelegationsListQuery to JSON.
func MapIdentitiesDelegationsListQueryToJSON(v *IdentitiesDelegationsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
