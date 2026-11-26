package magicmcptokens

import (
	"encoding/json"
	"time"
)

// MagicMcpTokensListOutputItemsServer represents the magic mcp tokens list output items server type.
type MagicMcpTokensListOutputItemsServer struct {
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

// MagicMcpTokensListOutputItemsEndpoint represents the magic mcp tokens list output items endpoint type.
type MagicMcpTokensListOutputItemsEndpoint struct {
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

// MagicMcpTokensListOutputItemsGroups represents the magic mcp tokens list output items groups type.
type MagicMcpTokensListOutputItemsGroups struct {
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

// MagicMcpTokensListOutputItems represents the magic mcp tokens list output items type.
type MagicMcpTokensListOutputItems struct {
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
	Server      *MagicMcpTokensListOutputItemsServer   `json:"server,omitempty"`
	Endpoint    *MagicMcpTokensListOutputItemsEndpoint `json:"endpoint,omitempty"`
	// Groups - Groups the token is scoped to. Empty means the token is not scoped.
	Groups []MagicMcpTokensListOutputItemsGroups `json:"groups"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpTokensListOutputPagination represents the magic mcp tokens list output pagination type.
type MagicMcpTokensListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// MagicMcpTokensListOutput represents the magic mcp tokens list output type.
type MagicMcpTokensListOutput struct {
	// Object - String representing the object’s type.
	Object     string                             `json:"object"`
	Items      []MagicMcpTokensListOutputItems    `json:"items"`
	Pagination MagicMcpTokensListOutputPagination `json:"pagination"`
}

// MapMagicMcpTokensListOutputFromJSON deserializes JSON data into a MagicMcpTokensListOutput.
func MapMagicMcpTokensListOutputFromJSON(data []byte) (*MagicMcpTokensListOutput, error) {
	var v MagicMcpTokensListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensListOutputToJSON serializes a MagicMcpTokensListOutput to JSON.
func MapMagicMcpTokensListOutputToJSON(v *MagicMcpTokensListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpTokensListQuery represents the magic mcp tokens list query type.
type MagicMcpTokensListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Status - Only return tokens with one of these statuses.
	Status *any `json:"status,omitempty"`
	// MagicMcpGroupId - Only return tokens scoped to one of these groups.
	MagicMcpGroupId *any `json:"magic_mcp_group_id,omitempty"`
	// MagicMcpServerId - Only return tokens bound to one of these servers.
	MagicMcpServerId *any `json:"magic_mcp_server_id,omitempty"`
	// MagicMcpEndpointId - Only return tokens bound to one of these endpoints.
	MagicMcpEndpointId *any `json:"magic_mcp_endpoint_id,omitempty"`
}

// MapMagicMcpTokensListQueryFromJSON deserializes JSON data into a MagicMcpTokensListQuery.
func MapMagicMcpTokensListQueryFromJSON(data []byte) (*MagicMcpTokensListQuery, error) {
	var v MagicMcpTokensListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpTokensListQueryToJSON serializes a MagicMcpTokensListQuery to JSON.
func MapMagicMcpTokensListQueryToJSON(v *MagicMcpTokensListQuery) ([]byte, error) {
	return json.Marshal(v)
}
