package portals

import (
	"encoding/json"
	"time"
)

// PortalsCreateOutputSkillConfiguration - Skill settings for the portal.
type PortalsCreateOutputSkillConfiguration struct {
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

// PortalsCreateOutputAuthAllowedRedirectUrlFilters represents the portals create output auth allowed redirect url filters type.
type PortalsCreateOutputAuthAllowedRedirectUrlFilters struct {
	// Url - A redirect URL pattern consumers may return to after signing in.
	Url string `json:"url"`
}

// PortalsCreateOutputAuth - Consumer authentication settings for the portal.
type PortalsCreateOutputAuth struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// SessionExpiryTimeInSeconds - How long consumer sessions stay valid, in seconds.
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
	// AllowedRedirectUrlFilters - Redirect URL patterns consumers may return to after signing in.
	AllowedRedirectUrlFilters []PortalsCreateOutputAuthAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters"`
}

// PortalsCreateOutputUrls represents the portals create output urls type.
type PortalsCreateOutputUrls struct {
	// Type - `namespace` for namespace-specific URLs, `default` otherwise.
	Type string `json:"type"`
	// Url - Where consumers open the portal.
	Url string `json:"url"`
}

// PortalsCreateOutput represents the portals create output type.
type PortalsCreateOutput struct {
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
	SkillConfiguration PortalsCreateOutputSkillConfiguration `json:"skill_configuration"`
	// Auth - Consumer authentication settings for the portal.
	Auth PortalsCreateOutputAuth `json:"auth"`
	// Urls - Addresses consumers open the portal at.
	Urls []PortalsCreateOutputUrls `json:"urls"`
	// MagicMcpUrl - The MCP connection URL for the portal.
	MagicMcpUrl string `json:"magic_mcp_url"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsCreateOutputFromJSON deserializes JSON data into a PortalsCreateOutput.
func MapPortalsCreateOutputFromJSON(data []byte) (*PortalsCreateOutput, error) {
	var v PortalsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsCreateOutputToJSON serializes a PortalsCreateOutput to JSON.
func MapPortalsCreateOutputToJSON(v *PortalsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsCreateBodyAllowedRedirectUrlFilters represents the portals create body allowed redirect url filters type.
type PortalsCreateBodyAllowedRedirectUrlFilters struct {
	// Url - A redirect URL pattern consumers may return to after signing in.
	Url string `json:"url"`
}

// PortalsCreateBody represents the portals create body type.
type PortalsCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// AllowedRedirectUrlFilters - Redirect URL patterns consumers may return to after signing in.
	AllowedRedirectUrlFilters *[]PortalsCreateBodyAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters,omitempty"`
	// SessionExpiryTimeInSeconds - How long consumer sessions stay valid, in seconds.
	SessionExpiryTimeInSeconds *float64 `json:"session_expiry_time_in_seconds,omitempty"`
	// AllowConsumerSkillAuthoring - Whether consumers can create skills on the portal.
	AllowConsumerSkillAuthoring *bool `json:"allow_consumer_skill_authoring,omitempty"`
	// AllowConsumerSkillPublishing - Whether consumers can publish skills on the portal.
	AllowConsumerSkillPublishing *bool `json:"allow_consumer_skill_publishing,omitempty"`
}

// MapPortalsCreateBodyFromJSON deserializes JSON data into a PortalsCreateBody.
func MapPortalsCreateBodyFromJSON(data []byte) (*PortalsCreateBody, error) {
	var v PortalsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsCreateBodyToJSON serializes a PortalsCreateBody to JSON.
func MapPortalsCreateBodyToJSON(v *PortalsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
