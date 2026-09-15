package workspaces

import (
	"encoding/json"
	"time"
)

// ChatWorkspacesGetOutput represents the chat workspaces get output type.
type ChatWorkspacesGetOutput struct {
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

// MapChatWorkspacesGetOutputFromJSON deserializes JSON data into a ChatWorkspacesGetOutput.
func MapChatWorkspacesGetOutputFromJSON(data []byte) (*ChatWorkspacesGetOutput, error) {
	var v ChatWorkspacesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatWorkspacesGetOutputToJSON serializes a ChatWorkspacesGetOutput to JSON.
func MapChatWorkspacesGetOutputToJSON(v *ChatWorkspacesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatWorkspacesGetQuery represents the chat workspaces get query type.
type ChatWorkspacesGetQuery struct {
	// ChatInstanceId - The chat instance this workspace belongs to
	ChatInstanceId string `json:"chat_instance_id"`
}

// MapChatWorkspacesGetQueryFromJSON deserializes JSON data into a ChatWorkspacesGetQuery.
func MapChatWorkspacesGetQueryFromJSON(data []byte) (*ChatWorkspacesGetQuery, error) {
	var v ChatWorkspacesGetQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatWorkspacesGetQueryToJSON serializes a ChatWorkspacesGetQuery to JSON.
func MapChatWorkspacesGetQueryToJSON(v *ChatWorkspacesGetQuery) ([]byte, error) {
	return json.Marshal(v)
}
