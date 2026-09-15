package dms

import (
	"encoding/json"
	"time"
)

// ChatsDmsOpenOutputContextAuthor represents the chats dms open output context author type.
type ChatsDmsOpenOutputContextAuthor struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsDmsOpenOutputContextAssignee represents the chats dms open output context assignee type.
type ChatsDmsOpenOutputContextAssignee struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsDmsOpenOutputContext - The external resource this conversation is attached to on the provider, such as an issue, ticket or page.
type ChatsDmsOpenOutputContext struct {
	// Type - What kind of external resource this is
	Type string `json:"type"`
	// Id - The resource's identifier on the provider
	Id string `json:"id"`
	// Description - Description of the resource
	Description *string `json:"description,omitempty"`
	// Status - Status of the resource, as reported by the provider
	Status *string `json:"status,omitempty"`
	// Url - Link to the resource
	Url      *string                            `json:"url,omitempty"`
	Author   *ChatsDmsOpenOutputContextAuthor   `json:"author,omitempty"`
	Assignee *ChatsDmsOpenOutputContextAssignee `json:"assignee,omitempty"`
	// Labels - Labels attached to the resource
	Labels *[]string `json:"labels,omitempty"`
}

// ChatsDmsOpenOutput represents the chats dms open output type.
type ChatsDmsOpenOutput struct {
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
	Context *ChatsDmsOpenOutputContext `json:"context,omitempty"`
	// CreatedAt - Timestamp when the channel was first seen by Metorial
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the channel was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// LastInteractionAt - Timestamp of the last activity Metorial saw in this channel
	LastInteractionAt *time.Time `json:"last_interaction_at,omitempty"`
}

// MapChatsDmsOpenOutputFromJSON deserializes JSON data into a ChatsDmsOpenOutput.
func MapChatsDmsOpenOutputFromJSON(data []byte) (*ChatsDmsOpenOutput, error) {
	var v ChatsDmsOpenOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsDmsOpenOutputToJSON serializes a ChatsDmsOpenOutput to JSON.
func MapChatsDmsOpenOutputToJSON(v *ChatsDmsOpenOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsDmsOpenBodyUser represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ChatsDmsOpenBodyUser struct {
	Type *string `json:"type,omitempty"`
	// Email - The user's email address
	Email *string `json:"email,omitempty"`
	// PhoneNumber - The user's phone number
	PhoneNumber *string `json:"phone_number,omitempty"`
	// UserId - The user's provider-specific ID
	UserId *string `json:"user_id,omitempty"`
}

// ChatsDmsOpenBodyUsers represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ChatsDmsOpenBodyUsers struct {
	Type *string `json:"type,omitempty"`
	// Emails - The users' email addresses
	Emails *[]string `json:"emails,omitempty"`
	// PhoneNumbers - The users' phone numbers
	PhoneNumbers *[]string `json:"phone_numbers,omitempty"`
	// UserIds - The users' provider-specific IDs
	UserIds *[]string `json:"user_ids,omitempty"`
}

// ChatsDmsOpenBody represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ChatsDmsOpenBody struct {
	Type  *string                `json:"type,omitempty"`
	User  *ChatsDmsOpenBodyUser  `json:"user,omitempty"`
	Users *ChatsDmsOpenBodyUsers `json:"users,omitempty"`
}

// MapChatsDmsOpenBodyFromJSON deserializes JSON data into a ChatsDmsOpenBody.
func MapChatsDmsOpenBodyFromJSON(data []byte) (*ChatsDmsOpenBody, error) {
	var v ChatsDmsOpenBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsDmsOpenBodyToJSON serializes a ChatsDmsOpenBody to JSON.
func MapChatsDmsOpenBodyToJSON(v *ChatsDmsOpenBody) ([]byte, error) {
	return json.Marshal(v)
}
