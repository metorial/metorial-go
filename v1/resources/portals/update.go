package portals

import (
	"encoding/json"
	"time"
)

// PortalsUpdateOutputSkillConfiguration - Skill settings for the portal.
type PortalsUpdateOutputSkillConfiguration struct {
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

// PortalsUpdateOutputAuthAllowedRedirectUrlFilters represents the portals update output auth allowed redirect url filters type.
type PortalsUpdateOutputAuthAllowedRedirectUrlFilters struct {
	// Url - A redirect URL pattern consumers may return to after signing in.
	Url string `json:"url"`
}

// PortalsUpdateOutputAuth - Consumer authentication settings for the portal.
type PortalsUpdateOutputAuth struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// SessionExpiryTimeInSeconds - How long consumer sessions stay valid, in seconds.
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
	// AllowedRedirectUrlFilters - Redirect URL patterns consumers may return to after signing in.
	AllowedRedirectUrlFilters []PortalsUpdateOutputAuthAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters"`
}

// PortalsUpdateOutputUrls represents the portals update output urls type.
type PortalsUpdateOutputUrls struct {
	// Type - `namespace` for namespace-specific URLs, `default` otherwise.
	Type string `json:"type"`
	// Url - Where consumers open the portal.
	Url string `json:"url"`
}

// PortalsUpdateOutput represents the portals update output type.
type PortalsUpdateOutput struct {
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
	SkillConfiguration PortalsUpdateOutputSkillConfiguration `json:"skill_configuration"`
	// Auth - Consumer authentication settings for the portal.
	Auth PortalsUpdateOutputAuth `json:"auth"`
	// Urls - Addresses consumers open the portal at.
	Urls []PortalsUpdateOutputUrls `json:"urls"`
	// MagicMcpUrl - The MCP connection URL for the portal.
	MagicMcpUrl string `json:"magic_mcp_url"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsUpdateOutputFromJSON deserializes JSON data into a PortalsUpdateOutput.
func MapPortalsUpdateOutputFromJSON(data []byte) (*PortalsUpdateOutput, error) {
	var v PortalsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsUpdateOutputToJSON serializes a PortalsUpdateOutput to JSON.
func MapPortalsUpdateOutputToJSON(v *PortalsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsUpdateBodyAllowedRedirectUrlFilters represents the portals update body allowed redirect url filters type.
type PortalsUpdateBodyAllowedRedirectUrlFilters struct {
	// Url - A redirect URL pattern consumers may return to after signing in.
	Url string `json:"url"`
}

// PortalsUpdateBodySkillConfiguration - Skill settings for the portal.
type PortalsUpdateBodySkillConfiguration struct {
	// AllowScripts - Whether skills on the portal can run scripts.
	AllowScripts *bool `json:"allow_scripts,omitempty"`
	// AllowedFileExtensions - File extensions skills on the portal can use.
	AllowedFileExtensions *[]string `json:"allowed_file_extensions,omitempty"`
	// AllowNonStandardDirectories - Whether skills can use directories outside the standard layout.
	AllowNonStandardDirectories *bool `json:"allow_non_standard_directories,omitempty"`
}

// PortalsUpdateBody represents the portals update body type.
type PortalsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// AllowedRedirectUrlFilters - Redirect URL patterns consumers may return to after signing in.
	AllowedRedirectUrlFilters *[]PortalsUpdateBodyAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters,omitempty"`
	// SessionExpiryTimeInSeconds - How long consumer sessions stay valid, in seconds.
	SessionExpiryTimeInSeconds *float64 `json:"session_expiry_time_in_seconds,omitempty"`
	// AllowConsumerSkillAuthoring - Whether consumers can create skills on the portal.
	AllowConsumerSkillAuthoring *bool `json:"allow_consumer_skill_authoring,omitempty"`
	// AllowConsumerSkillPublishing - Whether consumers can publish skills on the portal.
	AllowConsumerSkillPublishing *bool `json:"allow_consumer_skill_publishing,omitempty"`
	// SkillConfiguration - Skill settings for the portal.
	SkillConfiguration *PortalsUpdateBodySkillConfiguration `json:"skill_configuration,omitempty"`
}

// MapPortalsUpdateBodyFromJSON deserializes JSON data into a PortalsUpdateBody.
func MapPortalsUpdateBodyFromJSON(data []byte) (*PortalsUpdateBody, error) {
	var v PortalsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsUpdateBodyToJSON serializes a PortalsUpdateBody to JSON.
func MapPortalsUpdateBodyToJSON(v *PortalsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
