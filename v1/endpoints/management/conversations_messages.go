package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/conversations/messages"
)

// ConversationsMessagesEndpoint provides access to assistant and conversation endpoints
type ConversationsMessagesEndpoint struct {
	client *endpoint.Client
}

// NewConversationsMessagesEndpoint creates a new ConversationsMessagesEndpoint.
func NewConversationsMessagesEndpoint(client *endpoint.Client) *ConversationsMessagesEndpoint {
	return &ConversationsMessagesEndpoint{client: client}
}

// ConversationsMessagesEndpointListParams contains optional query parameters for List.
type ConversationsMessagesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// ConversationsMessagesEndpointCreateBody contains the request body for Create.
type ConversationsMessagesEndpointCreateBody struct {
	Message         map[string]any `json:"message"`
	ParentMessageId *string        `json:"parent_message_id,omitempty"`
	ModelId         *string        `json:"model_id,omitempty"`
}

// List list messages in a specific assistant conversation.
func (e *ConversationsMessagesEndpoint) List(instanceId string, assistantConversationId string, params *ConversationsMessagesEndpointListParams) (*messages.ConversationsMessagesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "conversations", assistantConversationId, "messages"},
		Query: query,
	}
	var result messages.ConversationsMessagesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create create a user message and assistant request in a specific conversation.
func (e *ConversationsMessagesEndpoint) Create(instanceId string, assistantConversationId string, body *ConversationsMessagesEndpointCreateBody) (*messages.ConversationsMessagesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "conversations", assistantConversationId, "messages"},
		Body: body,
	}
	var result messages.ConversationsMessagesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get get a specific assistant message.
func (e *ConversationsMessagesEndpoint) Get(instanceId string, assistantConversationId string, assistantMessageId string) (*messages.ConversationsMessagesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "conversations", assistantConversationId, "messages", assistantMessageId},
	}
	var result messages.ConversationsMessagesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
