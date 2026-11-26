package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providers/triggers"
)

// ProvidersTriggersEndpoint provides access to a provider trigger is an event a provider can emit for callbacks. Each provider version has its own set of triggers.
type ProvidersTriggersEndpoint struct {
	client *endpoint.Client
}

// NewProvidersTriggersEndpoint creates a new ProvidersTriggersEndpoint.
func NewProvidersTriggersEndpoint(client *endpoint.Client) *ProvidersTriggersEndpoint {
	return &ProvidersTriggersEndpoint{client: client}
}

// ProvidersTriggersEndpointListParams contains optional query parameters for List.
type ProvidersTriggersEndpointListParams struct {
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
	// ProviderVersionId - Provider version to list triggers for
	ProviderVersionId string `json:"provider_version_id"`
}

// List returns a paginated list of provider triggers for a specific provider version.
func (e *ProvidersTriggersEndpoint) List(params *ProvidersTriggersEndpointListParams) (*triggers.ProvidersTriggersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-triggers"},
		Query: query,
	}
	var result triggers.ProvidersTriggersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider trigger by ID.
func (e *ProvidersTriggersEndpoint) Get(providerTriggerId string) (*triggers.ProvidersTriggersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-triggers", providerTriggerId},
	}
	var result triggers.ProvidersTriggersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
