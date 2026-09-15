package provider

import (
	"encoding/json"
)

// ChatInstancesProviderAuthenticatedUserOutputWorkspace - The workspace the authenticated account belongs to, once one could be resolved
type ChatInstancesProviderAuthenticatedUserOutputWorkspace struct {
	// Id - Unique chat workspace identifier
	Id string `json:"id"`
	// ProviderWorkspaceId - The workspace's identifier on the chat provider
	ProviderWorkspaceId string `json:"provider_workspace_id"`
	// Name - Display name of the workspace
	Name *string `json:"name,omitempty"`
	// Domain - Domain the workspace is reachable under
	Domain *string `json:"domain,omitempty"`
	// ImageUrl - URL of the workspace's icon
	ImageUrl *string `json:"image_url,omitempty"`
}

// ChatInstancesProviderAuthenticatedUserOutput represents the chat instances provider authenticated user output type.
type ChatInstancesProviderAuthenticatedUserOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat author identifier. Null when Metorial has not yet persisted this author, which happens when the provider does not return a workspace for the authenticated account.
	Id *string `json:"id,omitempty"`
	// ChatId - The chat this author belongs to, once persisted
	ChatId *string `json:"chat_id,omitempty"`
	// Type - What kind of account this author is on the chat provider
	Type string `json:"type"`
	// Role - The role this author holds in the chat workspace
	Role string `json:"role"`
	// ProviderType - The provider's own name for this kind of author
	ProviderType *string `json:"provider_type,omitempty"`
	// ProviderAuthorId - The author's identifier on the chat provider
	ProviderAuthorId string `json:"provider_author_id"`
	// UserName - The author's handle on the chat provider
	UserName string `json:"user_name"`
	// FullName - The author's display name
	FullName string `json:"full_name"`
	// Email - Email address of the author, if the provider exposes one
	Email *string `json:"email,omitempty"`
	// ImageUrl - URL of the author's avatar, if the provider exposes one
	ImageUrl *string `json:"image_url,omitempty"`
	// IsSelf - Whether this author is the account the integration itself is authenticated as
	IsSelf bool `json:"is_self"`
	// Workspace - The workspace the authenticated account belongs to, once one could be resolved
	Workspace *ChatInstancesProviderAuthenticatedUserOutputWorkspace `json:"workspace,omitempty"`
}

// MapChatInstancesProviderAuthenticatedUserOutputFromJSON deserializes JSON data into a ChatInstancesProviderAuthenticatedUserOutput.
func MapChatInstancesProviderAuthenticatedUserOutputFromJSON(data []byte) (*ChatInstancesProviderAuthenticatedUserOutput, error) {
	var v ChatInstancesProviderAuthenticatedUserOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatInstancesProviderAuthenticatedUserOutputToJSON serializes a ChatInstancesProviderAuthenticatedUserOutput to JSON.
func MapChatInstancesProviderAuthenticatedUserOutputToJSON(v *ChatInstancesProviderAuthenticatedUserOutput) ([]byte, error) {
	return json.Marshal(v)
}
