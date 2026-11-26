package magicmcptokens

import (
	"encoding/json"
	"time"
)

// MagicMcpTokensGetOutputServer represents the magic mcp tokens get output server type.
type MagicMcpTokensGetOutputServer struct {
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

// MagicMcpTokensGetOutputEndpoint represents the magic mcp tokens get output endpoint type.
type MagicMcpTokensGetOutputEndpoint struct {
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

// MagicMcpTokensGetOutputGroups represents the magic mcp tokens get output groups type.
type MagicMcpTokensGetOutputGroups struct {
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

// MagicMcpTokensGetOutput represents the magic mcp tokens get output type.
type MagicMcpTokensGetOutput struct {
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
	Description *string                          `json:"description,omitempty"`
	Server      *MagicMcpTokensGetOutputServer   `json:"server,omitempty"`
	Endpoint    *MagicMcpTokensGetOutputEndpoint `json:"endpoint,omitempty"`
	// Groups - Groups the token is scoped to. Empty means the token is not scoped.
	Groups []MagicMcpTokensGetOutputGroups `json:"groups"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpTokensGetOutputFromJSON deserializes JSON data into a MagicMcpTokensGetOutput.
func MapMagicMcpTokensGetOutputFromJSON(data []byte) (*MagicMcpTokensGetOutput, error) {
	var v MagicMcpTokensGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensGetOutputToJSON serializes a MagicMcpTokensGetOutput to JSON.
func MapMagicMcpTokensGetOutputToJSON(v *MagicMcpTokensGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
