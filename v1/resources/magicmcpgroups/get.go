package magicmcpgroups

import (
	"encoding/json"
	"time"
)

// MagicMcpGroupsGetOutput represents the magic mcp groups get output type.
type MagicMcpGroupsGetOutput struct {
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

// MapMagicMcpGroupsGetOutputFromJSON deserializes JSON data into a MagicMcpGroupsGetOutput.
func MapMagicMcpGroupsGetOutputFromJSON(data []byte) (*MagicMcpGroupsGetOutput, error) {
	var v MagicMcpGroupsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpGroupsGetOutputToJSON serializes a MagicMcpGroupsGetOutput to JSON.
func MapMagicMcpGroupsGetOutputToJSON(v *MagicMcpGroupsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
