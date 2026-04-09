package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsRemoveServersOutput represents the magic mcp endpoints remove servers output type.
type MagicMcpEndpointsRemoveServersOutput struct {
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

// MapMagicMcpEndpointsRemoveServersOutputFromJSON deserializes JSON data into a MagicMcpEndpointsRemoveServersOutput.
func MapMagicMcpEndpointsRemoveServersOutputFromJSON(data []byte) (*MagicMcpEndpointsRemoveServersOutput, error) {
	var v MagicMcpEndpointsRemoveServersOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsRemoveServersOutputToJSON serializes a MagicMcpEndpointsRemoveServersOutput to JSON.
func MapMagicMcpEndpointsRemoveServersOutputToJSON(v *MagicMcpEndpointsRemoveServersOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpEndpointsRemoveServersBody represents the magic mcp endpoints remove servers body type.
type MagicMcpEndpointsRemoveServersBody struct {
	MagicMcpServerIds []string `json:"magic_mcp_server_ids"`
}

// MapMagicMcpEndpointsRemoveServersBodyFromJSON deserializes JSON data into a MagicMcpEndpointsRemoveServersBody.
func MapMagicMcpEndpointsRemoveServersBodyFromJSON(data []byte) (*MagicMcpEndpointsRemoveServersBody, error) {
	var v MagicMcpEndpointsRemoveServersBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsRemoveServersBodyToJSON serializes a MagicMcpEndpointsRemoveServersBody to JSON.
func MapMagicMcpEndpointsRemoveServersBodyToJSON(v *MagicMcpEndpointsRemoveServersBody) ([]byte, error) {
	return json.Marshal(v)
}
