package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/chat/instances"
)

// ChatInstancesEndpoint provides access to a chat instance is a chat connection running with its own configuration and identity.
type ChatInstancesEndpoint struct {
	client *endpoint.Client
}

// NewChatInstancesEndpoint creates a new ChatInstancesEndpoint.
func NewChatInstancesEndpoint(client *endpoint.Client) *ChatInstancesEndpoint {
	return &ChatInstancesEndpoint{client: client}
}

// ChatInstancesEndpointListParams contains optional query parameters for List.
type ChatInstancesEndpointListParams struct {
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
	// Status - Only return chat instances with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return chat instances with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ChatConnectionId - Only return chat instances for one of these chat connections.
	ChatConnectionId *any `json:"chat_connection_id,omitempty"`
	// CreatedAt - Only return results whose chat instance creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose chat instance last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ChatInstancesEndpointCreateBody contains the request body for Create.
type ChatInstancesEndpointCreateBody struct {
	// ChatConnectionId - Chat connection to create this instance for.
	ChatConnectionId string `json:"chat_connection_id"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the chat instance.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private data stored with the instance. Never returned by the API.
	PrivateMetadata *map[string]any `json:"private_metadata,omitempty"`
	// IdentityActorId - Identity actor to run this instance as
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// IdentityId - Identity to run this instance as
	IdentityId *string `json:"identity_id,omitempty"`
}

// ChatInstancesEndpointUpdateBody contains the request body for Update.
type ChatInstancesEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the chat instance. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private data stored with the instance. Never returned by the API.
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

// Get retrieves a chat instance.
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

// Update updates a chat instance. Omitted fields keep their current value.
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

// Delete archives a chat instance. It stays listable with `status=archived`.
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
