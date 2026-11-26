package magicmcptokens

import (
	"encoding/json"
	"time"
)

// MagicMcpTokensDeleteOutputServer represents the magic mcp tokens delete output server type.
type MagicMcpTokensDeleteOutputServer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the server is deleted.
	Status string `json:"status"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
}

// MagicMcpTokensDeleteOutputEndpoint represents the magic mcp tokens delete output endpoint type.
type MagicMcpTokensDeleteOutputEndpoint struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the endpoint is deleted.
	Status string `json:"status"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
}

// MagicMcpTokensDeleteOutputGroups represents the magic mcp tokens delete output groups type.
type MagicMcpTokensDeleteOutputGroups struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `deleted` once you delete the group.
	Status string `json:"status"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpTokensDeleteOutput represents the magic mcp tokens delete output type.
type MagicMcpTokensDeleteOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `deleted` once you delete the token.
	Status string `json:"status"`
	// Secret - Bearer token used against the connection API.
	Secret string `json:"secret"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string                             `json:"description,omitempty"`
	Server      *MagicMcpTokensDeleteOutputServer   `json:"server,omitempty"`
	Endpoint    *MagicMcpTokensDeleteOutputEndpoint `json:"endpoint,omitempty"`
	// Groups - Groups the token is scoped to. Empty means the token is not scoped.
	Groups []MagicMcpTokensDeleteOutputGroups `json:"groups"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
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
