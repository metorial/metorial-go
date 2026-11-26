package identities

import (
	"encoding/json"
	"time"
)

// IdentitiesDeleteOutputOwnerActor represents the identities delete output owner actor type.
type IdentitiesDeleteOutputOwnerActor struct {
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

// IdentitiesDeleteOutputOwner - Actor that owns the identity.
type IdentitiesDeleteOutputOwner struct {
	// Type - Owner type for the identity.
	Type  string                           `json:"type"`
	Actor IdentitiesDeleteOutputOwnerActor `json:"actor"`
}

// IdentitiesDeleteOutputCredentials represents the identities delete output credentials type.
type IdentitiesDeleteOutputCredentials struct {
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

// IdentitiesDeleteOutput represents the identities delete output type.
type IdentitiesDeleteOutput struct {
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
	Owner IdentitiesDeleteOutputOwner `json:"owner"`
	// Credentials - Credentials currently attached to the identity.
	Credentials []IdentitiesDeleteOutputCredentials `json:"credentials"`
	// DelegationConfigId - Default delegation config applied to the identity. `null` when none is set.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapIdentitiesDeleteOutputFromJSON deserializes JSON data into a IdentitiesDeleteOutput.
func MapIdentitiesDeleteOutputFromJSON(data []byte) (*IdentitiesDeleteOutput, error) {
	var v IdentitiesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDeleteOutputToJSON serializes a IdentitiesDeleteOutput to JSON.
func MapIdentitiesDeleteOutputToJSON(v *IdentitiesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
