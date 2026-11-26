package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/customproviders/environments"
)

// CustomProvidersEnvironmentsEndpoint provides access to a custom provider environment is a deployment target for custom provider versions, such as staging or production.
type CustomProvidersEnvironmentsEndpoint struct {
	client *endpoint.Client
}

// NewCustomProvidersEnvironmentsEndpoint creates a new CustomProvidersEnvironmentsEndpoint.
func NewCustomProvidersEnvironmentsEndpoint(client *endpoint.Client) *CustomProvidersEnvironmentsEndpoint {
	return &CustomProvidersEnvironmentsEndpoint{client: client}
}

// CustomProvidersEnvironmentsEndpointListParams contains optional query parameters for List.
type CustomProvidersEnvironmentsEndpointListParams struct {
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
	// Id - Only return environments with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CustomProviderVersionId - Only return environments running one of these versions.
	CustomProviderVersionId *any `json:"custom_provider_version_id,omitempty"`
	// CustomProviderId - Only return environments of one of these custom providers.
	CustomProviderId *any `json:"custom_provider_id,omitempty"`
	// CreatedAt - Only return results whose custom provider environment creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose custom provider environment last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of custom provider environments.
func (e *CustomProvidersEnvironmentsEndpoint) List(instanceId string, params *CustomProvidersEnvironmentsEndpointListParams) (*environments.CustomProvidersEnvironmentsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "custom-provider-environments"},
		Query: query,
	}
	var result environments.CustomProvidersEnvironmentsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a custom provider environment.
func (e *CustomProvidersEnvironmentsEndpoint) Get(instanceId string, customProviderEnvironmentId string) (*environments.CustomProvidersEnvironmentsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "custom-provider-environments", customProviderEnvironmentId},
	}
	var result environments.CustomProvidersEnvironmentsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
