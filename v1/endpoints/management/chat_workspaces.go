package management

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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
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
func (e *ChatWorkspacesEndpoint) List(instanceId string, params *ChatWorkspacesEndpointListParams) (*workspaces.ChatWorkspacesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "chat", "workspaces"},
		Query: query,
	}
	var result workspaces.ChatWorkspacesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific chat workspace.
func (e *ChatWorkspacesEndpoint) Get(instanceId string, chatWorkspaceId string, params *ChatWorkspacesEndpointGetParams) (*workspaces.ChatWorkspacesGetOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "chat", "workspaces", chatWorkspaceId},
		Query: query,
	}
	var result workspaces.ChatWorkspacesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
