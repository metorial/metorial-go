package magicmcpgroups

import (
	"encoding/json"
	"time"
)

// MagicMcpGroupsCreateOutput represents the magic mcp groups create output type.
type MagicMcpGroupsCreateOutput struct {
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

// MapMagicMcpGroupsCreateOutputFromJSON deserializes JSON data into a MagicMcpGroupsCreateOutput.
func MapMagicMcpGroupsCreateOutputFromJSON(data []byte) (*MagicMcpGroupsCreateOutput, error) {
	var v MagicMcpGroupsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpGroupsCreateOutputToJSON serializes a MagicMcpGroupsCreateOutput to JSON.
func MapMagicMcpGroupsCreateOutputToJSON(v *MagicMcpGroupsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpGroupsCreateBody represents the magic mcp groups create body type.
type MagicMcpGroupsCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the group. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapMagicMcpGroupsCreateBodyFromJSON deserializes JSON data into a MagicMcpGroupsCreateBody.
func MapMagicMcpGroupsCreateBodyFromJSON(data []byte) (*MagicMcpGroupsCreateBody, error) {
	var v MagicMcpGroupsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpGroupsCreateBodyToJSON serializes a MagicMcpGroupsCreateBody to JSON.
func MapMagicMcpGroupsCreateBodyToJSON(v *MagicMcpGroupsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
