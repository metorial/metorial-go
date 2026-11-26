package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/chat/connections"
)

// ChatConnectionsEndpoint provides access to chat connections link a chat provider, such as Slack or Microsoft Teams, to your instance.
type ChatConnectionsEndpoint struct {
	client *endpoint.Client
}

// NewChatConnectionsEndpoint creates a new ChatConnectionsEndpoint.
func NewChatConnectionsEndpoint(client *endpoint.Client) *ChatConnectionsEndpoint {
	return &ChatConnectionsEndpoint{client: client}
}

// ChatConnectionsEndpointListParams contains optional query parameters for List.
type ChatConnectionsEndpointListParams struct {
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
	// Status - Only return chat connections with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return chat connections with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose chat connection creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose chat connection last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ChatConnectionsEndpointCreateBody contains the request body for Create.
type ChatConnectionsEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the chat connection.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private data stored with the connection. Never returned by the API.
	PrivateMetadata *map[string]any `json:"private_metadata,omitempty"`
	// Provider - Chat provider to link this connection to.
	Provider map[string]any `json:"provider"`
}

// ChatConnectionsEndpointUpdateBody contains the request body for Update.
type ChatConnectionsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the chat connection. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private data stored with the connection. Never returned by the API.
	PrivateMetadata *map[string]any `json:"private_metadata,omitempty"`
}

// List returns a paginated list of chat connections.
func (e *ChatConnectionsEndpoint) List(instanceId string, params *ChatConnectionsEndpointListParams) (*connections.ChatConnectionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "chat", "connections"},
		Query: query,
	}
	var result connections.ChatConnectionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a chat connection.
func (e *ChatConnectionsEndpoint) Get(instanceId string, chatConnectionId string) (*connections.ChatConnectionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "connections", chatConnectionId},
	}
	var result connections.ChatConnectionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new chat connection, together with the single provider it is linked to.
func (e *ChatConnectionsEndpoint) Create(instanceId string, body *ChatConnectionsEndpointCreateBody) (*connections.ChatConnectionsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "connections"},
		Body: body,
	}
	var result connections.ChatConnectionsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a chat connection. Omitted fields keep their current value.
func (e *ChatConnectionsEndpoint) Update(instanceId string, chatConnectionId string, body *ChatConnectionsEndpointUpdateBody) (*connections.ChatConnectionsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "connections", chatConnectionId},
		Body: body,
	}
	var result connections.ChatConnectionsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a chat connection. It stays listable with `status=archived`.
func (e *ChatConnectionsEndpoint) Delete(instanceId string, chatConnectionId string) (*connections.ChatConnectionsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "connections", chatConnectionId},
	}
	var result connections.ChatConnectionsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
