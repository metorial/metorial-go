package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chats"
)

// ChatsEndpoint provides access to a chat represents a single connected chat surface, such as a Slack or Microsoft Teams tenant, running on a chat instance.
type ChatsEndpoint struct {
	client *endpoint.Client
}

// NewChatsEndpoint creates a new ChatsEndpoint.
func NewChatsEndpoint(client *endpoint.Client) *ChatsEndpoint {
	return &ChatsEndpoint{client: client}
}

// ChatsEndpointListParams contains optional query parameters for List.
type ChatsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Search - Search by name
	Search *string `json:"search,omitempty"`
	// Status - Filter by status
	Status *any `json:"status,omitempty"`
	// Id - Filter by chat ID(s)
	Id *any `json:"id,omitempty"`
	// ChatConnectionId - Filter by chat connection ID(s)
	ChatConnectionId *any `json:"chat_connection_id,omitempty"`
	// ChatInstanceId - Filter by chat instance ID(s)
	ChatInstanceId *any `json:"chat_instance_id,omitempty"`
	// ChatInstanceProviderId - Filter by chat instance provider ID(s)
	ChatInstanceProviderId *any `json:"chat_instance_provider_id,omitempty"`
	// CreatedAt - Filter chat creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter chat last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of chats.
func (e *ChatsEndpoint) List(instanceId string, params *ChatsEndpointListParams) (*chats.ChatsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "chats"},
		Query: query,
	}
	var result chats.ChatsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific chat.
func (e *ChatsEndpoint) Get(instanceId string, chatId string) (*chats.ChatsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chats", chatId},
	}
	var result chats.ChatsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
