package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Search - Search by name
	Search *string `json:"search,omitempty"`
	// Status - Filter by status
	Status *any `json:"status,omitempty"`
	// Id - Filter by chat connection ID(s)
	Id *any `json:"id,omitempty"`
	// CreatedAt - Filter chat connection creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter chat connection last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ChatConnectionsEndpointCreateBody contains the request body for Create.
type ChatConnectionsEndpointCreateBody struct {
	// Name - Display name of the chat connection
	Name string `json:"name"`
	// Description - Description of the connection
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata for the connection
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private metadata for the connection
	PrivateMetadata *map[string]any `json:"private_metadata,omitempty"`
	Provider        map[string]any  `json:"provider"`
}

// ChatConnectionsEndpointUpdateBody contains the request body for Update.
type ChatConnectionsEndpointUpdateBody struct {
	// Name - Display name of the chat connection
	Name *string `json:"name,omitempty"`
	// Description - Description of the connection
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private metadata
	PrivateMetadata *map[string]any `json:"private_metadata,omitempty"`
}

// List returns a paginated list of chat connections.
func (e *ChatConnectionsEndpoint) List(params *ChatConnectionsEndpointListParams) (*connections.ChatConnectionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chat", "connections"},
		Query: query,
	}
	var result connections.ChatConnectionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific chat connection.
func (e *ChatConnectionsEndpoint) Get(chatConnectionId string) (*connections.ChatConnectionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"chat", "connections", chatConnectionId},
	}
	var result connections.ChatConnectionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new chat connection, together with the single provider it is linked to.
func (e *ChatConnectionsEndpoint) Create(body *ChatConnectionsEndpointCreateBody) (*connections.ChatConnectionsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"chat", "connections"},
		Body: body,
	}
	var result connections.ChatConnectionsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific chat connection.
func (e *ChatConnectionsEndpoint) Update(chatConnectionId string, body *ChatConnectionsEndpointUpdateBody) (*connections.ChatConnectionsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"chat", "connections", chatConnectionId},
		Body: body,
	}
	var result connections.ChatConnectionsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a specific chat connection.
func (e *ChatConnectionsEndpoint) Delete(chatConnectionId string) (*connections.ChatConnectionsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"chat", "connections", chatConnectionId},
	}
	var result connections.ChatConnectionsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
