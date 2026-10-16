package commands

import (
	"encoding/json"
	"time"
)

// ChatsCommandsRespondOutputMessageAuthor represents the chats commands respond output message author type.
type ChatsCommandsRespondOutputMessageAuthor struct {
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

// ChatsCommandsRespondOutputMessageReactionsEmoji represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ChatsCommandsRespondOutputMessageReactionsEmoji struct {
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

// ChatsCommandsRespondOutputMessageReactionsAuthors represents the chats commands respond output message reactions authors type.
type ChatsCommandsRespondOutputMessageReactionsAuthors struct {
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

// ChatsCommandsRespondOutputMessageReactions represents the chats commands respond output message reactions type.
type ChatsCommandsRespondOutputMessageReactions struct {
	Emoji ChatsCommandsRespondOutputMessageReactionsEmoji `json:"emoji"`
	// Count - Number of times this emoji was used to react
	Count float64 `json:"count"`
	// Authors - The authors who left this reaction, if the provider reports them
	Authors *[]ChatsCommandsRespondOutputMessageReactionsAuthors `json:"authors,omitempty"`
}

// ChatsCommandsRespondOutputMessageUnfurls represents the chats commands respond output message unfurls type.
type ChatsCommandsRespondOutputMessageUnfurls struct {
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

// ChatsCommandsRespondOutputMessageAttachments represents the chats commands respond output message attachments type.
type ChatsCommandsRespondOutputMessageAttachments struct {
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

// ChatsCommandsRespondOutputMessage - The response message, when the provider identifies it. Null when the provider accepts the response without returning the message, as Slack does.
type ChatsCommandsRespondOutputMessage struct {
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
	ProviderReplyToMessageId *string                                  `json:"provider_reply_to_message_id,omitempty"`
	Author                   *ChatsCommandsRespondOutputMessageAuthor `json:"author,omitempty"`
	// Body - Structured content of the message. Null once the message has been deleted.
	Body *map[string]any `json:"body,omitempty"`
	// Reactions - Reactions left on the message, as reported by the provider
	Reactions *[]ChatsCommandsRespondOutputMessageReactions `json:"reactions,omitempty"`
	// Unfurls - Link previews the provider generated for the message
	Unfurls *[]ChatsCommandsRespondOutputMessageUnfurls `json:"unfurls,omitempty"`
	// Attachments - Files attached to the message
	Attachments []ChatsCommandsRespondOutputMessageAttachments `json:"attachments"`
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

// ChatsCommandsRespondOutput represents the chats commands respond output type.
type ChatsCommandsRespondOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Message - The response message, when the provider identifies it. Null when the provider accepts the response without returning the message, as Slack does.
	Message *ChatsCommandsRespondOutputMessage `json:"message,omitempty"`
}

// MapChatsCommandsRespondOutputFromJSON deserializes JSON data into a ChatsCommandsRespondOutput.
func MapChatsCommandsRespondOutputFromJSON(data []byte) (*ChatsCommandsRespondOutput, error) {
	var v ChatsCommandsRespondOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsCommandsRespondOutputToJSON serializes a ChatsCommandsRespondOutput to JSON.
func MapChatsCommandsRespondOutputToJSON(v *ChatsCommandsRespondOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsCommandsRespondBody represents the chats commands respond body type.
type ChatsCommandsRespondBody struct {
	// ChatEventId - The chat.command.invoked event to respond to. Providers only accept a response within their response window after the command was invoked.
	ChatEventId string `json:"chat_event_id"`
	// Parts - Structured content parts making up the message body
	Parts []any `json:"parts"`
	// AltText - Plain-text fallback for the response
	AltText *string `json:"alt_text,omitempty"`
	// Ephemeral - If true, only the user who invoked the command can see the response. Defaults to the provider behavior. Providers that already acknowledged the command publicly reject it.
	Ephemeral *bool `json:"ephemeral,omitempty"`
}

// MapChatsCommandsRespondBodyFromJSON deserializes JSON data into a ChatsCommandsRespondBody.
func MapChatsCommandsRespondBodyFromJSON(data []byte) (*ChatsCommandsRespondBody, error) {
	var v ChatsCommandsRespondBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsCommandsRespondBodyToJSON serializes a ChatsCommandsRespondBody to JSON.
func MapChatsCommandsRespondBodyToJSON(v *ChatsCommandsRespondBody) ([]byte, error) {
	return json.Marshal(v)
}
