package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsGetOutput represents the magic mcp endpoints get output type.
type MagicMcpEndpointsGetOutput struct {
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

// MapMagicMcpEndpointsGetOutputFromJSON deserializes JSON data into a MagicMcpEndpointsGetOutput.
func MapMagicMcpEndpointsGetOutputFromJSON(data []byte) (*MagicMcpEndpointsGetOutput, error) {
	var v MagicMcpEndpointsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsGetOutputToJSON serializes a MagicMcpEndpointsGetOutput to JSON.
func MapMagicMcpEndpointsGetOutputToJSON(v *MagicMcpEndpointsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
