package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/conversations"
)

// ConversationsEndpoint provides access to assistant and conversation endpoints
type ConversationsEndpoint struct {
	client *endpoint.Client
}

// NewConversationsEndpoint creates a new ConversationsEndpoint.
func NewConversationsEndpoint(client *endpoint.Client) *ConversationsEndpoint {
	return &ConversationsEndpoint{client: client}
}

// ConversationsEndpointListParams contains optional query parameters for List.
type ConversationsEndpointListParams struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order       *string `json:"order,omitempty"`
	AssistantId *any    `json:"assistant_id,omitempty"`
}

// ConversationsEndpointCreateBody contains the request body for Create.
type ConversationsEndpointCreateBody struct {
	AssistantId string          `json:"assistant_id"`
	Title       *string         `json:"title,omitempty"`
	Input       *map[string]any `json:"input,omitempty"`
}

// ConversationsEndpointUpdateBody contains the request body for Update.
type ConversationsEndpointUpdateBody struct {
	Title *string `json:"title,omitempty"`
}

// List list assistant conversations in an instance.
func (e *ConversationsEndpoint) List(params *ConversationsEndpointListParams) (*conversations.ConversationsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"conversations"},
		Query: query,
	}
	var result conversations.ConversationsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create create a new assistant conversation in an instance.
func (e *ConversationsEndpoint) Create(body *ConversationsEndpointCreateBody) (*conversations.ConversationsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"conversations"},
		Body: body,
	}
	var result conversations.ConversationsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get get a specific assistant conversation.
func (e *ConversationsEndpoint) Get(assistantConversationId string) (*conversations.ConversationsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"conversations", assistantConversationId},
	}
	var result conversations.ConversationsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update update a specific assistant conversation.
func (e *ConversationsEndpoint) Update(assistantConversationId string, body *ConversationsEndpointUpdateBody) (*conversations.ConversationsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"conversations", assistantConversationId},
		Body: body,
	}
	var result conversations.ConversationsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
