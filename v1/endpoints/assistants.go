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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
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
