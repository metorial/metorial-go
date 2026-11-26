package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/aigateway/virtualmodels/versions"
)

// AiGatewayVirtualModelsVersionsEndpoint provides access to every change to the definition of a virtual model is kept as a numbered version that can be rolled back to.
type AiGatewayVirtualModelsVersionsEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayVirtualModelsVersionsEndpoint creates a new AiGatewayVirtualModelsVersionsEndpoint.
func NewAiGatewayVirtualModelsVersionsEndpoint(client *endpoint.Client) *AiGatewayVirtualModelsVersionsEndpoint {
	return &AiGatewayVirtualModelsVersionsEndpoint{client: client}
}

// AiGatewayVirtualModelsVersionsEndpointListParams contains optional query parameters for List.
type AiGatewayVirtualModelsVersionsEndpointListParams struct {
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

// List returns a paginated list of the versions of a virtual model.
func (e *AiGatewayVirtualModelsVersionsEndpoint) List(instanceId string, virtualModelId string, params *AiGatewayVirtualModelsVersionsEndpointListParams) (*versions.AiGatewayVirtualModelsVersionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "ai-gateway", "virtual-models", virtualModelId, "versions"},
		Query: query,
	}
	var result versions.AiGatewayVirtualModelsVersionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a version of a virtual model by ID, with its definition.
func (e *AiGatewayVirtualModelsVersionsEndpoint) Get(instanceId string, virtualModelId string, versionId string) (*versions.AiGatewayVirtualModelsVersionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "ai-gateway", "virtual-models", virtualModelId, "versions", versionId},
	}
	var result versions.AiGatewayVirtualModelsVersionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
