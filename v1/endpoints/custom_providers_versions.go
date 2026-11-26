package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/customproviders/versions"
)

// CustomProvidersVersionsEndpoint provides access to a custom provider version is a release of a custom provider. It is built by a deployment and promoted into environments.
type CustomProvidersVersionsEndpoint struct {
	client *endpoint.Client
}

// NewCustomProvidersVersionsEndpoint creates a new CustomProvidersVersionsEndpoint.
func NewCustomProvidersVersionsEndpoint(client *endpoint.Client) *CustomProvidersVersionsEndpoint {
	return &CustomProvidersVersionsEndpoint{client: client}
}

// CustomProvidersVersionsEndpointListParams contains optional query parameters for List.
type CustomProvidersVersionsEndpointListParams struct {
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
	// Status - Only return versions with one of these deployment statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return versions with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider IDs (matches providers connected to sessions)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderVersionId - Filter by provider version IDs (matches providers connected to sessions)
	ProviderVersionId *any `json:"provider_version_id,omitempty"`
	// CustomProviderId - Filter by custom provider IDs (matches providers connected to sessions)
	CustomProviderId *any `json:"custom_provider_id,omitempty"`
	// CustomProviderDeploymentId - Filter by custom provider deployment IDs (matches providers connected to sessions)
	CustomProviderDeploymentId *any `json:"custom_provider_deployment_id,omitempty"`
	// CustomProviderEnvironmentId - Filter by custom provider environment IDs (matches providers connected to sessions)
	CustomProviderEnvironmentId *any `json:"custom_provider_environment_id,omitempty"`
	// CreatedAt - Only return results whose custom provider version creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose custom provider version last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// CustomProvidersVersionsEndpointCreateBody contains the request body for Create.
type CustomProvidersVersionsEndpointCreateBody struct {
	// CustomProviderId - Custom provider to release a new version for.
	CustomProviderId string `json:"custom_provider_id"`
	// From - Source the provider runs from.
	From any `json:"from"`
	// Config - Configuration schema and transformer for the provider.
	Config *map[string]any `json:"config,omitempty"`
}

// List returns a paginated list of custom provider versions.
func (e *CustomProvidersVersionsEndpoint) List(params *CustomProvidersVersionsEndpointListParams) (*versions.CustomProvidersVersionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"custom-provider-versions"},
		Query: query,
	}
	var result versions.CustomProvidersVersionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a custom provider version.
func (e *CustomProvidersVersionsEndpoint) Get(customProviderVersionId string) (*versions.CustomProvidersVersionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"custom-provider-versions", customProviderVersionId},
	}
	var result versions.CustomProvidersVersionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEnv retrieves the environment variables for a version of a custom provider.
func (e *CustomProvidersVersionsEndpoint) GetEnv(customProviderVersionId string) (*versions.CustomProvidersVersionsGetEnvOutput, error) {
	req := &endpoint.Request{
		Path: []string{"custom-provider-versions", customProviderVersionId, "env"},
	}
	var result versions.CustomProvidersVersionsGetEnvOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new version for a custom provider.
func (e *CustomProvidersVersionsEndpoint) Create(body *CustomProvidersVersionsEndpointCreateBody) (*versions.CustomProvidersVersionsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"custom-provider-versions"},
		Body: body,
	}
	var result versions.CustomProvidersVersionsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
