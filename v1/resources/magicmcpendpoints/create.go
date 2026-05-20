package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsCreateOutput represents the magic mcp endpoints create output type.
type MagicMcpEndpointsCreateOutput struct {
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

// MagicMcpEndpointsCreateBodyMagicMcpServers represents the magic mcp endpoints create body magic mcp servers type.
type MagicMcpEndpointsCreateBodyMagicMcpServers struct {
	MagicMcpServerId string `json:"magic_mcp_server_id"`
	ToolFilters      *any   `json:"tool_filters,omitempty"`
}

// MagicMcpEndpointsCreateBody represents the magic mcp endpoints create body type.
type MagicMcpEndpointsCreateBody struct {
	Name              *string                                       `json:"name,omitempty"`
	Description       *string                                       `json:"description,omitempty"`
	Metadata          *map[string]any                               `json:"metadata,omitempty"`
	ConsumerProfileId *string                                       `json:"consumer_profile_id,omitempty"`
	SkillPluginId     *string                                       `json:"skill_plugin_id,omitempty"`
	MagicMcpServers   *[]MagicMcpEndpointsCreateBodyMagicMcpServers `json:"magic_mcp_servers,omitempty"`
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
