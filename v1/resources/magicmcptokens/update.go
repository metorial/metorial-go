package magicmcptokens

import (
	"encoding/json"
	"time"
)

// MagicMcpTokensUpdateOutputServer represents the magic mcp tokens update output server type.
type MagicMcpTokensUpdateOutputServer struct {
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

// MagicMcpTokensUpdateOutputEndpoint represents the magic mcp tokens update output endpoint type.
type MagicMcpTokensUpdateOutputEndpoint struct {
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

// MagicMcpTokensUpdateOutputGroups represents the magic mcp tokens update output groups type.
type MagicMcpTokensUpdateOutputGroups struct {
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

// MagicMcpTokensUpdateOutput represents the magic mcp tokens update output type.
type MagicMcpTokensUpdateOutput struct {
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
	Server      *MagicMcpTokensUpdateOutputServer   `json:"server,omitempty"`
	Endpoint    *MagicMcpTokensUpdateOutputEndpoint `json:"endpoint,omitempty"`
	// Groups - Groups the token is scoped to. Empty means the token is not scoped.
	Groups []MagicMcpTokensUpdateOutputGroups `json:"groups"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpTokensUpdateOutputFromJSON deserializes JSON data into a MagicMcpTokensUpdateOutput.
func MapMagicMcpTokensUpdateOutputFromJSON(data []byte) (*MagicMcpTokensUpdateOutput, error) {
	var v MagicMcpTokensUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensUpdateOutputToJSON serializes a MagicMcpTokensUpdateOutput to JSON.
func MapMagicMcpTokensUpdateOutputToJSON(v *MagicMcpTokensUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpTokensUpdateBody represents the magic mcp tokens update body type.
type MagicMcpTokensUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the token. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapMagicMcpTokensUpdateBodyFromJSON deserializes JSON data into a MagicMcpTokensUpdateBody.
func MapMagicMcpTokensUpdateBodyFromJSON(data []byte) (*MagicMcpTokensUpdateBody, error) {
	var v MagicMcpTokensUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensUpdateBodyToJSON serializes a MagicMcpTokensUpdateBody to JSON.
func MapMagicMcpTokensUpdateBodyToJSON(v *MagicMcpTokensUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
