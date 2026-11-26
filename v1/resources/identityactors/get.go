package identityactors

import (
	"encoding/json"
	"time"
)

// IdentityActorsGetOutput represents the identity actors get output type.
type IdentityActorsGetOutput struct {
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

// MapIdentityActorsGetOutputFromJSON deserializes JSON data into a IdentityActorsGetOutput.
func MapIdentityActorsGetOutputFromJSON(data []byte) (*IdentityActorsGetOutput, error) {
	var v IdentityActorsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentityActorsGetOutputToJSON serializes a IdentityActorsGetOutput to JSON.
func MapIdentityActorsGetOutputToJSON(v *IdentityActorsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
