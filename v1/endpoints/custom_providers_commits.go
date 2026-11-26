package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/customproviders/commits"
)

// CustomProvidersCommitsEndpoint provides access to a custom provider commit is a promotion of a version into an environment, or a rollback to a previous version.
type CustomProvidersCommitsEndpoint struct {
	client *endpoint.Client
}

// NewCustomProvidersCommitsEndpoint creates a new CustomProvidersCommitsEndpoint.
func NewCustomProvidersCommitsEndpoint(client *endpoint.Client) *CustomProvidersCommitsEndpoint {
	return &CustomProvidersCommitsEndpoint{client: client}
}

// CustomProvidersCommitsEndpointListParams contains optional query parameters for List.
type CustomProvidersCommitsEndpointListParams struct {
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
	// Id - Only return commits with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CustomProviderVersionId - Only return commits for one of these custom provider versions.
	CustomProviderVersionId *any `json:"custom_provider_version_id,omitempty"`
	// CustomProviderEnvironmentId - Only return commits for one of these custom provider environments.
	CustomProviderEnvironmentId *any `json:"custom_provider_environment_id,omitempty"`
	// CustomProviderId - Only return commits for one of these custom providers.
	CustomProviderId *any `json:"custom_provider_id,omitempty"`
	// ProviderId - Only return commits for one of these provider IDs.
	ProviderId *any `json:"provider_id,omitempty"`
	// CreatedAt - Only return results whose custom provider commit creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose custom provider commit last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// CustomProvidersCommitsEndpointCreateBody contains the request body for Create.
type CustomProvidersCommitsEndpointCreateBody struct {
	// Message - Commit message
	Message string `json:"message"`
	// Action - The commit action to perform
	Action any `json:"action"`
}

// List returns a paginated list of custom provider commits.
func (e *CustomProvidersCommitsEndpoint) List(params *CustomProvidersCommitsEndpointListParams) (*commits.CustomProvidersCommitsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"custom-provider-commits"},
		Query: query,
	}
	var result commits.CustomProvidersCommitsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a custom provider commit.
func (e *CustomProvidersCommitsEndpoint) Get(customProviderCommitId string) (*commits.CustomProvidersCommitsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"custom-provider-commits", customProviderCommitId},
	}
	var result commits.CustomProvidersCommitsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new commit to promote or rollback a version in an environment.
func (e *CustomProvidersCommitsEndpoint) Create(body *CustomProvidersCommitsEndpointCreateBody) (*commits.CustomProvidersCommitsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"custom-provider-commits"},
		Body: body,
	}
	var result commits.CustomProvidersCommitsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
