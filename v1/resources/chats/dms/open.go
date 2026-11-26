package dms

import (
	"encoding/json"
	"time"
)

// ChatsDmsOpenOutputRecipient - The DM recipient, when the provider reports one. `null` otherwise.
type ChatsDmsOpenOutputRecipient struct {
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

// ChatsDmsOpenOutputContextAuthor - Author of the external resource, as reported by the provider.
type ChatsDmsOpenOutputContextAuthor struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsDmsOpenOutputContextAssignee - Assignee of the external resource, as reported by the provider.
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
	// Description - Description of the external resource, as reported by the provider.
	Description *string `json:"description,omitempty"`
	// Status - Status of the resource, as reported by the provider
	Status *string `json:"status,omitempty"`
	// Url - Link to the resource
	Url *string `json:"url,omitempty"`
	// Author - Author of the external resource, as reported by the provider.
	Author *ChatsDmsOpenOutputContextAuthor `json:"author,omitempty"`
	// Assignee - Assignee of the external resource, as reported by the provider.
	Assignee *ChatsDmsOpenOutputContextAssignee `json:"assignee,omitempty"`
	// Labels - Labels attached to the resource
	Labels *[]string `json:"labels,omitempty"`
}

// ChatsDmsOpenOutput represents the chats dms open output type.
type ChatsDmsOpenOutput struct {
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
	Recipient *ChatsDmsOpenOutputRecipient `json:"recipient,omitempty"`
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
	// Type - Identify the user by their provider-specific ID.
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
	// Type - Identify the users by their provider-specific IDs.
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
	// Type - Open a DM with several users.
	Type *string `json:"type,omitempty"`
	// User - User to open the DM with.
	User *ChatsDmsOpenBodyUser `json:"user,omitempty"`
	// Users - Users to open the group DM with.
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
