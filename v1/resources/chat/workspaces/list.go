package workspaces

import (
	"encoding/json"
	"time"
)

// ChatWorkspacesListOutputItems represents the chat workspaces list output items type.
type ChatWorkspacesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ChatId - The chat this workspace belongs to
	ChatId string `json:"chat_id"`
	// ProviderWorkspaceId - Identifier the chat provider uses for this workspace.
	ProviderWorkspaceId string `json:"provider_workspace_id"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Domain - Domain the workspace is reachable under on the chat provider
	Domain *string `json:"domain,omitempty"`
	// ImageUrl - URL of the workspace's icon, if the provider exposes one
	ImageUrl *string `json:"image_url,omitempty"`
	// CreatedAt - Timestamp when the workspace was first seen by Metorial
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the workspace was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatWorkspacesListOutputPagination represents the chat workspaces list output pagination type.
type ChatWorkspacesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ChatWorkspacesListOutput represents the chat workspaces list output type.
type ChatWorkspacesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                             `json:"object"`
	Items      []ChatWorkspacesListOutputItems    `json:"items"`
	Pagination ChatWorkspacesListOutputPagination `json:"pagination"`
}

// MapChatWorkspacesListOutputFromJSON deserializes JSON data into a ChatWorkspacesListOutput.
func MapChatWorkspacesListOutputFromJSON(data []byte) (*ChatWorkspacesListOutput, error) {
	var v ChatWorkspacesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatWorkspacesListOutputToJSON serializes a ChatWorkspacesListOutput to JSON.
func MapChatWorkspacesListOutputToJSON(v *ChatWorkspacesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatWorkspacesListQuery represents the chat workspaces list query type.
type ChatWorkspacesListQuery struct {
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
	// ChatInstanceId - The chat instance to list workspaces for
	ChatInstanceId string `json:"chat_instance_id"`
	// Search - Search by workspace name
	Search *string `json:"search,omitempty"`
}

// MapChatWorkspacesListQueryFromJSON deserializes JSON data into a ChatWorkspacesListQuery.
func MapChatWorkspacesListQueryFromJSON(data []byte) (*ChatWorkspacesListQuery, error) {
	var v ChatWorkspacesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatWorkspacesListQueryToJSON serializes a ChatWorkspacesListQuery to JSON.
func MapChatWorkspacesListQueryToJSON(v *ChatWorkspacesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
