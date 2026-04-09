package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsCreateOutput represents the magic mcp endpoints create output type.
type MagicMcpEndpointsCreateOutput struct {
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

// MapMagicMcpEndpointsCreateOutputFromJSON deserializes JSON data into a MagicMcpEndpointsCreateOutput.
func MapMagicMcpEndpointsCreateOutputFromJSON(data []byte) (*MagicMcpEndpointsCreateOutput, error) {
	var v MagicMcpEndpointsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsCreateOutputToJSON serializes a MagicMcpEndpointsCreateOutput to JSON.
func MapMagicMcpEndpointsCreateOutputToJSON(v *MagicMcpEndpointsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpEndpointsCreateBodyServers represents the magic mcp endpoints create body servers type.
type MagicMcpEndpointsCreateBodyServers struct {
	MagicMcpServerId string `json:"magic_mcp_server_id"`
	ToolFilters      *any   `json:"tool_filters,omitempty"`
}

// MagicMcpEndpointsCreateBody represents the magic mcp endpoints create body type.
type MagicMcpEndpointsCreateBody struct {
	Name              *string                               `json:"name,omitempty"`
	Description       *string                               `json:"description,omitempty"`
	Metadata          *map[string]any                       `json:"metadata,omitempty"`
	ConsumerProfileId *string                               `json:"consumer_profile_id,omitempty"`
	MagicMcpServerIds *[]string                             `json:"magic_mcp_server_ids,omitempty"`
	Servers           *[]MagicMcpEndpointsCreateBodyServers `json:"servers,omitempty"`
}

// MapMagicMcpEndpointsCreateBodyFromJSON deserializes JSON data into a MagicMcpEndpointsCreateBody.
func MapMagicMcpEndpointsCreateBodyFromJSON(data []byte) (*MagicMcpEndpointsCreateBody, error) {
	var v MagicMcpEndpointsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsCreateBodyToJSON serializes a MagicMcpEndpointsCreateBody to JSON.
func MapMagicMcpEndpointsCreateBodyToJSON(v *MagicMcpEndpointsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
