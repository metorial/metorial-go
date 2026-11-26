package identityactors

import (
	"encoding/json"
	"time"
)

// IdentityActorsUpdateOutput represents the identity actors update output type.
type IdentityActorsUpdateOutput struct {
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

// MapIdentityActorsUpdateOutputFromJSON deserializes JSON data into a IdentityActorsUpdateOutput.
func MapIdentityActorsUpdateOutputFromJSON(data []byte) (*IdentityActorsUpdateOutput, error) {
	var v IdentityActorsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentityActorsUpdateOutputToJSON serializes a IdentityActorsUpdateOutput to JSON.
func MapIdentityActorsUpdateOutputToJSON(v *IdentityActorsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentityActorsUpdateBody represents the identity actors update body type.
type IdentityActorsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the identity actor. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapIdentityActorsUpdateBodyFromJSON deserializes JSON data into a IdentityActorsUpdateBody.
func MapIdentityActorsUpdateBodyFromJSON(data []byte) (*IdentityActorsUpdateBody, error) {
	var v IdentityActorsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentityActorsUpdateBodyToJSON serializes a IdentityActorsUpdateBody to JSON.
func MapIdentityActorsUpdateBodyToJSON(v *IdentityActorsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
