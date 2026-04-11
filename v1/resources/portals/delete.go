package portals

import (
	"encoding/json"
	"time"
)

// PortalsDeleteOutputAuthAllowedRedirectUrlFilters represents the portals delete output auth allowed redirect url filters type.
type PortalsDeleteOutputAuthAllowedRedirectUrlFilters struct {
	Url string `json:"url"`
}

// PortalsDeleteOutputAuth represents the portals delete output auth type.
type PortalsDeleteOutputAuth struct {
	Object                     string                                             `json:"object"`
	SessionExpiryTimeInSeconds float64                                            `json:"session_expiry_time_in_seconds"`
	AllowedRedirectUrlFilters  []PortalsDeleteOutputAuthAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters"`
}

// PortalsDeleteOutputUrls represents the portals delete output urls type.
type PortalsDeleteOutputUrls struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// PortalsDeleteOutput represents the portals delete output type.
type PortalsDeleteOutput struct {
	Object      string                    `json:"object"`
	Id          string                    `json:"id"`
	Status      string                    `json:"status"`
	Name        string                    `json:"name"`
	Slug        string                    `json:"slug"`
	Description *string                   `json:"description,omitempty"`
	Auth        PortalsDeleteOutputAuth   `json:"auth"`
	Urls        []PortalsDeleteOutputUrls `json:"urls"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
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
