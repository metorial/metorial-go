package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chat/workspaces"
)

// ChatWorkspacesEndpoint provides access to chat workspaces group channels together, for chat providers that organize conversations that way.
type ChatWorkspacesEndpoint struct {
	client *endpoint.Client
}

// NewChatWorkspacesEndpoint creates a new ChatWorkspacesEndpoint.
func NewChatWorkspacesEndpoint(client *endpoint.Client) *ChatWorkspacesEndpoint {
	return &ChatWorkspacesEndpoint{client: client}
}

// ChatWorkspacesEndpointListParams contains optional query parameters for List.
type ChatWorkspacesEndpointListParams struct {
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
	// ChatInstanceId - The chat instance to list workspaces for
	ChatInstanceId string `json:"chat_instance_id"`
	// Search - Search by workspace name
	Search *string `json:"search,omitempty"`
}

// ChatWorkspacesEndpointGetParams contains optional query parameters for Get.
type ChatWorkspacesEndpointGetParams struct {
	// ChatInstanceId - The chat instance this workspace belongs to
	ChatInstanceId string `json:"chat_instance_id"`
}

// List returns a paginated list of workspaces for a chat instance.
func (e *ChatWorkspacesEndpoint) List(params *ChatWorkspacesEndpointListParams) (*workspaces.ChatWorkspacesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chat", "workspaces"},
		Query: query,
	}
	var result workspaces.ChatWorkspacesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a chat workspace.
func (e *ChatWorkspacesEndpoint) Get(chatWorkspaceId string, params *ChatWorkspacesEndpointGetParams) (*workspaces.ChatWorkspacesGetOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chat", "workspaces", chatWorkspaceId},
		Query: query,
	}
	var result workspaces.ChatWorkspacesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
