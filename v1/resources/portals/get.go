package portals

import (
	"encoding/json"
	"time"
)

// PortalsGetOutputSkillConfiguration - Skill settings for the portal.
type PortalsGetOutputSkillConfiguration struct {
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

// PortalsGetOutputAuthAllowedRedirectUrlFilters represents the portals get output auth allowed redirect url filters type.
type PortalsGetOutputAuthAllowedRedirectUrlFilters struct {
	// Url - A redirect URL pattern consumers may return to after signing in.
	Url string `json:"url"`
}

// PortalsGetOutputAuth - Consumer authentication settings for the portal.
type PortalsGetOutputAuth struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// SessionExpiryTimeInSeconds - How long consumer sessions stay valid, in seconds.
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
	// AllowedRedirectUrlFilters - Redirect URL patterns consumers may return to after signing in.
	AllowedRedirectUrlFilters []PortalsGetOutputAuthAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters"`
}

// PortalsGetOutputUrls represents the portals get output urls type.
type PortalsGetOutputUrls struct {
	// Type - `namespace` for namespace-specific URLs, `default` otherwise.
	Type string `json:"type"`
	// Url - Where consumers open the portal.
	Url string `json:"url"`
}

// PortalsGetOutput represents the portals get output type.
type PortalsGetOutput struct {
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
	SkillConfiguration PortalsGetOutputSkillConfiguration `json:"skill_configuration"`
	// Auth - Consumer authentication settings for the portal.
	Auth PortalsGetOutputAuth `json:"auth"`
	// Urls - Addresses consumers open the portal at.
	Urls []PortalsGetOutputUrls `json:"urls"`
	// MagicMcpUrl - The MCP connection URL for the portal.
	MagicMcpUrl string `json:"magic_mcp_url"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsGetOutputFromJSON deserializes JSON data into a PortalsGetOutput.
func MapPortalsGetOutputFromJSON(data []byte) (*PortalsGetOutput, error) {
	var v PortalsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsGetOutputToJSON serializes a PortalsGetOutput to JSON.
func MapPortalsGetOutputToJSON(v *PortalsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
