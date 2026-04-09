package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsUpdateOutput represents the magic mcp endpoints update output type.
type MagicMcpEndpointsUpdateOutput struct {
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

// MapMagicMcpEndpointsUpdateOutputFromJSON deserializes JSON data into a MagicMcpEndpointsUpdateOutput.
func MapMagicMcpEndpointsUpdateOutputFromJSON(data []byte) (*MagicMcpEndpointsUpdateOutput, error) {
	var v MagicMcpEndpointsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsUpdateOutputToJSON serializes a MagicMcpEndpointsUpdateOutput to JSON.
func MapMagicMcpEndpointsUpdateOutputToJSON(v *MagicMcpEndpointsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpEndpointsUpdateBody represents the magic mcp endpoints update body type.
type MagicMcpEndpointsUpdateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

// MapMagicMcpEndpointsUpdateBodyFromJSON deserializes JSON data into a MagicMcpEndpointsUpdateBody.
func MapMagicMcpEndpointsUpdateBodyFromJSON(data []byte) (*MagicMcpEndpointsUpdateBody, error) {
	var v MagicMcpEndpointsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsUpdateBodyToJSON serializes a MagicMcpEndpointsUpdateBody to JSON.
func MapMagicMcpEndpointsUpdateBodyToJSON(v *MagicMcpEndpointsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
