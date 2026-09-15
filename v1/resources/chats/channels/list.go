package channels

import (
	"encoding/json"
	"time"
)

// ChatsChannelsListOutputItemsContextAuthor represents the chats channels list output items context author type.
type ChatsChannelsListOutputItemsContextAuthor struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsChannelsListOutputItemsContextAssignee represents the chats channels list output items context assignee type.
type ChatsChannelsListOutputItemsContextAssignee struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsChannelsListOutputItemsContext - The external resource this conversation is attached to on the provider, such as an issue, ticket or page.
type ChatsChannelsListOutputItemsContext struct {
	// Type - What kind of external resource this is
	Type string `json:"type"`
	// Id - The resource's identifier on the provider
	Id string `json:"id"`
	// Description - Description of the resource
	Description *string `json:"description,omitempty"`
	// Status - Status of the resource, as reported by the provider
	Status *string `json:"status,omitempty"`
	// Url - Link to the resource
	Url      *string                                      `json:"url,omitempty"`
	Author   *ChatsChannelsListOutputItemsContextAuthor   `json:"author,omitempty"`
	Assignee *ChatsChannelsListOutputItemsContextAssignee `json:"assignee,omitempty"`
	// Labels - Labels attached to the resource
	Labels *[]string `json:"labels,omitempty"`
}

// ChatsChannelsListOutputItems represents the chats channels list output items type.
type ChatsChannelsListOutputItems struct {
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
	// MemberCount - Number of members in the channel, if the provider reports it
	MemberCount *float64 `json:"member_count,omitempty"`
	// Permalink - Link to the channel on the chat provider
	Permalink *string `json:"permalink,omitempty"`
	// Context - The external resource this conversation is attached to on the provider, such as an issue, ticket or page.
	Context *ChatsChannelsListOutputItemsContext `json:"context,omitempty"`
	// CreatedAt - Timestamp when the channel was first seen by Metorial
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the channel was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// LastInteractionAt - Timestamp of the last activity Metorial saw in this channel
	LastInteractionAt *time.Time `json:"last_interaction_at,omitempty"`
}

// ChatsChannelsListOutputPagination represents the chats channels list output pagination type.
type ChatsChannelsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ChatsChannelsListOutput represents the chats channels list output type.
type ChatsChannelsListOutput struct {
	Items      []ChatsChannelsListOutputItems    `json:"items"`
	Pagination ChatsChannelsListOutputPagination `json:"pagination"`
}

// MapChatsChannelsListOutputFromJSON deserializes JSON data into a ChatsChannelsListOutput.
func MapChatsChannelsListOutputFromJSON(data []byte) (*ChatsChannelsListOutput, error) {
	var v ChatsChannelsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsChannelsListOutputToJSON serializes a ChatsChannelsListOutput to JSON.
func MapChatsChannelsListOutputToJSON(v *ChatsChannelsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsChannelsListQuery represents the chats channels list query type.
type ChatsChannelsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// WorkspaceId - Filter to channels in this workspace
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// Type - Filter by channel type
	Type *string `json:"type,omitempty"`
	// Search - Search by channel name or topic
	Search *string `json:"search,omitempty"`
}

// MapChatsChannelsListQueryFromJSON deserializes JSON data into a ChatsChannelsListQuery.
func MapChatsChannelsListQueryFromJSON(data []byte) (*ChatsChannelsListQuery, error) {
	var v ChatsChannelsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsChannelsListQueryToJSON serializes a ChatsChannelsListQuery to JSON.
func MapChatsChannelsListQueryToJSON(v *ChatsChannelsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
