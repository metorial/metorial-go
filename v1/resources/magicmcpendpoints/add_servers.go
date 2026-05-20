package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsAddServersOutput represents the magic mcp endpoints add servers output type.
type MagicMcpEndpointsAddServersOutput struct {
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

// MapMagicMcpEndpointsAddServersOutputFromJSON deserializes JSON data into a MagicMcpEndpointsAddServersOutput.
func MapMagicMcpEndpointsAddServersOutputFromJSON(data []byte) (*MagicMcpEndpointsAddServersOutput, error) {
	var v MagicMcpEndpointsAddServersOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsAddServersOutputToJSON serializes a MagicMcpEndpointsAddServersOutput to JSON.
func MapMagicMcpEndpointsAddServersOutputToJSON(v *MagicMcpEndpointsAddServersOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpEndpointsAddServersBodyMagicMcpServers represents the magic mcp endpoints add servers body magic mcp servers type.
type MagicMcpEndpointsAddServersBodyMagicMcpServers struct {
	MagicMcpServerId string `json:"magic_mcp_server_id"`
	ToolFilters      *any   `json:"tool_filters,omitempty"`
}

// MagicMcpEndpointsAddServersBody represents the magic mcp endpoints add servers body type.
type MagicMcpEndpointsAddServersBody struct {
	MagicMcpServers *[]MagicMcpEndpointsAddServersBodyMagicMcpServers `json:"magic_mcp_servers,omitempty"`
}

// MapMagicMcpEndpointsAddServersBodyFromJSON deserializes JSON data into a MagicMcpEndpointsAddServersBody.
func MapMagicMcpEndpointsAddServersBodyFromJSON(data []byte) (*MagicMcpEndpointsAddServersBody, error) {
	var v MagicMcpEndpointsAddServersBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsAddServersBodyToJSON serializes a MagicMcpEndpointsAddServersBody to JSON.
func MapMagicMcpEndpointsAddServersBodyToJSON(v *MagicMcpEndpointsAddServersBody) ([]byte, error) {
	return json.Marshal(v)
}
