package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/aigateway/providers"
)

// AiGatewayProvidersEndpoint provides access to the providers that serve models in the AI gateway catalog, e.g. OpenAI, Amazon Bedrock or OpenRouter.
type AiGatewayProvidersEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayProvidersEndpoint creates a new AiGatewayProvidersEndpoint.
func NewAiGatewayProvidersEndpoint(client *endpoint.Client) *AiGatewayProvidersEndpoint {
	return &AiGatewayProvidersEndpoint{client: client}
}

// AiGatewayProvidersEndpointListParams contains optional query parameters for List.
type AiGatewayProvidersEndpointListParams struct {
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

// List returns a paginated list of the providers in the catalog, including whether you can connect your own account for each.
func (e *AiGatewayProvidersEndpoint) List(params *AiGatewayProvidersEndpointListParams) (*providers.AiGatewayProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"ai-gateway", "providers"},
		Query: query,
	}
	var result providers.AiGatewayProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
