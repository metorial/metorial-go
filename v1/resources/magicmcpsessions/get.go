package magicmcpsessions

import (
	"encoding/json"
	"time"
)

// MagicMcpSessionsGetOutputMagicMcpServerEndpoints represents the magic mcp sessions get output magic mcp server endpoints type.
type MagicMcpSessionsGetOutputMagicMcpServerEndpoints struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
	Url   string `json:"url"`
}

// MagicMcpSessionsGetOutputMagicMcpServer represents the magic mcp sessions get output magic mcp server type.
type MagicMcpSessionsGetOutputMagicMcpServer struct {
	Object             string                                             `json:"object"`
	Id                 string                                             `json:"id"`
	Status             string                                             `json:"status"`
	Source             string                                             `json:"source"`
	ProviderTemplateId *string                                            `json:"provider_template_id,omitempty"`
	Endpoints          []MagicMcpSessionsGetOutputMagicMcpServerEndpoints `json:"endpoints"`
	Name               *string                                            `json:"name,omitempty"`
	Description        *string                                            `json:"description,omitempty"`
	Metadata           map[string]any                                     `json:"metadata"`
	CreatedAt          time.Time                                          `json:"created_at"`
	UpdatedAt          time.Time                                          `json:"updated_at"`
}

// MagicMcpSessionsGetOutputMagicMcpEndpoint represents the magic mcp sessions get output magic mcp endpoint type.
type MagicMcpSessionsGetOutputMagicMcpEndpoint struct {
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

// MagicMcpSessionsGetOutput represents the magic mcp sessions get output type.
type MagicMcpSessionsGetOutput struct {
	Object           string                                     `json:"object"`
	Id               string                                     `json:"id"`
	MagicMcpServer   *MagicMcpSessionsGetOutputMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	MagicMcpEndpoint *MagicMcpSessionsGetOutputMagicMcpEndpoint `json:"magic_mcp_endpoint,omitempty"`
	SessionId        string                                     `json:"session_id"`
	CreatedAt        time.Time                                  `json:"created_at"`
	UpdatedAt        time.Time                                  `json:"updated_at"`
}

// MapMagicMcpSessionsGetOutputFromJSON deserializes JSON data into a MagicMcpSessionsGetOutput.
func MapMagicMcpSessionsGetOutputFromJSON(data []byte) (*MagicMcpSessionsGetOutput, error) {
	var v MagicMcpSessionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpSessionsGetOutputToJSON serializes a MagicMcpSessionsGetOutput to JSON.
func MapMagicMcpSessionsGetOutputToJSON(v *MagicMcpSessionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
