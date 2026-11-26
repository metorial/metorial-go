package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/customproviders/deployments"
)

// CustomProvidersDeploymentsEndpoint provides access to a custom provider deployment is a build and rollout of one custom provider version.
type CustomProvidersDeploymentsEndpoint struct {
	client *endpoint.Client
}

// NewCustomProvidersDeploymentsEndpoint creates a new CustomProvidersDeploymentsEndpoint.
func NewCustomProvidersDeploymentsEndpoint(client *endpoint.Client) *CustomProvidersDeploymentsEndpoint {
	return &CustomProvidersDeploymentsEndpoint{client: client}
}

// CustomProvidersDeploymentsEndpointListParams contains optional query parameters for List.
type CustomProvidersDeploymentsEndpointListParams struct {
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
	// Status - Only return deployments with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return deployments with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CustomProviderVersionId - Only return deployments of one of these custom provider versions.
	CustomProviderVersionId *any `json:"custom_provider_version_id,omitempty"`
	// CustomProviderId - Only return deployments of one of these custom providers.
	CustomProviderId *any `json:"custom_provider_id,omitempty"`
	// CreatedAt - Only return results whose custom provider deployment creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose custom provider deployment last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of custom provider deployments.
func (e *CustomProvidersDeploymentsEndpoint) List(instanceId string, params *CustomProvidersDeploymentsEndpointListParams) (*deployments.CustomProvidersDeploymentsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "custom-provider-deployments"},
		Query: query,
	}
	var result deployments.CustomProvidersDeploymentsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a custom provider deployment.
func (e *CustomProvidersDeploymentsEndpoint) Get(instanceId string, customProviderDeploymentId string) (*deployments.CustomProvidersDeploymentsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "custom-provider-deployments", customProviderDeploymentId},
	}
	var result deployments.CustomProvidersDeploymentsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLogs retrieves the build and deployment logs for a custom provider deployment.
func (e *CustomProvidersDeploymentsEndpoint) GetLogs(instanceId string, customProviderDeploymentId string) (*deployments.CustomProvidersDeploymentsGetLogsOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "custom-provider-deployments", customProviderDeploymentId, "logs"},
	}
	var result deployments.CustomProvidersDeploymentsGetLogsOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
