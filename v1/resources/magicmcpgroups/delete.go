package magicmcpgroups

import (
	"encoding/json"
	"time"
)

// MagicMcpGroupsDeleteOutput represents the magic mcp groups delete output type.
type MagicMcpGroupsDeleteOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `deleted` once you delete the group.
	Status string `json:"status"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpGroupsDeleteOutputFromJSON deserializes JSON data into a MagicMcpGroupsDeleteOutput.
func MapMagicMcpGroupsDeleteOutputFromJSON(data []byte) (*MagicMcpGroupsDeleteOutput, error) {
	var v MagicMcpGroupsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpGroupsDeleteOutputToJSON serializes a MagicMcpGroupsDeleteOutput to JSON.
func MapMagicMcpGroupsDeleteOutputToJSON(v *MagicMcpGroupsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
