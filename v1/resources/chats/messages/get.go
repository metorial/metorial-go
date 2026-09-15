package messages

import (
	"encoding/json"
	"time"
)

// ChatsMessagesGetOutputAuthor represents the chats messages get output author type.
type ChatsMessagesGetOutputAuthor struct {
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

// ChatsMessagesGetOutputReactionsEmoji represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ChatsMessagesGetOutputReactionsEmoji struct {
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

// ChatsMessagesGetOutputReactionsAuthors represents the chats messages get output reactions authors type.
type ChatsMessagesGetOutputReactionsAuthors struct {
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
	Raw      *any    `json:"raw,omitempty"`
}

// ChatsMessagesGetOutputReactions represents the chats messages get output reactions type.
type ChatsMessagesGetOutputReactions struct {
	Emoji ChatsMessagesGetOutputReactionsEmoji `json:"emoji"`
	// Count - Number of times this emoji was used to react
	Count float64 `json:"count"`
	// Authors - The authors who left this reaction, if the provider reports them
	Authors *[]ChatsMessagesGetOutputReactionsAuthors `json:"authors,omitempty"`
}

// ChatsMessagesGetOutputUnfurls represents the chats messages get output unfurls type.
type ChatsMessagesGetOutputUnfurls struct {
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

// ChatsMessagesGetOutputAttachments represents the chats messages get output attachments type.
type ChatsMessagesGetOutputAttachments struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat message attachment identifier
	Id string `json:"id"`
	// ProviderAttachmentId - The attachment's identifier on the chat provider
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
	// CreatedAt - Timestamp when the attachment was recorded
	CreatedAt time.Time `json:"created_at"`
}

// ChatsMessagesGetOutput represents the chats messages get output type.
type ChatsMessagesGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat message identifier
	Id string `json:"id"`
	// ChatId - The chat this message belongs to
	ChatId string `json:"chat_id"`
	// ChannelId - The channel this message was sent in
	ChannelId string `json:"channel_id"`
	// ThreadId - The thread this message was sent in, if it was sent in one
	ThreadId *string `json:"thread_id,omitempty"`
	// ProviderType - The provider's own name for this kind of message
	ProviderType string `json:"provider_type"`
	// ProviderMessageId - The message's identifier on the chat provider
	ProviderMessageId string `json:"provider_message_id"`
	// ProviderReplyToMessageId - Provider identifier of the message this one replies to, if it is a reply
	ProviderReplyToMessageId *string                       `json:"provider_reply_to_message_id,omitempty"`
	Author                   *ChatsMessagesGetOutputAuthor `json:"author,omitempty"`
	// Body - Structured content of the message. Null once the message has been deleted.
	Body *map[string]any `json:"body,omitempty"`
	// Reactions - Reactions left on the message, as reported by the provider
	Reactions *[]ChatsMessagesGetOutputReactions `json:"reactions,omitempty"`
	// Unfurls - Link previews the provider generated for the message
	Unfurls *[]ChatsMessagesGetOutputUnfurls `json:"unfurls,omitempty"`
	// Attachments - Files attached to the message
	Attachments []ChatsMessagesGetOutputAttachments `json:"attachments"`
	// SentAt - Timestamp when the message was sent on the chat provider
	SentAt time.Time `json:"sent_at"`
	// Edited - Whether the message has been edited since it was sent
	Edited bool `json:"edited"`
	// EditedAt - Timestamp when the message was last edited
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

// MapChatsMessagesGetOutputFromJSON deserializes JSON data into a ChatsMessagesGetOutput.
func MapChatsMessagesGetOutputFromJSON(data []byte) (*ChatsMessagesGetOutput, error) {
	var v ChatsMessagesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsMessagesGetOutputToJSON serializes a ChatsMessagesGetOutput to JSON.
func MapChatsMessagesGetOutputToJSON(v *ChatsMessagesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsMessagesGetQuery represents the chats messages get query type.
type ChatsMessagesGetQuery struct {
	// ChannelId - The channel this message belongs to
	ChannelId string `json:"channel_id"`
}

// MapChatsMessagesGetQueryFromJSON deserializes JSON data into a ChatsMessagesGetQuery.
func MapChatsMessagesGetQueryFromJSON(data []byte) (*ChatsMessagesGetQuery, error) {
	var v ChatsMessagesGetQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsMessagesGetQueryToJSON serializes a ChatsMessagesGetQuery to JSON.
func MapChatsMessagesGetQueryToJSON(v *ChatsMessagesGetQuery) ([]byte, error) {
	return json.Marshal(v)
}
