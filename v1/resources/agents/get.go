package agents

import (
	"encoding/json"
	"time"
)

// AgentsGetOutput represents the agents get output type.
type AgentsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the agent connects and runs.
	Type string `json:"type"`
	// Status - Current lifecycle status of the agent. `archived` once you archive the agent.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ActorId - Identity actor linked to the agent.
	ActorId string `json:"actor_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapAgentsGetOutputFromJSON deserializes JSON data into a AgentsGetOutput.
func MapAgentsGetOutputFromJSON(data []byte) (*AgentsGetOutput, error) {
	var v AgentsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAgentsGetOutputToJSON serializes a AgentsGetOutput to JSON.
func MapAgentsGetOutputToJSON(v *AgentsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
