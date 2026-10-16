package channels

import (
	"encoding/json"
	"time"
)

// ChatsChannelsGetOutputRecipient represents the chats channels get output recipient type.
type ChatsChannelsGetOutputRecipient struct {
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

// ChatsChannelsGetOutputContextAuthor represents the chats channels get output context author type.
type ChatsChannelsGetOutputContextAuthor struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsChannelsGetOutputContextAssignee represents the chats channels get output context assignee type.
type ChatsChannelsGetOutputContextAssignee struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsChannelsGetOutputContext - The external resource this conversation is attached to on the provider, such as an issue, ticket or page.
type ChatsChannelsGetOutputContext struct {
	// Type - What kind of external resource this is
	Type string `json:"type"`
	// Id - The resource's identifier on the provider
	Id string `json:"id"`
	// Description - Description of the resource
	Description *string `json:"description,omitempty"`
	// Status - Status of the resource, as reported by the provider
	Status *string `json:"status,omitempty"`
	// Url - Link to the resource
	Url      *string                                `json:"url,omitempty"`
	Author   *ChatsChannelsGetOutputContextAuthor   `json:"author,omitempty"`
	Assignee *ChatsChannelsGetOutputContextAssignee `json:"assignee,omitempty"`
	// Labels - Labels attached to the resource
	Labels *[]string `json:"labels,omitempty"`
}

// ChatsChannelsGetOutput represents the chats channels get output type.
type ChatsChannelsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat channel identifier
	Id string `json:"id"`
	// ChatId - The chat this channel belongs to
	ChatId string `json:"chat_id"`
	// WorkspaceId - The workspace this channel belongs to, for providers that group channels into workspaces
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// Type - How the channel is scoped on the chat provider
	Type string `json:"type"`
	// ProviderType - The provider's own name for this kind of channel
	ProviderType string `json:"provider_type"`
	// ProviderChannelId - The channel's identifier on the chat provider
	ProviderChannelId string `json:"provider_channel_id"`
	// Name - Display name of the channel
	Name *string `json:"name,omitempty"`
	// Topic - Topic set on the channel
	Topic *string `json:"topic,omitempty"`
	// Subject - Subject of the channel, for providers that separate it from the topic
	Subject *string `json:"subject,omitempty"`
	// HasAccess - Whether the chat integration can read content from this channel
	HasAccess bool                             `json:"has_access"`
	Recipient *ChatsChannelsGetOutputRecipient `json:"recipient,omitempty"`
	// MemberCount - Number of members in the channel, if the provider reports it
	MemberCount *float64 `json:"member_count,omitempty"`
	// Permalink - Link to the channel on the chat provider
	Permalink *string `json:"permalink,omitempty"`
	// Context - The external resource this conversation is attached to on the provider, such as an issue, ticket or page.
	Context *ChatsChannelsGetOutputContext `json:"context,omitempty"`
	// CreatedAt - Timestamp when the channel was first seen by Metorial
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the channel was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// LastInteractionAt - Timestamp of the last activity Metorial saw in this channel
	LastInteractionAt *time.Time `json:"last_interaction_at,omitempty"`
}

// MapChatsChannelsGetOutputFromJSON deserializes JSON data into a ChatsChannelsGetOutput.
func MapChatsChannelsGetOutputFromJSON(data []byte) (*ChatsChannelsGetOutput, error) {
	var v ChatsChannelsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsChannelsGetOutputToJSON serializes a ChatsChannelsGetOutput to JSON.
func MapChatsChannelsGetOutputToJSON(v *ChatsChannelsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
