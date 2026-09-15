package members

import (
	"encoding/json"
	"time"
)

// ChatsChannelsMembersGetOutput represents the chats channels members get output type.
type ChatsChannelsMembersGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat author identifier
	Id string `json:"id"`
	// ChatId - The chat this author belongs to
	ChatId string `json:"chat_id"`
	// Type - What kind of account this author is on the chat provider
	Type string `json:"type"`
	// Role - The role this author holds in the chat workspace
	Role string `json:"role"`
	// ProviderType - The provider's own name for this kind of author
	ProviderType string `json:"provider_type"`
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
	// CreatedAt - Timestamp when the author was first seen by Metorial
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the author was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// LastInteractionAt - Timestamp of the last activity Metorial saw from this author
	LastInteractionAt *time.Time `json:"last_interaction_at,omitempty"`
}

// MapChatsChannelsMembersGetOutputFromJSON deserializes JSON data into a ChatsChannelsMembersGetOutput.
func MapChatsChannelsMembersGetOutputFromJSON(data []byte) (*ChatsChannelsMembersGetOutput, error) {
	var v ChatsChannelsMembersGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsChannelsMembersGetOutputToJSON serializes a ChatsChannelsMembersGetOutput to JSON.
func MapChatsChannelsMembersGetOutputToJSON(v *ChatsChannelsMembersGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
