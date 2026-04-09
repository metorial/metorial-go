package magicmcptokens

import (
	"encoding/json"
	"time"
)

// MagicMcpTokensDeleteOutputServer represents the magic mcp tokens delete output server type.
type MagicMcpTokensDeleteOutputServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// MagicMcpTokensDeleteOutputEndpoint represents the magic mcp tokens delete output endpoint type.
type MagicMcpTokensDeleteOutputEndpoint struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Slug        string  `json:"slug"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// MagicMcpTokensDeleteOutputGroups represents the magic mcp tokens delete output groups type.
type MagicMcpTokensDeleteOutputGroups struct {
	Object      string         `json:"object"`
	Id          string         `json:"id"`
	Status      string         `json:"status"`
	Slug        string         `json:"slug"`
	Name        *string        `json:"name,omitempty"`
	Description *string        `json:"description,omitempty"`
	Metadata    map[string]any `json:"metadata"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// MagicMcpTokensDeleteOutput represents the magic mcp tokens delete output type.
type MagicMcpTokensDeleteOutput struct {
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Status      string                              `json:"status"`
	Secret      string                              `json:"secret"`
	Name        *string                             `json:"name,omitempty"`
	Description *string                             `json:"description,omitempty"`
	Server      *MagicMcpTokensDeleteOutputServer   `json:"server,omitempty"`
	Endpoint    *MagicMcpTokensDeleteOutputEndpoint `json:"endpoint,omitempty"`
	Groups      []MagicMcpTokensDeleteOutputGroups  `json:"groups"`
	Metadata    map[string]any                      `json:"metadata"`
	CreatedAt   time.Time                           `json:"created_at"`
	UpdatedAt   time.Time                           `json:"updated_at"`
}

// MapMagicMcpTokensDeleteOutputFromJSON deserializes JSON data into a MagicMcpTokensDeleteOutput.
func MapMagicMcpTokensDeleteOutputFromJSON(data []byte) (*MagicMcpTokensDeleteOutput, error) {
	var v MagicMcpTokensDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensDeleteOutputToJSON serializes a MagicMcpTokensDeleteOutput to JSON.
func MapMagicMcpTokensDeleteOutputToJSON(v *MagicMcpTokensDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
