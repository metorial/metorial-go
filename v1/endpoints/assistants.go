package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/assistants"
)

// AssistantsEndpoint provides access to assistant and conversation endpoints
type AssistantsEndpoint struct {
	client *endpoint.Client
}

// NewAssistantsEndpoint creates a new AssistantsEndpoint.
func NewAssistantsEndpoint(client *endpoint.Client) *AssistantsEndpoint {
	return &AssistantsEndpoint{client: client}
}

// AssistantsEndpointListParams contains optional query parameters for List.
type AssistantsEndpointListParams struct {
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
}

// List list assistants available in an instance.
func (e *AssistantsEndpoint) List(params *AssistantsEndpointListParams) (*assistants.AssistantsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"assistants"},
		Query: query,
	}
	var result assistants.AssistantsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get get an assistant available in an instance.
func (e *AssistantsEndpoint) Get(assistantId string) (*assistants.AssistantsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"assistants", assistantId},
	}
	var result assistants.AssistantsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
