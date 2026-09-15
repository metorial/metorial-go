package reactions

import (
	"encoding/json"
)

// ChatsMessagesReactionsListOutputReactionsEmoji represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ChatsMessagesReactionsListOutputReactionsEmoji struct {
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

// ChatsMessagesReactionsListOutputReactionsAuthors represents the chats messages reactions list output reactions authors type.
type ChatsMessagesReactionsListOutputReactionsAuthors struct {
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

// ChatsMessagesReactionsListOutputReactions represents the chats messages reactions list output reactions type.
type ChatsMessagesReactionsListOutputReactions struct {
	Emoji ChatsMessagesReactionsListOutputReactionsEmoji `json:"emoji"`
	// Count - Number of times this emoji was used to react
	Count float64 `json:"count"`
	// Authors - The authors who left this reaction, if the provider reports them
	Authors *[]ChatsMessagesReactionsListOutputReactionsAuthors `json:"authors,omitempty"`
}

// ChatsMessagesReactionsListOutput represents the chats messages reactions list output type.
type ChatsMessagesReactionsListOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Reactions - Reactions left on the message, as reported by the provider
	Reactions []ChatsMessagesReactionsListOutputReactions `json:"reactions"`
}

// MapChatsMessagesReactionsListOutputFromJSON deserializes JSON data into a ChatsMessagesReactionsListOutput.
func MapChatsMessagesReactionsListOutputFromJSON(data []byte) (*ChatsMessagesReactionsListOutput, error) {
	var v ChatsMessagesReactionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsMessagesReactionsListOutputToJSON serializes a ChatsMessagesReactionsListOutput to JSON.
func MapChatsMessagesReactionsListOutputToJSON(v *ChatsMessagesReactionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsMessagesReactionsListQuery represents the chats messages reactions list query type.
type ChatsMessagesReactionsListQuery struct {
	// ChannelId - The channel this message belongs to
	ChannelId string `json:"channel_id"`
}

// MapChatsMessagesReactionsListQueryFromJSON deserializes JSON data into a ChatsMessagesReactionsListQuery.
func MapChatsMessagesReactionsListQueryFromJSON(data []byte) (*ChatsMessagesReactionsListQuery, error) {
	var v ChatsMessagesReactionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsMessagesReactionsListQueryToJSON serializes a ChatsMessagesReactionsListQuery to JSON.
func MapChatsMessagesReactionsListQueryToJSON(v *ChatsMessagesReactionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
