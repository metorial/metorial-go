package magicmcptokens

import (
	"encoding/json"
	"time"
)

// MagicMcpTokensCreateOutputServer represents the magic mcp tokens create output server type.
type MagicMcpTokensCreateOutputServer struct {
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

// MagicMcpTokensCreateOutputEndpoint represents the magic mcp tokens create output endpoint type.
type MagicMcpTokensCreateOutputEndpoint struct {
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

// MagicMcpTokensCreateOutputGroups represents the magic mcp tokens create output groups type.
type MagicMcpTokensCreateOutputGroups struct {
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

// MagicMcpTokensCreateOutput represents the magic mcp tokens create output type.
type MagicMcpTokensCreateOutput struct {
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
	Server      *MagicMcpTokensCreateOutputServer   `json:"server,omitempty"`
	Endpoint    *MagicMcpTokensCreateOutputEndpoint `json:"endpoint,omitempty"`
	// Groups - Groups the token is scoped to. Empty means the token is not scoped.
	Groups []MagicMcpTokensCreateOutputGroups `json:"groups"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpTokensCreateOutputFromJSON deserializes JSON data into a MagicMcpTokensCreateOutput.
func MapMagicMcpTokensCreateOutputFromJSON(data []byte) (*MagicMcpTokensCreateOutput, error) {
	var v MagicMcpTokensCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensCreateOutputToJSON serializes a MagicMcpTokensCreateOutput to JSON.
func MapMagicMcpTokensCreateOutputToJSON(v *MagicMcpTokensCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpTokensCreateBody represents the magic mcp tokens create body type.
type MagicMcpTokensCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the token. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// MagicMcpGroupIds - Groups to scope the token to.
	MagicMcpGroupIds *[]string `json:"magic_mcp_group_ids,omitempty"`
	// MagicMcpServerId - Server to bind the token to.
	MagicMcpServerId *string `json:"magic_mcp_server_id,omitempty"`
	// MagicMcpEndpointId - Endpoint to bind the token to.
	MagicMcpEndpointId *string `json:"magic_mcp_endpoint_id,omitempty"`
}

// MapMagicMcpTokensCreateBodyFromJSON deserializes JSON data into a MagicMcpTokensCreateBody.
func MapMagicMcpTokensCreateBodyFromJSON(data []byte) (*MagicMcpTokensCreateBody, error) {
	var v MagicMcpTokensCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensCreateBodyToJSON serializes a MagicMcpTokensCreateBody to JSON.
func MapMagicMcpTokensCreateBodyToJSON(v *MagicMcpTokensCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
