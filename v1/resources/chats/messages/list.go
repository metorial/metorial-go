package messages

import (
	"encoding/json"
	"time"
)

// ChatsMessagesListOutputItemsAuthor - Author who sent the message. `null` when the message was first seen as a deletion, which carries no author.
type ChatsMessagesListOutputItemsAuthor struct {
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

// ChatsMessagesListOutputItemsReactionsEmoji represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ChatsMessagesListOutputItemsReactionsEmoji struct {
	// Type - Provider-specific custom emoji.
	Type *string `json:"type,omitempty"`
	// Value - The unicode emoji shortcode
	Value *string `json:"value,omitempty"`
	// Name - The custom emoji name
	Name *string `json:"name,omitempty"`
	// Url - URL of the custom emoji image
	Url *string `json:"url,omitempty"`
	// Id - The custom emoji's identifier on the provider
	Id *string `json:"id,omitempty"`
}

// ChatsMessagesListOutputItemsReactionsAuthors represents the chats messages list output items reactions authors type.
type ChatsMessagesListOutputItemsReactionsAuthors struct {
	// UserId - The author's identifier on the chat provider
	UserId string `json:"userId"`
	// UserName - The author's handle on the chat provider
	UserName string `json:"userName"`
	// FullName - The author's display name
	FullName string `json:"fullName"`
	// Type - What kind of account this author is on the chat provider
	Type string `json:"type"`
	// Role - The role this author holds in the chat workspace
	Role *string `json:"role,omitempty"`
	// ProviderType - The provider's own name for this kind of author
	ProviderType *string `json:"providerType,omitempty"`
	// IsMe - Whether this author is the account the integration itself is authenticated as
	IsMe bool `json:"isMe"`
	// Email - Email address of the author
	Email *string `json:"email,omitempty"`
	// ImageUrl - URL of the author's avatar
	ImageUrl *string `json:"imageUrl,omitempty"`
	// Raw - Raw reaction payload from the provider.
	Raw *any `json:"raw,omitempty"`
}

// ChatsMessagesListOutputItemsReactions represents the chats messages list output items reactions type.
type ChatsMessagesListOutputItemsReactions struct {
	// Emoji - Emoji used for the reaction.
	Emoji ChatsMessagesListOutputItemsReactionsEmoji `json:"emoji"`
	// Count - Number of times this emoji was used to react
	Count float64 `json:"count"`
	// Authors - The authors who left this reaction, if the provider reports them
	Authors *[]ChatsMessagesListOutputItemsReactionsAuthors `json:"authors,omitempty"`
}

// ChatsMessagesListOutputItemsUnfurls represents the chats messages list output items unfurls type.
type ChatsMessagesListOutputItemsUnfurls struct {
	// Url - URL the preview was generated for
	Url string `json:"url"`
	// Title - Title of the linked page
	Title *string `json:"title,omitempty"`
	// Description - Description of the linked page
	Description *string `json:"description,omitempty"`
	// ImageUrl - Preview image for the linked page
	ImageUrl *string `json:"imageUrl,omitempty"`
	// SiteName - Name of the site the link belongs to
	SiteName *string `json:"siteName,omitempty"`
	// MessageId - Provider identifier of the message this unfurl is attached to
	MessageId *string `json:"messageId,omitempty"`
}

