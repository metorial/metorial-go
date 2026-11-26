package delegationrequests

import (
	"encoding/json"
	"time"
)

// IdentitiesDelegationRequestsGetOutputRequester represents the identities delegation requests get output requester type.
type IdentitiesDelegationRequestsGetOutputRequester struct {
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

// IdentitiesDelegationRequestsGetOutputDelegationAttestation - How the delegation was approved. `null` when not recorded.
type IdentitiesDelegationRequestsGetOutputDelegationAttestation struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the delegation was approved.
	Type string `json:"type"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsGetOutputDelegationIdentity - Identity the delegation grants access to.
type IdentitiesDelegationRequestsGetOutputDelegationIdentity struct {
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

// IdentitiesDelegationRequestsGetOutputDelegationPartiesActor represents the identities delegation requests get output delegation parties actor type.
type IdentitiesDelegationRequestsGetOutputDelegationPartiesActor struct {
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

// IdentitiesDelegationRequestsGetOutputDelegationParties represents the identities delegation requests get output delegation parties type.
type IdentitiesDelegationRequestsGetOutputDelegationParties struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Roles - Roles this actor has in the delegation.
	Roles []string                                                    `json:"roles"`
	Actor IdentitiesDelegationRequestsGetOutputDelegationPartiesActor `json:"actor"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsGetOutputDelegationRequestRequester represents the identities delegation requests get output delegation request requester type.
type IdentitiesDelegationRequestsGetOutputDelegationRequestRequester struct {
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

// IdentitiesDelegationRequestsGetOutputDelegationRequest - Request that created this delegation. `null` for direct delegations.
type IdentitiesDelegationRequestsGetOutputDelegationRequest struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current status of the related delegation request.
	Status string `json:"status"`
	// DeniedReason - Reason the request was denied, if applicable.
	DeniedReason *string                                                         `json:"denied_reason,omitempty"`
	Requester    IdentitiesDelegationRequestsGetOutputDelegationRequestRequester `json:"requester"`
	// IdentityId - Identity targeted by the request.
	IdentityId string `json:"identity_id"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsGetOutputDelegationCredentialOverrides represents the identities delegation requests get output delegation credential overrides type.
type IdentitiesDelegationRequestsGetOutputDelegationCredentialOverrides struct {
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

// IdentitiesDelegationRequestsGetOutputDelegation represents the identities delegation requests get output delegation type.
type IdentitiesDelegationRequestsGetOutputDelegation struct {
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
	Attestation *IdentitiesDelegationRequestsGetOutputDelegationAttestation `json:"attestation,omitempty"`
	// Note - Optional note explaining the delegation.
	Note *string `json:"note,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Identity - Identity the delegation grants access to.
	Identity IdentitiesDelegationRequestsGetOutputDelegationIdentity `json:"identity"`
	// DelegationConfigId - Delegation config used to evaluate this delegation.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// Parties - Actors involved in the delegation and their roles.
	Parties []IdentitiesDelegationRequestsGetOutputDelegationParties `json:"parties"`
	// Request - Request that created this delegation. `null` for direct delegations.
	Request *IdentitiesDelegationRequestsGetOutputDelegationRequest `json:"request,omitempty"`
	// CredentialOverrides - Per-credential permission overrides attached to the delegation.
	CredentialOverrides []IdentitiesDelegationRequestsGetOutputDelegationCredentialOverrides `json:"credential_overrides"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - When the delegation expires. `null` when it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// RevokedAt - When the delegation was revoked. `null` unless revoked.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

// IdentitiesDelegationRequestsGetOutput represents the identities delegation requests get output type.
type IdentitiesDelegationRequestsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Where the request is in the approval workflow.
	Status string `json:"status"`
	// DeniedReason - Reason the request ultimately resulted in a denied delegation.
	DeniedReason *string                                        `json:"denied_reason,omitempty"`
	Requester    IdentitiesDelegationRequestsGetOutputRequester `json:"requester"`
	// IdentityId - Identity targeted by the delegation request.
	IdentityId string                                          `json:"identity_id"`
	Delegation IdentitiesDelegationRequestsGetOutputDelegation `json:"delegation"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapIdentitiesDelegationRequestsGetOutputFromJSON deserializes JSON data into a IdentitiesDelegationRequestsGetOutput.
func MapIdentitiesDelegationRequestsGetOutputFromJSON(data []byte) (*IdentitiesDelegationRequestsGetOutput, error) {
	var v IdentitiesDelegationRequestsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationRequestsGetOutputToJSON serializes a IdentitiesDelegationRequestsGetOutput to JSON.
func MapIdentitiesDelegationRequestsGetOutputToJSON(v *IdentitiesDelegationRequestsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentitiesDelegationRequestsGetQuery represents the identities delegation requests get query type.
type IdentitiesDelegationRequestsGetQuery struct {
	// AllowDeleted - Return the request even if it has been deleted.
	AllowDeleted *bool `json:"allow_deleted,omitempty"`
}

// MapIdentitiesDelegationRequestsGetQueryFromJSON deserializes JSON data into a IdentitiesDelegationRequestsGetQuery.
func MapIdentitiesDelegationRequestsGetQueryFromJSON(data []byte) (*IdentitiesDelegationRequestsGetQuery, error) {
	var v IdentitiesDelegationRequestsGetQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationRequestsGetQueryToJSON serializes a IdentitiesDelegationRequestsGetQuery to JSON.
func MapIdentitiesDelegationRequestsGetQueryToJSON(v *IdentitiesDelegationRequestsGetQuery) ([]byte, error) {
	return json.Marshal(v)
}
