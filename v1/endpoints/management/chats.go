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
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Search - Search by name
	Search *string `json:"search,omitempty"`
	// Status - Only return chats with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return chats with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ChatConnectionId - Only return chats for one of these chat connections.
	ChatConnectionId *any `json:"chat_connection_id,omitempty"`
	// ChatInstanceId - Only return chats for one of these chat instances.
	ChatInstanceId *any `json:"chat_instance_id,omitempty"`
	// ChatInstanceProviderId - Only return chats for one of these chat instance providers.
	ChatInstanceProviderId *any `json:"chat_instance_provider_id,omitempty"`
	// CreatedAt - Only return results whose chat creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose chat last update time is in this range. Set `gt`, `lt`, or both.
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

// Get retrieves a chat.
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
