package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chat/instances"
)

// ChatInstancesEndpoint provides access to chat instances materialize a chat connection for a specific runtime configuration.
type ChatInstancesEndpoint struct {
	client *endpoint.Client
}

// NewChatInstancesEndpoint creates a new ChatInstancesEndpoint.
func NewChatInstancesEndpoint(client *endpoint.Client) *ChatInstancesEndpoint {
	return &ChatInstancesEndpoint{client: client}
}

// ChatInstancesEndpointListParams contains optional query parameters for List.
type ChatInstancesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Search - Search by name
	Search *string `json:"search,omitempty"`
	// Status - Filter by status
	Status *any `json:"status,omitempty"`
	// Id - Filter by chat instance ID(s)
	Id *any `json:"id,omitempty"`
	// ChatConnectionId - Filter by chat connection ID(s)
	ChatConnectionId *any `json:"chat_connection_id,omitempty"`
	// CreatedAt - Filter chat instance creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter chat instance last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ChatInstancesEndpointCreateBody contains the request body for Create.
type ChatInstancesEndpointCreateBody struct {
	// ChatConnectionId - The chat connection to create this instance for
	ChatConnectionId string `json:"chat_connection_id"`
	// Name - Display name of the chat instance
	Name *string `json:"name,omitempty"`
	// Description - Description of the instance
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata for the instance
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private metadata for the instance
	PrivateMetadata *map[string]any `json:"private_metadata,omitempty"`
	// IdentityActorId - Identity actor to run this instance as
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// IdentityId - Identity to run this instance as
	IdentityId *string `json:"identity_id,omitempty"`
}

// ChatInstancesEndpointUpdateBody contains the request body for Update.
type ChatInstancesEndpointUpdateBody struct {
	// Name - Display name of the chat instance
	Name *string `json:"name,omitempty"`
	// Description - Description of the instance
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private metadata
	PrivateMetadata *map[string]any `json:"private_metadata,omitempty"`
}

// List returns a paginated list of chat instances.
func (e *ChatInstancesEndpoint) List(instanceId string, params *ChatInstancesEndpointListParams) (*instances.ChatInstancesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "chat", "instances"},
		Query: query,
	}
	var result instances.ChatInstancesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific chat instance.
func (e *ChatInstancesEndpoint) Get(instanceId string, chatInstanceId string) (*instances.ChatInstancesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "instances", chatInstanceId},
	}
	var result instances.ChatInstancesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new chat instance for a chat connection.
func (e *ChatInstancesEndpoint) Create(instanceId string, body *ChatInstancesEndpointCreateBody) (*instances.ChatInstancesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "instances"},
		Body: body,
	}
	var result instances.ChatInstancesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific chat instance.
func (e *ChatInstancesEndpoint) Update(instanceId string, chatInstanceId string, body *ChatInstancesEndpointUpdateBody) (*instances.ChatInstancesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "instances", chatInstanceId},
		Body: body,
	}
	var result instances.ChatInstancesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a specific chat instance.
func (e *ChatInstancesEndpoint) Delete(instanceId string, chatInstanceId string) (*instances.ChatInstancesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "instances", chatInstanceId},
	}
	var result instances.ChatInstancesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Sync triggers a sync of the workspaces available on this chat instance.
func (e *ChatInstancesEndpoint) Sync(instanceId string, chatInstanceId string) (*instances.ChatInstancesSyncOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "instances", chatInstanceId, "sync"},
	}
	var result instances.ChatInstancesSyncOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
