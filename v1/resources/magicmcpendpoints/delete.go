package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsDeleteOutput represents the magic mcp endpoints delete output type.
type MagicMcpEndpointsDeleteOutput struct {
	Object      string           `json:"object"`
	Id          string           `json:"id"`
	Status      string           `json:"status"`
	Slug        string           `json:"slug"`
	Url         string           `json:"url"`
	Servers     []map[string]any `json:"servers"`
	Name        *string          `json:"name,omitempty"`
	Description *string          `json:"description,omitempty"`
	Metadata    map[string]any   `json:"metadata"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

// MapMagicMcpEndpointsDeleteOutputFromJSON deserializes JSON data into a MagicMcpEndpointsDeleteOutput.
func MapMagicMcpEndpointsDeleteOutputFromJSON(data []byte) (*MagicMcpEndpointsDeleteOutput, error) {
	var v MagicMcpEndpointsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsDeleteOutputToJSON serializes a MagicMcpEndpointsDeleteOutput to JSON.
func MapMagicMcpEndpointsDeleteOutputToJSON(v *MagicMcpEndpointsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
