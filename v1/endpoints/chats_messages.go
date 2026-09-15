package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chats/messages"
)

// ChatsMessagesEndpoint provides access to chat messages are the individual messages sent within a chat channel.
type ChatsMessagesEndpoint struct {
	client *endpoint.Client
}

// NewChatsMessagesEndpoint creates a new ChatsMessagesEndpoint.
func NewChatsMessagesEndpoint(client *endpoint.Client) *ChatsMessagesEndpoint {
	return &ChatsMessagesEndpoint{client: client}
}

// ChatsMessagesEndpointListParams contains optional query parameters for List.
type ChatsMessagesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// ChannelId - The channel to list messages for
	ChannelId string `json:"channel_id"`
	// ThreadId - Filter to messages in this thread
	ThreadId *string `json:"thread_id,omitempty"`
	// Search - Search message content
	Search *string `json:"search,omitempty"`
}

// ChatsMessagesEndpointGetParams contains optional query parameters for Get.
type ChatsMessagesEndpointGetParams struct {
	// ChannelId - The channel this message belongs to
	ChannelId string `json:"channel_id"`
}

// ChatsMessagesEndpointCreateBody contains the request body for Create.
type ChatsMessagesEndpointCreateBody struct {
	// ChannelId - The channel to send the message to
	ChannelId string `json:"channel_id"`
	// ThreadId - Send the message within this thread
	ThreadId *string `json:"thread_id,omitempty"`
	// Parts - Structured content parts making up the message body
	Parts []any `json:"parts"`
	// AltText - Plain-text fallback for the message
	AltText     *string           `json:"alt_text,omitempty"`
	Attachments *[]map[string]any `json:"attachments,omitempty"`
	// ReplyMessageId - Reply to this message
	ReplyMessageId *string `json:"reply_message_id,omitempty"`
	// EphemeralTargetUserId - Send this message so only this user can see it
	EphemeralTargetUserId *string `json:"ephemeral_target_user_id,omitempty"`
}

// ChatsMessagesEndpointUpdateBody contains the request body for Update.
type ChatsMessagesEndpointUpdateBody struct {
	// ChannelId - The channel this message belongs to
	ChannelId string `json:"channel_id"`
	// Parts - Structured content parts making up the message body
	Parts []any `json:"parts"`
	// AltText - Plain-text fallback for the message
	AltText *string `json:"alt_text,omitempty"`
}

// ChatsMessagesEndpointDeleteParams contains optional query parameters for Delete.
type ChatsMessagesEndpointDeleteParams struct {
	// ChannelId - The channel this message belongs to
	ChannelId string `json:"channel_id"`
}

// ChatsMessagesEndpointReadBody contains the request body for Read.
type ChatsMessagesEndpointReadBody struct {
	// ChannelId - The channel this message belongs to
	ChannelId string `json:"channel_id"`
	// ThreadId - The thread this message belongs to
	ThreadId *string `json:"thread_id,omitempty"`
}

// List returns a paginated list of messages in a chat channel.
func (e *ChatsMessagesEndpoint) List(chatId string, params *ChatsMessagesEndpointListParams) (*messages.ChatsMessagesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chats", chatId, "messages"},
		Query: query,
	}
	var result messages.ChatsMessagesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific chat message.
func (e *ChatsMessagesEndpoint) Get(chatId string, messageId string, params *ChatsMessagesEndpointGetParams) (*messages.ChatsMessagesGetOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chats", chatId, "messages", messageId},
		Query: query,
	}
	var result messages.ChatsMessagesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create sends a new message to a chat channel.
func (e *ChatsMessagesEndpoint) Create(chatId string, body *ChatsMessagesEndpointCreateBody) (*messages.ChatsMessagesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"chats", chatId, "messages"},
		Body: body,
	}
	var result messages.ChatsMessagesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update edits a specific chat message.
func (e *ChatsMessagesEndpoint) Update(chatId string, messageId string, body *ChatsMessagesEndpointUpdateBody) (*messages.ChatsMessagesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"chats", chatId, "messages", messageId},
		Body: body,
	}
	var result messages.ChatsMessagesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a specific chat message.
func (e *ChatsMessagesEndpoint) Delete(chatId string, messageId string, params *ChatsMessagesEndpointDeleteParams) (*messages.ChatsMessagesDeleteOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chats", chatId, "messages", messageId},
		Query: query,
	}
	var result messages.ChatsMessagesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Read marks a specific chat message as read.
func (e *ChatsMessagesEndpoint) Read(chatId string, messageId string, body *ChatsMessagesEndpointReadBody) (*messages.ChatsMessagesReadOutput, error) {
	req := &endpoint.Request{
		Path: []string{"chats", chatId, "messages", messageId, "read"},
		Body: body,
	}
	var result messages.ChatsMessagesReadOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
