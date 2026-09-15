package workspaces

import (
	"encoding/json"
	"time"
)

// ChatWorkspacesListOutputItems represents the chat workspaces list output items type.
type ChatWorkspacesListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat workspace identifier
	Id string `json:"id"`
	// ChatId - The chat this workspace belongs to
	ChatId string `json:"chat_id"`
	// ProviderWorkspaceId - The workspace's identifier on the chat provider
	ProviderWorkspaceId string `json:"provider_workspace_id"`
	// Name - Display name of the workspace
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
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ChatWorkspacesListOutput represents the chat workspaces list output type.
type ChatWorkspacesListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
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
