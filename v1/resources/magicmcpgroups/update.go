package magicmcpgroups

import (
	"encoding/json"
	"time"
)

// MagicMcpGroupsUpdateOutput represents the magic mcp groups update output type.
type MagicMcpGroupsUpdateOutput struct {
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

// MapMagicMcpGroupsUpdateOutputFromJSON deserializes JSON data into a MagicMcpGroupsUpdateOutput.
func MapMagicMcpGroupsUpdateOutputFromJSON(data []byte) (*MagicMcpGroupsUpdateOutput, error) {
	var v MagicMcpGroupsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpGroupsUpdateOutputToJSON serializes a MagicMcpGroupsUpdateOutput to JSON.
func MapMagicMcpGroupsUpdateOutputToJSON(v *MagicMcpGroupsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpGroupsUpdateBody represents the magic mcp groups update body type.
type MagicMcpGroupsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the group. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapMagicMcpGroupsUpdateBodyFromJSON deserializes JSON data into a MagicMcpGroupsUpdateBody.
func MapMagicMcpGroupsUpdateBodyFromJSON(data []byte) (*MagicMcpGroupsUpdateBody, error) {
	var v MagicMcpGroupsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpGroupsUpdateBodyToJSON serializes a MagicMcpGroupsUpdateBody to JSON.
func MapMagicMcpGroupsUpdateBodyToJSON(v *MagicMcpGroupsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
