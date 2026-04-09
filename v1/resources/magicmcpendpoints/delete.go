package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsDeleteOutput represents the magic mcp endpoints delete output type.
type MagicMcpEndpointsDeleteOutput struct {
	Object            string           `json:"object"`
	Id                string           `json:"id"`
	Status            string           `json:"status"`
	Slug              string           `json:"slug"`
	Url               string           `json:"url"`
	ConsumerProfileId *string          `json:"consumer_profile_id,omitempty"`
	SessionTemplateId *string          `json:"session_template_id,omitempty"`
	SessionId         *string          `json:"session_id,omitempty"`
	Servers           []map[string]any `json:"servers"`
	Name              *string          `json:"name,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Metadata          map[string]any   `json:"metadata"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
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
