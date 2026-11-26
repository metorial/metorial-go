package delegationrequests

import (
	"encoding/json"
	"time"
)

// IdentitiesDelegationRequestsCreateOutputRequester represents the identities delegation requests create output requester type.
type IdentitiesDelegationRequestsCreateOutputRequester struct {
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

// IdentitiesDelegationRequestsCreateOutputDelegationAttestation - How the delegation was approved. `null` when not recorded.
type IdentitiesDelegationRequestsCreateOutputDelegationAttestation struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the delegation was approved.
	Type string `json:"type"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsCreateOutputDelegationIdentity - Identity the delegation grants access to.
type IdentitiesDelegationRequestsCreateOutputDelegationIdentity struct {
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

// IdentitiesDelegationRequestsCreateOutputDelegationPartiesActor represents the identities delegation requests create output delegation parties actor type.
type IdentitiesDelegationRequestsCreateOutputDelegationPartiesActor struct {
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

// IdentitiesDelegationRequestsCreateOutputDelegationParties represents the identities delegation requests create output delegation parties type.
type IdentitiesDelegationRequestsCreateOutputDelegationParties struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Roles - Roles this actor has in the delegation.
	Roles []string                                                       `json:"roles"`
	Actor IdentitiesDelegationRequestsCreateOutputDelegationPartiesActor `json:"actor"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsCreateOutputDelegationRequestRequester represents the identities delegation requests create output delegation request requester type.
type IdentitiesDelegationRequestsCreateOutputDelegationRequestRequester struct {
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

// IdentitiesDelegationRequestsCreateOutputDelegationRequest - Request that created this delegation. `null` for direct delegations.
type IdentitiesDelegationRequestsCreateOutputDelegationRequest struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current status of the related delegation request.
	Status string `json:"status"`
	// DeniedReason - Reason the request was denied, if applicable.
	DeniedReason *string                                                            `json:"denied_reason,omitempty"`
	Requester    IdentitiesDelegationRequestsCreateOutputDelegationRequestRequester `json:"requester"`
	// IdentityId - Identity targeted by the request.
	IdentityId string `json:"identity_id"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsCreateOutputDelegationCredentialOverrides represents the identities delegation requests create output delegation credential overrides type.
type IdentitiesDelegationRequestsCreateOutputDelegationCredentialOverrides struct {
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

// IdentitiesDelegationRequestsCreateOutputDelegation represents the identities delegation requests create output delegation type.
type IdentitiesDelegationRequestsCreateOutputDelegation struct {
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
	Attestation *IdentitiesDelegationRequestsCreateOutputDelegationAttestation `json:"attestation,omitempty"`
	// Note - Optional note explaining the delegation.
	Note *string `json:"note,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Identity - Identity the delegation grants access to.
	Identity IdentitiesDelegationRequestsCreateOutputDelegationIdentity `json:"identity"`
	// DelegationConfigId - Delegation config used to evaluate this delegation.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// Parties - Actors involved in the delegation and their roles.
	Parties []IdentitiesDelegationRequestsCreateOutputDelegationParties `json:"parties"`
	// Request - Request that created this delegation. `null` for direct delegations.
	Request *IdentitiesDelegationRequestsCreateOutputDelegationRequest `json:"request,omitempty"`
	// CredentialOverrides - Per-credential permission overrides attached to the delegation.
	CredentialOverrides []IdentitiesDelegationRequestsCreateOutputDelegationCredentialOverrides `json:"credential_overrides"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - When the delegation expires. `null` when it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// RevokedAt - When the delegation was revoked. `null` unless revoked.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

// IdentitiesDelegationRequestsCreateOutput represents the identities delegation requests create output type.
type IdentitiesDelegationRequestsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Where the request is in the approval workflow.
	Status string `json:"status"`
	// DeniedReason - Reason the request ultimately resulted in a denied delegation.
	DeniedReason *string                                           `json:"denied_reason,omitempty"`
	Requester    IdentitiesDelegationRequestsCreateOutputRequester `json:"requester"`
	// IdentityId - Identity targeted by the delegation request.
	IdentityId string                                             `json:"identity_id"`
	Delegation IdentitiesDelegationRequestsCreateOutputDelegation `json:"delegation"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapIdentitiesDelegationRequestsCreateOutputFromJSON deserializes JSON data into a IdentitiesDelegationRequestsCreateOutput.
func MapIdentitiesDelegationRequestsCreateOutputFromJSON(data []byte) (*IdentitiesDelegationRequestsCreateOutput, error) {
	var v IdentitiesDelegationRequestsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationRequestsCreateOutputToJSON serializes a IdentitiesDelegationRequestsCreateOutput to JSON.
func MapIdentitiesDelegationRequestsCreateOutputToJSON(v *IdentitiesDelegationRequestsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentitiesDelegationRequestsCreateBodyCredentialOverrides represents the identities delegation requests create body credential overrides type.
type IdentitiesDelegationRequestsCreateBodyCredentialOverrides struct {
	// CredentialId - Credential that should receive override permissions.
	CredentialId string `json:"credential_id"`
	// Permissions - Permissions to grant on the credential override.
	Permissions *[]string `json:"permissions,omitempty"`
	// ExpiresAt - Optional expiration timestamp for the credential override.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// IdentitiesDelegationRequestsCreateBody represents the identities delegation requests create body type.
type IdentitiesDelegationRequestsCreateBody struct {
	// IdentityId - Identity to request delegation for.
	IdentityId string `json:"identity_id"`
	// RequesterActorId - Actor requesting the delegation.
	RequesterActorId string `json:"requester_actor_id"`
	// DelegatorActorId - Actor submitting the request on behalf of the requester.
	DelegatorActorId *string `json:"delegator_actor_id,omitempty"`
	// Permissions - Permissions being requested.
	Permissions *[]string `json:"permissions,omitempty"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// DelegationConfigId - Delegation config to use for the resulting delegation.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// CredentialOverrides - Optional per-credential permission overrides.
	CredentialOverrides *[]IdentitiesDelegationRequestsCreateBodyCredentialOverrides `json:"credential_overrides,omitempty"`
	// Note - Optional human-readable note for the request.
	Note *string `json:"note,omitempty"`
	// Metadata - Custom data to identify or organize the identity delegation request.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapIdentitiesDelegationRequestsCreateBodyFromJSON deserializes JSON data into a IdentitiesDelegationRequestsCreateBody.
func MapIdentitiesDelegationRequestsCreateBodyFromJSON(data []byte) (*IdentitiesDelegationRequestsCreateBody, error) {
	var v IdentitiesDelegationRequestsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationRequestsCreateBodyToJSON serializes a IdentitiesDelegationRequestsCreateBody to JSON.
func MapIdentitiesDelegationRequestsCreateBodyToJSON(v *IdentitiesDelegationRequestsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
