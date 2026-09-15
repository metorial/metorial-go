package reactions

import (
	"encoding/json"
)

// ChatsMessagesReactionsDeleteOutputReactionsEmoji represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ChatsMessagesReactionsDeleteOutputReactionsEmoji struct {
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

// ChatsMessagesReactionsDeleteOutputReactionsAuthors represents the chats messages reactions delete output reactions authors type.
type ChatsMessagesReactionsDeleteOutputReactionsAuthors struct {
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

// ChatsMessagesReactionsDeleteOutputReactions represents the chats messages reactions delete output reactions type.
type ChatsMessagesReactionsDeleteOutputReactions struct {
	Emoji ChatsMessagesReactionsDeleteOutputReactionsEmoji `json:"emoji"`
	// Count - Number of times this emoji was used to react
	Count float64 `json:"count"`
	// Authors - The authors who left this reaction, if the provider reports them
	Authors *[]ChatsMessagesReactionsDeleteOutputReactionsAuthors `json:"authors,omitempty"`
}

// ChatsMessagesReactionsDeleteOutput represents the chats messages reactions delete output type.
type ChatsMessagesReactionsDeleteOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Reactions - Reactions left on the message, as reported by the provider
	Reactions []ChatsMessagesReactionsDeleteOutputReactions `json:"reactions"`
}

// MapChatsMessagesReactionsDeleteOutputFromJSON deserializes JSON data into a ChatsMessagesReactionsDeleteOutput.
func MapChatsMessagesReactionsDeleteOutputFromJSON(data []byte) (*ChatsMessagesReactionsDeleteOutput, error) {
	var v ChatsMessagesReactionsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsMessagesReactionsDeleteOutputToJSON serializes a ChatsMessagesReactionsDeleteOutput to JSON.
func MapChatsMessagesReactionsDeleteOutputToJSON(v *ChatsMessagesReactionsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsMessagesReactionsDeleteQuery represents the chats messages reactions delete query type.
type ChatsMessagesReactionsDeleteQuery struct {
	// ChannelId - The channel this message belongs to
	ChannelId string `json:"channel_id"`
	// Emoji - The unicode emoji shortcode to remove
	Emoji string `json:"emoji"`
}

// MapChatsMessagesReactionsDeleteQueryFromJSON deserializes JSON data into a ChatsMessagesReactionsDeleteQuery.
func MapChatsMessagesReactionsDeleteQueryFromJSON(data []byte) (*ChatsMessagesReactionsDeleteQuery, error) {
	var v ChatsMessagesReactionsDeleteQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsMessagesReactionsDeleteQueryToJSON serializes a ChatsMessagesReactionsDeleteQuery to JSON.
func MapChatsMessagesReactionsDeleteQueryToJSON(v *ChatsMessagesReactionsDeleteQuery) ([]byte, error) {
	return json.Marshal(v)
}
