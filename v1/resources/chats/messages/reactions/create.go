package reactions

import (
	"encoding/json"
)

// ChatsMessagesReactionsCreateOutputReactionsEmoji represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ChatsMessagesReactionsCreateOutputReactionsEmoji struct {
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

// ChatsMessagesReactionsCreateOutputReactionsAuthors represents the chats messages reactions create output reactions authors type.
type ChatsMessagesReactionsCreateOutputReactionsAuthors struct {
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

// ChatsMessagesReactionsCreateOutputReactions represents the chats messages reactions create output reactions type.
type ChatsMessagesReactionsCreateOutputReactions struct {
	Emoji ChatsMessagesReactionsCreateOutputReactionsEmoji `json:"emoji"`
	// Count - Number of times this emoji was used to react
	Count float64 `json:"count"`
	// Authors - The authors who left this reaction, if the provider reports them
	Authors *[]ChatsMessagesReactionsCreateOutputReactionsAuthors `json:"authors,omitempty"`
}

// ChatsMessagesReactionsCreateOutput represents the chats messages reactions create output type.
type ChatsMessagesReactionsCreateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Reactions - Reactions left on the message, as reported by the provider
	Reactions []ChatsMessagesReactionsCreateOutputReactions `json:"reactions"`
}

// MapChatsMessagesReactionsCreateOutputFromJSON deserializes JSON data into a ChatsMessagesReactionsCreateOutput.
func MapChatsMessagesReactionsCreateOutputFromJSON(data []byte) (*ChatsMessagesReactionsCreateOutput, error) {
	var v ChatsMessagesReactionsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsMessagesReactionsCreateOutputToJSON serializes a ChatsMessagesReactionsCreateOutput to JSON.
func MapChatsMessagesReactionsCreateOutputToJSON(v *ChatsMessagesReactionsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsMessagesReactionsCreateBodyEmoji represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ChatsMessagesReactionsCreateBodyEmoji struct {
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

// ChatsMessagesReactionsCreateBody represents the chats messages reactions create body type.
type ChatsMessagesReactionsCreateBody struct {
	// ChannelId - The channel this message belongs to
	ChannelId string                                `json:"channel_id"`
	Emoji     ChatsMessagesReactionsCreateBodyEmoji `json:"emoji"`
}

// MapChatsMessagesReactionsCreateBodyFromJSON deserializes JSON data into a ChatsMessagesReactionsCreateBody.
func MapChatsMessagesReactionsCreateBodyFromJSON(data []byte) (*ChatsMessagesReactionsCreateBody, error) {
	var v ChatsMessagesReactionsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsMessagesReactionsCreateBodyToJSON serializes a ChatsMessagesReactionsCreateBody to JSON.
func MapChatsMessagesReactionsCreateBodyToJSON(v *ChatsMessagesReactionsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
