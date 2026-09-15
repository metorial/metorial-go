package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chats/threads"
)

// ChatsThreadsEndpoint provides access to chat threads group replies to a message within a chat channel.
type ChatsThreadsEndpoint struct {
	client *endpoint.Client
}

// NewChatsThreadsEndpoint creates a new ChatsThreadsEndpoint.
func NewChatsThreadsEndpoint(client *endpoint.Client) *ChatsThreadsEndpoint {
	return &ChatsThreadsEndpoint{client: client}
}

// ChatsThreadsEndpointListParams contains optional query parameters for List.
type ChatsThreadsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// ChannelId - The channel to list threads for
	ChannelId string `json:"channel_id"`
	// Type - Filter by thread type
	Type *string `json:"type,omitempty"`
}

// ChatsThreadsEndpointGetParams contains optional query parameters for Get.
type ChatsThreadsEndpointGetParams struct {
	// ChannelId - The channel this thread belongs to
	ChannelId string `json:"channel_id"`
}

// List returns a paginated list of threads in a chat channel.
func (e *ChatsThreadsEndpoint) List(chatId string, params *ChatsThreadsEndpointListParams) (*threads.ChatsThreadsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chats", chatId, "threads"},
		Query: query,
	}
	var result threads.ChatsThreadsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific chat thread.
func (e *ChatsThreadsEndpoint) Get(chatId string, threadId string, params *ChatsThreadsEndpointGetParams) (*threads.ChatsThreadsGetOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chats", chatId, "threads", threadId},
		Query: query,
	}
	var result threads.ChatsThreadsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
