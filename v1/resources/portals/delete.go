package portals

import (
	"encoding/json"
	"time"
)

// PortalsDeleteOutputSkillConfiguration - Skill settings for the portal.
type PortalsDeleteOutputSkillConfiguration struct {
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

// PortalsDeleteOutputAuthAllowedRedirectUrlFilters represents the portals delete output auth allowed redirect url filters type.
type PortalsDeleteOutputAuthAllowedRedirectUrlFilters struct {
	// Url - A redirect URL pattern consumers may return to after signing in.
	Url string `json:"url"`
}

// PortalsDeleteOutputAuth - Consumer authentication settings for the portal.
type PortalsDeleteOutputAuth struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// SessionExpiryTimeInSeconds - How long consumer sessions stay valid, in seconds.
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
	// AllowedRedirectUrlFilters - Redirect URL patterns consumers may return to after signing in.
	AllowedRedirectUrlFilters []PortalsDeleteOutputAuthAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters"`
}

// PortalsDeleteOutputUrls represents the portals delete output urls type.
type PortalsDeleteOutputUrls struct {
	// Type - `namespace` for namespace-specific URLs, `default` otherwise.
	Type string `json:"type"`
	// Url - Where consumers open the portal.
	Url string `json:"url"`
}

// PortalsDeleteOutput represents the portals delete output type.
type PortalsDeleteOutput struct {
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
	SkillConfiguration PortalsDeleteOutputSkillConfiguration `json:"skill_configuration"`
	// Auth - Consumer authentication settings for the portal.
	Auth PortalsDeleteOutputAuth `json:"auth"`
	// Urls - Addresses consumers open the portal at.
	Urls []PortalsDeleteOutputUrls `json:"urls"`
	// MagicMcpUrl - The MCP connection URL for the portal.
	MagicMcpUrl string `json:"magic_mcp_url"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsDeleteOutputFromJSON deserializes JSON data into a PortalsDeleteOutput.
func MapPortalsDeleteOutputFromJSON(data []byte) (*PortalsDeleteOutput, error) {
	var v PortalsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsDeleteOutputToJSON serializes a PortalsDeleteOutput to JSON.
func MapPortalsDeleteOutputToJSON(v *PortalsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
