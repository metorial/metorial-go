package identityactors

import (
	"encoding/json"
	"time"
)

// IdentityActorsCreateOutput represents the identity actors create output type.
type IdentityActorsCreateOutput struct {
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

// MapIdentityActorsCreateOutputFromJSON deserializes JSON data into a IdentityActorsCreateOutput.
func MapIdentityActorsCreateOutputFromJSON(data []byte) (*IdentityActorsCreateOutput, error) {
	var v IdentityActorsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentityActorsCreateOutputToJSON serializes a IdentityActorsCreateOutput to JSON.
func MapIdentityActorsCreateOutputToJSON(v *IdentityActorsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentityActorsCreateBody represents the identity actors create body type.
type IdentityActorsCreateBody struct {
	// Type - Whether this actor is a person or an agent.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the identity actor.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapIdentityActorsCreateBodyFromJSON deserializes JSON data into a IdentityActorsCreateBody.
func MapIdentityActorsCreateBodyFromJSON(data []byte) (*IdentityActorsCreateBody, error) {
	var v IdentityActorsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentityActorsCreateBodyToJSON serializes a IdentityActorsCreateBody to JSON.
func MapIdentityActorsCreateBodyToJSON(v *IdentityActorsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
