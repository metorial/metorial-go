package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/aigateway/labs"
)

// AiGatewayLabsEndpoint provides access to the labs that make the models in the AI gateway catalog, e.g. Anthropic or OpenAI.
type AiGatewayLabsEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayLabsEndpoint creates a new AiGatewayLabsEndpoint.
func NewAiGatewayLabsEndpoint(client *endpoint.Client) *AiGatewayLabsEndpoint {
	return &AiGatewayLabsEndpoint{client: client}
}

// AiGatewayLabsEndpointListParams contains optional query parameters for List.
type AiGatewayLabsEndpointListParams struct {
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

// List returns a paginated list of the labs in the catalog.
func (e *AiGatewayLabsEndpoint) List(params *AiGatewayLabsEndpointListParams) (*labs.AiGatewayLabsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"ai-gateway", "labs"},
		Query: query,
	}
	var result labs.AiGatewayLabsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
