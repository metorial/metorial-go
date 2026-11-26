package identities

import (
	"encoding/json"
	"time"
)

// IdentitiesCreateOutputOwnerActor represents the identities create output owner actor type.
type IdentitiesCreateOutputOwnerActor struct {
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

// IdentitiesCreateOutputOwner - Actor that owns the identity.
type IdentitiesCreateOutputOwner struct {
	// Type - Owner type for the identity.
	Type  string                           `json:"type"`
	Actor IdentitiesCreateOutputOwnerActor `json:"actor"`
}

// IdentitiesCreateOutputCredentials represents the identities create output credentials type.
type IdentitiesCreateOutputCredentials struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current lifecycle status of the credential. `archived` once you archive the credential.
	Status string `json:"status"`
	// IdentityId - Identity that owns this credential.
	IdentityId string `json:"identity_id"`
	// ProviderId - Provider this credential holds access for.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Provider deployment used by this credential.
	DeploymentId *string `json:"deployment_id,omitempty"`
	// ConfigId - Provider config used by this credential.
	ConfigId *string `json:"config_id,omitempty"`
	// AuthConfigId - Provider auth config used by this credential.
	AuthConfigId *string `json:"auth_config_id,omitempty"`
	// DelegationConfigId - Delegation config applied to this credential.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IdentitiesCreateOutput represents the identities create output type.
type IdentitiesCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current lifecycle status of the identity. `archived` once you archive the identity.
	Status string `json:"status"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Owner - Actor that owns the identity.
	Owner IdentitiesCreateOutputOwner `json:"owner"`
	// Credentials - Credentials currently attached to the identity.
	Credentials []IdentitiesCreateOutputCredentials `json:"credentials"`
	// DelegationConfigId - Default delegation config applied to the identity. `null` when none is set.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapIdentitiesCreateOutputFromJSON deserializes JSON data into a IdentitiesCreateOutput.
func MapIdentitiesCreateOutputFromJSON(data []byte) (*IdentitiesCreateOutput, error) {
	var v IdentitiesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesCreateOutputToJSON serializes a IdentitiesCreateOutput to JSON.
func MapIdentitiesCreateOutputToJSON(v *IdentitiesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentitiesCreateBodyCredentials represents the identities create body credentials type.
type IdentitiesCreateBodyCredentials struct {
	// DeploymentId - Provider deployment to attach to the identity.
	DeploymentId *string `json:"deployment_id,omitempty"`
	// ConfigId - Provider config to attach to the identity.
	ConfigId *string `json:"config_id,omitempty"`
	// AuthConfigId - Provider auth config to attach to the identity.
	AuthConfigId *string `json:"auth_config_id,omitempty"`
	// DelegationConfigId - Delegation config to apply to the new credential.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
}

// IdentitiesCreateBody represents the identities create body type.
type IdentitiesCreateBody struct {
	// ActorId - Identity actor that will own the new identity.
	ActorId string `json:"actor_id"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the identity.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Credentials - Credentials to create and attach as part of identity creation.
	Credentials *[]IdentitiesCreateBodyCredentials `json:"credentials,omitempty"`
}

// MapIdentitiesCreateBodyFromJSON deserializes JSON data into a IdentitiesCreateBody.
func MapIdentitiesCreateBodyFromJSON(data []byte) (*IdentitiesCreateBody, error) {
	var v IdentitiesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesCreateBodyToJSON serializes a IdentitiesCreateBody to JSON.
func MapIdentitiesCreateBodyToJSON(v *IdentitiesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
