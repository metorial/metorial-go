package channels

import (
	"encoding/json"
	"time"
)

// ChatsChannelsListOutputItemsRecipient - The DM recipient, when the provider reports one. `null` otherwise.
type ChatsChannelsListOutputItemsRecipient struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ChatId - The chat this author belongs to
	ChatId string `json:"chat_id"`
	// Type - What kind of account this author is on the chat provider
	Type string `json:"type"`
	// Role - The role this author holds in the chat workspace
	Role string `json:"role"`
	// ProviderType - The provider's own name for this kind of author
	ProviderType string `json:"provider_type"`
	// ProviderAuthorId - Identifier the chat provider uses for this author.
	ProviderAuthorId string `json:"provider_author_id"`
	// UserName - The author's handle on the chat provider
	UserName string `json:"user_name"`
	// FullName - Display name.
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

// ChatsChannelsListOutputItemsContextAuthor - Author of the external resource, as reported by the provider.
type ChatsChannelsListOutputItemsContextAuthor struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsChannelsListOutputItemsContextAssignee - Assignee of the external resource, as reported by the provider.
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
	// Description - Description of the external resource, as reported by the provider.
	Description *string `json:"description,omitempty"`
	// Status - Status of the resource, as reported by the provider
	Status *string `json:"status,omitempty"`
	// Url - Link to the resource
	Url *string `json:"url,omitempty"`
	// Author - Author of the external resource, as reported by the provider.
	Author *ChatsChannelsListOutputItemsContextAuthor `json:"author,omitempty"`
	// Assignee - Assignee of the external resource, as reported by the provider.
	Assignee *ChatsChannelsListOutputItemsContextAssignee `json:"assignee,omitempty"`
	// Labels - Labels attached to the resource
	Labels *[]string `json:"labels,omitempty"`
}

// ChatsChannelsListOutputItems represents the chats channels list output items type.
type ChatsChannelsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ChatId - The chat this channel belongs to
	ChatId string `json:"chat_id"`
	// WorkspaceId - The workspace this channel belongs to, for providers that group channels into workspaces
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// Type - How the channel is scoped on the chat provider
	Type string `json:"type"`
	// ProviderType - The provider's own name for this kind of channel
	ProviderType string `json:"provider_type"`
	// ProviderChannelId - Identifier the chat provider uses for this channel.
	ProviderChannelId string `json:"provider_channel_id"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Topic - Topic set on the channel
	Topic *string `json:"topic,omitempty"`
	// Subject - Subject of the channel, for providers that separate it from the topic
	Subject *string `json:"subject,omitempty"`
	// HasAccess - Whether the chat integration can read content from this channel
	HasAccess bool `json:"has_access"`
	// Recipient - The DM recipient, when the provider reports one. `null` otherwise.
	Recipient *ChatsChannelsListOutputItemsRecipient `json:"recipient,omitempty"`
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
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ChatsChannelsListOutput represents the chats channels list output type.
type ChatsChannelsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                            `json:"object"`
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
	// WorkspaceId - Only return channels in this workspace.
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// Type - Only return channels of this type.
	Type *string `json:"type,omitempty"`
	// Search - Search by channel name or topic
	Search *string `json:"search,omitempty"`
	// HasAccess - Whether to return channels the chat integration can or cannot access; defaults to true
	HasAccess *bool `json:"has_access,omitempty"`
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
