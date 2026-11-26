package customproviders

import (
	"encoding/json"
	"time"
)

// CustomProvidersListOutputItemsDraftContainerImage - Container image this provider deploys. Only set for `container` providers.
type CustomProvidersListOutputItemsDraftContainerImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersListOutputItemsDraftRemoteMcpServer - Remote MCP server this provider connects to. Only set for `remote` providers.
type CustomProvidersListOutputItemsDraftRemoteMcpServer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersListOutputItemsDraftConfigSchema - Configuration fields for the custom provider.
type CustomProvidersListOutputItemsDraftConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersListOutputItemsDraftConfig - Configuration for the custom provider.
type CustomProvidersListOutputItemsDraftConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersListOutputItemsDraftConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersListOutputItemsDraft - Draft source and configuration of the custom provider.
type CustomProvidersListOutputItemsDraft struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// ContainerImage - Container image this provider deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersListOutputItemsDraftContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this provider connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersListOutputItemsDraftRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// Config - Configuration for the custom provider.
	Config CustomProvidersListOutputItemsDraftConfig `json:"config"`
}

// CustomProvidersListOutputItemsScmRepoProvider - Provider the repository lives on.
type CustomProvidersListOutputItemsScmRepoProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Provider the repository lives on.
	Type string `json:"type"`
	// Id - ID of the repository on the provider.
	Id string `json:"id"`
	// Name - Name of the repository on the provider.
	Name string `json:"name"`
	// Owner - Owner of the repository on the provider.
	Owner string `json:"owner"`
}

// CustomProvidersListOutputItemsScmRepo - Repository the provider code is deployed from. `null` when not linked.
type CustomProvidersListOutputItemsScmRepo struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersListOutputItemsScmRepoProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersListOutputItemsProviderPublisher represents the custom providers list output items provider publisher type.
type CustomProvidersListOutputItemsProviderPublisher struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ImageUrl - URL of the publisher logo
	ImageUrl string `json:"image_url"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersListOutputItemsProviderCurrentVersion represents the custom providers list output items provider current version type.
type CustomProvidersListOutputItemsProviderCurrentVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release label for the version.
	Version string `json:"version"`
	// ProviderId - Provider this version belongs to.
	ProviderId string `json:"provider_id"`
	// IsCurrent - The version deployments follow unless locked to another.
	IsCurrent bool `json:"is_current"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SpecificationId - Specification this version belongs to.
	SpecificationId *string `json:"specification_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersListOutputItemsProviderOauthAutoRegistration represents the custom providers list output items provider oauth auto registration type.
type CustomProvidersListOutputItemsProviderOauthAutoRegistration struct {
	// Status - Auto-registration status
	Status string `json:"status"`
}

// CustomProvidersListOutputItemsProviderOauth represents the custom providers list output items provider oauth type.
type CustomProvidersListOutputItemsProviderOauth struct {
	// Status - OAuth status
	Status string `json:"status"`
	// CallbackUrl - Callback URL to register on your OAuth app.
	CallbackUrl      *string                                                     `json:"callback_url,omitempty"`
	AutoRegistration CustomProvidersListOutputItemsProviderOauthAutoRegistration `json:"auto_registration"`
}

// CustomProvidersListOutputItemsProvider - Session provider backed by this custom provider. `null` when none is linked.
type CustomProvidersListOutputItemsProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Access - `public` providers are available to every instance. `tenant` providers are your own.
	Access string `json:"access"`
	// Status - Archived once the provider is archived.
	Status         string                                                `json:"status"`
	Publisher      CustomProvidersListOutputItemsProviderPublisher       `json:"publisher"`
	CurrentVersion *CustomProvidersListOutputItemsProviderCurrentVersion `json:"current_version,omitempty"`
	Oauth          *CustomProvidersListOutputItemsProviderOauth          `json:"oauth,omitempty"`
	// Identifier - Provider identifier
	Identifier string `json:"identifier"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersListOutputItems represents the custom providers list output items type.
type CustomProvidersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once archived. Archived providers reject new connections.
	Status string `json:"status"`
	// Type - How the provider runs: `function`, `container`, or `remote`.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Draft - Draft source and configuration of the custom provider.
	Draft CustomProvidersListOutputItemsDraft `json:"draft"`
	// ScmRepo - Repository the provider code is deployed from. `null` when not linked.
	ScmRepo *CustomProvidersListOutputItemsScmRepo `json:"scm_repo,omitempty"`
	// Provider - Session provider backed by this custom provider. `null` when none is linked.
	Provider *CustomProvidersListOutputItemsProvider `json:"provider,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersListOutputPagination represents the custom providers list output pagination type.
type CustomProvidersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// CustomProvidersListOutput represents the custom providers list output type.
type CustomProvidersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                              `json:"object"`
	Items      []CustomProvidersListOutputItems    `json:"items"`
	Pagination CustomProvidersListOutputPagination `json:"pagination"`
}

// MapCustomProvidersListOutputFromJSON deserializes JSON data into a CustomProvidersListOutput.
func MapCustomProvidersListOutputFromJSON(data []byte) (*CustomProvidersListOutput, error) {
	var v CustomProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersListOutputToJSON serializes a CustomProvidersListOutput to JSON.
func MapCustomProvidersListOutputToJSON(v *CustomProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CustomProvidersListQueryCreatedAt - Only return results whose custom provider creation time is in this range. Set `gt`, `lt`, or both.
type CustomProvidersListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CustomProvidersListQueryUpdatedAt - Only return results whose custom provider last update time is in this range. Set `gt`, `lt`, or both.
type CustomProvidersListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CustomProvidersListQuery represents the custom providers list query type.
type CustomProvidersListQuery struct {
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
	// Status - Only return custom providers with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Type - Only return custom providers with one of these types.
	Type *any `json:"type,omitempty"`
	// Id - Only return custom providers with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ProviderId - Only return custom providers linked to one of these provider IDs (matches providers connected to sessions).
	ProviderId *any `json:"provider_id,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose custom provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *CustomProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose custom provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *CustomProvidersListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapCustomProvidersListQueryFromJSON deserializes JSON data into a CustomProvidersListQuery.
func MapCustomProvidersListQueryFromJSON(data []byte) (*CustomProvidersListQuery, error) {
	var v CustomProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersListQueryToJSON serializes a CustomProvidersListQuery to JSON.
func MapCustomProvidersListQueryToJSON(v *CustomProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
