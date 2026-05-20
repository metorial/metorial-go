package agents

import (
	"encoding/json"
	"time"
)

// AgentsGetOutput represents the agents get output type.
type AgentsGetOutput struct {
	Object      string          `json:"object"`
	Id          string          `json:"id"`
	Type        string          `json:"type"`
	Status      string          `json:"status"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Slug        string          `json:"slug"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	ActorId     string          `json:"actor_id"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	ArchivedAt  *time.Time      `json:"archived_at,omitempty"`
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
