package providers

import (
	"encoding/json"
	"time"
)

// ProvidersListOutputItemsPublisher represents the providers list output items publisher type.
type ProvidersListOutputItemsPublisher struct {
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

// ProvidersListOutputItemsCurrentVersion represents the providers list output items current version type.
type ProvidersListOutputItemsCurrentVersion struct {
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

// ProvidersListOutputItemsOauthAutoRegistration represents the providers list output items oauth auto registration type.
type ProvidersListOutputItemsOauthAutoRegistration struct {
	// Status - Auto-registration status
	Status string `json:"status"`
}

// ProvidersListOutputItemsOauth represents the providers list output items oauth type.
type ProvidersListOutputItemsOauth struct {
	// Status - OAuth status
	Status string `json:"status"`
	// CallbackUrl - Callback URL to register on your OAuth app.
	CallbackUrl      *string                                       `json:"callback_url,omitempty"`
	AutoRegistration ProvidersListOutputItemsOauthAutoRegistration `json:"auto_registration"`
}

// ProvidersListOutputItems represents the providers list output items type.
type ProvidersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Access - `public` providers are available to every instance. `tenant` providers are your own.
	Access string `json:"access"`
	// Status - Archived once the provider is archived.
	Status         string                                  `json:"status"`
	Publisher      ProvidersListOutputItemsPublisher       `json:"publisher"`
	CurrentVersion *ProvidersListOutputItemsCurrentVersion `json:"current_version,omitempty"`
	Oauth          *ProvidersListOutputItemsOauth          `json:"oauth,omitempty"`
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

// ProvidersListOutputPagination represents the providers list output pagination type.
type ProvidersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProvidersListOutput represents the providers list output type.
type ProvidersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                        `json:"object"`
	Items      []ProvidersListOutputItems    `json:"items"`
	Pagination ProvidersListOutputPagination `json:"pagination"`
}

// MapProvidersListOutputFromJSON deserializes JSON data into a ProvidersListOutput.
func MapProvidersListOutputFromJSON(data []byte) (*ProvidersListOutput, error) {
	var v ProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersListOutputToJSON serializes a ProvidersListOutput to JSON.
func MapProvidersListOutputToJSON(v *ProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProvidersListQuery represents the providers list query type.
type ProvidersListQuery struct {
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
	// Id - Filter by provider ID(s)
	Id *any `json:"id,omitempty"`
	// Search - Search providers by name, description, or readme
	Search *string `json:"search,omitempty"`
	// AuthMethod - Filter by auth method — matches an auth method ID, auth method global ID, key, name, or type (oauth, token, service_account, custom)
	AuthMethod *any `json:"auth_method,omitempty"`
	// AuthSetup - Filter by auth setup status. "configured" matches providers with a token or custom auth method, or with auth credentials already configured. "not_configured" matches providers with only OAuth auth methods and no auth credentials configured.
	AuthSetup *any `json:"auth_setup,omitempty"`
}

// MapProvidersListQueryFromJSON deserializes JSON data into a ProvidersListQuery.
func MapProvidersListQueryFromJSON(data []byte) (*ProvidersListQuery, error) {
	var v ProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersListQueryToJSON serializes a ProvidersListQuery to JSON.
func MapProvidersListQueryToJSON(v *ProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
