package portals

import (
	"encoding/json"
	"time"
)

// PortalsListOutputItemsSkillConfiguration - Skill settings for the portal.
type PortalsListOutputItemsSkillConfiguration struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Whether the portal uses the default skill settings.
	IsDefault bool `json:"is_default"`
	// AllowScripts - Whether skills on the portal can run scripts.
	AllowScripts bool `json:"allow_scripts"`
	// AllowedFileExtensions - File extensions skills on the portal can use.
	AllowedFileExtensions []string `json:"allowed_file_extensions"`
	// AllowNonStandardDirectories - Whether skills can use directories outside the standard layout.
	AllowNonStandardDirectories bool `json:"allow_non_standard_directories"`
}

// PortalsListOutputItemsAuthAllowedRedirectUrlFilters represents the portals list output items auth allowed redirect url filters type.
type PortalsListOutputItemsAuthAllowedRedirectUrlFilters struct {
	// Url - A redirect URL pattern consumers may return to after signing in.
	Url string `json:"url"`
}

// PortalsListOutputItemsAuth - Consumer authentication settings for the portal.
type PortalsListOutputItemsAuth struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// SessionExpiryTimeInSeconds - How long consumer sessions stay valid, in seconds.
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
	// AllowedRedirectUrlFilters - Redirect URL patterns consumers may return to after signing in.
	AllowedRedirectUrlFilters []PortalsListOutputItemsAuthAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters"`
}

// PortalsListOutputItemsUrls represents the portals list output items urls type.
type PortalsListOutputItemsUrls struct {
	// Type - `namespace` for namespace-specific URLs, `default` otherwise.
	Type string `json:"type"`
	// Url - Where consumers open the portal.
	Url string `json:"url"`
}

// PortalsListOutputItems represents the portals list output items type.
type PortalsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you delete the portal.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Slug - URL-safe name, unique across Metorial.
	Slug string `json:"slug"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// AllowConsumerSkillAuthoring - Whether consumers can create skills on the portal.
	AllowConsumerSkillAuthoring bool `json:"allow_consumer_skill_authoring"`
	// AllowConsumerSkillPublishing - Whether consumers can publish skills on the portal.
	AllowConsumerSkillPublishing bool `json:"allow_consumer_skill_publishing"`
	// SkillConfiguration - Skill settings for the portal.
	SkillConfiguration PortalsListOutputItemsSkillConfiguration `json:"skill_configuration"`
	// Auth - Consumer authentication settings for the portal.
	Auth PortalsListOutputItemsAuth `json:"auth"`
	// Urls - Addresses consumers open the portal at.
	Urls []PortalsListOutputItemsUrls `json:"urls"`
	// MagicMcpUrl - The MCP connection URL for the portal.
	MagicMcpUrl string `json:"magic_mcp_url"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// PortalsListOutputPagination represents the portals list output pagination type.
type PortalsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// PortalsListOutput represents the portals list output type.
type PortalsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                      `json:"object"`
	Items      []PortalsListOutputItems    `json:"items"`
	Pagination PortalsListOutputPagination `json:"pagination"`
}

// MapPortalsListOutputFromJSON deserializes JSON data into a PortalsListOutput.
func MapPortalsListOutputFromJSON(data []byte) (*PortalsListOutput, error) {
	var v PortalsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListOutputToJSON serializes a PortalsListOutput to JSON.
func MapPortalsListOutputToJSON(v *PortalsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsListQuery represents the portals list query type.
type PortalsListQuery struct {
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
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
}

// MapPortalsListQueryFromJSON deserializes JSON data into a PortalsListQuery.
func MapPortalsListQueryFromJSON(data []byte) (*PortalsListQuery, error) {
	var v PortalsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListQueryToJSON serializes a PortalsListQuery to JSON.
func MapPortalsListQueryToJSON(v *PortalsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
