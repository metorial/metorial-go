package delegationrequests

import (
	"encoding/json"
	"time"
)

// IdentitiesDelegationRequestsListOutputItemsRequester represents the identities delegation requests list output items requester type.
type IdentitiesDelegationRequestsListOutputItemsRequester struct {
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

// IdentitiesDelegationRequestsListOutputItemsDelegationAttestation - How the delegation was approved. `null` when not recorded.
type IdentitiesDelegationRequestsListOutputItemsDelegationAttestation struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the delegation was approved.
	Type string `json:"type"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsListOutputItemsDelegationIdentity - Identity the delegation grants access to.
type IdentitiesDelegationRequestsListOutputItemsDelegationIdentity struct {
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

// IdentitiesDelegationRequestsListOutputItemsDelegationPartiesActor represents the identities delegation requests list output items delegation parties actor type.
type IdentitiesDelegationRequestsListOutputItemsDelegationPartiesActor struct {
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

// IdentitiesDelegationRequestsListOutputItemsDelegationParties represents the identities delegation requests list output items delegation parties type.
type IdentitiesDelegationRequestsListOutputItemsDelegationParties struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Roles - Roles this actor has in the delegation.
	Roles []string                                                          `json:"roles"`
	Actor IdentitiesDelegationRequestsListOutputItemsDelegationPartiesActor `json:"actor"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsListOutputItemsDelegationRequestRequester represents the identities delegation requests list output items delegation request requester type.
type IdentitiesDelegationRequestsListOutputItemsDelegationRequestRequester struct {
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

// IdentitiesDelegationRequestsListOutputItemsDelegationRequest - Request that created this delegation. `null` for direct delegations.
type IdentitiesDelegationRequestsListOutputItemsDelegationRequest struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current status of the related delegation request.
	Status string `json:"status"`
	// DeniedReason - Reason the request was denied, if applicable.
	DeniedReason *string                                                               `json:"denied_reason,omitempty"`
	Requester    IdentitiesDelegationRequestsListOutputItemsDelegationRequestRequester `json:"requester"`
	// IdentityId - Identity targeted by the request.
	IdentityId string `json:"identity_id"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsListOutputItemsDelegationCredentialOverrides represents the identities delegation requests list output items delegation credential overrides type.
type IdentitiesDelegationRequestsListOutputItemsDelegationCredentialOverrides struct {
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

// IdentitiesDelegationRequestsListOutputItemsDelegation represents the identities delegation requests list output items delegation type.
type IdentitiesDelegationRequestsListOutputItemsDelegation struct {
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
	Attestation *IdentitiesDelegationRequestsListOutputItemsDelegationAttestation `json:"attestation,omitempty"`
	// Note - Optional note explaining the delegation.
	Note *string `json:"note,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Identity - Identity the delegation grants access to.
	Identity IdentitiesDelegationRequestsListOutputItemsDelegationIdentity `json:"identity"`
	// DelegationConfigId - Delegation config used to evaluate this delegation.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// Parties - Actors involved in the delegation and their roles.
	Parties []IdentitiesDelegationRequestsListOutputItemsDelegationParties `json:"parties"`
	// Request - Request that created this delegation. `null` for direct delegations.
	Request *IdentitiesDelegationRequestsListOutputItemsDelegationRequest `json:"request,omitempty"`
	// CredentialOverrides - Per-credential permission overrides attached to the delegation.
	CredentialOverrides []IdentitiesDelegationRequestsListOutputItemsDelegationCredentialOverrides `json:"credential_overrides"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - When the delegation expires. `null` when it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// RevokedAt - When the delegation was revoked. `null` unless revoked.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

// IdentitiesDelegationRequestsListOutputItems represents the identities delegation requests list output items type.
type IdentitiesDelegationRequestsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Where the request is in the approval workflow.
	Status string `json:"status"`
	// DeniedReason - Reason the request ultimately resulted in a denied delegation.
	DeniedReason *string                                              `json:"denied_reason,omitempty"`
	Requester    IdentitiesDelegationRequestsListOutputItemsRequester `json:"requester"`
	// IdentityId - Identity targeted by the delegation request.
	IdentityId string                                                `json:"identity_id"`
	Delegation IdentitiesDelegationRequestsListOutputItemsDelegation `json:"delegation"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsListOutputPagination represents the identities delegation requests list output pagination type.
type IdentitiesDelegationRequestsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IdentitiesDelegationRequestsListOutput represents the identities delegation requests list output type.
type IdentitiesDelegationRequestsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                           `json:"object"`
	Items      []IdentitiesDelegationRequestsListOutputItems    `json:"items"`
	Pagination IdentitiesDelegationRequestsListOutputPagination `json:"pagination"`
}

// MapIdentitiesDelegationRequestsListOutputFromJSON deserializes JSON data into a IdentitiesDelegationRequestsListOutput.
func MapIdentitiesDelegationRequestsListOutputFromJSON(data []byte) (*IdentitiesDelegationRequestsListOutput, error) {
	var v IdentitiesDelegationRequestsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationRequestsListOutputToJSON serializes a IdentitiesDelegationRequestsListOutput to JSON.
func MapIdentitiesDelegationRequestsListOutputToJSON(v *IdentitiesDelegationRequestsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentitiesDelegationRequestsListQueryCreatedAt - Only return results whose identity delegation request creation time is in this range. Set `gt`, `lt`, or both.
type IdentitiesDelegationRequestsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentitiesDelegationRequestsListQueryUpdatedAt - Only return results whose identity delegation request last update time is in this range. Set `gt`, `lt`, or both.
type IdentitiesDelegationRequestsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentitiesDelegationRequestsListQuery represents the identities delegation requests list query type.
type IdentitiesDelegationRequestsListQuery struct {
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
	// Status - Only return identity delegation requests with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return identity delegation requests with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ActorId - Only return identity delegation requests from one of these requesters.
	ActorId *any `json:"actor_id,omitempty"`
	// IdentityId - Only return identity delegation requests for one of these identities.
	IdentityId *any `json:"identity_id,omitempty"`
	// CreatedAt - Only return results whose identity delegation request creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IdentitiesDelegationRequestsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity delegation request last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IdentitiesDelegationRequestsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIdentitiesDelegationRequestsListQueryFromJSON deserializes JSON data into a IdentitiesDelegationRequestsListQuery.
func MapIdentitiesDelegationRequestsListQueryFromJSON(data []byte) (*IdentitiesDelegationRequestsListQuery, error) {
	var v IdentitiesDelegationRequestsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationRequestsListQueryToJSON serializes a IdentitiesDelegationRequestsListQuery to JSON.
func MapIdentitiesDelegationRequestsListQueryToJSON(v *IdentitiesDelegationRequestsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