// ChatsMessagesListOutputItemsAttachments represents the chats messages list output items attachments type.
type ChatsMessagesListOutputItemsAttachments struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderAttachmentId - Identifier the chat provider uses for this attachment.
	ProviderAttachmentId *string `json:"provider_attachment_id,omitempty"`
	// Type - The kind of content the attachment holds
	Type string `json:"type"`
	// Name - File name of the attachment
	Name *string `json:"name,omitempty"`
	// MimeType - MIME type of the attachment
	MimeType *string `json:"mime_type,omitempty"`
	// Size - Size of the attachment in bytes
	Size *float64 `json:"size,omitempty"`
	// Width - Width of the attachment in pixels, for visual attachments
	Width *float64 `json:"width,omitempty"`
	// Height - Height of the attachment in pixels, for visual attachments
	Height *float64 `json:"height,omitempty"`
	// Position - Position of the attachment within its message, starting at zero
	Position float64 `json:"position"`
	// FileId - The Metorial file the attachment content is served through
	FileId string `json:"file_id"`
	// DownloadUrl - Temporary URL the attachment content can be downloaded from. Null while the content is unavailable.
	DownloadUrl *string `json:"download_url,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// ChatsMessagesListOutputItems represents the chats messages list output items type.
type ChatsMessagesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ChatId - The chat this message belongs to
	ChatId string `json:"chat_id"`
	// ChannelId - The channel this message was sent in
	ChannelId string `json:"channel_id"`
	// ThreadId - The thread this message was sent in, if it was sent in one
	ThreadId *string `json:"thread_id,omitempty"`
	// ProviderType - The provider's own name for this kind of message
	ProviderType string `json:"provider_type"`
	// ProviderMessageId - Identifier the chat provider uses for this message.
	ProviderMessageId string `json:"provider_message_id"`
	// ProviderReplyToMessageId - Provider identifier of the message this one replies to, if it is a reply
	ProviderReplyToMessageId *string `json:"provider_reply_to_message_id,omitempty"`
	// Author - Author who sent the message. `null` when the message was first seen as a deletion, which carries no author.
	Author *ChatsMessagesListOutputItemsAuthor `json:"author,omitempty"`
	// Body - Structured content of the message. Null once the message has been deleted.
	Body *map[string]any `json:"body,omitempty"`
	// Reactions - Reactions left on the message, as reported by the provider
	Reactions *[]ChatsMessagesListOutputItemsReactions `json:"reactions,omitempty"`
	// Unfurls - Link previews the provider generated for the message
	Unfurls *[]ChatsMessagesListOutputItemsUnfurls `json:"unfurls,omitempty"`
	// Attachments - Files attached to the message
	Attachments []ChatsMessagesListOutputItemsAttachments `json:"attachments"`
	// SentAt - Timestamp when the message was sent on the chat provider
	SentAt time.Time `json:"sent_at"`
	// Edited - Whether the message has been edited since it was sent
	Edited bool `json:"edited"`
	// EditedAt - When the message was last edited. `null` if never edited.
	EditedAt *time.Time `json:"edited_at,omitempty"`
	// CreatedAt - Timestamp when the message was first seen by Metorial
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the message was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// LastInteractionAt - Timestamp of the last activity Metorial saw on this message
	LastInteractionAt *time.Time `json:"last_interaction_at,omitempty"`
	// DeletedAt - Timestamp when the message was deleted on the chat provider. Deleted messages keep their identity but lose their content.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// ChatsMessagesListOutputPagination represents the chats messages list output pagination type.
type ChatsMessagesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ChatsMessagesListOutput represents the chats messages list output type.
type ChatsMessagesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                            `json:"object"`
	Items      []ChatsMessagesListOutputItems    `json:"items"`
	Pagination ChatsMessagesListOutputPagination `json:"pagination"`
}

// MapChatsMessagesListOutputFromJSON deserializes JSON data into a ChatsMessagesListOutput.
func MapChatsMessagesListOutputFromJSON(data []byte) (*ChatsMessagesListOutput, error) {
	var v ChatsMessagesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsMessagesListOutputToJSON serializes a ChatsMessagesListOutput to JSON.
func MapChatsMessagesListOutputToJSON(v *ChatsMessagesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsMessagesListQuery represents the chats messages list query type.
type ChatsMessagesListQuery struct {
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
	// ChannelId - The channel to list messages for
	ChannelId string `json:"channel_id"`
	// ThreadId - Only return messages in this thread.
	ThreadId *string `json:"thread_id,omitempty"`
	// Search - Search message content
	Search *string `json:"search,omitempty"`
}

// MapChatsMessagesListQueryFromJSON deserializes JSON data into a ChatsMessagesListQuery.
func MapChatsMessagesListQueryFromJSON(data []byte) (*ChatsMessagesListQuery, error) {
	var v ChatsMessagesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsMessagesListQueryToJSON serializes a ChatsMessagesListQuery to JSON.
func MapChatsMessagesListQueryToJSON(v *ChatsMessagesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
