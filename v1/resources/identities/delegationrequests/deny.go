package delegationrequests

import (
	"encoding/json"
	"time"
)

// IdentitiesDelegationRequestsDenyOutputRequester represents the identities delegation requests deny output requester type.
type IdentitiesDelegationRequestsDenyOutputRequester struct {
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

// IdentitiesDelegationRequestsDenyOutputDelegationAttestation - How the delegation was approved. `null` when not recorded.
type IdentitiesDelegationRequestsDenyOutputDelegationAttestation struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the delegation was approved.
	Type string `json:"type"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsDenyOutputDelegationIdentity - Identity the delegation grants access to.
type IdentitiesDelegationRequestsDenyOutputDelegationIdentity struct {
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

// IdentitiesDelegationRequestsDenyOutputDelegationPartiesActor represents the identities delegation requests deny output delegation parties actor type.
type IdentitiesDelegationRequestsDenyOutputDelegationPartiesActor struct {
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

// IdentitiesDelegationRequestsDenyOutputDelegationParties represents the identities delegation requests deny output delegation parties type.
type IdentitiesDelegationRequestsDenyOutputDelegationParties struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Roles - Roles this actor has in the delegation.
	Roles []string                                                     `json:"roles"`
	Actor IdentitiesDelegationRequestsDenyOutputDelegationPartiesActor `json:"actor"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsDenyOutputDelegationRequestRequester represents the identities delegation requests deny output delegation request requester type.
type IdentitiesDelegationRequestsDenyOutputDelegationRequestRequester struct {
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

// IdentitiesDelegationRequestsDenyOutputDelegationRequest - Request that created this delegation. `null` for direct delegations.
type IdentitiesDelegationRequestsDenyOutputDelegationRequest struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current status of the related delegation request.
	Status string `json:"status"`
	// DeniedReason - Reason the request was denied, if applicable.
	DeniedReason *string                                                          `json:"denied_reason,omitempty"`
	Requester    IdentitiesDelegationRequestsDenyOutputDelegationRequestRequester `json:"requester"`
	// IdentityId - Identity targeted by the request.
	IdentityId string `json:"identity_id"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsDenyOutputDelegationCredentialOverrides represents the identities delegation requests deny output delegation credential overrides type.
type IdentitiesDelegationRequestsDenyOutputDelegationCredentialOverrides struct {
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

// IdentitiesDelegationRequestsDenyOutputDelegation represents the identities delegation requests deny output delegation type.
type IdentitiesDelegationRequestsDenyOutputDelegation struct {
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
	Attestation *IdentitiesDelegationRequestsDenyOutputDelegationAttestation `json:"attestation,omitempty"`
	// Note - Optional note explaining the delegation.
	Note *string `json:"note,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Identity - Identity the delegation grants access to.
	Identity IdentitiesDelegationRequestsDenyOutputDelegationIdentity `json:"identity"`
	// DelegationConfigId - Delegation config used to evaluate this delegation.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// Parties - Actors involved in the delegation and their roles.
	Parties []IdentitiesDelegationRequestsDenyOutputDelegationParties `json:"parties"`
	// Request - Request that created this delegation. `null` for direct delegations.
	Request *IdentitiesDelegationRequestsDenyOutputDelegationRequest `json:"request,omitempty"`
	// CredentialOverrides - Per-credential permission overrides attached to the delegation.
	CredentialOverrides []IdentitiesDelegationRequestsDenyOutputDelegationCredentialOverrides `json:"credential_overrides"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - When the delegation expires. `null` when it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// RevokedAt - When the delegation was revoked. `null` unless revoked.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

// IdentitiesDelegationRequestsDenyOutput represents the identities delegation requests deny output type.
type IdentitiesDelegationRequestsDenyOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Where the request is in the approval workflow.
	Status string `json:"status"`
	// DeniedReason - Reason the request ultimately resulted in a denied delegation.
	DeniedReason *string                                         `json:"denied_reason,omitempty"`
	Requester    IdentitiesDelegationRequestsDenyOutputRequester `json:"requester"`
	// IdentityId - Identity targeted by the delegation request.
	IdentityId string                                           `json:"identity_id"`
	Delegation IdentitiesDelegationRequestsDenyOutputDelegation `json:"delegation"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapIdentitiesDelegationRequestsDenyOutputFromJSON deserializes JSON data into a IdentitiesDelegationRequestsDenyOutput.
func MapIdentitiesDelegationRequestsDenyOutputFromJSON(data []byte) (*IdentitiesDelegationRequestsDenyOutput, error) {
	var v IdentitiesDelegationRequestsDenyOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationRequestsDenyOutputToJSON serializes a IdentitiesDelegationRequestsDenyOutput to JSON.
func MapIdentitiesDelegationRequestsDenyOutputToJSON(v *IdentitiesDelegationRequestsDenyOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentitiesDelegationRequestsDenyQuery represents the identities delegation requests deny query type.
type IdentitiesDelegationRequestsDenyQuery struct {
	// AllowDeleted - Allow denying a request that is already deleted.
	AllowDeleted *bool `json:"allow_deleted,omitempty"`
}

// MapIdentitiesDelegationRequestsDenyQueryFromJSON deserializes JSON data into a IdentitiesDelegationRequestsDenyQuery.
func MapIdentitiesDelegationRequestsDenyQueryFromJSON(data []byte) (*IdentitiesDelegationRequestsDenyQuery, error) {
	var v IdentitiesDelegationRequestsDenyQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationRequestsDenyQueryToJSON serializes a IdentitiesDelegationRequestsDenyQuery to JSON.
func MapIdentitiesDelegationRequestsDenyQueryToJSON(v *IdentitiesDelegationRequestsDenyQuery) ([]byte, error) {
	return json.Marshal(v)
}
