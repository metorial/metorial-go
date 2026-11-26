package delegationrequests

import (
	"encoding/json"
	"time"
)

// IdentitiesDelegationRequestsApproveOutputRequester represents the identities delegation requests approve output requester type.
type IdentitiesDelegationRequestsApproveOutputRequester struct {
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

// IdentitiesDelegationRequestsApproveOutputDelegationAttestation - How the delegation was approved. `null` when not recorded.
type IdentitiesDelegationRequestsApproveOutputDelegationAttestation struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the delegation was approved.
	Type string `json:"type"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsApproveOutputDelegationIdentity - Identity the delegation grants access to.
type IdentitiesDelegationRequestsApproveOutputDelegationIdentity struct {
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

// IdentitiesDelegationRequestsApproveOutputDelegationPartiesActor represents the identities delegation requests approve output delegation parties actor type.
type IdentitiesDelegationRequestsApproveOutputDelegationPartiesActor struct {
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

// IdentitiesDelegationRequestsApproveOutputDelegationParties represents the identities delegation requests approve output delegation parties type.
type IdentitiesDelegationRequestsApproveOutputDelegationParties struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Roles - Roles this actor has in the delegation.
	Roles []string                                                        `json:"roles"`
	Actor IdentitiesDelegationRequestsApproveOutputDelegationPartiesActor `json:"actor"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsApproveOutputDelegationRequestRequester represents the identities delegation requests approve output delegation request requester type.
type IdentitiesDelegationRequestsApproveOutputDelegationRequestRequester struct {
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

// IdentitiesDelegationRequestsApproveOutputDelegationRequest - Request that created this delegation. `null` for direct delegations.
type IdentitiesDelegationRequestsApproveOutputDelegationRequest struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current status of the related delegation request.
	Status string `json:"status"`
	// DeniedReason - Reason the request was denied, if applicable.
	DeniedReason *string                                                             `json:"denied_reason,omitempty"`
	Requester    IdentitiesDelegationRequestsApproveOutputDelegationRequestRequester `json:"requester"`
	// IdentityId - Identity targeted by the request.
	IdentityId string `json:"identity_id"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IdentitiesDelegationRequestsApproveOutputDelegationCredentialOverrides represents the identities delegation requests approve output delegation credential overrides type.
type IdentitiesDelegationRequestsApproveOutputDelegationCredentialOverrides struct {
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

// IdentitiesDelegationRequestsApproveOutputDelegation represents the identities delegation requests approve output delegation type.
type IdentitiesDelegationRequestsApproveOutputDelegation struct {
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
	Attestation *IdentitiesDelegationRequestsApproveOutputDelegationAttestation `json:"attestation,omitempty"`
	// Note - Optional note explaining the delegation.
	Note *string `json:"note,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Identity - Identity the delegation grants access to.
	Identity IdentitiesDelegationRequestsApproveOutputDelegationIdentity `json:"identity"`
	// DelegationConfigId - Delegation config used to evaluate this delegation.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// Parties - Actors involved in the delegation and their roles.
	Parties []IdentitiesDelegationRequestsApproveOutputDelegationParties `json:"parties"`
	// Request - Request that created this delegation. `null` for direct delegations.
	Request *IdentitiesDelegationRequestsApproveOutputDelegationRequest `json:"request,omitempty"`
	// CredentialOverrides - Per-credential permission overrides attached to the delegation.
	CredentialOverrides []IdentitiesDelegationRequestsApproveOutputDelegationCredentialOverrides `json:"credential_overrides"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - When the delegation expires. `null` when it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// RevokedAt - When the delegation was revoked. `null` unless revoked.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

// IdentitiesDelegationRequestsApproveOutput represents the identities delegation requests approve output type.
type IdentitiesDelegationRequestsApproveOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Where the request is in the approval workflow.
	Status string `json:"status"`
	// DeniedReason - Reason the request ultimately resulted in a denied delegation.
	DeniedReason *string                                            `json:"denied_reason,omitempty"`
	Requester    IdentitiesDelegationRequestsApproveOutputRequester `json:"requester"`
	// IdentityId - Identity targeted by the delegation request.
	IdentityId string                                              `json:"identity_id"`
	Delegation IdentitiesDelegationRequestsApproveOutputDelegation `json:"delegation"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapIdentitiesDelegationRequestsApproveOutputFromJSON deserializes JSON data into a IdentitiesDelegationRequestsApproveOutput.
func MapIdentitiesDelegationRequestsApproveOutputFromJSON(data []byte) (*IdentitiesDelegationRequestsApproveOutput, error) {
	var v IdentitiesDelegationRequestsApproveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationRequestsApproveOutputToJSON serializes a IdentitiesDelegationRequestsApproveOutput to JSON.
func MapIdentitiesDelegationRequestsApproveOutputToJSON(v *IdentitiesDelegationRequestsApproveOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentitiesDelegationRequestsApproveQuery represents the identities delegation requests approve query type.
type IdentitiesDelegationRequestsApproveQuery struct {
	// AllowDeleted - Allow approving a request that is already deleted.
	AllowDeleted *bool `json:"allow_deleted,omitempty"`
}

// MapIdentitiesDelegationRequestsApproveQueryFromJSON deserializes JSON data into a IdentitiesDelegationRequestsApproveQuery.
func MapIdentitiesDelegationRequestsApproveQueryFromJSON(data []byte) (*IdentitiesDelegationRequestsApproveQuery, error) {
	var v IdentitiesDelegationRequestsApproveQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationRequestsApproveQueryToJSON serializes a IdentitiesDelegationRequestsApproveQuery to JSON.
func MapIdentitiesDelegationRequestsApproveQueryToJSON(v *IdentitiesDelegationRequestsApproveQuery) ([]byte, error) {
	return json.Marshal(v)
}
