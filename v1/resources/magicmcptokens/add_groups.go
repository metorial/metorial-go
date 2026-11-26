package magicmcptokens

import (
	"encoding/json"
	"time"
)

// MagicMcpTokensAddGroupsOutputServer represents the magic mcp tokens add groups output server type.
type MagicMcpTokensAddGroupsOutputServer struct {
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

// MagicMcpTokensAddGroupsOutputEndpoint represents the magic mcp tokens add groups output endpoint type.
type MagicMcpTokensAddGroupsOutputEndpoint struct {
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

// MagicMcpTokensAddGroupsOutputGroups represents the magic mcp tokens add groups output groups type.
type MagicMcpTokensAddGroupsOutputGroups struct {
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

// MagicMcpTokensAddGroupsOutput represents the magic mcp tokens add groups output type.
type MagicMcpTokensAddGroupsOutput struct {
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
	Description *string                                `json:"description,omitempty"`
	Server      *MagicMcpTokensAddGroupsOutputServer   `json:"server,omitempty"`
	Endpoint    *MagicMcpTokensAddGroupsOutputEndpoint `json:"endpoint,omitempty"`
	// Groups - Groups the token is scoped to. Empty means the token is not scoped.
	Groups []MagicMcpTokensAddGroupsOutputGroups `json:"groups"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpTokensAddGroupsOutputFromJSON deserializes JSON data into a MagicMcpTokensAddGroupsOutput.
func MapMagicMcpTokensAddGroupsOutputFromJSON(data []byte) (*MagicMcpTokensAddGroupsOutput, error) {
	var v MagicMcpTokensAddGroupsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensAddGroupsOutputToJSON serializes a MagicMcpTokensAddGroupsOutput to JSON.
func MapMagicMcpTokensAddGroupsOutputToJSON(v *MagicMcpTokensAddGroupsOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpTokensAddGroupsBody represents the magic mcp tokens add groups body type.
type MagicMcpTokensAddGroupsBody struct {
	// MagicMcpGroupIds - Groups to scope the token to.
	MagicMcpGroupIds []string `json:"magic_mcp_group_ids"`
}

// MapMagicMcpTokensAddGroupsBodyFromJSON deserializes JSON data into a MagicMcpTokensAddGroupsBody.
func MapMagicMcpTokensAddGroupsBodyFromJSON(data []byte) (*MagicMcpTokensAddGroupsBody, error) {
	var v MagicMcpTokensAddGroupsBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensAddGroupsBodyToJSON serializes a MagicMcpTokensAddGroupsBody to JSON.
func MapMagicMcpTokensAddGroupsBodyToJSON(v *MagicMcpTokensAddGroupsBody) ([]byte, error) {
	return json.Marshal(v)
}
