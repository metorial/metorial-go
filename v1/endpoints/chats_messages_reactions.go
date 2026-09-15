package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chats/messages/reactions"
)

// ChatsMessagesReactionsEndpoint provides access to chat messages are the individual messages sent within a chat channel.
type ChatsMessagesReactionsEndpoint struct {
	client *endpoint.Client
}

// NewChatsMessagesReactionsEndpoint creates a new ChatsMessagesReactionsEndpoint.
func NewChatsMessagesReactionsEndpoint(client *endpoint.Client) *ChatsMessagesReactionsEndpoint {
	return &ChatsMessagesReactionsEndpoint{client: client}
}

// ChatsMessagesReactionsEndpointListParams contains optional query parameters for List.
type ChatsMessagesReactionsEndpointListParams struct {
	// ChannelId - The channel this message belongs to
	ChannelId string `json:"channel_id"`
}

// ChatsMessagesReactionsEndpointCreateBody contains the request body for Create.
type ChatsMessagesReactionsEndpointCreateBody struct {
	// ChannelId - The channel this message belongs to
	ChannelId string `json:"channel_id"`
	Emoji     any    `json:"emoji"`
}

// ChatsMessagesReactionsEndpointDeleteParams contains optional query parameters for Delete.
type ChatsMessagesReactionsEndpointDeleteParams struct {
	// ChannelId - The channel this message belongs to
	ChannelId string `json:"channel_id"`
	// Emoji - The unicode emoji shortcode to remove
	Emoji string `json:"emoji"`
}

// List returns the reactions left on a specific chat message.
func (e *ChatsMessagesReactionsEndpoint) List(chatId string, messageId string, params *ChatsMessagesReactionsEndpointListParams) (*reactions.ChatsMessagesReactionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chats", chatId, "messages", messageId, "reactions"},
		Query: query,
	}
	var result reactions.ChatsMessagesReactionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create adds a reaction to a specific chat message.
func (e *ChatsMessagesReactionsEndpoint) Create(chatId string, messageId string, body *ChatsMessagesReactionsEndpointCreateBody) (*reactions.ChatsMessagesReactionsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"chats", chatId, "messages", messageId, "reactions"},
		Body: body,
	}
	var result reactions.ChatsMessagesReactionsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a reaction from a specific chat message.
func (e *ChatsMessagesReactionsEndpoint) Delete(chatId string, messageId string, params *ChatsMessagesReactionsEndpointDeleteParams) (*reactions.ChatsMessagesReactionsDeleteOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chats", chatId, "messages", messageId, "reactions"},
		Query: query,
	}
	var result reactions.ChatsMessagesReactionsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
