package magicmcptokens

import (
	"encoding/json"
	"time"
)

// MagicMcpTokensRemoveGroupsOutputServer represents the magic mcp tokens remove groups output server type.
type MagicMcpTokensRemoveGroupsOutputServer struct {
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

// MagicMcpTokensRemoveGroupsOutputEndpoint represents the magic mcp tokens remove groups output endpoint type.
type MagicMcpTokensRemoveGroupsOutputEndpoint struct {
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

// MagicMcpTokensRemoveGroupsOutputGroups represents the magic mcp tokens remove groups output groups type.
type MagicMcpTokensRemoveGroupsOutputGroups struct {
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

// MagicMcpTokensRemoveGroupsOutput represents the magic mcp tokens remove groups output type.
type MagicMcpTokensRemoveGroupsOutput struct {
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
	Description *string                                   `json:"description,omitempty"`
	Server      *MagicMcpTokensRemoveGroupsOutputServer   `json:"server,omitempty"`
	Endpoint    *MagicMcpTokensRemoveGroupsOutputEndpoint `json:"endpoint,omitempty"`
	// Groups - Groups the token is scoped to. Empty means the token is not scoped.
	Groups []MagicMcpTokensRemoveGroupsOutputGroups `json:"groups"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpTokensRemoveGroupsOutputFromJSON deserializes JSON data into a MagicMcpTokensRemoveGroupsOutput.
func MapMagicMcpTokensRemoveGroupsOutputFromJSON(data []byte) (*MagicMcpTokensRemoveGroupsOutput, error) {
	var v MagicMcpTokensRemoveGroupsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensRemoveGroupsOutputToJSON serializes a MagicMcpTokensRemoveGroupsOutput to JSON.
func MapMagicMcpTokensRemoveGroupsOutputToJSON(v *MagicMcpTokensRemoveGroupsOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpTokensRemoveGroupsBody represents the magic mcp tokens remove groups body type.
type MagicMcpTokensRemoveGroupsBody struct {
	// MagicMcpGroupIds - Groups to remove from the token.
	MagicMcpGroupIds []string `json:"magic_mcp_group_ids"`
}

// MapMagicMcpTokensRemoveGroupsBodyFromJSON deserializes JSON data into a MagicMcpTokensRemoveGroupsBody.
func MapMagicMcpTokensRemoveGroupsBodyFromJSON(data []byte) (*MagicMcpTokensRemoveGroupsBody, error) {
	var v MagicMcpTokensRemoveGroupsBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensRemoveGroupsBodyToJSON serializes a MagicMcpTokensRemoveGroupsBody to JSON.
func MapMagicMcpTokensRemoveGroupsBodyToJSON(v *MagicMcpTokensRemoveGroupsBody) ([]byte, error) {
	return json.Marshal(v)
}
