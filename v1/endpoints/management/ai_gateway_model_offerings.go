package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/aigateway/modelofferings"
)

// AiGatewayModelOfferingsEndpoint provides access to how each provider serves a model: its id for the model, pricing, limits and capabilities.
type AiGatewayModelOfferingsEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayModelOfferingsEndpoint creates a new AiGatewayModelOfferingsEndpoint.
func NewAiGatewayModelOfferingsEndpoint(client *endpoint.Client) *AiGatewayModelOfferingsEndpoint {
	return &AiGatewayModelOfferingsEndpoint{client: client}
}

// AiGatewayModelOfferingsEndpointListParams contains optional query parameters for List.
type AiGatewayModelOfferingsEndpointListParams struct {
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
	// ModelId - Only offerings of the model with this ID or slug
	ModelId *string `json:"model_id,omitempty"`
	// Provider - Only offerings of the provider with this slug
	Provider *string `json:"provider,omitempty"`
}

// List returns a paginated list of the offerings in the catalog that are not retired.
func (e *AiGatewayModelOfferingsEndpoint) List(instanceId string, params *AiGatewayModelOfferingsEndpointListParams) (*modelofferings.AiGatewayModelOfferingsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "ai-gateway", "model-offerings"},
		Query: query,
	}
	var result modelofferings.AiGatewayModelOfferingsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
